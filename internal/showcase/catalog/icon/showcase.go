package iconcatalog

import (
	"context"
	"io"

	appicon "github.com/fastygo/ui/internal/ui/components/icon"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	return appicon.Icon(appicon.IconProps{Name: "home"}).Render(ctx, w)
}

func previewSizes(ctx context.Context, w io.Writer) error {
	for _, sz := range []string{"xs", "sm", "md", "lg"} {
		if err := appicon.Icon(appicon.IconProps{Name: "settings", Size: sz, Class: "mr-2"}).Render(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
