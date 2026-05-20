package registry

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

// VariantFromFunc builds a Variant with a ComponentFunc preview.
func VariantFromFunc(id, title, description, code string, render func(ctx context.Context, w io.Writer) error) Variant {
	return Variant{
		ID:          id,
		Title:       title,
		Description: description,
		Code:        code,
		Preview: templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if render == nil {
				return nil
			}
			return render(ctx, w)
		}),
	}
}
