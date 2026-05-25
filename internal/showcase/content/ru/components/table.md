---
slug: table
section: components
title: "Table"
description: "Таблица с секциями, строками и ячейками."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Пагинация"
    href: /docs/components/pagination/
  - label: "Форма"
    href: /docs/components/form/
api:
  - name: "TableHeadCell"
    type: "component"
    description: "Заголовок колонки (Scope: col | row)"
  - name: "TableCell"
    type: "component"
    description: "Ячейка данных"
  - name: "ColSpan"
    type: "int"
    description: "Объединение колонок"
---

Таблица с секциями, строками и ячейками.

## Базовая таблица

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Table(ui.TableProps{Class: "w-full text-sm"}) {
		@ui.TableHead(ui.TableSectionProps{}) {
			@ui.TableRow(ui.TableRowProps{}) {
				@ui.TableHeadCell(ui.TableCellProps{Scope: "col"}) { Имя }
				@ui.TableHeadCell(ui.TableCellProps{Scope: "col"}) { Роль }
			}
		}
		@ui.TableBody(ui.TableSectionProps{}) {
			@ui.TableRow(ui.TableRowProps{}) {
				@ui.TableCell(ui.TableCellProps{}) { Алекс }
				@ui.TableCell(ui.TableCellProps{}) { Администратор }
			}
		}
	}
}
```

## ColSpan

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Table(ui.TableProps{}) {
		@ui.TableBody(ui.TableSectionProps{}) {
			@ui.TableRow(ui.TableRowProps{}) {
				@ui.TableCell(ui.TableCellProps{ColSpan: 2}) { Объединяет две колонки }
			}
		}
	}
}
```
