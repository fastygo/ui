---
slug: grid
section: primitives
title: "Сетка"
description: "CSS Grid контейнер с ячейками GridCol."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Стек"
    href: /docs/primitives/stack/
  - label: "Группа"
    href: /docs/primitives/group/
  - label: "Box"
    href: /docs/primitives/box/
api:
  - name: "Class"
    type: "string"
    description: "Utility-классы сетки (grid-cols-*, gap-*)"
  - name: "GridCol"
    type: "component"
    description: "Одна ячейка колонки внутри Grid"
---

Grid размещает дочерние элементы в CSS Grid контейнере. GridCol оборачивает одну ячейку колонки.

## Две колонки

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Grid(ui.GridProps{Class: "grid-cols-2 gap-4"}) {
		@ui.GridCol(ui.GridColProps{}) {
			@ui.Text(ui.TextProps{}, "Колонка A")
		}
		@ui.GridCol(ui.GridColProps{}) {
			@ui.Text(ui.TextProps{}, "Колонка B")
		}
	}
}
```

## Три колонки

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Grid(ui.GridProps{Class: "grid-cols-3 gap-4"}) {
		@ui.GridCol(ui.GridColProps{}) { A }
		@ui.GridCol(ui.GridColProps{}) { B }
		@ui.GridCol(ui.GridColProps{}) { C }
	}
}
```
