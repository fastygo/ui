package docgen

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestHighlightSource_go(t *testing.T) {
	html, err := HighlightSource("go", `package main

func main() {}
`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(html, `<pre`) {
		t.Fatalf("expected pre wrapper: %q", html)
	}
	if !strings.Contains(html, `class="`) {
		t.Fatalf("expected chroma classes: %q", html)
	}
	if strings.Contains(html, "package main") {
		t.Fatalf("expected escaped/tokenized output, got raw source")
	}
}

func TestHighlightSource_templAtTokenNotError(t *testing.T) {
	html, err := HighlightSource("templ", `@ui.TableBody(ui.TableSectionProps{}) {
}
`)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(html, `class="err">@`) {
		t.Fatalf("templ @ should not use error token: %q", html)
	}
	if !strings.Contains(html, `class="x">@`) {
		t.Fatalf("templ @ should use neutral token: %q", html)
	}
}

func TestHighlightSource_templUsesGoLexer(t *testing.T) {
	html, err := HighlightSource("templ", `import "github.com/fastygo/templ/ui"

templ Example() {}
`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(html, "chroma") {
		t.Fatalf("expected chroma markup: %q", html)
	}
}

func TestHighlightCodeBlocks_fillsPreviewAndCode(t *testing.T) {
	pages := []DocPage{{
		SourceFile: "test.md",
		Blocks: []Block{
			CodeBlock{Language: "go", Source: "package main\n"},
			PreviewCodeBlock{Language: "templ", Source: "templ Example() {}\n"},
		},
	}}
	if err := HighlightCodeBlocks(pages); err != nil {
		t.Fatal(err)
	}
	cb := pages[0].Blocks[0].(CodeBlock)
	if cb.HighlightedHTML == "" {
		t.Fatal("code block missing highlight html")
	}
	pb := pages[0].Blocks[1].(PreviewCodeBlock)
	if pb.HighlightedHTML == "" {
		t.Fatal("preview block missing highlight html")
	}
}

func TestWriteChromaCSS_usesModusSemanticTokens(t *testing.T) {
	var buf bytes.Buffer
	if err := WriteChromaCSS(&buf); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	for _, want := range []string{
		"var(--code-keyword)",
		"var(--code-string)",
		"var(--code-comment)",
		"var(--code-function)",
		"var(--code-type)",
		"var(--code-at)",
		".docs-code .chroma",
		"Modus",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("css missing %q", want)
		}
	}
	if strings.Contains(out, "#5317ac") || strings.Contains(out, "#b6a0ff") {
		t.Fatal("css should reference --code-* tokens, not hardcoded modus hex")
	}
}

func TestExportChromaCSS(t *testing.T) {
	if os.Getenv("EXPORT_CHROMA_CSS") == "" {
		t.Skip("set EXPORT_CHROMA_CSS=1 to regenerate web/static/css/code.css")
	}
	root, err := findRepoRoot(t)
	if err != nil {
		t.Fatal(err)
	}
	var buf bytes.Buffer
	if err := WriteChromaCSS(&buf); err != nil {
		t.Fatal(err)
	}
	out := filepath.Join(root, "web", "static", "css", "code.css")
	if err := os.WriteFile(out, buf.Bytes(), 0o644); err != nil {
		t.Fatal(err)
	}
}
