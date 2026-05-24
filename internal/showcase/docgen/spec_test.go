package docgen_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/fastygo/ui/internal/showcase/docgen"
)

func TestParseFile_specBlock(t *testing.T) {
	raw, err := os.ReadFile(filepath.Join("..", "content", "en", "primitives", "block.md"))
	if err != nil {
		t.Fatal(err)
	}
	page, err := docgen.ParseFile(docgen.DefaultRouting(), "en", "en/primitives/block.md", raw)
	if err != nil {
		t.Fatal(err)
	}
	if page.Meta.Slug != "block" {
		t.Fatalf("slug: %q", page.Meta.Slug)
	}
	if page.Meta.Section != "primitives" {
		t.Fatalf("section: %q", page.Meta.Section)
	}
	if page.Meta.Title != "Block" {
		t.Fatalf("title: %q", page.Meta.Title)
	}
	if page.PublicPath != "/docs/primitives/block/" {
		t.Fatalf("public: %q", page.PublicPath)
	}
	if len(page.Meta.API) == 0 {
		t.Fatal("expected flattened API fields")
	}
}
