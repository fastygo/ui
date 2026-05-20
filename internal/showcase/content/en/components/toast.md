---
slug: toast
section: components
title: "Toast"
description: "Transient notification wireframe."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Alert"
    href: /docs/components/alert/
  - label: "Dialog"
    href: /docs/components/dialog/
api:
  - name: "Class"
    type: "string"
    description: "Card-like surface"
---

Transient notification wireframe.

## Default

Wireframe composition from ui primitives.

{{demo id="toast.default"}}

```templ
@ui.Box { message + action }
```
