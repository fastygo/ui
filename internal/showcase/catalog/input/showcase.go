package input

import (
	"context"
	"io"

	"github.com/fastygo/templ/ui"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	return ui.Input(ui.InputProps{Placeholder: "Email"}).Render(ctx, w)
}
func previewDisabled(ctx context.Context, w io.Writer) error {
	return ui.Input(ui.InputProps{Placeholder: "Disabled", Disabled: true}).Render(ctx, w)
}
func previewFile(ctx context.Context, w io.Writer) error {
	return ui.Input(ui.InputProps{Type: "file"}).Render(ctx, w)
}
