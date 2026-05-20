package separator

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
	return wfStack(ctx, w, nil, func(ctx context.Context, w io.Writer) error {
		if err := ui.Text(ui.TextProps{}, "Above").Render(ctx, w); err != nil {
			return err
		}
		if err := ui.Box(ui.BoxProps{Class: "h-px w-full bg-border", Attrs: templ.Attributes{"role": "separator"}}).Render(ctx, w); err != nil {
			return err
		}
		return ui.Text(ui.TextProps{}, "Below").Render(ctx, w)
	})
}
