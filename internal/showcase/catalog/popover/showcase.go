package popover

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
		if err := showcaseutil.Button(ui.ButtonProps{Variant: "outline"}, "Open popover").Render(ctx, w); err != nil {
			return err
		}
		return ui.Box(ui.BoxProps{Class: "w-56 rounded-lg border border-border bg-card p-3 text-sm shadow"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
			return ui.Text(ui.TextProps{}, "Popover body copy.").Render(ctx, w)
		})), w)
	})
}
