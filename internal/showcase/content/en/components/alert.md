---
slug: alert
section: components
title: "Alert"
description: "Callout for important messages."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Card"
    href: /docs/components/card/
  - label: "Badge"
    href: /docs/components/badge/
api:
  - name: "Variant"
    type: "string"
    description: "default | destructive"
  - name: "Class"
    type: "string"
    description: "Extra utilities"
---

Callout for important messages.

## Default

```templ
import cmp "github.com/fastygo/templ/components"
import "github.com/fastygo/templ/ui"

templ Example() {
	@cmp.Alert(cmp.AlertProps{}) {
		@ui.Title(ui.TitleProps{Order: 4, Class: "text-sm font-semibold"}, "Heads up")
		@ui.Text(ui.TextProps{Class: "text-sm"}, "You can add components from the gallery.")
	}
}
```

## Destructive

```templ
import cmp "github.com/fastygo/templ/components"
import "github.com/fastygo/templ/ui"

templ Example() {
	@cmp.Alert(cmp.AlertProps{Variant: "destructive"}) {
		@ui.Title(ui.TitleProps{Order: 4, Class: "text-sm font-semibold"}, "Error")
		@ui.Text(ui.TextProps{Class: "text-sm"}, "Something went wrong.")
	}
}
```
