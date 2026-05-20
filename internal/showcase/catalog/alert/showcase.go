package alert

import (
	"context"
	"io"

	cmp "github.com/fastygo/templ/components"
	"github.com/fastygo/templ/ui"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func alertBody(ctx context.Context, w io.Writer, p cmp.AlertProps, title, body string) error {
	return showcaseutil.RenderWithChildren(ctx, w, cmp.Alert(p), func(ctx context.Context, w io.Writer) error {
		if err := ui.Title(ui.TitleProps{Order: 4, Class: "text-sm font-semibold"}, title).Render(ctx, w); err != nil {
			return err
		}
		return ui.Text(ui.TextProps{Class: "text-sm"}, body).Render(ctx, w)
	})
}
func previewDefault(ctx context.Context, w io.Writer) error {
	return alertBody(ctx, w, cmp.AlertProps{}, "Heads up", "You can add components from the gallery.")
}
func previewDestructive(ctx context.Context, w io.Writer) error {
	return alertBody(ctx, w, cmp.AlertProps{Variant: "destructive"}, "Error", "Something went wrong.")
}
