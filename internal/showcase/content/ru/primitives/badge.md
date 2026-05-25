---
slug: badge
section: primitives
title: "Badge"
description: "Короткая метка статуса с вариантами и размерами."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Текст"
    href: /docs/primitives/text/
  - label: "Кнопка"
    href: /docs/primitives/button/
api:
  - name: "Variant"
    type: "string"
    description: "default | secondary | destructive | outline"
  - name: "Size"
    type: "string"
    description: "default | sm | lg"
  - name: "Class"
    type: "string"
    description: "Дополнительные utility-классы"
---

Бейдж показывает короткую метку статуса. Используйте варианты и размеры для визуальной иерархии.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Badge(ui.BadgeProps{Variant: "default"}) {
		По умолчанию
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

## Деструктивный

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Badge(ui.BadgeProps{Variant: "destructive"}) {
		Ошибка
	}
}
```

## Малый размер

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Badge(ui.BadgeProps{Size: "sm"}) {
		Малый
	}
}
```
