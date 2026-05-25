---
slug: dialog
section: components
title: "Dialog"
description: "Wireframe модального диалога (data-ui8kit dialog)."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Alert Dialog"
    href: /docs/components/alert-dialog/
  - label: "Drawer"
    href: /docs/components/drawer/
api:
  - name: "Attrs"
    type: "templ.Attributes"
    description: "data-ui8kit=dialog"
  - name: "Title"
    type: "Title"
    description: "Цель для aria-labelledby"
---

Wireframe модального диалога (data-ui8kit dialog).

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Box(ui.BoxProps{Class: "max-w-sm rounded-lg border border-border bg-card p-4 shadow-lg", Attrs: templ.Attributes{"data-ui8kit": "dialog"}}) {
		@ui.Title(ui.TitleProps{Order: 3, Class: "text-base font-semibold"}, "Заголовок диалога")
		@ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "Описание диалога.")
		@ui.Group(ui.GroupProps{Class: "mt-4 flex justify-end gap-2"}) {
			@ui.Button(ui.ButtonProps{Variant: "outline", Attrs: templ.Attributes{"data-ui8kit-close": ""}}) {
				Отмена
			}
			@ui.Button(ui.ButtonProps{Attrs: templ.Attributes{"data-ui8kit-close": ""}}) {
				Подтвердить
			}
		}
	}
}
```
