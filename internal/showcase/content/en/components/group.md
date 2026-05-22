---
slug: group
section: components
title: "Group"
description: "Horizontal flex row for grouping controls."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Stack"
    href: /docs/components/stack/
  - label: "Box"
    href: /docs/components/box/
api:
  - name: "Class"
    type: "string"
    description: "Tailwind utilities"
  - name: "Tag"
    type: "string"
    description: "div | span"
  - name: "Attrs"
    type: "templ.Attributes"
    description: "Extra attributes"
---

Horizontal flex row for grouping controls.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "flex items-center gap-2"}) {
		@ui.Button(ui.ButtonProps{Size: "sm"}) {
			One
		}
		@ui.Button(ui.ButtonProps{Size: "sm", Variant: "outline"}) {
			Two
		}
	}
}
```

## Wrap

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "flex flex-wrap gap-2 max-w-xs"}) {
		@ui.Badge(ui.BadgeProps{}) {
			1
		}
		@ui.Badge(ui.BadgeProps{}) {
			2
		}
		@ui.Badge(ui.BadgeProps{}) {
			3
		}
		@ui.Badge(ui.BadgeProps{}) {
			4
		}
	}
}
```
