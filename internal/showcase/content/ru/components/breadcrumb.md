---
slug: breadcrumb
section: components
title: "Breadcrumb"
description: "Цепочка навигации до текущей страницы."
source: github.com/fastygo/templ/components
package: github.com/fastygo/templ/components
related:
  - label: "Меню навигации"
    href: /docs/components/navigation-menu/
  - label: "Пагинация"
    href: /docs/components/pagination/
api:
  - name: "Items"
    type: "[]BreadcrumbItem"
    description: "Элементы цепочки"
  - name: "Label"
    type: "string"
    description: "Подпись элемента"
  - name: "Href"
    type: "string"
    description: "Ссылка на раздел"
  - name: "Current"
    type: "bool"
    description: "Текущая страница (aria-current)"
  - name: "Disabled"
    type: "bool"
    description: "Недоступный элемент"
---

Цепочка навигации до текущей страницы.

## Стандартная цепочка

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.Breadcrumb(cmp.BreadcrumbProps{
		Items: []cmp.BreadcrumbItem{
			{Label: "Главная", Href: "/"},
			{Label: "Документация", Href: "/docs"},
			{Label: "Кнопка", Current: true},
		},
	})
}
```

## Недоступный элемент

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.Breadcrumb(cmp.BreadcrumbProps{
		Items: []cmp.BreadcrumbItem{
			{Label: "Главная", Href: "/"},
			{Label: "Заблокировано", Disabled: true},
		},
	})
}
```
