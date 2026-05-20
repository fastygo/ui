package toast

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
	return ui.Box(ui.BoxProps{Class: "flex max-w-sm items-center justify-between gap-4 rounded-lg border border-border bg-card p-4 shadow"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
		if err := ui.Text(ui.TextProps{Class: "text-sm"}, "Saved successfully.").Render(ctx, w); err != nil {
			return err
		}
		return showcaseutil.Button(ui.ButtonProps{Variant: "outline", Size: "sm"}, "Undo").Render(ctx, w)
	})), w)
}
