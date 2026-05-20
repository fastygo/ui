package navigationmenu

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
	return showcaseutil.RenderWithChildren(ctx, w, ui.List(ui.ListProps{Tag: "menu", Class: "flex gap-4 text-sm"}), func(ctx context.Context, w io.Writer) error {
		for _, label := range []string{"Home", "Docs", "Blog"} {
			if err := ui.ListItem(ui.ListItemProps{}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
				return showcaseutil.Button(ui.ButtonProps{Variant: "link"}, label).Render(ctx, w)
			})), w); err != nil {
				return err
			}
		}
		return nil
	})
}
