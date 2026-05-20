---
slug: alert
section: components
title: "Alert"
description: "Callout for important messages."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Card"
    href: /docs/components/card/
  - label: "Badge"
    href: /docs/components/badge/
api:
  - name: "Variant"
    type: "string"
    description: "default | destructive"
  - name: "Class"
    type: "string"
    description: "Extra utilities"
---

Callout for important messages.

## Default

{{demo id="alert.default"}}

```templ
@cmp.Alert(cmp.AlertProps{}) { … }
```

## Destructive

{{demo id="alert.destructive"}}

```templ
@cmp.Alert(cmp.AlertProps{Variant: "destructive"}) { … }
```
