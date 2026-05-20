package docsarticle

import (
	"context"
	"io"

	"github.com/fastygo/templ/ui"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	return blockBody(ctx, w, "Wireframe article body for documentation blocks.", false)
}
func previewCompact(ctx context.Context, w io.Writer) error {
	return blockBody(ctx, w, "Wireframe article body for documentation blocks.", true)
}
func blockBody(ctx context.Context, w io.Writer, body string, compact bool) error {
	gap := "gap-4"
	if compact {
		gap = "gap-2"
	}
	return showcaseutil.RenderWithChildren(ctx, w, ui.Stack(ui.StackProps{Class: gap + " max-w-2xl"}), func(ctx context.Context, w io.Writer) error {
		if err := ui.Title(ui.TitleProps{Order: 2}, "Docs Article").Render(ctx, w); err != nil {
			return err
		}
		if err := ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground leading-relaxed"}, body).Render(ctx, w); err != nil {
			return err
		}
		return showcaseutil.Button(ui.ButtonProps{Variant: "outline", Size: "sm"}, "Action").Render(ctx, w)
	})
}
