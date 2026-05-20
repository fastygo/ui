---
slug: radio
section: components
title: "Radio"
description: "Single choice within a group."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Checkbox"
    href: /docs/components/checkbox/
  - label: "Select"
    href: /docs/components/select/
api:
  - name: "Name"
    type: "string"
    description: "Group name"
  - name: "Value"
    type: "string"
    description: "Option value"
  - name: "Checked"
    type: "bool"
    description: "Selected state"
---

Single choice within a group.

## Default

{{demo id="radio.default"}}

```templ
@ui.Radio(ui.RadioProps{Name: "plan", Value: "free"})
```

## Selected

{{demo id="radio.checked"}}

```templ
@ui.Radio(ui.RadioProps{Checked: true})
```
