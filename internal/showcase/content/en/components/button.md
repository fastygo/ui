---
slug: button
section: components
title: "Button"
description: "Triggers an action or navigates when rendered as a link. Built on github.com/fastygo/templ/ui."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Toggle"
    href: /docs/components/toggle/
  - label: "Form"
    href: /docs/components/form/
api:
  - name: "Variant"
    type: "string"
    description: "default | secondary | destructive | outline | ghost | link | unstyled"
  - name: "Size"
    type: "string"
    description: "default | sm | lg | icon"
  - name: "Type"
    type: "string"
    description: "button | submit | reset"
  - name: "Href"
    type: "string"
    description: "When set, renders an anchor instead of button"
  - name: "Disabled"
    type: "bool"
    description: "Disables interaction"
  - name: "Class"
    type: "string"
    description: "Additional Tailwind utilities"
  - name: "AriaLabel"
    type: "string"
    description: "Accessible name when visible text is insufficient"
---

Triggers an action or navigates when rendered as a link. Built on github.com/fastygo/templ/ui.

## Default

{{demo id="button.default"}}

```templ
@showcaseutil.Button(ui.ButtonProps{}, "Button").Render(ctx, w)
```

## Secondary

{{demo id="button.secondary"}}

```templ
@showcaseutil.Button(ui.ButtonProps{Variant: "secondary"}, "Secondary").Render(ctx, w)
```

## Outline

{{demo id="button.outline"}}

```templ
@showcaseutil.Button(ui.ButtonProps{Variant: "outline"}, "Outline").Render(ctx, w)
```

## Destructive

{{demo id="button.destructive"}}

```templ
@showcaseutil.Button(ui.ButtonProps{Variant: "destructive"}, "Destructive").Render(ctx, w)
```

## Ghost

{{demo id="button.ghost"}}

```templ
@showcaseutil.Button(ui.ButtonProps{Variant: "ghost"}, "Ghost").Render(ctx, w)
```

## Link

{{demo id="button.link"}}

```templ
@showcaseutil.Button(ui.ButtonProps{Variant: "link"}, "Link").Render(ctx, w)
```

## Sizes

{{demo id="button.sizes"}}

```go
sm / default / lg via ButtonProps.Size
```
