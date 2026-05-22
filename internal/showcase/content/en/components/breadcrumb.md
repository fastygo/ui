---
slug: breadcrumb
section: components
title: "Breadcrumb"
description: "Hierarchy navigation trail."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Navigation Menu"
    href: /docs/components/navigation-menu/
  - label: "Tabs"
    href: /docs/components/tabs/
api:
  - name: "Items"
    type: "[]BreadcrumbItem"
    description: "Label, Href, Current, Disabled"
  - name: "Class"
    type: "string"
    description: "Nav wrapper utilities"
---

Hierarchy navigation trail.

## Default

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.Breadcrumb(cmp.BreadcrumbProps{Items: []cmp.BreadcrumbItem{
		{Label: "Home", Href: "/"},
		{Label: "Docs", Href: "/docs"},
		{Label: "Button", Current: true},
	}})
}
```

## Current page

```templ
import cmp "github.com/fastygo/templ/components"

templ Example() {
	@cmp.Breadcrumb(cmp.BreadcrumbProps{Items: []cmp.BreadcrumbItem{
		{Label: "Home", Href: "/"},
		{Label: "Settings", Current: true},
	}})
}
```
