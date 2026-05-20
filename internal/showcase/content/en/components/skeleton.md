---
slug: skeleton
section: components
title: "Skeleton"
description: "Loading placeholder blocks."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Progress"
    href: /docs/components/progress/
  - label: "Card"
    href: /docs/components/card/
api:
  - name: "Class"
    type: "string"
    description: "animate-pulse bg-muted shapes"
---

Loading placeholder blocks.

## Default

Wireframe composition from ui primitives.

{{demo id="skeleton.default"}}

```templ
@ui.Box(ui.BoxProps{Class: "animate-pulse bg-muted"})
```
