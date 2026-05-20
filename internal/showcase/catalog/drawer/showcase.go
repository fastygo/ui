package drawer

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
	return ui.Box(ui.BoxProps{Class: "w-full max-w-md rounded-t-xl border border-border bg-card p-4"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
		if err := ui.Box(ui.BoxProps{Class: "mx-auto mb-3 h-1 w-10 rounded-full bg-muted"}).Render(ctx, w); err != nil {
			return err
		}
		return ui.Text(ui.TextProps{}, "Drawer content.").Render(ctx, w)
	})), w)
}
