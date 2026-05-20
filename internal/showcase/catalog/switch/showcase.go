package formswitch

import (
	"context"
	"io"

	"github.com/fastygo/templ/ui"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	return ui.Switch(ui.SwitchProps{Name: "airplane", AriaLabel: "Airplane mode"}).Render(ctx, w)
}
func previewChecked(ctx context.Context, w io.Writer) error {
	return ui.Switch(ui.SwitchProps{Name: "airplane", Checked: true, AriaLabel: "Airplane mode on"}).Render(ctx, w)
}
