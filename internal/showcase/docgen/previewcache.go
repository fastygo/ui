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
	"strings"
)

var templExampleRe = regexp.MustCompile(`(?m)^\s*templ\s+Example\s*\(\s*\)\s*\{`)

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
		if err := os.RemoveAll(storeAbs); err != nil {
			return stats, fmt.Errorf("preview store: clear: %w", err)
		}
	}
	if err := os.MkdirAll(storeAbs, 0o755); err != nil {
		return stats, fmt.Errorf("preview store: mkdir: %w", err)
	}

	type job struct {
		pageIndex  int
		blockIndex int
		block      PreviewCodeBlock
		dir        string
		srcHash    string
	}
	var jobs []job

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
			dir := filepath.Join(storeAbs, pb.ID)

			if !cfg.Force {
				if html, ok := loadCachedPreviewHTML(dir, srcHash); ok {
					pb.HTML = html
					pages[pi].Blocks[bi] = pb
					stats.Cached++
					continue
				}
			}

			if err := os.MkdirAll(dir, 0o755); err != nil {
				return stats, err
			}
			if err := writePreviewPackage(dir, pages[pi].SourceFile, fenceIndex, pb.Source); err != nil {
				return stats, err
			}
			jobs = append(jobs, job{
				pageIndex:  pi,
				blockIndex: bi,
				block:      pb,
				dir:        dir,
				srcHash:    srcHash,
			})
		}
	}

	if len(jobs) == 0 {
		return stats, nil
	}

	if err := runCmd(root, "go", "tool", "templ", "generate", "-path", filepath.ToSlash(previewStoreRoot)); err != nil {
		return stats, fmt.Errorf("preview store: templ generate: %w", err)
	}

	for _, j := range jobs {
		if err := runCmd(j.dir, "go", "run", "."); err != nil {
			return stats, fmt.Errorf("preview store: render %s: %w", j.block.ID, err)
		}
		htmlPath := filepath.Join(j.dir, "preview.html")
		html, err := os.ReadFile(htmlPath)
		if err != nil {
			return stats, fmt.Errorf("preview store: read %s: %w", htmlPath, err)
		}
		if err := writePreviewSourceStamp(j.dir, j.srcHash); err != nil {
			return stats, fmt.Errorf("preview store: stamp %s: %w", j.block.ID, err)
		}
		j.block.HTML = string(html)
		pages[j.pageIndex].Blocks[j.blockIndex] = j.block
		stats.Compiled++
	}

	return stats, nil
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

func writePreviewPackage(dir, sourceFile string, fenceIndex int, authorSource string) error {
	templPath := filepath.Join(dir, "preview.templ")
	content := buildPreviewTempl(sourceFile, fenceIndex, authorSource)
	if err := os.WriteFile(templPath, []byte(content), 0o644); err != nil {
		return err
	}
	renderPath := filepath.Join(dir, "render.go")
	return os.WriteFile(renderPath, []byte(renderMainGo), 0o644)
}

func buildPreviewTempl(sourceFile string, fenceIndex int, authorSource string) string {
	var b strings.Builder
	b.WriteString("package main\n\n")
	b.WriteString(fmt.Sprintf("// Source: %s fence #%d\n\n", sourceFile, fenceIndex))
	b.WriteString(strings.TrimSpace(authorSource))
	b.WriteString("\n")
	return b.String()
}

const renderMainGo = `package main

import (
	"bytes"
	"context"
	"os"
)

func main() {
	var buf bytes.Buffer
	if err := Example().Render(context.Background(), &buf); err != nil {
		panic(err)
	}
	if err := os.WriteFile("preview.html", buf.Bytes(), 0o644); err != nil {
		panic(err)
	}
}
`

// ValidateTemplExample ensures a templ fence defines Example().
func ValidateTemplExample(sourceFile, source string) error {
	if !templExampleRe.MatchString(source) {
		return fmt.Errorf("%s: templ fence must define templ Example()", sourceFile)
	}
	return nil
}
