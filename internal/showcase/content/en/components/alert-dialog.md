---
slug: alert-dialog
section: components
title: "Alert Dialog"
description: "Modal that interrupts flow."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Dialog"
    href: /docs/components/dialog/
  - label: "Alert"
    href: /docs/components/alert/
api:
  - name: "Title"
    type: "Title"
    description: "Alert heading"
---

Modal that interrupts flow.

## Default

Wireframe composition from ui primitives.

{{demo id="alert-dialog.default"}}

```templ
@ui.Box { alert copy + confirm }
```
