package selectfield

import (
	"context"
	"io"

	"github.com/fastygo/templ/ui"
	"github.com/fastygo/templ/ui/selectfield"
)

var selectOpts = []selectfield.Option{{Value: "viewer", Label: "Viewer"}, {Value: "editor", Label: "Editor"}}

func previewDefault(ctx context.Context, w io.Writer) error {
	return ui.Select(ui.SelectProps{Name: "role", Options: selectOpts, Value: "viewer"}).Render(ctx, w)
}
func previewDisabled(ctx context.Context, w io.Writer) error {
	return ui.Select(ui.SelectProps{Name: "role", Options: selectOpts, Disabled: true}).Render(ctx, w)
}
