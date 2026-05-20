package button

import (
	"context"
	"io"

	"github.com/fastygo/ui/internal/showcase/showcaseutil"

	"github.com/fastygo/templ/ui"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	return showcaseutil.Button(ui.ButtonProps{}, "Button").Render(ctx, w)
}

func previewSecondary(ctx context.Context, w io.Writer) error {
	return showcaseutil.Button(ui.ButtonProps{Variant: "secondary"}, "Secondary").Render(ctx, w)
}

func previewOutline(ctx context.Context, w io.Writer) error {
	return showcaseutil.Button(ui.ButtonProps{Variant: "outline"}, "Outline").Render(ctx, w)
}

func previewDestructive(ctx context.Context, w io.Writer) error {
	return showcaseutil.Button(ui.ButtonProps{Variant: "destructive"}, "Destructive").Render(ctx, w)
}

func previewGhost(ctx context.Context, w io.Writer) error {
	return showcaseutil.Button(ui.ButtonProps{Variant: "ghost"}, "Ghost").Render(ctx, w)
}

func previewLink(ctx context.Context, w io.Writer) error {
	return showcaseutil.Button(ui.ButtonProps{Variant: "link"}, "Link").Render(ctx, w)
}

func previewSizes(ctx context.Context, w io.Writer) error {
	if err := showcaseutil.Button(ui.ButtonProps{Size: "sm"}, "Small").Render(ctx, w); err != nil {
		return err
	}
	if err := showcaseutil.Button(ui.ButtonProps{}, "Default").Render(ctx, w); err != nil {
		return err
	}
	return showcaseutil.Button(ui.ButtonProps{Size: "lg"}, "Large").Render(ctx, w)
}
