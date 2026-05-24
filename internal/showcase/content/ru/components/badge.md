---
slug: badge
section: components
title: "Badge"
description: "Небольшая метка статуса."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Кнопка"
    href: /docs/primitives/button/
  - label: "Alert"
    href: /docs/components/alert/
api:
  - name: "Variant"
    type: "string"
    description: "default | secondary | destructive | outline"
  - name: "Size"
    type: "string"
    description: "default | sm | lg"
  - name: "Class"
    type: "string"
    description: "Дополнительные утилиты"
---

Небольшая метка статуса.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Badge(ui.BadgeProps{}) {
		Badge
	}
}
```

## Вторичный

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Badge(ui.BadgeProps{Variant: "secondary"}) {
		Вторичный
	}
}
```

## Контурный

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Badge(ui.BadgeProps{Variant: "outline"}) {
		Контурный
	}
}
```

## Деструктивный

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Badge(ui.BadgeProps{Variant: "destructive"}) {
		Внимание
	}
}
```
