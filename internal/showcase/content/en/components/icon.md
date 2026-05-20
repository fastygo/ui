---
slug: icon
section: components
title: "Icon"
description: "Latty icon mask (app registry)."
source: github.com/fastygo/ui/internal/ui/components/icon
package: github.com/fastygo/ui/internal/ui/components/icon
related:
  - label: "Button"
    href: /docs/components/button/
  - label: "Badge"
    href: /docs/components/badge/
api:
  - name: "Name"
    type: "string"
    description: "Latty icon name"
  - name: "Size"
    type: "string"
    description: "xs | sm | md | lg"
  - name: "Class"
    type: "string"
    description: "Extra utilities"
---

Latty icon mask (app registry).

## Default

{{demo id="icon.default"}}

```templ
@icon.Icon(icon.IconProps{Name: "home"})
```

## Sizes

{{demo id="icon.sizes"}}

```go
Size: xs | sm | md | lg
```
