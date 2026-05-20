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

{{demo id="breadcrumb.default"}}

```templ
@cmp.Breadcrumb(cmp.BreadcrumbProps{Items: items})
```

## Current page

{{demo id="breadcrumb.current"}}

```go
Last item with Current: true
```
