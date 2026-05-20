package combobox

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
	return wfStack(ctx, w, templ.Attributes{"data-ui8kit": "combobox"}, func(ctx context.Context, w io.Writer) error {
		if err := ui.Input(ui.InputProps{Placeholder: "Search framework…", Attrs: templ.Attributes{"role": "combobox", "aria-expanded": "true"}}).Render(ctx, w); err != nil {
			return err
		}
		return showcaseutil.RenderWithChildren(ctx, w, ui.List(ui.ListProps{Class: "rounded-md border border-border bg-card p-1 text-sm", Attrs: templ.Attributes{"role": "listbox"}}), func(ctx context.Context, w io.Writer) error {
			for _, opt := range []string{"Go", "TypeScript", "Rust"} {
				if err := ui.ListItem(ui.ListItemProps{}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
					return showcaseutil.Button(ui.ButtonProps{Variant: "ghost", Class: "h-8 w-full justify-start"}, opt).Render(ctx, w)
				})), w); err != nil {
					return err
				}
			}
			return nil
		})
	})
}
