package box

import (
	"context"
	"io"

	"github.com/fastygo/templ/ui"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	return showcaseutil.RenderWithChildren(ctx, w, ui.Box(ui.BoxProps{Class: "rounded-lg border border-border p-4"}), func(ctx context.Context, w io.Writer) error {
		return ui.Text(ui.TextProps{}, "Box content.").Render(ctx, w)
	})
}
func previewPre(ctx context.Context, w io.Writer) error {
	return showcaseutil.RenderWithChildren(ctx, w, ui.Box(ui.BoxProps{Tag: "pre", Class: "rounded-md border border-border bg-muted/30 p-3 text-xs font-mono"}), func(ctx context.Context, w io.Writer) error {
		return ui.Text(ui.TextProps{Tag: "code"}, "code snippet").Render(ctx, w)
	})
}
