package block

import (
	"context"
	"io"

	"github.com/fastygo/templ/ui"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func previewMain(ctx context.Context, w io.Writer) error {
	return showcaseutil.RenderWithChildren(ctx, w, ui.Block(ui.BlockProps{Tag: "main", Class: "rounded-lg border border-border p-4"}), func(ctx context.Context, w io.Writer) error {
		return ui.Text(ui.TextProps{}, "Main landmark block.").Render(ctx, w)
	})
}
func previewAside(ctx context.Context, w io.Writer) error {
	return showcaseutil.RenderWithChildren(ctx, w, ui.Block(ui.BlockProps{Tag: "aside", Class: "rounded-lg border border-border p-4 w-48"}), func(ctx context.Context, w io.Writer) error {
		return ui.Text(ui.TextProps{}, "Aside block.").Render(ctx, w)
	})
}
