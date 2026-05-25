---
slug: dropdown-menu
section: components
title: "Dropdown Menu"
description: "Menu triggered by a button."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Navigation Menu"
    href: /docs/components/navigation-menu/
  - label: "Menubar"
    href: /docs/components/menubar/
api:
  - name: "List"
    type: "List"
    description: "menu tag for items"
---

Menu triggered by a button.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2"}) {
		@ui.Button(ui.ButtonProps{Variant: "outline"}) {
			Open menu
		}
		@ui.List(ui.ListProps{Tag: "menu", Class: "w-40 rounded-md border border-border bg-card p-1 text-sm"}) {
			@ui.ListItem(ui.ListItemProps{Class: "rounded px-2 py-1 hover:bg-accent"}) {
				@ui.Button(ui.ButtonProps{Variant: "ghost", Class: "h-8 w-full justify-start px-2"}) {
					Profile
				}
			}
			@ui.ListItem(ui.ListItemProps{Class: "rounded px-2 py-1 hover:bg-accent"}) {
				@ui.Button(ui.ButtonProps{Variant: "ghost", Class: "h-8 w-full justify-start px-2"}) {
					Settings
				}
			}
			@ui.ListItem(ui.ListItemProps{Class: "rounded px-2 py-1 hover:bg-accent"}) {
				@ui.Button(ui.ButtonProps{Variant: "ghost", Class: "h-8 w-full justify-start px-2"}) {
					Sign out
				}
			}
		}
	}
}
```
