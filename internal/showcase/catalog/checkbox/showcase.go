package checkbox

import (
	"context"
	"io"

	"github.com/fastygo/templ/ui"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	return ui.Checkbox(ui.CheckboxProps{Name: "terms", AriaLabel: "Accept terms"}).Render(ctx, w)
}
func previewChecked(ctx context.Context, w io.Writer) error {
	return ui.Checkbox(ui.CheckboxProps{Name: "terms", Checked: true, AriaLabel: "Accepted"}).Render(ctx, w)
}
