package title

import (
	"context"
	"io"

	"github.com/fastygo/templ/ui"
)

func previewH1(ctx context.Context, w io.Writer) error {
	return ui.Title(ui.TitleProps{Order: 1}, "Page title").Render(ctx, w)
}
func previewH2(ctx context.Context, w io.Writer) error {
	return ui.Title(ui.TitleProps{Order: 2}, "Section").Render(ctx, w)
}
func previewH3(ctx context.Context, w io.Writer) error {
	return ui.Title(ui.TitleProps{Order: 3}, "Subsection").Render(ctx, w)
}
