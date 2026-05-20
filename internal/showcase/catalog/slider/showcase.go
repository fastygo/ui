package slider

import (
	"context"
	"io"

	"github.com/fastygo/templ/ui"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	return ui.Input(ui.InputProps{Type: "range", Min: "0", Max: "100", AriaLabel: "Volume"}).Render(ctx, w)
}
func previewValue(ctx context.Context, w io.Writer) error {
	return ui.Input(ui.InputProps{Type: "range", Min: "0", Max: "100", Value: "50", AriaLabel: "Brightness"}).Render(ctx, w)
}
