---
slug: slider
section: components
title: "Slider"
description: "Нативный range input через ui.Input."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Switch"
    href: /docs/primitives/switch/
  - label: "Поле ввода"
    href: /docs/primitives/input/
api:
  - name: "Min"
    type: "string"
    description: "Минимальное значение"
  - name: "Max"
    type: "string"
    description: "Максимальное значение"
  - name: "Value"
    type: "string"
    description: "Текущее значение"
---

Нативный range input через ui.Input.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Input(ui.InputProps{Type: "range", Min: "0", Max: "100", AriaLabel: "Громкость"})
}
```

## Со значением

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Input(ui.InputProps{Type: "range", Min: "0", Max: "100", Value: "50", AriaLabel: "Яркость"})
}
```
