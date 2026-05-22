---
slug: stack
section: components
title: "Stack"
description: "Vertical flex column for stacking children with gap utilities."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Group"
    href: /docs/components/group/
  - label: "Box"
    href: /docs/components/box/
api:
  - name: "Class"
    type: "string"
    description: "Tailwind utilities including gap"
  - name: "Tag"
    type: "string"
    description: "div | nav | section | ul | …"
  - name: "Attrs"
    type: "templ.Attributes"
    description: "Extra HTML attributes"
---

Vertical flex column for stacking children with gap utilities.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2"}) {
		@ui.Title(ui.TitleProps{Order: 3}, "Stack")
		@ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "Children stack vertically.")
	}
}
```

## Row

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "flex-row items-center gap-2"}) {
		@ui.Title(ui.TitleProps{Order: 3}, "Stack")
		@ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "Children stack vertically.")
	}
}
```

## Nav tag

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Tag: "nav", Class: "gap-1"}) {
		@ui.Title(ui.TitleProps{Order: 3}, "Stack")
		@ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "Children stack vertically.")
	}
}
```
