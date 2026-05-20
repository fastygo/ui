---
slug: select
section: components
title: "Select"
description: "Native select dropdown (ui.Select / selectfield)."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Combobox"
    href: /docs/components/combobox/
  - label: "Radio"
    href: /docs/components/radio/
api:
  - name: "Options"
    type: "[]ui.Option"
    description: "Value/label pairs"
  - name: "Name"
    type: "string"
    description: "Form field name"
  - name: "Value"
    type: "string"
    description: "Selected value"
---

Native select dropdown (ui.Select / selectfield).

## Default

{{demo id="select.default"}}

```templ
@ui.Select(ui.SelectProps{Name: "role", Options: opts})
```

## Disabled

{{demo id="select.disabled"}}

```templ
@ui.Select(ui.SelectProps{Disabled: true})
```
