package calendar

import (
	"context"
	"fmt"
	"io"

	"github.com/a-h/templ"
	"github.com/fastygo/templ/ui"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func wfStack(ctx context.Context, w io.Writer, attrs templ.Attributes, body func(context.Context, io.Writer) error) error {
	props := ui.StackProps{Class: "gap-2 max-w-md"}
	if attrs != nil {
		props.Attrs = attrs
	}
	return showcaseutil.RenderWithChildren(ctx, w, ui.Stack(props), body)
}

func previewDefault(ctx context.Context, w io.Writer) error {
	return showcaseutil.RenderWithChildren(ctx, w, ui.Box(ui.BoxProps{Class: "w-64 rounded-lg border border-border p-3"}), func(ctx context.Context, w io.Writer) error {
		if err := ui.Text(ui.TextProps{Class: "text-sm font-medium"}, "May 2026 — placeholder").Render(ctx, w); err != nil {
			return err
		}
		return ui.Grid(ui.GridProps{Class: "mt-2 grid-cols-7 gap-1 text-center text-xs text-muted-foreground"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
			for d := 1; d <= 7; d++ {
				if err := ui.Text(ui.TextProps{}, fmt.Sprint(d)).Render(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})), w)
	})
}
