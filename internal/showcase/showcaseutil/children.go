package showcaseutil

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/fastygo/templ/ui"
)

// TextChild is a templ child that writes plain text (for Button, Badge, Label bodies in previews).
func TextChild(s string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, s)
		return err
	})
}

// RenderButton renders ui.Button with a text child (Go previews cannot use @ui.Button { } syntax).
func RenderButton(ctx context.Context, w io.Writer, p ui.ButtonProps, label string) error {
	return ui.Button(p).Render(templ.WithChildren(ctx, TextChild(label)), w)
}

// Button returns a component for use with Render(ctx, w, …).
func Button(p ui.ButtonProps, label string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return RenderButton(ctx, w, p, label)
	})
}

// RenderLabel renders ui.Label with a text child.
func RenderLabel(ctx context.Context, w io.Writer, p ui.LabelProps, label string) error {
	return ui.Label(p).Render(templ.WithChildren(ctx, TextChild(label)), w)
}

// RenderBadge renders ui.Badge with a text child.
func RenderBadge(ctx context.Context, w io.Writer, p ui.BadgeProps, label string) error {
	return ui.Badge(p).Render(templ.WithChildren(ctx, TextChild(label)), w)
}

// RenderTableHeadCell renders ui.TableHeadCell with a text child.
func RenderTableHeadCell(ctx context.Context, w io.Writer, p ui.TableCellProps, label string) error {
	return ui.TableHeadCell(p).Render(templ.WithChildren(ctx, TextChild(label)), w)
}

// RenderTableCell renders ui.TableCell with a text child.
func RenderTableCell(ctx context.Context, w io.Writer, p ui.TableCellProps, label string) error {
	return ui.TableCell(p).Render(templ.WithChildren(ctx, TextChild(label)), w)
}
