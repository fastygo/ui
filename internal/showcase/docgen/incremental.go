package docgen

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"

	"github.com/fastygo/ui/internal/doclocale"
)

const (
	previewStoreRoot = ".internal/docgen/docpreviews/store"
	pageStampRoot    = ".internal/docgen/stamps"
)

func contentHash(raw []byte) string {
	sum := sha256.Sum256(raw)
	return hex.EncodeToString(sum[:])
}

func previewSourceHash(source string) string {
	sum := sha256.Sum256([]byte(source))
	return hex.EncodeToString(sum[:])
}

func globalBuildInput(root string) (string, error) {
	files := []string{
		"internal/fixtures/locale/en.json",
		"internal/fixtures/locale/ru.json",
		"internal/views/docsstatic/render.go",
		"internal/views/docsstatic/types.go",
		"internal/views/partials/shell_head.templ",
		"internal/views/layout.templ",
		"internal/ui/layout/shell.templ",
		"internal/showcase/docgen/nav.go",
		"internal/showcase/docgen/convert.go",
		"internal/showcase/docgen/highlight.go",
		"web/static/css/tweakcn.css",
		"web/static/css/code.css",
		"web/static/css/docs-preview.css",
		"web/static/css/prose.css",
		"web/static/js/ui8kit.js",
	}
	h := sha256.New()
	for _, rel := range files {
		path := filepath.Join(root, filepath.FromSlash(rel))
		raw, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}
		h.Write([]byte(rel))
		h.Write(raw)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func pageBuildInput(page DocPage, globalHash string) string {
	h := sha256.New()
	h.Write([]byte(page.Locale))
	h.Write([]byte(page.ContentHash))
	h.Write([]byte(globalHash))
	for _, b := range page.Blocks {
		if pb, ok := b.(PreviewCodeBlock); ok {
			h.Write([]byte(pb.ID))
			h.Write([]byte(previewSourceHash(pb.Source)))
		}
	}
	return hex.EncodeToString(h.Sum(nil))
}

func pageStampPath(root, outputHTMLPath string) string {
	sum := sha256.Sum256([]byte(filepath.ToSlash(outputHTMLPath)))
	key := hex.EncodeToString(sum[:])
	return filepath.Join(root, filepath.FromSlash(pageStampRoot), key+".stamp")
}

func pageUpToDate(root, outputHTMLPath, inputHash string) bool {
	if _, err := os.Stat(outputHTMLPath); err != nil {
		return false
	}
	stampPath := pageStampPath(root, outputHTMLPath)
	raw, err := os.ReadFile(stampPath)
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(raw)) == inputHash
}

func writePageStamp(root, outputHTMLPath, inputHash string) error {
	stampPath := pageStampPath(root, outputHTMLPath)
	if err := os.MkdirAll(filepath.Dir(stampPath), 0o755); err != nil {
		return err
	}
	return os.WriteFile(stampPath, []byte(inputHash), 0o644)
}

func loadCachedPreviewHTML(dir, wantSourceHash string) (string, bool) {
	sumPath := filepath.Join(dir, "source.sha256")
	htmlPath := filepath.Join(dir, "preview.html")
	sumRaw, err := os.ReadFile(sumPath)
	if err != nil {
		return "", false
	}
	if strings.TrimSpace(string(sumRaw)) != wantSourceHash {
		return "", false
	}
	html, err := os.ReadFile(htmlPath)
	if err != nil {
		return "", false
	}
	return string(html), true
}

func writePreviewSourceStamp(dir, sourceHash string) error {
	return os.WriteFile(filepath.Join(dir, "source.sha256"), []byte(sourceHash), 0o644)
}

func indexBuildInput(locale string, pages []DocPage, globalHash string) string {
	h := sha256.New()
	h.Write([]byte(locale))
	h.Write([]byte(globalHash))
	for _, p := range pages {
		h.Write([]byte(pageBuildInput(p, globalHash)))
	}
	return hex.EncodeToString(h.Sum(nil))
}

func indexOutputPath(outRoot string, routing doclocale.Routing, locale string) string {
	return filepath.Join(outRoot, filepath.FromSlash(routing.IndexOutputRelPath(locale)))
}
