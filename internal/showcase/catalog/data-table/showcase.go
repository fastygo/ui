package datatable

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/fastygo/templ/ui"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func wfStack(ctx context.Context, w io.Writer, attrs templ.Attributes, body func(context.Context, io.Writer) error) error {
	props := ui.StackProps{Class: "gap-2 max-w-md"}
	if attrs != nil {
		props.Attrs = attrs
	}
	return showcaseutil.RenderWithChildren(ctx, w, ui.Stack(props), body)
}

func previewDefault(ctx context.Context, w io.Writer) error {
	return showcaseutil.RenderWithChildren(ctx, w, ui.Stack(ui.StackProps{Class: "gap-3 max-w-lg"}), func(ctx context.Context, w io.Writer) error {
		if err := ui.Group(ui.GroupProps{Class: "flex items-center justify-between"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
			if err := ui.Input(ui.InputProps{Placeholder: "Filter rows…", Class: "max-w-xs"}).Render(ctx, w); err != nil {
				return err
			}
			return showcaseutil.Button(ui.ButtonProps{Size: "sm", Variant: "outline"}, "Columns").Render(ctx, w)
		})), w); err != nil {
			return err
		}
		return showcaseutil.RenderWithChildren(ctx, w, ui.Table(ui.TableProps{Class: "w-full text-sm border border-border"}), func(ctx context.Context, w io.Writer) error {
			return showcaseutil.RenderWithChildren(ctx, w, ui.TableBody(ui.TableSectionProps{}), func(ctx context.Context, w io.Writer) error {
				return showcaseutil.RenderWithChildren(ctx, w, ui.TableRow(ui.TableRowProps{}), func(ctx context.Context, w io.Writer) error {
					if err := showcaseutil.RenderTableCell(ctx, w, ui.TableCellProps{}, "Row A"); err != nil {
						return err
					}
					return showcaseutil.RenderTableCell(ctx, w, ui.TableCellProps{}, "Active")
				})
			})
		})
	})
}
