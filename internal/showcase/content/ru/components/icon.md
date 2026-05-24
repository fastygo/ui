---
slug: icon
section: components
title: "Иконка"
description: "Latty icon mask (реестр приложения)."
source: github.com/fastygo/ui/internal/ui/components/icon
package: github.com/fastygo/ui/internal/ui/components/icon
related:
  - label: "Кнопка"
    href: /docs/primitives/button/
  - label: "Badge"
    href: /docs/primitives/badge/
api:
  - name: "Name"
    type: "string"
    description: "Имя иконки Latty"
  - name: "Size"
    type: "string"
    description: "xs | sm | md | lg"
  - name: "Class"
    type: "string"
    description: "Дополнительные утилиты"
---

Latty icon mask (реестр приложения).

## По умолчанию

```templ
import appicon "github.com/fastygo/ui/internal/ui/components/icon"

templ Example() {
	@appicon.Icon(appicon.IconProps{Name: "home"})
}
```

## Размеры

```templ
import appicon "github.com/fastygo/ui/internal/ui/components/icon"
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{}) {
		@appicon.Icon(appicon.IconProps{Name: "settings", Size: "xs", Class: "mr-2"})
		@appicon.Icon(appicon.IconProps{Name: "settings", Size: "sm", Class: "mr-2"})
		@appicon.Icon(appicon.IconProps{Name: "settings", Size: "md", Class: "mr-2"})
		@appicon.Icon(appicon.IconProps{Name: "settings", Size: "lg", Class: "mr-2"})
	}
}
```
