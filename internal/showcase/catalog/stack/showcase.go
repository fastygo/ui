package stack

import (
	"context"
	"io"

	"github.com/fastygo/templ/ui"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	return stackPreview(ctx, w, ui.StackProps{Class: "gap-2"})
}
func previewHorizontal(ctx context.Context, w io.Writer) error {
	return stackPreview(ctx, w, ui.StackProps{Class: "flex-row items-center gap-2"})
}
func previewNav(ctx context.Context, w io.Writer) error {
	return stackPreview(ctx, w, ui.StackProps{Tag: "nav", Class: "gap-1"})
}
func stackPreview(ctx context.Context, w io.Writer, props ui.StackProps) error {
	return showcaseutil.RenderWithChildren(ctx, w, ui.Stack(props), func(ctx context.Context, w io.Writer) error {
		if err := ui.Title(ui.TitleProps{Order: 3}, "Stack").Render(ctx, w); err != nil {
			return err
		}
		return ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "Children stack vertically.").Render(ctx, w)
	})
}
