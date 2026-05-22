---
slug: toggle-group
section: components
title: "Toggle Group"
description: "Grouped toggle buttons."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Toggle"
    href: /docs/components/toggle/
  - label: "Tabs"
    href: /docs/components/tabs/
api:
  - name: "Class"
    type: "string"
    description: "Group layout utilities"
---

Grouped toggle buttons.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "inline-flex rounded-md border border-border"}) {
		@ui.Button(ui.ButtonProps{Variant: "secondary", Size: "sm", Attrs: templ.Attributes{"aria-pressed": "true"}}) {
			Left
		}
		@ui.Button(ui.ButtonProps{Variant: "ghost", Size: "sm", Attrs: templ.Attributes{"aria-pressed": "false"}}) {
			Right
		}
	}
}
```
