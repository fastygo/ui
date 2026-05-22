---
slug: navigation-menu
section: components
title: "Navigation Menu"
description: "Site navigation with sections."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Menubar"
    href: /docs/components/menubar/
  - label: "Breadcrumb"
    href: /docs/components/breadcrumb/
api:
  - name: "Tag"
    type: "string"
    description: "menu for menubar-style lists"
---

Site navigation with sections.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.List(ui.ListProps{Tag: "menu", Class: "flex gap-4 text-sm"}) {
		@ui.ListItem(ui.ListItemProps{}) {
			@ui.Button(ui.ButtonProps{Variant: "link"}) {
				Home
			}
		}
		@ui.ListItem(ui.ListItemProps{}) {
			@ui.Button(ui.ButtonProps{Variant: "link"}) {
				Docs
			}
		}
		@ui.ListItem(ui.ListItemProps{}) {
			@ui.Button(ui.ButtonProps{Variant: "link"}) {
				Blog
			}
		}
	}
}
```
