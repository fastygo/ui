---
slug: grid
section: components
title: "Grid"
description: "CSS grid layout with columns."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Stack"
    href: /docs/components/stack/
  - label: "Container"
    href: /docs/components/container/
api:
  - name: "Class"
    type: "string"
    description: "grid-cols-* and gap utilities"
  - name: "GridCol"
    type: "component"
    description: "Column cell wrapper"
---

CSS grid layout with columns.

## Two columns

{{demo id="grid.default"}}

```templ
@ui.Grid(ui.GridProps{Class: "grid-cols-2 gap-4"}) { … }
```

## Three columns

{{demo id="grid.three"}}

```templ
@ui.Grid(ui.GridProps{Class: "grid-cols-3 gap-2"}) { … }
```
