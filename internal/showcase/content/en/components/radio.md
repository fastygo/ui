---
slug: radio
section: components
title: "Radio"
description: "Single-choice radio input."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Checkbox"
    href: /docs/components/checkbox/
  - label: "Switch"
    href: /docs/components/switch/
api:
  - name: "Name"
    type: "string"
    description: "Form field name"
  - name: "Value"
    type: "string"
    description: "Option value"
  - name: "Checked"
    type: "bool"
    description: "Initial checked state"
---

Single-choice radio input.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Radio(ui.RadioProps{Name: "plan", Value: "free", AriaLabel: "Free plan"})
}
```

## Checked

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Radio(ui.RadioProps{Name: "plan", Value: "pro", Checked: true, AriaLabel: "Pro plan"})
}
```
