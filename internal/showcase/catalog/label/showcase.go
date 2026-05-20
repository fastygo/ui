package label

import (
	"context"
	"io"

	"github.com/fastygo/templ/ui"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	if err := showcaseutil.RenderLabel(ctx, w, ui.LabelProps{HTMLFor: "showcase-email"}, "Email"); err != nil {
		return err
	}
	return ui.Input(ui.InputProps{ID: "showcase-email", Placeholder: "you@example.com"}).Render(ctx, w)
}
func previewRequired(ctx context.Context, w io.Writer) error {
	if err := showcaseutil.RenderLabel(ctx, w, ui.LabelProps{HTMLFor: "showcase-name"}, "Name"); err != nil {
		return err
	}
	return ui.Input(ui.InputProps{ID: "showcase-name", Required: true, Placeholder: "Required"}).Render(ctx, w)
}
