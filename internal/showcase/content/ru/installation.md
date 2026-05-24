---
slug: installation
section: getting-started
title: "Установка"
description: "Используйте этот репозиторий как эталонное приложение или копируйте отдельные компоненты из галереи в свой проект (в будущем — workflow fastygo add)."
source: github.com/fastygo/ui
package: cmd/server
related:
  - label: "Темизация"
    href: /docs/theming/
  - label: "Кнопка"
    href: /docs/primitives/button/
---

Используйте этот репозиторий как эталонное приложение или копируйте отдельные компоненты из галереи в свой проект (в будущем — workflow fastygo add).

## Из git

Клонируйте приложение, установите Bun для CSS, сгенерируйте templ и запустите сервер.

```go
bun install
go mod download
bun run build:css
go tool templ generate ./...
bun run go
```
