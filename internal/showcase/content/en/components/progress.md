---
slug: progress
section: components
title: "Progress"
description: "Progress indicator wireframe."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Skeleton"
    href: /docs/components/skeleton/
  - label: "Slider"
    href: /docs/components/slider/
api:
  - name: "Role"
    type: "string"
    description: "progressbar"
  - name: "AriaValuenow"
    type: "string"
    description: "Current value"
---

Progress indicator wireframe.

## Default

Wireframe composition from ui primitives.

{{demo id="progress.default"}}

```templ
@ui.Box { track + fill }
```
