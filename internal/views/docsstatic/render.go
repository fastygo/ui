package docsstatic

import (
	"context"
	"html"
	"io"

	"github.com/a-h/templ"
	cmp "github.com/fastygo/templ/components"
	"github.com/fastygo/templ/ui"
)

// Page renders a static documentation page.
func Page(data PageData) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderDocsLayout(ctx, w, data.TOC, data.TOCLabel, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return renderPageBody(ctx, w, data)
		}))
	})
}

func renderPageBody(ctx context.Context, w io.Writer, data PageData) error {
	if err := ui.Title(ui.TitleProps{Order: 1}, data.Title).Render(ctx, w); err != nil {
		return err
	}
	if err := ui.Text(ui.TextProps{Class: "docs-lead"}, data.Description).Render(ctx, w); err != nil {
		return err
	}
	if data.Source != "" {
		if err := ui.Text(ui.TextProps{Class: "docs-source"}, data.Source).Render(ctx, w); err != nil {
			return err
		}
	}
	for _, block := range data.Blocks {
		if err := renderBlock(ctx, w, block); err != nil {
			return err
		}
	}
	if err := renderAPI(ctx, w, data.API, data.APISectionTitle); err != nil {
		return err
	}
	return renderRelated(ctx, w, data.Related, data.RelatedSectionTitle)
}

func renderDocsLayout(ctx context.Context, w io.Writer, toc []TOCHeading, tocLabel string, article templ.Component) error {
	return ui.Box(ui.BoxProps{Class: "docs-layout"}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := ui.Stack(ui.StackProps{
				Tag:   "article",
				Class: "prose prose-docs docs-article",
			}).Render(templ.WithChildren(ctx, article), w); err != nil {
				return err
			}
			if len(toc) == 0 {
				return nil
			}
			return renderTOCAside(ctx, w, toc, tocLabel)
		})), w)
}

func renderTOCAside(ctx context.Context, w io.Writer, toc []TOCHeading, label string) error {
	if label == "" {
		label = "On this page"
	}
	return ui.Block(ui.BlockProps{
		Tag:   "aside",
		Class: "docs-toc",
		Attrs: templ.Attributes{"aria-label": label},
	}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return ui.Box(ui.BoxProps{Class: "sticky top-20 max-h-[calc(100vh-5rem)] overflow-y-auto pb-4 pl-2"}).Render(
			templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				if err := ui.Text(ui.TextProps{Class: "mb-3 text-xs font-semibold uppercase tracking-wide text-muted-foreground"}, label).Render(ctx, w); err != nil {
					return err
				}
				return ui.Stack(ui.StackProps{Tag: "nav", Class: "gap-1 border-l border-border"}).Render(
					templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
						for _, h := range toc {
							if err := renderTOCLink(ctx, w, h); err != nil {
								return err
							}
						}
						return nil
					})), w)
			})), w)
	})), w)
}

func renderTOCLink(ctx context.Context, w io.Writer, h TOCHeading) error {
	class := "block h-auto justify-start rounded-none border-0 py-1 pl-3 pr-2 text-left text-xs font-normal text-muted-foreground hover:text-foreground"
	if h.Level >= 3 {
		class += " pl-6 text-xs"
	}
	return renderButtonLabel(ctx, w, ui.ButtonProps{
		Href:    "#" + h.ID,
		Variant: "link",
		Class:   class,
	}, h.Text)
}

func renderBlock(ctx context.Context, w io.Writer, block Block) error {
	switch b := block.(type) {
	case HeadingBlock:
		order := b.Level
		if order < 1 {
			order = 2
		}
		if order > 6 {
			order = 6
		}
		title := ui.Title(ui.TitleProps{Order: order}, b.Text)
		if b.ID == "" {
			return title.Render(ctx, w)
		}
		return ui.Box(ui.BoxProps{
			Tag:   "div",
			Class: "scroll-mt-24",
			Attrs: templ.Attributes{"id": b.ID},
		}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return title.Render(ctx, w)
		})), w)
	case ParagraphBlock:
		return ui.Text(ui.TextProps{}, b.Text).Render(ctx, w)
	case ListBlock:
		return renderList(ctx, w, b.Items)
	case PreviewCodeBlock:
		return renderPreviewCode(ctx, w, b)
	case CodeBlock:
		return renderCodeBlock(ctx, w, b)
	default:
		return nil
	}
}

