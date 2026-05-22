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
		if err := ui.Title(ui.TitleProps{Order: 1}, data.Title).Render(ctx, w); err != nil {
			return err
		}
		if err := ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground leading-relaxed max-w-3xl"}, data.Description).Render(ctx, w); err != nil {
			return err
		}
		if data.Source != "" {
			if err := ui.Text(ui.TextProps{Class: "text-xs font-mono text-muted-foreground"}, data.Source).Render(ctx, w); err != nil {
				return err
			}
		}
		for _, block := range data.Blocks {
			if err := renderBlock(ctx, w, block); err != nil {
				return err
			}
		}
		if err := renderAPI(ctx, w, data.API); err != nil {
			return err
		}
		return renderRelated(ctx, w, data.Related)
	})
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
		class := "text-base font-semibold pt-8"
		if order <= 2 {
			class = "text-lg font-semibold pt-6"
		}
		return ui.Title(ui.TitleProps{Order: order, Class: class}, b.Text).Render(ctx, w)
	case ParagraphBlock:
		return ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground leading-relaxed max-w-3xl"}, b.Text).Render(ctx, w)
	case ListBlock:
		return renderList(ctx, w, b.Items)
	case PreviewCodeBlock:
		return renderPreviewCode(ctx, w, b)
	case CodeBlock:
		return renderCodeBox(ctx, w, b.Source)
	default:
		return nil
	}
}

func renderList(ctx context.Context, w io.Writer, items []string) error {
	return ui.List(ui.ListProps{Class: "list-disc pl-6 text-sm text-muted-foreground max-w-3xl"}).Render(
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
	if err := cmp.Card(cmp.CardProps{Class: "p-6 my-3"}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if b.HTML == "" {
				return nil
			}
			_, err := io.WriteString(w, b.HTML)
			return err
		})), w); err != nil {
		return err
	}
	return renderPreviewCodePanel(ctx, w, b.Source)
}

func renderPreviewCodePanel(ctx context.Context, w io.Writer, source string) error {
	escaped := html.EscapeString(source)
	markup := `<details class="group relative my-3 rounded-md border border-border bg-muted/30">
<summary class="list-none cursor-pointer [&::-webkit-details-marker]:hidden">
<div class="relative max-h-16 overflow-hidden p-4 group-open:max-h-none group-open:overflow-visible">
<pre class="text-xs leading-relaxed"><code class="font-mono whitespace-pre text-foreground">` + escaped + `</code></pre>
<div class="pointer-events-none absolute inset-x-0 bottom-0 h-12 bg-gradient-to-b from-transparent to-card group-open:hidden"></div>
<div class="pointer-events-none absolute inset-x-0 bottom-2 flex justify-center">
<span class="rounded-md border border-border bg-background px-3 py-1 text-xs font-medium group-open:hidden">Show code</span>
<span class="hidden rounded-md border border-border bg-background px-3 py-1 text-xs font-medium group-open:inline">Hide code</span>
</div>
</div>
</summary>
</details>`
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, markup)
		return err
	}).Render(ctx, w)
}

func renderCodeBox(ctx context.Context, w io.Writer, source string) error {
	return ui.Box(ui.BoxProps{
		Tag:   "pre",
		Class: "overflow-x-auto rounded-md border border-border bg-muted/30 p-4 text-xs leading-relaxed",
	}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return ui.Text(ui.TextProps{Tag: "code", Class: "font-mono whitespace-pre text-foreground"}, source).Render(ctx, w)
		})), w)
}

func renderAPI(ctx context.Context, w io.Writer, fields []APIField) error {
	if len(fields) == 0 {
		return nil
	}
	if err := ui.Title(ui.TitleProps{Order: 2, Class: "text-lg font-semibold pt-10"}, "API").Render(ctx, w); err != nil {
		return err
	}
	for _, f := range fields {
		line := f.Name + " · " + f.Type
		if f.Description != "" {
			line += " — " + f.Description
		}
		if err := ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground py-1"}, line).Render(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func renderRelated(ctx context.Context, w io.Writer, links []RelatedLink) error {
	if len(links) == 0 {
		return nil
	}
	if err := ui.Title(ui.TitleProps{Order: 2, Class: "text-lg font-semibold pt-8"}, "Related").Render(ctx, w); err != nil {
		return err
	}
	for _, link := range links {
		if err := renderButtonLabel(ctx, w, ui.ButtonProps{
			Href:    link.Href,
			Variant: "link",
			Class:   "block h-auto justify-start p-0 text-sm",
		}, link.Label); err != nil {
			return err
		}
	}
	return nil
}

// Index renders the docs landing page.
func Index(title, description string, sections []IndexSection) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		if err := ui.Title(ui.TitleProps{Order: 1}, title).Render(ctx, w); err != nil {
			return err
		}
		if err := ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground leading-relaxed"}, description).Render(ctx, w); err != nil {
			return err
		}
		for _, sec := range sections {
			if err := renderIndexSection(ctx, w, sec); err != nil {
				return err
			}
		}
		return nil
	})
}

// IndexSection groups index links.
type IndexSection struct {
	Label string
	Links []IndexLink
}

// IndexLink is one docs index card.
type IndexLink struct {
	Title       string
	Description string
	Href        string
}

func renderIndexSection(ctx context.Context, w io.Writer, sec IndexSection) error {
	if err := ui.Title(ui.TitleProps{Order: 2, Class: "text-lg font-semibold pt-6"}, sec.Label).Render(ctx, w); err != nil {
		return err
	}
	for _, link := range sec.Links {
		if err := cmp.Card(cmp.CardProps{Class: "p-4 mb-3"}).Render(
			templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				if err := renderButtonLabel(ctx, w, ui.ButtonProps{
					Href:    link.Href,
					Variant: "link",
					Class:   "h-auto justify-start p-0 text-base font-semibold",
				}, link.Title); err != nil {
					return err
				}
				return ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, link.Description).Render(ctx, w)
			})), w); err != nil {
			return err
		}
	}
	return nil
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

// FormatPageTitle builds the shell title for a docs page.
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
