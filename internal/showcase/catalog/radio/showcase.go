package radio

import (
	"context"
	"io"

	"github.com/fastygo/templ/ui"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	return ui.Radio(ui.RadioProps{Name: "plan", Value: "free", AriaLabel: "Free plan"}).Render(ctx, w)
}
func previewChecked(ctx context.Context, w io.Writer) error {
	return ui.Radio(ui.RadioProps{Name: "plan", Value: "pro", Checked: true, AriaLabel: "Pro plan"}).Render(ctx, w)
}
