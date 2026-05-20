package togglegroup

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/fastygo/templ/ui"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	return showcaseutil.RenderWithChildren(ctx, w, ui.Group(ui.GroupProps{Class: "inline-flex rounded-md border border-border"}), func(ctx context.Context, w io.Writer) error {
		if err := showcaseutil.Button(ui.ButtonProps{Variant: "secondary", Size: "sm", Attrs: templ.Attributes{"aria-pressed": "true"}}, "Left").Render(ctx, w); err != nil {
			return err
		}
		return showcaseutil.Button(ui.ButtonProps{Variant: "ghost", Size: "sm", Attrs: templ.Attributes{"aria-pressed": "false"}}, "Right").Render(ctx, w)
	})
}
