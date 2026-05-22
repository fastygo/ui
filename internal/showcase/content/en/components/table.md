---
slug: table
section: components
title: "Table"
description: "Semantic data table structure."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Data Table"
    href: /docs/components/data-table/
  - label: "List"
    href: /docs/components/list/
api:
  - name: "Class"
    type: "string"
    description: "Table wrapper utilities"
  - name: "TableHead"
    type: "component"
    description: "thead section"
  - name: "TableBody"
    type: "component"
    description: "tbody section"
---

Semantic data table structure.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Table(ui.TableProps{Class: "w-full text-sm border border-border"}) {
		@ui.TableHead(ui.TableSectionProps{}) {
			@ui.TableRow(ui.TableRowProps{}) {
				@ui.TableHeadCell(ui.TableCellProps{}) {
					Name
				}
				@ui.TableHeadCell(ui.TableCellProps{}) {
					Role
				}
			}
		}
		@ui.TableBody(ui.TableSectionProps{}) {
			@ui.TableRow(ui.TableRowProps{}) {
				@ui.TableCell(ui.TableCellProps{}) {
					Ada
				}
				@ui.TableCell(ui.TableCellProps{}) {
					Admin
				}
			}
			@ui.TableRow(ui.TableRowProps{}) {
				@ui.TableCell(ui.TableCellProps{}) {
					Lin
				}
				@ui.TableCell(ui.TableCellProps{}) {
					Editor
				}
			}
		}
	}
}
```

## Compact

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Table(ui.TableProps{Class: "w-full text-sm border border-border"}) {
		@ui.TableHead(ui.TableSectionProps{}) {
			@ui.TableRow(ui.TableRowProps{}) {
				@ui.TableHeadCell(ui.TableCellProps{Class: "text-xs"}) {
					Name
				}
				@ui.TableHeadCell(ui.TableCellProps{Class: "text-xs"}) {
					Role
				}
			}
		}
		@ui.TableBody(ui.TableSectionProps{}) {
			@ui.TableRow(ui.TableRowProps{}) {
				@ui.TableCell(ui.TableCellProps{Class: "text-xs"}) {
					Ada
				}
				@ui.TableCell(ui.TableCellProps{Class: "text-xs"}) {
					Admin
				}
			}
			@ui.TableRow(ui.TableRowProps{}) {
				@ui.TableCell(ui.TableCellProps{Class: "text-xs"}) {
					Lin
				}
				@ui.TableCell(ui.TableCellProps{Class: "text-xs"}) {
					Editor
				}
			}
		}
	}
}
```
