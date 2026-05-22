package docgen

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFile_blogCard(t *testing.T) {
	raw, err := os.ReadFile(filepath.Join("..", "content", "en", "components", "blog-card.md"))
	if err != nil {
		t.Fatal(err)
	}
	page, err := ParseFile("en", "en/components/blog-card.md", raw)
	if err != nil {
		t.Fatal(err)
	}
	if page.Meta.Slug != "blog-card" {
		t.Fatalf("slug: got %q", page.Meta.Slug)
	}
	var previews int
	for _, b := range page.Blocks {
		if _, ok := b.(PreviewCodeBlock); ok {
			previews++
		}
	}
	if previews != 2 {
		t.Fatalf("preview blocks: got %d want 2", previews)
	}
}

func TestValidateBodyRules_rejectsClassInProse(t *testing.T) {
	body := "Use `class=\"flex\"` in prose.\n"
	err := ValidateBodyRules("test.md", body)
	if err == nil {
		t.Fatal("expected error for class in prose")
	}
}

func TestValidateBodyRules_allowsCodeFence(t *testing.T) {
	body := "Example:\n\n```go\n<body class=\"bg-background\">\n</body>\n```\n"
	if err := ValidateBodyRules("test.md", body); err != nil {
		t.Fatal(err)
	}
}
