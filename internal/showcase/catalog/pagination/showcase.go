package pagination

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
	return showcaseutil.RenderWithChildren(ctx, w, ui.Group(ui.GroupProps{Class: "flex items-center gap-1"}), func(ctx context.Context, w io.Writer) error {
		if err := showcaseutil.Button(ui.ButtonProps{Variant: "outline", Size: "sm"}, "Prev").Render(ctx, w); err != nil {
			return err
		}
		if err := showcaseutil.Button(ui.ButtonProps{Variant: "secondary", Size: "sm"}, "1").Render(ctx, w); err != nil {
			return err
		}
		return showcaseutil.Button(ui.ButtonProps{Variant: "outline", Size: "sm"}, "Next").Render(ctx, w)
	})
}