func renderList(ctx context.Context, w io.Writer, items []string) error {
	return ui.List(ui.ListProps{}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			for _, item := range items {
				if err := ui.Text(ui.TextProps{Tag: "li"}, item).Render(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})), w)
}

func renderPreviewCode(ctx context.Context, w io.Writer, b PreviewCodeBlock) error {
	preview := b.HTML
	code := buildCodeBlockInner(b.Source, b.HighlightedHTML, codeBlockEmbedded)
	markup := `<div class="docs-preview my-3 w-full overflow-hidden rounded-md border border-border bg-card text-card-foreground shadow-sm">` +
		`<div class="flex w-full items-center justify-center p-6 min-h-32">` + preview + `</div>` +
		`<details class="border-t border-border bg-muted/30">` +
		`<summary class="docs-preview-summary list-none cursor-pointer [&::-webkit-details-marker]:hidden">` +
		`<div class="docs-preview-code-panel">` + code +
		`<div class="docs-preview-fade" aria-hidden="true"></div>` +
		`<div class="docs-preview-view" aria-hidden="true">` +
		`<span class="docs-preview-toggle">View code</span>` +
		`</div>` +
		`</div>` +
		`</summary>` +
		`<button type="button" class="docs-preview-hide" data-docs-preview-hide aria-label="Hide code">` +
		`<span class="docs-preview-toggle">Hide code</span>` +
		`</button>` +
		`</details></div>`
	return writeRawHTML(ctx, w, markup)
}

func renderCodeBlock(ctx context.Context, w io.Writer, b CodeBlock) error {
	return writeRawHTML(ctx, w, buildCodeBlockInner(b.Source, b.HighlightedHTML, codeBlockStandalone))
}

type codeBlockLayout int

const (
	codeBlockStandalone codeBlockLayout = iota
	codeBlockEmbedded
)

func buildCodeBlockInner(source, highlighted string, layout codeBlockLayout) string {
	body := highlighted
	if body == "" {
		body = `<pre class="overflow-x-auto text-xs leading-relaxed"><code class="font-mono whitespace-pre text-foreground">` +
			html.EscapeString(source) + `</code></pre>`
	}
	wrapperClass := "docs-code relative font-mono"
	if layout == codeBlockStandalone {
		wrapperClass += " my-3 rounded-md border border-border bg-muted/30"
	}
	return `<div class="` + wrapperClass + `" data-ui8kit="copy-button">` +
		`<button type="button" class="docs-copy-trigger absolute right-2 top-2 z-20 inline-flex items-center justify-center rounded-md p-1 text-muted-foreground hover:text-foreground" data-copy-trigger aria-label="Copy code" data-copy-label="Copy code" data-copied-label="Copied"><span class="inline-block shrink-0 latty latty-copy h-4 w-4" aria-hidden="true"></span></button>` +
		`<textarea hidden readonly aria-hidden="true" tabindex="-1" data-copy-source>` + html.EscapeString(source) + `</textarea>` +
		`<div class="overflow-x-auto p-4 text-xs leading-relaxed">` + body + `</div>` +
		`</div>`
}

func writeRawHTML(ctx context.Context, w io.Writer, markup string) error {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, markup)
		return err
	}).Render(ctx, w)
}

func renderAPI(ctx context.Context, w io.Writer, fields []APIField, title string) error {
	if len(fields) == 0 {
		return nil
	}
	if title == "" {
		title = "API"
	}
	return ui.Box(ui.BoxProps{
		Tag:   "div",
		Class: "scroll-mt-24",
		Attrs: templ.Attributes{"id": "api"},
	}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		if err := ui.Title(ui.TitleProps{Order: 2}, title).Render(ctx, w); err != nil {
			return err
		}
		for _, f := range fields {
			line := f.Name + " · " + f.Type
			if f.Description != "" {
				line += " — " + f.Description
			}
			if err := ui.Text(ui.TextProps{Class: "docs-api-item"}, line).Render(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})), w)
}

