---
slug: accordion
section: components
title: "Accordion"
description: "Вертикально сложенные раскрывающиеся секции (wireframe; data-ui8kit accordion)."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Разделитель"
    href: /docs/components/separator/
  - label: "Вкладки"
    href: /docs/components/tabs/
api:
  - name: "Attrs"
    type: "templ.Attributes"
    description: "data-ui8kit на корне"
  - name: "Trigger"
    type: "Button"
    description: "data-ui8kit-trigger"
  - name: "Panel"
    type: "Box"
    description: "data-ui8kit-panel"
---

Вертикально сложенные раскрывающиеся секции (wireframe; data-ui8kit accordion).

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2 max-w-md", Attrs: templ.Attributes{"data-ui8kit": "accordion"}}) {
		@ui.Button(ui.ButtonProps{Attrs: templ.Attributes{"data-ui8kit-trigger": "item-1", "aria-expanded": "false"}}) {
			Раздел 1
		}
		@ui.Box(ui.BoxProps{Class: "rounded border border-border p-3 text-sm", Attrs: templ.Attributes{"data-ui8kit-panel": "item-1", "hidden": true}}) {
			@ui.Text(ui.TextProps{}, "Содержимое панели.")
		}
	}
}
```
