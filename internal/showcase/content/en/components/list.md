---
slug: list
section: components
title: "List"
description: "Semantic ul/ol/dl list containers."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Table"
    href: /docs/components/table/
  - label: "Breadcrumb"
    href: /docs/components/breadcrumb/
api:
  - name: "Tag"
    type: "string"
    description: "ul | ol | dl | menu"
  - name: "Class"
    type: "string"
    description: "Tailwind utilities"
  - name: "Attrs"
    type: "templ.Attributes"
    description: "Extra attributes"
---

Semantic ul/ol/dl list containers.

## Unordered

{{demo id="list.unordered"}}

```templ
@ui.List(ui.ListProps{Class: "list-disc pl-5"}) { … }
```

## Ordered

{{demo id="list.ordered"}}

```templ
@ui.List(ui.ListProps{Tag: "ol", Class: "list-decimal pl-5"}) { … }
```
