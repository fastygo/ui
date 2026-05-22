---
slug: badge
section: components
title: "Badge"
description: "Small status label chip."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Button"
    href: /docs/components/button/
  - label: "Alert"
    href: /docs/components/alert/
api:
  - name: "Variant"
    type: "string"
    description: "default | secondary | destructive | outline"
  - name: "Size"
    type: "string"
    description: "default | sm | lg"
  - name: "Class"
    type: "string"
    description: "Extra utilities"
---

Small status label chip.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Badge(ui.BadgeProps{}) {
		Badge
	}
}
```

## Secondary

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Badge(ui.BadgeProps{Variant: "secondary"}) {
		Secondary
	}
}
```

## Outline

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Badge(ui.BadgeProps{Variant: "outline"}) {
		Outline
	}
}
```

## Destructive

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Badge(ui.BadgeProps{Variant: "destructive"}) {
		Alert
	}
}
```
