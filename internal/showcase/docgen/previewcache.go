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

const previewCacheRoot = ".internal/docgen/docpreviews/cache"

var templExampleRe = regexp.MustCompile(`(?m)^\s*templ\s+Example\s*\(\s*\)\s*\{`)

// PreviewCacheConfig controls disposable preview compilation.
type PreviewCacheConfig struct {
	ModuleRoot string
	KeepCache  bool
}

// CompilePreviews renders every PreviewCodeBlock on pages through a temporary cache.
func CompilePreviews(pages []DocPage, cfg PreviewCacheConfig) error {
	root, err := resolveModuleRoot(cfg.ModuleRoot)
	if err != nil {
		return err
	}
	cacheAbs := filepath.Join(root, filepath.FromSlash(previewCacheRoot))
	if err := os.RemoveAll(cacheAbs); err != nil {
		return fmt.Errorf("preview cache: clear: %w", err)
	}
	if err := os.MkdirAll(cacheAbs, 0o755); err != nil {
		return fmt.Errorf("preview cache: mkdir: %w", err)
	}

	type job struct {
		pageIndex  int
		blockIndex int
		block      PreviewCodeBlock
		dir        string
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
			pb.FenceIndex = fenceIndex
			pb.ID = previewCacheID(pages[pi].SourceFile, fenceIndex, pb.Source)
			dir := filepath.Join(cacheAbs, pb.ID)
			if err := os.MkdirAll(dir, 0o755); err != nil {
				return err
			}
			if err := writePreviewPackage(dir, pages[pi].SourceFile, fenceIndex, pb.Source); err != nil {
				return err
			}
			jobs = append(jobs, job{pageIndex: pi, blockIndex: bi, block: pb, dir: dir})
		}
	}

	if len(jobs) == 0 {
		if !cfg.KeepCache {
			_ = os.RemoveAll(cacheAbs)
		}
		return nil
	}

	if err := runCmd(root, "go", "tool", "templ", "generate", "-path", filepath.ToSlash(previewCacheRoot)); err != nil {
		return fmt.Errorf("preview cache: templ generate: %w", err)
	}

	for _, j := range jobs {
		if err := runCmd(j.dir, "go", "run", "."); err != nil {
			return fmt.Errorf("preview cache: render %s: %w", j.block.ID, err)
		}
		htmlPath := filepath.Join(j.dir, "preview.html")
		html, err := os.ReadFile(htmlPath)
		if err != nil {
			return fmt.Errorf("preview cache: read %s: %w", htmlPath, err)
		}
		j.block.HTML = string(html)
		pages[j.pageIndex].Blocks[j.blockIndex] = j.block
	}

	if !cfg.KeepCache {
		if err := os.RemoveAll(cacheAbs); err != nil {
			return fmt.Errorf("preview cache: cleanup: %w", err)
		}
	}
	return nil
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
