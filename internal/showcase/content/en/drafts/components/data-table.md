---
slug: data-table
section: components
title: "Data Table"
description: "Table with toolbar (wireframe)."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Table"
    href: /docs/components/table/
  - label: "Pagination"
    href: /docs/components/pagination/
api:
  - name: "Table"
    type: "Table"
    description: "Semantic table"
  - name: "Input"
    type: "Input"
    description: "Filter field"
---

Table with toolbar (wireframe).

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-3 max-w-lg"}) {
		@ui.Group(ui.GroupProps{Class: "flex items-center justify-between"}) {
			@ui.Input(ui.InputProps{Placeholder: "Filter rows…", Class: "max-w-xs"})
			@ui.Button(ui.ButtonProps{Size: "sm", Variant: "outline"}) {
				Columns
			}
		}
		@ui.Table(ui.TableProps{Class: "w-full text-sm border border-border"}) {
			@ui.TableBody(ui.TableSectionProps{}) {
				@ui.TableRow(ui.TableRowProps{}) {
					@ui.TableCell(ui.TableCellProps{}) {
						Row A
					}
					@ui.TableCell(ui.TableCellProps{}) {
						Active
					}
				}
			}
		}
	}
}
```
