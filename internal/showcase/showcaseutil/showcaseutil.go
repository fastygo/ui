package showcaseutil

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/fastygo/ui/internal/registry"
)

// Variant builds a registry variant with a ComponentFunc preview.
func Variant(id, title, description, code string, render func(ctx context.Context, w io.Writer) error) registry.Variant {
	return registry.VariantFromFunc(id, title, description, code, render)
}

// Render writes a child templ component to w (helper for ComponentFunc bodies).
func Render(ctx context.Context, w io.Writer, c templ.Component) error {
	if c == nil {
		return nil
	}
	return c.Render(ctx, w)
}

// Child wraps a render func as a templ child for WithChildren.
func Child(render func(ctx context.Context, w io.Writer) error) templ.Component {
	if render == nil {
		return templ.NopComponent
	}
	return templ.ComponentFunc(render)
}

// RenderWithChildren renders parent with a single child body.
func RenderWithChildren(ctx context.Context, w io.Writer, parent templ.Component, body func(ctx context.Context, w io.Writer) error) error {
	if parent == nil {
		return nil
	}
	return parent.Render(templ.WithChildren(ctx, Child(body)), w)
}
