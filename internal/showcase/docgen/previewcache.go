package docgen

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const previewBatchRoot = ".internal/docgen/docpreviews/batch"

var (
	templExampleRe = regexp.MustCompile(`(?m)^\s*templ\s+Example\s*\(\s*\)\s*\{`)
	importBlockRe  = regexp.MustCompile(`(?ms)^import\s+(?:\(\s*[\s\S]*?\)|[\w.]+\s+"[^"]+"|"[^"]+")\s*`)
)

// PreviewCacheConfig controls persistent preview compilation.
type PreviewCacheConfig struct {
	ModuleRoot string
	Force      bool // recompile every preview even when cached
	CleanStore bool // wipe persistent preview store before build
}

// CompilePreviewsStats reports preview cache hits and misses.
type CompilePreviewsStats struct {
	Total    int
	Cached   int
	Compiled int
	Unique   int // distinct source hashes compiled in the last batch run
}

type previewRef struct {
	pageIndex  int
	blockIndex int
	block      PreviewCodeBlock
}

type batchCompileEntry struct {
	componentName string
	srcHash       string
	source        string
	refs          []previewRef
}

// CompilePreviews renders every PreviewCodeBlock on pages through a persistent store.
func CompilePreviews(pages []DocPage, cfg PreviewCacheConfig) (CompilePreviewsStats, error) {
	var stats CompilePreviewsStats
	root, err := resolveModuleRoot(cfg.ModuleRoot)
	if err != nil {
		return stats, err
	}
	storeAbs := filepath.Join(root, filepath.FromSlash(previewStoreRoot))
	if cfg.CleanStore {
		if err := os.RemoveAll(filepath.Join(root, filepath.FromSlash(".internal/docgen/docpreviews"))); err != nil {
			return stats, fmt.Errorf("preview store: clear: %w", err)
		}
	}
	if err := os.MkdirAll(storeAbs, 0o755); err != nil {
		return stats, fmt.Errorf("preview store: mkdir: %w", err)
	}

	pending := map[string]*batchCompileEntry{}
	var pendingOrder []string

	for pi := range pages {
		fenceIndex := 0
		for bi, b := range pages[pi].Blocks {
			pb, ok := b.(PreviewCodeBlock)
			if !ok {
				continue
			}
			fenceIndex++
			stats.Total++
			pb.FenceIndex = fenceIndex
			pb.ID = previewCacheID(pages[pi].SourceFile, fenceIndex, pb.Source)
			srcHash := previewSourceHash(pb.Source)
			cacheDir := previewCacheDir(storeAbs, srcHash)

			if !cfg.Force {
				if html, ok := loadCachedPreviewHTML(cacheDir, srcHash); ok {
					pb.HTML = html
					pages[pi].Blocks[bi] = pb
					stats.Cached++
					continue
				}
			}

			ref := previewRef{pageIndex: pi, blockIndex: bi, block: pb}
			entry, ok := pending[srcHash]
			if !ok {
				name := fmt.Sprintf("Preview%03d", len(pendingOrder)+1)
				entry = &batchCompileEntry{
					componentName: name,
					srcHash:       srcHash,
					source:        pb.Source,
				}
				pending[srcHash] = entry
				pendingOrder = append(pendingOrder, srcHash)
			}
			entry.refs = append(entry.refs, ref)
		}
	}

	if len(pendingOrder) == 0 {
		return stats, nil
	}

	entries := make([]batchCompileEntry, 0, len(pendingOrder))
	for _, hash := range pendingOrder {
		entries = append(entries, *pending[hash])
	}
	stats.Unique = len(entries)

	if err := compilePreviewBatch(root, storeAbs, entries); err != nil {
		return stats, err
	}

	for _, entry := range entries {
		cacheDir := previewCacheDir(storeAbs, entry.srcHash)
		html, ok := loadCachedPreviewHTML(cacheDir, entry.srcHash)
		if !ok {
			return stats, fmt.Errorf("preview store: missing compiled output for hash %s", entry.srcHash)
		}
		for _, ref := range entry.refs {
			ref.block.HTML = html
			pages[ref.pageIndex].Blocks[ref.blockIndex] = ref.block
			stats.Compiled++
		}
	}

	return stats, nil
}

func compilePreviewBatch(root, storeAbs string, entries []batchCompileEntry) error {
	batchAbs := filepath.Join(root, filepath.FromSlash(previewBatchRoot))
	if err := os.RemoveAll(batchAbs); err != nil {
		return fmt.Errorf("preview batch: clear: %w", err)
	}
	if err := os.MkdirAll(batchAbs, 0o755); err != nil {
		return fmt.Errorf("preview batch: mkdir: %w", err)
	}

	templPath := filepath.Join(batchAbs, "batch.templ")
	if err := os.WriteFile(templPath, []byte(buildBatchTempl(entries)), 0o644); err != nil {
		return err
	}
	renderPath := filepath.Join(batchAbs, "render.go")
	if err := os.WriteFile(renderPath, []byte(buildBatchRenderGo(entries)), 0o644); err != nil {
		return err
	}

	if err := runCmd(root, "go", "tool", "templ", "generate", "-path", filepath.ToSlash(previewBatchRoot)); err != nil {
		return fmt.Errorf("preview batch: templ generate: %w", err)
	}

	cmd := exec.Command("go", "run", ".")
	cmd.Dir = batchAbs
	cmd.Env = append(os.Environ(), "PREVIEW_STORE_ROOT="+storeAbs)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		msg := strings.TrimSpace(stderr.String())
		if msg != "" {
			return fmt.Errorf("preview batch: render: %v: %s", err, msg)
		}
		return fmt.Errorf("preview batch: render: %w", err)
	}
	return nil
}

