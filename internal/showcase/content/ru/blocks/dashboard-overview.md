---
slug: dashboard-overview
section: blocks
title: "Dashboard Overview"
description: "Каркас блока — wireframe-секция с placeholder-текстом для будущего извлечения в github.com/fastygo/blocks."
source: internal/ui/blocks
package: internal/ui/blocks
related:
  - label: "Карточка"
    href: /docs/components/card/
  - label: "Стек"
    href: /docs/primitives/stack/
api:
  - name: "Title"
    type: "string"
    description: "Заголовок секции"
  - name: "Body"
    type: "string"
    description: "Вспомогательный текст"
---

Каркас блока — wireframe-секция с placeholder-текстом для будущего извлечения в github.com/fastygo/blocks.

## Wireframe

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-4 max-w-2xl"}) {
		@ui.Title(ui.TitleProps{Order: 2}, "Обзор дашборда")
		@ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground leading-relaxed"}, "Wireframe-метрики и строка сводки для dashboard-блоков.")
		@ui.Button(ui.ButtonProps{Variant: "outline", Size: "sm"}) {
			Действие
		}
	}
}
```

## Компактный

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2 max-w-2xl"}) {
		@ui.Title(ui.TitleProps{Order: 2}, "Обзор дашборда")
		@ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground leading-relaxed"}, "Wireframe-метрики и строка сводки для dashboard-блоков.")
		@ui.Button(ui.ButtonProps{Variant: "outline", Size: "sm"}) {
			Действие
		}
	}
}
```
