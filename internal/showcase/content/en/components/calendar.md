---
slug: calendar
section: components
title: "Calendar"
description: "Date grid placeholder."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Command"
    href: /docs/components/command/
  - label: "Popover"
    href: /docs/components/popover/
api:
  - name: "Class"
    type: "string"
    description: "Calendar frame utilities"
---

Date grid placeholder.

## Default

Wireframe composition from ui primitives.

{{demo id="calendar.default"}}

```templ
@ui.Box { month label + day grid }
```
