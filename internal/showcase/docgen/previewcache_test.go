package docgen

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseFile_templFenceBecomesPreviewBlock(t *testing.T) {
	raw := []byte(`---
slug: sample
section: components
title: Sample
description: Sample page.
---

## Example

` + "```templ\n" + `import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{}) {
		Button
	}
}
` + "```\n")

	page, err := ParseFile("en", "en/components/sample.md", raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Blocks) != 2 {
		t.Fatalf("blocks: got %d", len(page.Blocks))
	}
	pb, ok := page.Blocks[1].(PreviewCodeBlock)
	if !ok {
		t.Fatalf("expected PreviewCodeBlock, got %T", page.Blocks[1])
	}
	if pb.Language != "templ" {
		t.Fatalf("language: %q", pb.Language)
	}
	if !strings.Contains(pb.Source, "templ Example()") {
		t.Fatalf("source missing Example(): %q", pb.Source)
	}
}

func TestParseFile_nonTemplFenceStaysCodeBlock(t *testing.T) {
	raw := []byte(`---
slug: sample
section: components
title: Sample
description: Sample page.
---

` + "```go\n" + `package main
` + "```\n")

	page, err := ParseFile("en", "en/components/sample.md", raw)
	if err != nil {
		t.Fatal(err)
	}
	cb, ok := page.Blocks[0].(CodeBlock)
	if !ok {
		t.Fatalf("expected CodeBlock, got %T", page.Blocks[0])
	}
	if cb.Language != "go" {
		t.Fatalf("language: %q", cb.Language)
	}
}

func TestValidateTemplExample_missingExample(t *testing.T) {
	err := ValidateTemplExample("test.md", `import "github.com/fastygo/templ/ui"

templ ButtonDemo() {
}
`)
	if err == nil {
		t.Fatal("expected error for missing Example()")
	}
}

func TestParseFile_buttonMdHasNoDemoDirectives(t *testing.T) {
	raw, err := os.ReadFile(filepath.Join("..", "content", "en", "components", "button.md"))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), "{{demo") {
		t.Fatal("button.md still contains legacy demo directives")
	}
	page, err := ParseFile("en", "en/components/button.md", raw)
	if err != nil {
		t.Fatal(err)
	}
	var previews int
	for _, b := range page.Blocks {
		if _, ok := b.(PreviewCodeBlock); ok {
			previews++
		}
	}
	if previews != 7 {
		t.Fatalf("preview blocks: got %d want 7", previews)
	}
}

func TestCompilePreviews_rendersButtonExample(t *testing.T) {
	root, err := findRepoRoot(t)
	if err != nil {
		t.Fatal(err)
	}
	pages := []DocPage{{
		Locale:     "en",
		SourceFile: "test/button.md",
		Blocks: []Block{
			PreviewCodeBlock{
				Language: "templ",
				Source: `import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{}) {
		Button
	}
}`,
			},
		},
	}}
	stats, err := CompilePreviews(pages, PreviewCacheConfig{ModuleRoot: root})
	if err != nil {
		t.Fatal(err)
	}
	if stats.Compiled != 1 {
		t.Fatalf("compiled: got %d want 1", stats.Compiled)
	}
	pb, ok := pages[0].Blocks[0].(PreviewCodeBlock)
	if !ok {
		t.Fatalf("expected PreviewCodeBlock, got %T", pages[0].Blocks[0])
	}
	if strings.TrimSpace(pb.HTML) == "" {
		t.Fatal("expected non-empty preview HTML")
	}
	if !strings.Contains(pb.HTML, "button") {
		t.Fatalf("preview HTML missing button element: %q", pb.HTML)
	}
}

func TestCompilePreviews_cacheHit(t *testing.T) {
	root, err := findRepoRoot(t)
	if err != nil {
		t.Fatal(err)
	}
	pages := []DocPage{{
		Locale:     "en",
		SourceFile: "test/cache-hit.md",
		Blocks: []Block{
			PreviewCodeBlock{
				Language: "templ",
				Source: `import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{}) {
		Cache hit probe
	}
}`,
			},
		},
	}}
	first, err := CompilePreviews(pages, PreviewCacheConfig{ModuleRoot: root, CleanStore: true})
	if err != nil {
		t.Fatal(err)
	}
	if first.Compiled != 1 {
		t.Fatalf("first compiled: got %d want 1", first.Compiled)
	}
	second, err := CompilePreviews(pages, PreviewCacheConfig{ModuleRoot: root})
	if err != nil {
		t.Fatal(err)
	}
	if second.Cached != 1 {
		t.Fatalf("second cached: got %d want 1", second.Cached)
	}
	if second.Compiled != 0 {
		t.Fatalf("second compiled: got %d want 0", second.Compiled)
	}
}

func findRepoRoot(t *testing.T) (string, error) {
	t.Helper()
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
			return "", os.ErrNotExist
		}
		dir = parent
	}
}
