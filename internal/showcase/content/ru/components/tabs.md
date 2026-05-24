---
slug: tabs
section: components
title: "Вкладки"
description: "Интерфейс с вкладками (wireframe; data-ui8kit tabs)."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Аккордеон"
    href: /docs/components/accordion/
  - label: "Меню навигации"
    href: /docs/components/navigation-menu/
api:
  - name: "Attrs"
    type: "templ.Attributes"
    description: "data-ui8kit=tabs на корне"
  - name: "Tab"
    type: "Button"
    description: "data-ui8kit-tab"
  - name: "Panel"
    type: "Box"
    description: "data-ui8kit-panel"
---

Интерфейс с вкладками (wireframe; data-ui8kit tabs).

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2 max-w-md", Attrs: templ.Attributes{"data-ui8kit": "tabs"}}) {
		@ui.Group(ui.GroupProps{Class: "flex gap-1", Attrs: templ.Attributes{"role": "tablist"}}) {
			@ui.Button(ui.ButtonProps{Variant: "secondary", Size: "sm", Attrs: templ.Attributes{"data-ui8kit-tab": "a", "aria-selected": "true"}}) {
				Вкладка A
			}
			@ui.Button(ui.ButtonProps{Variant: "ghost", Size: "sm", Attrs: templ.Attributes{"data-ui8kit-tab": "b"}}) {
				Вкладка B
			}
		}
		@ui.Box(ui.BoxProps{Class: "rounded border border-border p-3 text-sm", Attrs: templ.Attributes{"data-ui8kit-panel": "a"}}) {
			@ui.Text(ui.TextProps{}, "Панель вкладки A.")
		}
	}
}
```
