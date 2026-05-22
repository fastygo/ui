---
slug: combobox
section: components
title: "Combobox"
description: "Searchable select wireframe (data-ui8kit combobox)."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Select"
    href: /docs/components/select/
  - label: "Command"
    href: /docs/components/command/
api:
  - name: "Input"
    type: "Input"
    description: "Filter field"
  - name: "List"
    type: "List"
    description: "Options listbox"
---

Searchable select wireframe (data-ui8kit combobox).

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2 max-w-md", Attrs: templ.Attributes{"data-ui8kit": "combobox"}}) {
		@ui.Input(ui.InputProps{Placeholder: "Search framework…", Attrs: templ.Attributes{"role": "combobox", "aria-expanded": "true"}})
		@ui.List(ui.ListProps{Class: "rounded-md border border-border bg-card p-1 text-sm", Attrs: templ.Attributes{"role": "listbox"}}) {
			@ui.ListItem(ui.ListItemProps{}) {
				@ui.Button(ui.ButtonProps{Variant: "ghost", Class: "h-8 w-full justify-start"}) {
					Go
				}
			}
			@ui.ListItem(ui.ListItemProps{}) {
				@ui.Button(ui.ButtonProps{Variant: "ghost", Class: "h-8 w-full justify-start"}) {
					TypeScript
				}
			}
			@ui.ListItem(ui.ListItemProps{}) {
				@ui.Button(ui.ButtonProps{Variant: "ghost", Class: "h-8 w-full justify-start"}) {
					Rust
				}
			}
		}
	}
}
```
