---
slug: separator
section: components
title: "Separator"
description: "Visual divider between sections."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Stack"
    href: /docs/components/stack/
  - label: "Card"
    href: /docs/components/card/
api:
  - name: "Class"
    type: "string"
    description: "Typically h-px bg-border"
  - name: "Role"
    type: "string"
    description: "separator"
---

Visual divider between sections.

## Default

Wireframe composition from ui primitives.

{{demo id="separator.default"}}

```templ
@ui.Box(ui.BoxProps{Class: "h-px bg-border"})
```
