---
slug: checkbox
section: components
title: "Checkbox"
description: "Boolean checkbox input."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Radio"
    href: /docs/components/radio/
  - label: "Switch"
    href: /docs/components/switch/
api:
  - name: "Name"
    type: "string"
    description: "Form field name"
  - name: "Checked"
    type: "bool"
    description: "Initial checked state"
  - name: "Disabled"
    type: "bool"
    description: "Disables control"
---

Boolean checkbox input.

## Default

{{demo id="checkbox.default"}}

```templ
@ui.Checkbox(ui.CheckboxProps{Name: "terms"})
```

## Checked

{{demo id="checkbox.checked"}}

```templ
@ui.Checkbox(ui.CheckboxProps{Checked: true})
```
