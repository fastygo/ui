---
slug: sheet
section: components
title: "Sheet"
description: "Slide-over panel wireframe."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Dialog"
    href: /docs/components/dialog/
  - label: "Drawer"
    href: /docs/components/drawer/
api:
  - name: "Class"
    type: "string"
    description: "Panel surface utilities"
---

Slide-over panel wireframe.

## Default

Wireframe composition from ui primitives.

{{demo id="sheet.default"}}

```templ
@ui.Box { header + body }
```
