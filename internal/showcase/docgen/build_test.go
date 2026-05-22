package docgen_test

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/fastygo/ui/internal/showcase/docgen"
	"github.com/fastygo/ui/internal/showcase/previews"

	_ "github.com/fastygo/ui/internal/showcase"
)

func TestBuild_blogCard(t *testing.T) {
	previews.Reset()
	if err := previews.RegisterFromRegistry(); err != nil {
		t.Fatal(err)
	}
	pages, err := docgen.LoadAll(docgen.LoadOptions{Locales: []string{"en"}})
	if err != nil {
		t.Fatal(err)
	}
	if err := docgen.ResolveDemos(pages); err != nil {
		t.Fatal(err)
	}
	if err := docgen.CompilePreviews(pages, docgen.PreviewCacheConfig{}); err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	if err := docgen.Build(context.Background(), pages, docgen.BuildConfig{
		OutputDir: dir,
		Locales:   []string{"en"},
	}); err != nil {
		t.Fatal(err)
	}
	htmlPath := filepath.Join(dir, "components", "blog-card", "index.html")
	raw, err := os.ReadFile(htmlPath)
	if err != nil {
		t.Fatal(err)
	}
	html := string(raw)
	for _, want := range []string{"<!doctype html>", "Blog Card", "Read more"} {
		if !strings.Contains(html, want) {
			t.Fatalf("html missing %q (%d bytes)", want, len(html))
		}
	}
}
