package carousel

import (
	"context"
	"fmt"
	"io"

	"github.com/a-h/templ"
	"github.com/fastygo/templ/ui"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	return showcaseutil.RenderWithChildren(ctx, w, ui.Group(ui.GroupProps{Class: "flex gap-2 overflow-x-auto max-w-md"}), func(ctx context.Context, w io.Writer) error {
		for i := 1; i <= 3; i++ {
			label := fmt.Sprintf("Slide %d", i)
			if err := ui.Box(ui.BoxProps{Class: "min-w-[8rem] shrink-0 rounded-lg border border-border bg-card p-4 text-sm"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
				return ui.Text(ui.TextProps{}, label).Render(ctx, w)
			})), w); err != nil {
				return err
			}
		}
		return nil
	})
}
