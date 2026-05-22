---
slug: card
section: components
title: "Card"
description: "Grouped content with header, body, and footer."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Alert"
    href: /docs/components/alert/
  - label: "Dialog"
    href: /docs/components/dialog/
api:
  - name: "Variant"
    type: "string"
    description: "Surface variant"
  - name: "CardHeader"
    type: "component"
    description: "Title area"
  - name: "CardContent"
    type: "component"
    description: "Main body"
---

Grouped content with header, body, and footer.

## Default

```templ
import cmp "github.com/fastygo/templ/components"
import "github.com/fastygo/templ/ui"

templ Example() {
	@cmp.Card(cmp.CardProps{Class: "max-w-sm"}) {
		@cmp.CardHeader(cmp.CardHeaderProps{}) {
			@cmp.CardTitle(cmp.CardTitleProps{}, "Card title")
			@cmp.CardDescription(cmp.CardDescriptionProps{}, "Wireframe card description.")
		}
		@cmp.CardContent(cmp.CardContentProps{}) {
			@ui.Text(ui.TextProps{}, "Card body copy.")
		}
	}
}
```

## With footer

```templ
import cmp "github.com/fastygo/templ/components"
import "github.com/fastygo/templ/ui"

templ Example() {
	@cmp.Card(cmp.CardProps{Class: "max-w-sm"}) {
		@cmp.CardHeader(cmp.CardHeaderProps{}) {
			@cmp.CardTitle(cmp.CardTitleProps{}, "Card title")
			@cmp.CardDescription(cmp.CardDescriptionProps{}, "Wireframe card description.")
		}
		@cmp.CardContent(cmp.CardContentProps{}) {
			@ui.Text(ui.TextProps{}, "Card body copy.")
		}
		@cmp.CardFooter(cmp.CardFooterProps{}) {
			@ui.Button(ui.ButtonProps{Size: "sm"}) {
				Action
			}
		}
	}
}
```
