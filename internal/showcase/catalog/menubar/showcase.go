package menubar

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
	return showcaseutil.RenderWithChildren(ctx, w, ui.Group(ui.GroupProps{Class: "flex gap-1 rounded-md border border-border bg-card p-1 text-sm"}), func(ctx context.Context, w io.Writer) error {
		for _, label := range []string{"File", "Edit", "View"} {
			if err := showcaseutil.Button(ui.ButtonProps{Variant: "ghost", Size: "sm"}, label).Render(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
}
