package textarea

import (
	"context"
	"io"

	"github.com/fastygo/templ/ui"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	return ui.Textarea(ui.TextareaProps{Placeholder: "Your message"}).Render(ctx, w)
}
func previewDisabled(ctx context.Context, w io.Writer) error {
	return ui.Textarea(ui.TextareaProps{Placeholder: "Disabled", Disabled: true}).Render(ctx, w)
}
