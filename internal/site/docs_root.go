package site

import (
	"log/slog"
	"os"
	"path/filepath"
)

// resolveDocsRoot returns the directory that contains locale subfolders (en/, ru/, …).
func resolveDocsRoot(primary string) string {
	candidates := []string{}
	if primary != "" {
		candidates = append(candidates, filepath.Clean(primary))
	}
	if wd, err := os.Getwd(); err == nil {
		candidates = append(candidates, filepath.Join(wd, "web", "static", "docs"))
	}
	seen := map[string]struct{}{}
	for _, root := range candidates {
		if root == "" {
			continue
		}
		if abs, err := filepath.Abs(root); err == nil {
			root = abs
		}
		if _, ok := seen[root]; ok {
			continue
		}
		seen[root] = struct{}{}
		if docsRootValid(root) {
			return root
		}
	}
	if primary != "" {
		if abs, err := filepath.Abs(primary); err == nil {
			return abs
		}
		return filepath.Clean(primary)
	}
	return ""
}

func docsRootValid(root string) bool {
	for _, locale := range []string{"en", "ru"} {
		if _, err := os.Stat(filepath.Join(root, locale, "index.html")); err == nil {
			return true
		}
	}
	return false
}

func logDocsRoot(root string) {
	if root == "" {
		slog.Warn("docs: output directory not found; /docs/ routes disabled (run go run ./cmd/docgen)")
		return
	}
	index := filepath.Join(root, "en", "index.html")
	if _, err := os.Stat(index); err != nil {
		slog.Warn("docs: locale index missing", "root", root, "file", index, "err", err)
		return
	}
	slog.Info("docs: serving static HTML", "root", root)
}
