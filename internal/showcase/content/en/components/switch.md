---
slug: switch
section: components
title: "Switch"
description: "Toggle switch (formswitch / ui.Switch)."
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
    description: "On state"
  - name: "AriaLabel"
    type: "string"
    description: "Accessible name"
---

Toggle switch (formswitch / ui.Switch).

## Default

{{demo id="switch.default"}}

```templ
@ui.Switch(ui.SwitchProps{Name: "airplane"})
```

## On

{{demo id="switch.checked"}}

```templ
@ui.Switch(ui.SwitchProps{Checked: true})
```
