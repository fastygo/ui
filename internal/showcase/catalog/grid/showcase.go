package grid

import (
	"context"
	"fmt"
	"io"

	"github.com/a-h/templ"
	"github.com/fastygo/templ/ui"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	return gridDemo(ctx, w, "grid-cols-2 gap-4", 2)
}
func previewThree(ctx context.Context, w io.Writer) error {
	return gridDemo(ctx, w, "grid-cols-3 gap-2", 3)
}
func gridDemo(ctx context.Context, w io.Writer, class string, n int) error {
	return showcaseutil.RenderWithChildren(ctx, w, ui.Grid(ui.GridProps{Class: class + " max-w-md"}), func(ctx context.Context, w io.Writer) error {
		for i := 1; i <= n; i++ {
			label := fmt.Sprint(i)
			if err := ui.GridCol(ui.GridColProps{}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
				return ui.Box(ui.BoxProps{Class: "rounded border border-border p-3 text-center text-sm"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
					return ui.Text(ui.TextProps{}, label).Render(ctx, w)
				})), w)
			})), w); err != nil {
				return err
			}
		}
		return nil
	})
}
