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

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Input(ui.InputProps{Type: "range", Min: "0", Max: "100", AriaLabel: "Volume"})
}
```

## With value

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Input(ui.InputProps{Type: "range", Min: "0", Max: "100", Value: "50", AriaLabel: "Brightness"})
}
```
