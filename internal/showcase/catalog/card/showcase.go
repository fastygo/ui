package card

import (
	"context"
	"io"

	cmp "github.com/fastygo/templ/components"
	"github.com/fastygo/templ/ui"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func previewDefault(ctx context.Context, w io.Writer) error { return cardPreview(ctx, w, false) }
func previewFooter(ctx context.Context, w io.Writer) error  { return cardPreview(ctx, w, true) }
func cardPreview(ctx context.Context, w io.Writer, footer bool) error {
	return showcaseutil.RenderWithChildren(ctx, w, cmp.Card(cmp.CardProps{Class: "max-w-sm"}), func(ctx context.Context, w io.Writer) error {
		h := func(ctx context.Context, w io.Writer) error {
			return showcaseutil.RenderWithChildren(ctx, w, cmp.CardHeader(cmp.CardHeaderProps{}), func(ctx context.Context, w io.Writer) error {
				if err := cmp.CardTitle(cmp.CardTitleProps{}, "Card title").Render(ctx, w); err != nil {
					return err
				}
				return cmp.CardDescription(cmp.CardDescriptionProps{}, "Wireframe card description.").Render(ctx, w)
			})
		}
		c := func(ctx context.Context, w io.Writer) error {
			return showcaseutil.RenderWithChildren(ctx, w, cmp.CardContent(cmp.CardContentProps{}), func(ctx context.Context, w io.Writer) error {
				return ui.Text(ui.TextProps{}, "Card body copy.").Render(ctx, w)
			})
		}
		if err := h(ctx, w); err != nil {
			return err
		}
		if err := c(ctx, w); err != nil {
			return err
		}
		if footer {
			return showcaseutil.RenderWithChildren(ctx, w, cmp.CardFooter(cmp.CardFooterProps{}), func(ctx context.Context, w io.Writer) error {
				return showcaseutil.Button(ui.ButtonProps{Size: "sm"}, "Action").Render(ctx, w)
			})
		}
		return nil
	})
}
