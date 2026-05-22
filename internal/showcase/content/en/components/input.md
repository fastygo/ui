---
slug: input
section: components
title: "Input"
description: "Text input control."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Textarea"
    href: /docs/components/textarea/
  - label: "Form"
    href: /docs/components/form/
api:
  - name: "Type"
    type: "string"
    description: "HTML input type"
  - name: "Placeholder"
    type: "string"
    description: "Placeholder text"
  - name: "Disabled"
    type: "bool"
    description: "Disables control"
---

Text input control.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Input(ui.InputProps{Placeholder: "Email"})
}
```

## Disabled

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Input(ui.InputProps{Placeholder: "Disabled", Disabled: true})
}
```

## File

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Input(ui.InputProps{Type: "file"})
}
```
