---
slug: toggle
section: components
title: "Toggle"
description: "Pressable toggle button (wireframe)."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Toggle Group"
    href: /docs/components/toggle-group/
  - label: "Switch"
    href: /docs/components/switch/
api:
  - name: "AriaPressed"
    type: "string"
    description: "true | false"
  - name: "Variant"
    type: "string"
    description: "Button variant"
---

Pressable toggle button (wireframe).

## Default

{{demo id="toggle.default"}}

```templ
@ui.Button(ui.ButtonProps{Variant: "outline", Attrs: templ.Attributes{"aria-pressed": "false"}})
```

## Pressed

{{demo id="toggle.pressed"}}

```go
aria-pressed="true"
```
