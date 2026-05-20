---
slug: combobox
section: components
title: "Combobox"
description: "Searchable select wireframe (data-ui8kit combobox)."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Select"
    href: /docs/components/select/
  - label: "Command"
    href: /docs/components/command/
api:
  - name: "Input"
    type: "Input"
    description: "Filter field"
  - name: "List"
    type: "List"
    description: "Options listbox"
---

Searchable select wireframe (data-ui8kit combobox).

## Default

{{demo id="combobox.default"}}

```templ
@ui.Input + @ui.List { data-ui8kit=combobox }
```
