---
slug: textarea
section: components
title: "Textarea"
description: "Multi-line text input."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Input"
    href: /docs/components/input/
  - label: "Form"
    href: /docs/components/form/
api:
  - name: "Placeholder"
    type: "string"
    description: "Placeholder text"
  - name: "Rows"
    type: "int"
    description: "Visible row count"
  - name: "Disabled"
    type: "bool"
    description: "Disables control"
---

Multi-line text input.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Textarea(ui.TextareaProps{Placeholder: "Your message"})
}
```

## Disabled

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Textarea(ui.TextareaProps{Placeholder: "Disabled", Disabled: true})
}
```
