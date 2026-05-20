---
slug: card
section: components
title: "Card"
description: "Grouped content with header, body, and footer."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Alert"
    href: /docs/components/alert/
  - label: "Dialog"
    href: /docs/components/dialog/
api:
  - name: "Variant"
    type: "string"
    description: "Surface variant"
  - name: "CardHeader"
    type: "component"
    description: "Title area"
  - name: "CardContent"
    type: "component"
    description: "Main body"
---

Grouped content with header, body, and footer.

## Default

{{demo id="card.default"}}

```templ
@cmp.Card(cmp.CardProps{}) { @cmp.CardHeader … }
```

## With footer

{{demo id="card.footer"}}

```templ
@cmp.CardFooter …
```
