---
slug: menubar
section: components
title: "Menubar"
description: "Horizontal menu bar."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Navigation Menu"
    href: /docs/components/navigation-menu/
  - label: "Dropdown Menu"
    href: /docs/components/dropdown-menu/
api:
  - name: "Group"
    type: "Group"
    description: "Horizontal button row"
---

Horizontal menu bar.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "flex gap-1 rounded-md border border-border bg-card p-1 text-sm"}) {
		@ui.Button(ui.ButtonProps{Variant: "ghost", Size: "sm"}) {
			File
		}
		@ui.Button(ui.ButtonProps{Variant: "ghost", Size: "sm"}) {
			Edit
		}
		@ui.Button(ui.ButtonProps{Variant: "ghost", Size: "sm"}) {
			View
		}
	}
}
```
