package tabs

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
	return wfStack(ctx, w, templ.Attributes{"data-ui8kit": "tabs"}, func(ctx context.Context, w io.Writer) error {
		if err := ui.Group(ui.GroupProps{Class: "flex gap-1", Attrs: templ.Attributes{"role": "tablist"}}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
			if err := showcaseutil.Button(ui.ButtonProps{Variant: "secondary", Size: "sm", Attrs: templ.Attributes{"data-ui8kit-tab": "a", "aria-selected": "true"}}, "Tab A").Render(ctx, w); err != nil {
				return err
			}
			return showcaseutil.Button(ui.ButtonProps{Variant: "ghost", Size: "sm", Attrs: templ.Attributes{"data-ui8kit-tab": "b"}}, "Tab B").Render(ctx, w)
		})), w); err != nil {
			return err
		}
		return ui.Box(ui.BoxProps{Class: "rounded border border-border p-3 text-sm", Attrs: templ.Attributes{"data-ui8kit-panel": "a"}}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
			return ui.Text(ui.TextProps{}, "Tab A panel.").Render(ctx, w)
		})), w)
	})
}
