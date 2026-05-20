---
slug: badge
section: components
title: "Badge"
description: "Small status label chip."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Button"
    href: /docs/components/button/
  - label: "Alert"
    href: /docs/components/alert/
api:
  - name: "Variant"
    type: "string"
    description: "default | secondary | destructive | outline"
  - name: "Size"
    type: "string"
    description: "default | sm | lg"
  - name: "Class"
    type: "string"
    description: "Extra utilities"
---

Small status label chip.

## Default

{{demo id="badge.default"}}

```templ
@showcaseutil.RenderBadge(ctx, w, ui.BadgeProps{}, "Badge")
```

## Secondary

{{demo id="badge.secondary"}}

```templ
@showcaseutil.RenderBadge(ctx, w, ui.BadgeProps{Variant: "secondary"}, "Secondary")
```

## Outline

{{demo id="badge.outline"}}

```templ
@showcaseutil.RenderBadge(ctx, w, ui.BadgeProps{Variant: "outline"}, "Outline")
```

## Destructive

{{demo id="badge.destructive"}}

```templ
@showcaseutil.RenderBadge(ctx, w, ui.BadgeProps{Variant: "destructive"}, "Alert")
```
