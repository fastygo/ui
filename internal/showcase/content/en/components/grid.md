---
slug: grid
section: components
title: "Grid"
description: "CSS grid layout with columns."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Stack"
    href: /docs/components/stack/
  - label: "Container"
    href: /docs/components/container/
api:
  - name: "Class"
    type: "string"
    description: "grid-cols-* and gap utilities"
  - name: "GridCol"
    type: "component"
    description: "Column cell wrapper"
---

CSS grid layout with columns.

## Two columns

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Grid(ui.GridProps{Class: "grid-cols-2 gap-4 max-w-md"}) {
		@ui.GridCol(ui.GridColProps{}) {
			@ui.Box(ui.BoxProps{Class: "rounded border border-border p-3 text-center text-sm"}) {
				@ui.Text(ui.TextProps{}, "1")
			}
		}
		@ui.GridCol(ui.GridColProps{}) {
			@ui.Box(ui.BoxProps{Class: "rounded border border-border p-3 text-center text-sm"}) {
				@ui.Text(ui.TextProps{}, "2")
			}
		}
	}
}
```

## Three columns

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Grid(ui.GridProps{Class: "grid-cols-3 gap-2 max-w-md"}) {
		@ui.GridCol(ui.GridColProps{}) {
			@ui.Box(ui.BoxProps{Class: "rounded border border-border p-3 text-center text-sm"}) {
				@ui.Text(ui.TextProps{}, "1")
			}
		}
		@ui.GridCol(ui.GridColProps{}) {
			@ui.Box(ui.BoxProps{Class: "rounded border border-border p-3 text-center text-sm"}) {
				@ui.Text(ui.TextProps{}, "2")
			}
		}
		@ui.GridCol(ui.GridColProps{}) {
			@ui.Box(ui.BoxProps{Class: "rounded border border-border p-3 text-center text-sm"}) {
				@ui.Text(ui.TextProps{}, "3")
			}
		}
	}
}
```
