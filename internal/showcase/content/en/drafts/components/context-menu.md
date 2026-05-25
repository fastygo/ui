---
slug: context-menu
section: components
title: "Context Menu"
description: "Stub menu list on right-click target."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Dropdown Menu"
    href: /docs/components/dropdown-menu/
  - label: "Menubar"
    href: /docs/components/menubar/
api:
  - name: "List"
    type: "List"
    description: "menu tag"
---

Stub menu list on right-click target.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2"}) {
		@ui.Box(ui.BoxProps{Class: "rounded-lg border border-dashed border-border p-8 text-center text-sm text-muted-foreground"}) {
			@ui.Text(ui.TextProps{}, "Right-click area (stub)")
		}
		@ui.List(ui.ListProps{Tag: "menu", Class: "w-40 rounded-md border border-border bg-card p-1 text-sm"}) {
			@ui.ListItem(ui.ListItemProps{}) {
				@ui.Text(ui.TextProps{}, "Copy")
			}
			@ui.ListItem(ui.ListItemProps{}) {
				@ui.Text(ui.TextProps{}, "Paste")
			}
		}
	}
}
```
