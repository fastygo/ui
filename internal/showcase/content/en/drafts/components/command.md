---
slug: command
section: components
title: "Command"
description: "Command palette: search + list."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Combobox"
    href: /docs/components/combobox/
  - label: "Dialog"
    href: /docs/components/dialog/
api:
  - name: "Input"
    type: "Input"
    description: "Search field"
  - name: "List"
    type: "List"
    description: "Commands"
---

Command palette: search + list.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2 max-w-sm rounded-lg border border-border bg-card p-2"}) {
		@ui.Input(ui.InputProps{Placeholder: "Type a command…"})
		@ui.List(ui.ListProps{Class: "text-sm"}) {
			@ui.ListItem(ui.ListItemProps{}) {
				@ui.Button(ui.ButtonProps{Variant: "ghost", Class: "h-8 w-full justify-start"}) {
					Open docs
				}
			}
			@ui.ListItem(ui.ListItemProps{}) {
				@ui.Button(ui.ButtonProps{Variant: "ghost", Class: "h-8 w-full justify-start"}) {
					Toggle theme
				}
			}
			@ui.ListItem(ui.ListItemProps{}) {
				@ui.Button(ui.ButtonProps{Variant: "ghost", Class: "h-8 w-full justify-start"}) {
					Go home
				}
			}
		}
	}
}
```
