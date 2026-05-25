---
slug: introduction
section: getting-started
title: "Introduction"
description: "FastyGo UI — wireframe-галерея компонентов на Go и templ. Страницы показывают живые превью, фрагменты templ для копирования и таблицы свойств — по аналогии с shadcn/ui, без монолитного UI-пакета."
source: github.com/fastygo/ui
package: internal/site
related:
  - label: "Установка"
    href: /docs/installation/
  - label: "Каталог компонентов"
    href: /docs/
---

FastyGo UI — wireframe-галерея компонентов на Go и templ. Страницы показывают живые превью, фрагменты templ для копирования и таблицы свойств — по аналогии с shadcn/ui, без монолитного UI-пакета.

## Wireframe-этап

Сначала структура, семантика и доступность; визуальный бренд и полировка — на следующем этапе.

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2"}) {
		@ui.Title(ui.TitleProps{Order: 2}, "Галерея компонентов")
		@ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "Wireframe-превью и таблицы API.")
	}
}
```
