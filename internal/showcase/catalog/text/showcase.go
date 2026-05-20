package text

import (
	"context"
	"io"

	"github.com/fastygo/templ/ui"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	return ui.Text(ui.TextProps{}, "Body copy.").Render(ctx, w)
}
func previewMuted(ctx context.Context, w io.Writer) error {
	return ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "Muted supporting text.").Render(ctx, w)
}
func previewCode(ctx context.Context, w io.Writer) error {
	return ui.Text(ui.TextProps{Tag: "code", Class: "font-mono text-xs"}, "npm install").Render(ctx, w)
}