func previewCacheDir(storeAbs, srcHash string) string {
	return filepath.Join(storeAbs, srcHash)
}

func buildBatchTempl(entries []batchCompileEntry) string {
	imports := collectUniqueImports(entries)
	var b strings.Builder
	b.WriteString("package main\n\n")
	for _, imp := range imports {
		b.WriteString(imp)
		b.WriteByte('\n')
	}
	if len(imports) > 0 {
		b.WriteByte('\n')
	}
	for _, entry := range entries {
		body, err := extractTemplExampleBody(entry.source)
		if err != nil {
			body = fmt.Sprintf("// compile error: %v", err)
		}
		if ref := firstPreviewRef(entry); ref != nil {
			b.WriteString(fmt.Sprintf("// Source: %s fence #%d hash %s\n", ref.block.SourceFile, ref.block.FenceIndex, entry.srcHash))
		}
		b.WriteString(fmt.Sprintf("templ %s() {\n", entry.componentName))
		b.WriteString(body)
		b.WriteByte('\n')
		b.WriteString("}\n\n")
	}
	return b.String()
}

func firstPreviewRef(entry batchCompileEntry) *previewRef {
	if len(entry.refs) == 0 {
		return nil
	}
	return &entry.refs[0]
}

func collectUniqueImports(entries []batchCompileEntry) []string {
	seen := map[string]struct{}{}
	var out []string
	for _, entry := range entries {
		for _, block := range importBlockRe.FindAllString(entry.source, -1) {
			block = strings.TrimSpace(block)
			if block == "" {
				continue
			}
			if _, ok := seen[block]; ok {
				continue
			}
			seen[block] = struct{}{}
			out = append(out, block)
		}
	}
	sort.Strings(out)
	return out
}

func extractTemplExampleBody(source string) (string, error) {
	loc := templExampleRe.FindStringIndex(source)
	if loc == nil {
		return "", fmt.Errorf("missing templ Example()")
	}
	rest := source[loc[0]:]
	open := strings.Index(rest, "{")
	if open < 0 {
		return "", fmt.Errorf("missing opening brace")
	}
	depth := 0
	for i := open; i < len(rest); i++ {
		switch rest[i] {
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return strings.TrimSpace(rest[open+1 : i]), nil
			}
		}
	}
	return "", fmt.Errorf("unbalanced braces in templ Example()")
}

func buildBatchRenderGo(entries []batchCompileEntry) string {
	var b strings.Builder
	b.WriteString(`package main

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
)

func main() {
	storeRoot := os.Getenv("PREVIEW_STORE_ROOT")
	if storeRoot == "" {
		panic("PREVIEW_STORE_ROOT required")
	}
	ctx := context.Background()
`)
	for _, entry := range entries {
		b.WriteString(fmt.Sprintf("\t{\n\t\tvar buf bytes.Buffer\n\t\tif err := %s().Render(ctx, &buf); err != nil {\n\t\t\tpanic(%q + \": \" + err.Error())\n\t\t}\n\t\tif err := persistPreview(storeRoot, %q, buf.Bytes()); err != nil {\n\t\t\tpanic(err)\n\t\t}\n\t}\n", entry.componentName, entry.srcHash, entry.srcHash))
	}
	b.WriteString("}\n\n")
	b.WriteString(`func persistPreview(storeRoot, hash string, html []byte) error {
	dir := filepath.Join(storeRoot, hash)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(dir, "preview.html"), html, 0o644); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "source.sha256"), []byte(hash), 0o644)
}
`)
	return b.String()
}

// StubPreviewHTML fills empty preview HTML with a minimal placeholder for fast build tests.
func StubPreviewHTML(pages []DocPage) {
	const stub = `<span data-preview-stub="1"></span>`
	for pi := range pages {
		for bi, b := range pages[pi].Blocks {
			pb, ok := b.(PreviewCodeBlock)
			if !ok || strings.TrimSpace(pb.HTML) != "" {
				continue
			}
			pb.HTML = stub
			pages[pi].Blocks[bi] = pb
		}
	}
}

func runCmd(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Env = os.Environ()
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		msg := strings.TrimSpace(stderr.String())
		if msg != "" {
			return fmt.Errorf("%v: %s", err, msg)
		}
		return err
	}
	return nil
}

func resolveModuleRoot(explicit string) (string, error) {
	if explicit != "" {
		return filepath.Abs(explicit)
	}
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod not found from %s", dir)
		}
		dir = parent
	}
}

func previewCacheID(sourceFile string, fenceIndex int, source string) string {
	sum := sha256.Sum256([]byte(source))
	hash := hex.EncodeToString(sum[:4])
	safe := strings.Map(func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			return r
		default:
			return '_'
		}
	}, strings.ToLower(strings.ReplaceAll(sourceFile, "/", "_")))
	return fmt.Sprintf("%s_%03d_%s", strings.Trim(safe, "_"), fenceIndex, hash)
}

// ValidateTemplExample ensures a templ fence defines Example().
func ValidateTemplExample(sourceFile, source string) error {
	if !templExampleRe.MatchString(source) {
		return fmt.Errorf("%s: templ fence must define templ Example()", sourceFile)
	}
	return nil
}
