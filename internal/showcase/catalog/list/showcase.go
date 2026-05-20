package list

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/fastygo/templ/ui"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func previewUnordered(ctx context.Context, w io.Writer) error {
	return listPreview(ctx, w, ui.ListProps{Class: "list-disc pl-5 text-sm"})
}
func previewOrdered(ctx context.Context, w io.Writer) error {
	return listPreview(ctx, w, ui.ListProps{Tag: "ol", Class: "list-decimal pl-5 text-sm"})
}
func listPreview(ctx context.Context, w io.Writer, props ui.ListProps) error {
	return showcaseutil.RenderWithChildren(ctx, w, ui.List(props), func(ctx context.Context, w io.Writer) error {
		for _, item := range []string{"First item", "Second item"} {
			if err := ui.ListItem(ui.ListItemProps{}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
				return ui.Text(ui.TextProps{}, item).Render(ctx, w)
			})), w); err != nil {
				return err
			}
		}
		return nil
	})
}
