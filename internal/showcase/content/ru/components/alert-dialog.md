---
slug: alert-dialog
section: components
title: "Alert Dialog"
description: "Модальное окно, прерывающее поток действий."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Диалог"
    href: /docs/components/dialog/
  - label: "Alert"
    href: /docs/components/alert/
api:
  - name: "Title"
    type: "Title"
    description: "Заголовок предупреждения"
---

Модальное окно, прерывающее поток действий.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Box(ui.BoxProps{Class: "max-w-sm rounded-lg border border-border bg-card p-4"}) {
		@ui.Title(ui.TitleProps{Order: 3, Class: "text-base font-semibold"}, "Вы уверены?")
		@ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "Это действие нельзя отменить.")
		@ui.Group(ui.GroupProps{Class: "mt-4 flex justify-end gap-2"}) {
			@ui.Button(ui.ButtonProps{Variant: "outline"}) {
				Отмена
			}
			@ui.Button(ui.ButtonProps{Variant: "destructive"}) {
				Удалить
			}
		}
	}
}
```
