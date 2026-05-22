---
slug: icon
section: components
title: "Icon"
description: "Latty icon mask (app registry)."
source: github.com/fastygo/ui/internal/ui/components/icon
package: github.com/fastygo/ui/internal/ui/components/icon
related:
  - label: "Button"
    href: /docs/components/button/
  - label: "Badge"
    href: /docs/components/badge/
api:
  - name: "Name"
    type: "string"
    description: "Latty icon name"
  - name: "Size"
    type: "string"
    description: "xs | sm | md | lg"
  - name: "Class"
    type: "string"
    description: "Extra utilities"
---

Latty icon mask (app registry).

## Default

```templ
import appicon "github.com/fastygo/ui/internal/ui/components/icon"

templ Example() {
	@appicon.Icon(appicon.IconProps{Name: "home"})
}
```

## Sizes

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
