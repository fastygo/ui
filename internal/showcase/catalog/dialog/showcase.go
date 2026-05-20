package dialog

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
	return showcaseutil.RenderWithChildren(ctx, w, ui.Box(ui.BoxProps{Class: "max-w-sm rounded-lg border border-border bg-card p-4 shadow-lg", Attrs: templ.Attributes{"data-ui8kit": "dialog"}}), func(ctx context.Context, w io.Writer) error {
		if err := ui.Title(ui.TitleProps{Order: 3, Class: "text-base font-semibold"}, "Dialog title").Render(ctx, w); err != nil {
			return err
		}
		if err := ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "Dialog description copy.").Render(ctx, w); err != nil {
			return err
		}
		return ui.Group(ui.GroupProps{Class: "mt-4 flex justify-end gap-2"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
			if err := showcaseutil.Button(ui.ButtonProps{Variant: "outline", Attrs: templ.Attributes{"data-ui8kit-close": ""}}, "Cancel").Render(ctx, w); err != nil {
				return err
			}
			return showcaseutil.Button(ui.ButtonProps{Attrs: templ.Attributes{"data-ui8kit-close": ""}}, "Confirm").Render(ctx, w)
		})), w)
	})
}
