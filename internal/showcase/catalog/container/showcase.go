package container

import (
	"context"
	"io"

	"github.com/fastygo/templ/ui"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	return showcaseutil.RenderWithChildren(ctx, w, ui.Container(ui.ContainerProps{Class: "mx-auto max-w-3xl border border-border rounded-lg p-4"}), func(ctx context.Context, w io.Writer) error {
		return ui.Text(ui.TextProps{}, "Container content.").Render(ctx, w)
	})
}
func previewSection(ctx context.Context, w io.Writer) error {
	return showcaseutil.RenderWithChildren(ctx, w, ui.Container(ui.ContainerProps{Tag: "section", Class: "mx-auto max-w-2xl border border-border rounded-lg p-4"}), func(ctx context.Context, w io.Writer) error {
		return ui.Text(ui.TextProps{}, "Section landmark.").Render(ctx, w)
	})
}
