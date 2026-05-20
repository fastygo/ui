package toggle

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/fastygo/templ/ui"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	return showcaseutil.Button(ui.ButtonProps{Variant: "outline", Attrs: templ.Attributes{"aria-pressed": "false"}}, "Bold").Render(ctx, w)
}
func previewPressed(ctx context.Context, w io.Writer) error {
	return showcaseutil.Button(ui.ButtonProps{Variant: "secondary", Attrs: templ.Attributes{"aria-pressed": "true"}}, "Bold").Render(ctx, w)
}
