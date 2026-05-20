package group

import (
	"context"
	"fmt"
	"io"

	"github.com/a-h/templ"
	"github.com/fastygo/templ/ui"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	return showcaseutil.RenderWithChildren(ctx, w, ui.Group(ui.GroupProps{Class: "flex items-center gap-2"}), func(ctx context.Context, w io.Writer) error {
		if err := showcaseutil.Button(ui.ButtonProps{Size: "sm"}, "One").Render(ctx, w); err != nil {
			return err
		}
		return showcaseutil.Button(ui.ButtonProps{Size: "sm", Variant: "outline"}, "Two").Render(ctx, w)
	})
}
func previewWrap(ctx context.Context, w io.Writer) error {
	return showcaseutil.RenderWithChildren(ctx, w, ui.Group(ui.GroupProps{Class: "flex flex-wrap gap-2 max-w-xs"}), func(ctx context.Context, w io.Writer) error {
		for i := 1; i <= 4; i++ {
			if err := ui.Badge(ui.BadgeProps{}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
				return ui.Text(ui.TextProps{}, fmt.Sprint(i)).Render(ctx, w)
			})), w); err != nil {
				return err
			}
		}
		return nil
	})
}
