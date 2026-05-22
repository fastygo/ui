---
slug: card
section: components
title: "Карточка"
description: "Сгруппированный контент с header, body и footer."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Alert"
    href: /docs/components/alert/
  - label: "Диалог"
    href: /docs/components/dialog/
api:
  - name: "Variant"
    type: "string"
    description: "Вариант поверхности"
  - name: "CardHeader"
    type: "component"
    description: "Область заголовка"
  - name: "CardContent"
    type: "component"
    description: "Основное содержимое"
---

Сгруппированный контент с header, body и footer.

## По умолчанию

```templ
import cmp "github.com/fastygo/templ/components"
import "github.com/fastygo/templ/ui"

templ Example() {
	@cmp.Card(cmp.CardProps{Class: "max-w-sm"}) {
		@cmp.CardHeader(cmp.CardHeaderProps{}) {
			@cmp.CardTitle(cmp.CardTitleProps{}, "Заголовок карточки")
			@cmp.CardDescription(cmp.CardDescriptionProps{}, "Описание карточки в wireframe.")
		}
		@cmp.CardContent(cmp.CardContentProps{}) {
			@ui.Text(ui.TextProps{}, "Текст в теле карточки.")
		}
	}
}
```

## С footer

```templ
import cmp "github.com/fastygo/templ/components"
import "github.com/fastygo/templ/ui"

templ Example() {
	@cmp.Card(cmp.CardProps{Class: "max-w-sm"}) {
		@cmp.CardHeader(cmp.CardHeaderProps{}) {
			@cmp.CardTitle(cmp.CardTitleProps{}, "Заголовок карточки")
			@cmp.CardDescription(cmp.CardDescriptionProps{}, "Описание карточки в wireframe.")
		}
		@cmp.CardContent(cmp.CardContentProps{}) {
			@ui.Text(ui.TextProps{}, "Текст в теле карточки.")
		}
		@cmp.CardFooter(cmp.CardFooterProps{}) {
			@ui.Button(ui.ButtonProps{Size: "sm"}) {
				Действие
			}
		}
	}
}
```