func renderRelated(ctx context.Context, w io.Writer, links []RelatedLink, title string) error {
	if len(links) == 0 {
		return nil
	}
	if title == "" {
		title = "Related"
	}
	return ui.Box(ui.BoxProps{
		Tag:   "div",
		Class: "scroll-mt-24",
		Attrs: templ.Attributes{"id": "related"},
	}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		if err := ui.Title(ui.TitleProps{Order: 2}, title).Render(ctx, w); err != nil {
			return err
		}
		return ui.Box(ui.BoxProps{Class: "docs-related-links"}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			for _, link := range links {
				if err := renderButtonLabel(ctx, w, ui.ButtonProps{
					Href:    link.Href,
					Variant: "link",
					Class:   "h-auto justify-start p-0 text-xs font-normal",
				}, link.Label); err != nil {
					return err
				}
			}
			return nil
		})), w)
	})), w)
}

// Index renders the docs landing page.
func Index(title, description string, sections []IndexSection) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderDocsLayout(ctx, w, nil, "", templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := ui.Title(ui.TitleProps{Order: 1}, title).Render(ctx, w); err != nil {
				return err
			}
			if err := ui.Text(ui.TextProps{Class: "docs-lead"}, description).Render(ctx, w); err != nil {
				return err
			}
			for _, sec := range sections {
				if err := renderIndexSection(ctx, w, sec); err != nil {
					return err
				}
			}
			return nil
		}))
	})
}

// IndexSection groups index links.
type IndexSection struct {
	Section string
	Label   string
	Links   []IndexLink
}

// IndexLink is one docs index card.
type IndexLink struct {
	Title       string
	Description string
	Href        string
}

func renderIndexSection(ctx context.Context, w io.Writer, sec IndexSection) error {
	if err := ui.Title(ui.TitleProps{Order: 2}, sec.Label).Render(ctx, w); err != nil {
		return err
	}
	return ui.Grid(ui.GridProps{Class: indexSectionGridClass(sec.Section)}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			for _, link := range sec.Links {
				if err := cmp.Card(cmp.CardProps{Class: "docs-index-card p-3"}).Render(
					templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
						if err := renderButtonLabel(ctx, w, ui.ButtonProps{
							Href:    link.Href,
							Variant: "link",
							Class:   "h-auto justify-start p-0 text-base font-semibold",
						}, link.Title); err != nil {
							return err
						}
						return ui.Text(ui.TextProps{Class: "docs-index-card-desc text-muted-foreground"}, link.Description).Render(ctx, w)
					})), w); err != nil {
					return err
				}
			}
			return nil
		})), w)
}

func indexSectionGridClass(sectionID string) string {
	switch sectionID {
	case "primitives", "components":
		return "docs-index-grid docs-index-grid-3"
	default:
		return "docs-index-grid docs-index-grid-2"
	}
}

// SectionLabel returns a human label for a docs section id.
func SectionLabel(id string) string {
	switch id {
	case "getting-started":
		return "Getting Started"
	case "components":
		return "Components"
	case "blocks":
		return "Blocks"
	default:
		return id
	}
}

// Truncate shortens descriptions for index cards.
func Truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max] + "…"
}

// FormatPageTitle builds the document title for a docs page.
func FormatPageTitle(pageTitle, brand string) string {
	if brand == "" {
		brand = "FastyGo UI"
	}
	return pageTitle + " · " + brand
}

// StaticAssetPaths are fixed app static URLs.
func StaticAssetPaths() (css, themeJS, appJS string) {
	return "/static/css/app.css", "/static/js/theme.js", "/static/js/ui8kit.js"
}

func renderButtonLabel(ctx context.Context, w io.Writer, props ui.ButtonProps, label string) error {
	return ui.Button(props).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, label)
		return err
	})), w)
}
