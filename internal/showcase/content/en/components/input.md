---
slug: input
section: components
title: "Input"
description: "Single-line text input control."
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
    description: "text | email | password | file | range | …"
  - name: "Placeholder"
    type: "string"
    description: "Placeholder text"
  - name: "Disabled"
    type: "bool"
    description: "Disables input"
  - name: "Class"
    type: "string"
    description: "Tailwind utilities"
---

Single-line text input control.

## Default

{{demo id="input.default"}}

```templ
@ui.Input(ui.InputProps{Placeholder: "Email"})
```

## Disabled

{{demo id="input.disabled"}}

```templ
@ui.Input(ui.InputProps{Disabled: true})
```

## File

{{demo id="input.file"}}

```templ
@ui.Input(ui.InputProps{Type: "file"})
```
