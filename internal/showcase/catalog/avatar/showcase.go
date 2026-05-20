package avatar

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
	return ui.Group(ui.GroupProps{Class: "flex items-center gap-3"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
		if err := ui.Box(ui.BoxProps{Class: "flex h-10 w-10 items-center justify-center rounded-full bg-muted text-sm font-medium"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
			return ui.Text(ui.TextProps{}, "AB").Render(ctx, w)
		})), w); err != nil {
			return err
		}
		return ui.Text(ui.TextProps{Class: "text-sm font-medium"}, "Ada Lovelace").Render(ctx, w)
	})), w)
}
