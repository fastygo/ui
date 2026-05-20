package command

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/fastygo/templ/ui"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	return showcaseutil.RenderWithChildren(ctx, w, ui.Stack(ui.StackProps{Class: "gap-2 max-w-sm rounded-lg border border-border bg-card p-2"}), func(ctx context.Context, w io.Writer) error {
		if err := ui.Input(ui.InputProps{Placeholder: "Type a command…"}).Render(ctx, w); err != nil {
			return err
		}
		return showcaseutil.RenderWithChildren(ctx, w, ui.List(ui.ListProps{Class: "text-sm"}), func(ctx context.Context, w io.Writer) error {
			for _, cmd := range []string{"Open docs", "Toggle theme", "Go home"} {
				if err := ui.ListItem(ui.ListItemProps{}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
					return showcaseutil.Button(ui.ButtonProps{Variant: "ghost", Class: "h-8 w-full justify-start"}, cmd).Render(ctx, w)
				})), w); err != nil {
					return err
				}
			}
			return nil
		})
	})
}
