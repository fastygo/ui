---
slug: pagination
section: components
title: "Pagination"
description: "Page navigation controls."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Table"
    href: /docs/components/table/
  - label: "Breadcrumb"
    href: /docs/components/breadcrumb/
api:
  - name: "Buttons"
    type: "Button"
    description: "Prev / page / Next"
---

Page navigation controls.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "flex items-center gap-1"}) {
		@ui.Button(ui.ButtonProps{Variant: "outline", Size: "sm"}) {
			Prev
		}
		@ui.Button(ui.ButtonProps{Variant: "secondary", Size: "sm"}) {
			1
		}
		@ui.Button(ui.ButtonProps{Variant: "outline", Size: "sm"}) {
			Next
		}
	}
}
```
