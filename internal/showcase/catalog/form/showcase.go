package form

import (
	"context"
	"io"

	"github.com/fastygo/templ/ui"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	return showcaseutil.RenderWithChildren(ctx, w, ui.Form(ui.FormProps{Class: "max-w-sm"}), func(ctx context.Context, w io.Writer) error {
		item := func(ctx context.Context, w io.Writer) error {
			return showcaseutil.RenderWithChildren(ctx, w, ui.FormItem(ui.FormItemProps{}), func(ctx context.Context, w io.Writer) error {
				if err := showcaseutil.RenderLabel(ctx, w, ui.LabelProps{HTMLFor: "login-email"}, "Email"); err != nil {
					return err
				}
				return ui.Input(ui.InputProps{ID: "login-email", Type: "email", Placeholder: "you@example.com"}).Render(ctx, w)
			})
		}
		if err := item(ctx, w); err != nil {
			return err
		}
		return showcaseutil.Button(ui.ButtonProps{Type: "submit"}, "Sign in").Render(ctx, w)
	})
}
func previewInline(ctx context.Context, w io.Writer) error {
	return showcaseutil.RenderWithChildren(ctx, w, ui.Form(ui.FormProps{Class: "max-w-md"}), func(ctx context.Context, w io.Writer) error {
		return showcaseutil.RenderWithChildren(ctx, w, ui.Group(ui.GroupProps{Class: "flex items-end gap-2"}), func(ctx context.Context, w io.Writer) error {
			if err := ui.Input(ui.InputProps{Placeholder: "Search"}).Render(ctx, w); err != nil {
				return err
			}
			return showcaseutil.Button(ui.ButtonProps{}, "Go").Render(ctx, w)
		})
	})
}
