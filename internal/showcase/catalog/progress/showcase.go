package progress

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
	return ui.Box(ui.BoxProps{Class: "h-2 w-full max-w-xs overflow-hidden rounded-full bg-muted", Attrs: templ.Attributes{"role": "progressbar", "aria-valuenow": "60", "aria-valuemin": "0", "aria-valuemax": "100"}}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
		return ui.Box(ui.BoxProps{Class: "h-full w-3/5 bg-primary"}).Render(ctx, w)
	})), w)
}
