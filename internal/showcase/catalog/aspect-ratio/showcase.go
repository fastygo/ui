package aspectratio

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
	return ui.Box(ui.BoxProps{Class: "aspect-video w-full max-w-xs overflow-hidden rounded-lg border border-border bg-muted"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
		return ui.Box(ui.BoxProps{Class: "flex h-full items-center justify-center text-sm text-muted-foreground"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
			return ui.Text(ui.TextProps{}, "16:9").Render(ctx, w)
		})), w)
	})), w)
}
