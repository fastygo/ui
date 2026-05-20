---
slug: dialog
section: components
title: "Dialog"
description: "Modal dialog wireframe (data-ui8kit dialog)."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Alert Dialog"
    href: /docs/components/alert-dialog/
  - label: "Sheet"
    href: /docs/components/sheet/
api:
  - name: "Attrs"
    type: "templ.Attributes"
    description: "data-ui8kit=dialog"
  - name: "Title"
    type: "Title"
    description: "aria-labelledby target"
---

Modal dialog wireframe (data-ui8kit dialog).

## Default

Wireframe composition from ui primitives.

{{demo id="dialog.default"}}

```templ
@ui.Box[data-ui8kit=dialog] { title + actions }
```
