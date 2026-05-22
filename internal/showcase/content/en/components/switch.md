---
slug: switch
section: components
title: "Switch"
description: "Toggle switch input."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Checkbox"
    href: /docs/components/checkbox/
  - label: "Toggle"
    href: /docs/components/toggle/
api:
  - name: "Name"
    type: "string"
    description: "Form field name"
  - name: "Checked"
    type: "bool"
    description: "Initial on state"
  - name: "Disabled"
    type: "bool"
    description: "Disables control"
---

Toggle switch input.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Switch(ui.SwitchProps{Name: "airplane", AriaLabel: "Airplane mode"})
}
```

## Checked

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Switch(ui.SwitchProps{Name: "airplane", Checked: true, AriaLabel: "Airplane mode on"})
}
```
