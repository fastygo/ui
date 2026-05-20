---
slug: theming
section: getting-started
title: "Theming"
description: "Semantic tokens (background, foreground, primary, …) come from tweakcn-style CSS variables. The header theme toggle switches light/dark via theme.js."
source: web/static/css/tweakcn.css
package: web/static/js/theme.js
related:
  - label: "Introduction"
    href: /docs/introduction/
---

Semantic tokens (background, foreground, primary, …) come from tweakcn-style CSS variables. The header theme toggle switches light/dark via theme.js.

## Semantic tokens

Use token utilities such as bg-background and text-foreground in templ class strings.

{{demo id="theming.tokens"}}

```go
<body class="bg-background text-foreground">
  <!-- theme.js toggles .dark on html -->
</body>
```
