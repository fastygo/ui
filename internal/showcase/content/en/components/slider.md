---
slug: slider
section: components
title: "Slider"
description: "Native range input styled via ui.Input."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Switch"
    href: /docs/components/switch/
  - label: "Input"
    href: /docs/components/input/
api:
  - name: "Min"
    type: "string"
    description: "Minimum value"
  - name: "Max"
    type: "string"
    description: "Maximum value"
  - name: "Value"
    type: "string"
    description: "Current value"
---

Native range input styled via ui.Input.

## Default

{{demo id="slider.default"}}

```templ
@ui.Input(ui.InputProps{Type: "range", Min: "0", Max: "100"})
```

## With value

{{demo id="slider.value"}}

```templ
@ui.Input(ui.InputProps{Type: "range", Value: "50"})
```
