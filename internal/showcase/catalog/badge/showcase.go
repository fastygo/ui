package badge

import (
	"context"
	"io"

	"github.com/fastygo/templ/ui"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func badge(ctx context.Context, w io.Writer, p ui.BadgeProps, label string) error {
	return showcaseutil.RenderWithChildren(ctx, w, ui.Badge(p), func(ctx context.Context, w io.Writer) error {
		return ui.Text(ui.TextProps{}, label).Render(ctx, w)
	})
}
func previewDefault(ctx context.Context, w io.Writer) error {
	return badge(ctx, w, ui.BadgeProps{}, "Badge")
}
func previewSecondary(ctx context.Context, w io.Writer) error {
	return badge(ctx, w, ui.BadgeProps{Variant: "secondary"}, "Secondary")
}
func previewOutline(ctx context.Context, w io.Writer) error {
	return badge(ctx, w, ui.BadgeProps{Variant: "outline"}, "Outline")
}
func previewDestructive(ctx context.Context, w io.Writer) error {
	return badge(ctx, w, ui.BadgeProps{Variant: "destructive"}, "Alert")
}
