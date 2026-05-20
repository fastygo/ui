package docs

import (
	"context"
	"io"

	"github.com/fastygo/ui/internal/showcase/showcaseutil"

	"github.com/a-h/templ"
	cmp "github.com/fastygo/templ/components"
	"github.com/fastygo/templ/ui"
	"github.com/fastygo/ui/internal/registry"
)

// Index renders the /docs landing page.
func Index(data IndexData) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		if err := ui.Title(ui.TitleProps{Order: 1}, data.Title).Render(ctx, w); err != nil {
			return err
		}
		if err := ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground leading-relaxed"}, data.Description).Render(ctx, w); err != nil {
			return err
		}
		for _, sec := range data.Sections {
			if err := renderIndexSection(ctx, w, sec); err != nil {
				return err
			}
		}
		return nil
	})
}

// ComponentDoc renders a single registry component page.
func ComponentDoc(page registry.Page) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		if err := ui.Title(ui.TitleProps{Order: 1}, page.Title).Render(ctx, w); err != nil {
			return err
		}
		if err := ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground leading-relaxed max-w-3xl"}, page.Description).Render(ctx, w); err != nil {
			return err
		}
		if page.Source != "" {
			if err := ui.Text(ui.TextProps{Class: "text-xs font-mono text-muted-foreground"}, page.Source).Render(ctx, w); err != nil {
				return err
			}
		}
		for _, v := range page.Variants {
			if err := renderVariant(ctx, w, v); err != nil {
				return err
			}
		}
		if err := renderAPI(ctx, w, page.API); err != nil {
			return err
		}
		return renderRelated(ctx, w, page.Related)
	})
}

func renderIndexSection(ctx context.Context, w io.Writer, sec IndexSection) error {
	if err := ui.Title(ui.TitleProps{Order: 2, Class: "text-lg font-semibold pt-6"}, sec.Label).Render(ctx, w); err != nil {
		return err
	}
	for _, link := range sec.Links {
		if err := cmp.Card(cmp.CardProps{Class: "p-4 mb-3"}).Render(
			templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				if err := showcaseutil.Button(ui.ButtonProps{
					Href:    link.Href,
					Variant: "link",
					Class:   "h-auto justify-start p-0 text-base font-semibold",
				}, link.Title).Render(ctx, w); err != nil {
					return err
				}
				return ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, link.Description).Render(ctx, w)
			})), w); err != nil {
			return err
		}
	}
	return nil
}

func renderVariant(ctx context.Context, w io.Writer, v registry.Variant) error {
	if err := ui.Title(ui.TitleProps{Order: 3, Class: "text-base font-semibold pt-8"}, v.Title).Render(ctx, w); err != nil {
		return err
	}
	if v.Description != "" {
		if err := ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, v.Description).Render(ctx, w); err != nil {
			return err
		}
	}
	if err := cmp.Card(cmp.CardProps{Class: "p-6 my-3"}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return v.Preview.Render(ctx, w)
		})), w); err != nil {
		return err
	}
	return ui.Box(ui.BoxProps{
		Tag:   "pre",
		Class: "overflow-x-auto rounded-md border border-border bg-muted/30 p-4 text-xs leading-relaxed",
	}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return ui.Text(ui.TextProps{Tag: "code", Class: "font-mono whitespace-pre text-foreground"}, v.Code).Render(ctx, w)
		})), w)
}

func renderAPI(ctx context.Context, w io.Writer, fields []registry.APIField) error {
	if len(fields) == 0 {
		return nil
	}
	if err := ui.Title(ui.TitleProps{Order: 2, Class: "text-lg font-semibold pt-10"}, "API").Render(ctx, w); err != nil {
		return err
	}
	for _, f := range fields {
		line := f.Name + " · " + f.Type + " — " + f.Description
		if err := ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground py-1"}, line).Render(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func renderRelated(ctx context.Context, w io.Writer, links []registry.RelatedLink) error {
	if len(links) == 0 {
		return nil
	}
	if err := ui.Title(ui.TitleProps{Order: 2, Class: "text-lg font-semibold pt-8"}, "Related").Render(ctx, w); err != nil {
		return err
	}
	for _, link := range links {
		if err := showcaseutil.Button(ui.ButtonProps{
			Href:    link.Href,
			Variant: "link",
			Class:   "block h-auto justify-start p-0 text-sm",
		}, link.Label).Render(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
