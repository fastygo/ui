package docgen_test

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/fastygo/ui/internal/showcase/docgen"
)

func TestBuild_blogCard(t *testing.T) {
	pages, err := docgen.LoadAll(docgen.LoadOptions{Locales: []string{"en"}})
	if err != nil {
		t.Fatal(err)
	}
	if err := docgen.HighlightCodeBlocks(pages); err != nil {
		t.Fatal(err)
	}
	docgen.StubPreviewHTML(pages)
	dir := t.TempDir()
	if _, err := docgen.Build(context.Background(), pages, docgen.BuildConfig{
		OutputDir:     dir,
		Locales:       []string{"en"},
		DefaultLocale: "en",
	}); err != nil {
		t.Fatal(err)
	}
	htmlPath := filepath.Join(dir, "en", "components", "blog-card", "index.html")
	raw, err := os.ReadFile(htmlPath)
	if err != nil {
		t.Fatal(err)
	}
	html := string(raw)
	for _, want := range []string{"<!doctype html>", "Blog Card", `aria-label="On this page"`, `href="#`, `data-preview-stub="1"`} {
		if !strings.Contains(html, want) {
			t.Fatalf("html missing %q (%d bytes)", want, len(html))
		}
	}
}

func TestBuild_incrementalSkipsUnchanged(t *testing.T) {
	pages, err := docgen.LoadAll(docgen.LoadOptions{Locales: []string{"en"}})
	if err != nil {
		t.Fatal(err)
	}
	if err := docgen.HighlightCodeBlocks(pages); err != nil {
		t.Fatal(err)
	}
	docgen.StubPreviewHTML(pages)
	dir := t.TempDir()
	first, err := docgen.Build(context.Background(), pages, docgen.BuildConfig{
		OutputDir:     dir,
		Locales:       []string{"en"},
		DefaultLocale: "en",
		Incremental:   true,
	})
	if err != nil {
		t.Fatal(err)
	}
	if first.PagesWritten == 0 {
		t.Fatal("expected pages written on first build")
	}

	second, err := docgen.Build(context.Background(), pages, docgen.BuildConfig{
		OutputDir:     dir,
		Locales:       []string{"en"},
		DefaultLocale: "en",
		Incremental:   true,
	})
	if err != nil {
		t.Fatal(err)
	}
	if second.PagesWritten != 0 {
		t.Fatalf("second build wrote %d pages, want 0", second.PagesWritten)
	}
	if second.PagesSkipped <= 0 {
		t.Fatalf("second build skipped %d pages, want > 0", second.PagesSkipped)
	}
}
