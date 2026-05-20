package contextmenu

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
		if err := ui.Box(ui.BoxProps{Class: "rounded-lg border border-dashed border-border p-8 text-center text-sm text-muted-foreground"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
			return ui.Text(ui.TextProps{}, "Right-click area (stub)").Render(ctx, w)
		})), w); err != nil {
			return err
		}
		return showcaseutil.RenderWithChildren(ctx, w, ui.List(ui.ListProps{Tag: "menu", Class: "w-40 rounded-md border border-border bg-card p-1 text-sm"}), func(ctx context.Context, w io.Writer) error {
			for _, label := range []string{"Copy", "Paste"} {
				if err := ui.ListItem(ui.ListItemProps{}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
					return ui.Text(ui.TextProps{}, label).Render(ctx, w)
				})), w); err != nil {
					return err
				}
			}
			return nil
		})
	})
}
