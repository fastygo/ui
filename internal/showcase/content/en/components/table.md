---
slug: table
section: components
title: "Table"
description: "Semantic data table structure."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Data Table"
    href: /docs/components/data-table/
  - label: "List"
    href: /docs/components/list/
api:
  - name: "Class"
    type: "string"
    description: "Table wrapper utilities"
  - name: "TableHead"
    type: "component"
    description: "thead section"
  - name: "TableBody"
    type: "component"
    description: "tbody section"
---

Semantic data table structure.

## Default

{{demo id="table.default"}}

```templ
@ui.Table(ui.TableProps{}) { @ui.TableHead … }
```

## Compact

{{demo id="table.striped"}}

```go
Dense row styling via TableCell Class
```
