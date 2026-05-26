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
		`overflow-x-auto p-4 text-xs`,
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

func TestIndex_rendersSectionGrids(t *testing.T) {
	page := docsstatic.Index("Components", "Wireframe showcase.", []docsstatic.IndexSection{
		{
			Section: "getting-started",
			Label:   "Getting Started",
			Links: []docsstatic.IndexLink{
				{Title: "Introduction", Description: "Overview.", Href: "/docs/introduction/"},
				{Title: "Installation", Description: "Setup steps.", Href: "/docs/installation/"},
			},
		},
		{
			Section: "primitives",
			Label:   "Primitives",
			Links: []docsstatic.IndexLink{
				{Title: "Badge", Description: "Status label.", Href: "/docs/primitives/badge/"},
				{Title: "Box", Description: "Internal wrapper.", Href: "/docs/primitives/box/"},
				{Title: "Button", Description: "Primary actions.", Href: "/docs/primitives/button/"},
				{Title: "Radio", Description: "Single radio option.", Href: "/docs/primitives/radio/"},
			},
		},
		{
			Section: "components",
			Label:   "Components",
			Links: []docsstatic.IndexLink{
				{Title: "Accordion", Description: "Expandable sections.", Href: "/docs/components/accordion/"},
				{Title: "Alert", Description: "Status message.", Href: "/docs/components/alert/"},
				{Title: "Alert Dialog", Description: "Blocking modal.", Href: "/docs/components/alert-dialog/"},
				{Title: "Card", Description: "Grouped content.", Href: "/docs/components/card/"},
				{Title: "Dialog", Description: "Modal surface.", Href: "/docs/components/dialog/"},
				{Title: "Form", Description: "Field layout.", Href: "/docs/components/form/"},
				{Title: "Pagination", Description: "Page controls.", Href: "/docs/components/pagination/"},
				{Title: "Table", Description: "Data grid.", Href: "/docs/components/table/"},
				{Title: "Tabs", Description: "Tabbed views.", Href: "/docs/components/tabs/"},
			},
		},
	})

	var buf bytes.Buffer
	if err := page.Render(context.Background(), &buf); err != nil {
		t.Fatal(err)
	}
	html := buf.String()
	for _, want := range []string{
		"docs-index-grid-2",
		"docs-index-grid-3",
		"docs-index-card-desc",
		"docs-index-card p-3",
		"docs-index-card-illustrated",
		"docs-index-illus-button",
	} {
		if !strings.Contains(html, want) {
			t.Fatalf("html missing %q", want)
		}
	}
	if strings.Contains(html, "data:image/svg+xml;base64,") {
		t.Fatal("index cards should use CSS illustrations, not inline SVG data URLs")
	}
	if strings.Count(html, "docs-index-grid-2") != 1 {
		t.Fatalf("expected one 2-column grid, got %d", strings.Count(html, "docs-index-grid-2"))
	}
	if strings.Count(html, "docs-index-grid-3") != 2 {
		t.Fatalf("expected two 3-column grids, got %d", strings.Count(html, "docs-index-grid-3"))
	}
	if got := strings.Count(html, "docs-index-card-illustrated"); got != 13 {
		t.Fatalf("expected thirteen illustrated index cards, got %d", got)
	}
}
