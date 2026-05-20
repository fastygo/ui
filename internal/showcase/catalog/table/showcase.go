package table

import (
	"context"
	"io"

	"github.com/fastygo/templ/ui"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func previewDefault(ctx context.Context, w io.Writer) error { return renderSampleTable(ctx, w, "") }
func previewCompact(ctx context.Context, w io.Writer) error {
	return renderSampleTable(ctx, w, "text-xs")
}
func renderSampleTable(ctx context.Context, w io.Writer, cellClass string) error {
	return showcaseutil.RenderWithChildren(ctx, w, ui.Table(ui.TableProps{Class: "w-full text-sm border border-border"}), func(ctx context.Context, w io.Writer) error {
		head := func(ctx context.Context, w io.Writer) error {
			return showcaseutil.RenderWithChildren(ctx, w, ui.TableHead(ui.TableSectionProps{}), func(ctx context.Context, w io.Writer) error {
				return showcaseutil.RenderWithChildren(ctx, w, ui.TableRow(ui.TableRowProps{}), func(ctx context.Context, w io.Writer) error {
					if err := showcaseutil.RenderTableHeadCell(ctx, w, ui.TableCellProps{Class: cellClass}, "Name"); err != nil {
						return err
					}
					return showcaseutil.RenderTableHeadCell(ctx, w, ui.TableCellProps{Class: cellClass}, "Role")
				})
			})
		}
		body := func(ctx context.Context, w io.Writer) error {
			return showcaseutil.RenderWithChildren(ctx, w, ui.TableBody(ui.TableSectionProps{}), func(ctx context.Context, w io.Writer) error {
				rows := [][2]string{{"Ada", "Admin"}, {"Lin", "Editor"}}
				for _, row := range rows {
					if err := showcaseutil.RenderWithChildren(ctx, w, ui.TableRow(ui.TableRowProps{}), func(ctx context.Context, w io.Writer) error {
						if err := showcaseutil.RenderTableCell(ctx, w, ui.TableCellProps{Class: cellClass}, row[0]); err != nil {
							return err
						}
						return showcaseutil.RenderTableCell(ctx, w, ui.TableCellProps{Class: cellClass}, row[1])
					}); err != nil {
						return err
					}
				}
				return nil
			})
		}
		if err := head(ctx, w); err != nil {
			return err
		}
		return body(ctx, w)
	})
}
