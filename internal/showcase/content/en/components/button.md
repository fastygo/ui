---
slug: button
section: components
title: "Button"
description: "Triggers an action or navigates when rendered as a link. Built on github.com/fastygo/templ/ui."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Toggle"
    href: /docs/components/toggle/
  - label: "Form"
    href: /docs/components/form/
api:
  - name: "Variant"
    type: "string"
    description: "default | secondary | destructive | outline | ghost | link | unstyled"
  - name: "Size"
    type: "string"
    description: "default | sm | lg | icon"
  - name: "Type"
    type: "string"
    description: "button | submit | reset"
  - name: "Href"
    type: "string"
    description: "When set, renders an anchor instead of button"
  - name: "Disabled"
    type: "bool"
    description: "Disables interaction"
  - name: "Class"
    type: "string"
    description: "Additional Tailwind utilities"
  - name: "AriaLabel"
    type: "string"
    description: "Accessible name when visible text is insufficient"
---

Triggers an action or navigates when rendered as a link. Built on github.com/fastygo/templ/ui.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{}) {
		Button
	}
}
```

## Secondary

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Variant: "secondary"}) {
		Secondary
	}
}
```

## Outline

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Variant: "outline"}) {
		Outline
	}
}
```

## Destructive

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Variant: "destructive"}) {
		Destructive
	}
}
```

## Ghost

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Variant: "ghost"}) {
		Ghost
	}
}
```

## Link

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Variant: "link"}) {
		Link
	}
}
```

## Sizes

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "gap-2"}) {
		@ui.Button(ui.ButtonProps{Size: "sm"}) {
			Small
		}
		@ui.Button(ui.ButtonProps{}) {
			Default
		}
		@ui.Button(ui.ButtonProps{Size: "lg"}) {
			Large
		}
	}
}
```
