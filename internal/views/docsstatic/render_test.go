package docsstatic_test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/fastygo/ui/internal/views/docsstatic"
)

func TestPage_rendersPreviewWidget(t *testing.T) {
	page := docsstatic.Page(docsstatic.PageData{
		Title:       "Badge",
		Description: "Status labels.",
		Blocks: []docsstatic.Block{
			docsstatic.PreviewCodeBlock{
				Source:          "templ Example() {}\n",
				HTML:            `<span class="text-sm">Badge</span>`,
				HighlightedHTML: `<pre class="chroma"><code><span class="line"><span class="cl">templ Example() {}</span></span></code></pre>`,
			},
		},
	})

	var buf bytes.Buffer
	if err := page.Render(context.Background(), &buf); err != nil {
		t.Fatal(err)
	}
	html := buf.String()
	for _, want := range []string{
		`docs-preview`,
		`docs-preview-code-panel`,
		`docs-preview-fade`,
		`w-full`,
		`border-t border-border`,
		`View code`,
		`z-20`,
		`overflow-x-auto px-4 pt-4 pb-5 text-xs`,
		`data-docs-preview-hide`,
		`docs-preview-toggle`,
	} {
		if !strings.Contains(html, want) {
			t.Fatalf("html missing %q", want)
		}
	}
	if strings.Contains(html, `max-h-20`) {
		t.Fatal("preview collapse should use docs-preview-code-panel CSS, not max-h-20 utility")
	}
	if strings.Contains(html, `pt-10`) {
		t.Fatal("preview code should not reserve top padding for copy button")
	}
}

func TestPage_rendersFloatingTOC(t *testing.T) {
	page := docsstatic.Page(docsstatic.PageData{
		Title:       "Button",
		Description: "Primary actions.",
		TOC: []docsstatic.TOCHeading{
			{Level: 2, Text: "Default", ID: "default"},
			{Level: 2, Text: "Variants", ID: "variants"},
		},
		TOCLabel: "On this page",
		Blocks: []docsstatic.Block{
			docsstatic.HeadingBlock{Level: 2, Text: "Default", ID: "default"},
		},
	})

	var buf bytes.Buffer
	if err := page.Render(context.Background(), &buf); err != nil {
		t.Fatal(err)
	}
	html := buf.String()
	for _, want := range []string{
		`<aside`,
		`docs-toc`,
		`aria-label="On this page"`,
		`href="#default"`,
		`href="#variants"`,
		`id="default"`,
		`<article`,
		"prose-docs",
		"docs-article",
		"docs-layout",
	} {
		if !strings.Contains(html, want) {
			t.Fatalf("html missing %q\n%s", want, html)
		}
	}
}
