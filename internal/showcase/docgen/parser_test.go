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
	if len(page.DemoIDs) != 2 {
		t.Fatalf("demos: got %d", len(page.DemoIDs))
	}
	if page.DemoIDs[0] != "blog-card.vertical" {
		t.Fatalf("first demo: %q", page.DemoIDs[0])
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
