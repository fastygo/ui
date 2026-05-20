package sheet

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
		if err := ui.Box(ui.BoxProps{Class: "w-64 rounded-l-lg border border-border bg-card p-4"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
			if err := ui.Title(ui.TitleProps{Order: 3, Class: "text-sm font-semibold"}, "Sheet").Render(ctx, w); err != nil {
				return err
			}
			return ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "Side panel content.").Render(ctx, w)
		})), w); err != nil {
			return err
		}
		return nil
	})
}
