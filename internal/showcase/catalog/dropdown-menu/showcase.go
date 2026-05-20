package dropdownmenu

import (
	"context"
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
	return showcaseutil.RenderWithChildren(ctx, w, ui.Stack(ui.StackProps{Class: "gap-2"}), func(ctx context.Context, w io.Writer) error {
		if err := showcaseutil.Button(ui.ButtonProps{Variant: "outline"}, "Open menu").Render(ctx, w); err != nil {
			return err
		}
		return showcaseutil.RenderWithChildren(ctx, w, ui.List(ui.ListProps{Tag: "menu", Class: "w-40 rounded-md border border-border bg-card p-1 text-sm"}), func(ctx context.Context, w io.Writer) error {
			for _, label := range []string{"Profile", "Settings", "Sign out"} {
				if err := ui.ListItem(ui.ListItemProps{Class: "rounded px-2 py-1 hover:bg-accent"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
					return showcaseutil.Button(ui.ButtonProps{Variant: "ghost", Class: "h-8 w-full justify-start px-2"}, label).Render(ctx, w)
				})), w); err != nil {
					return err
				}
			}
			return nil
		})
	})
}
