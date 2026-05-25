---
slug: utils
section: utils
title: "Utils"
description: "Общие Go-хелперы для registry-кирпичей: классы, атрибуты, ARIA и рецепты вариантов."
source: github.com/fastygo/templ/utils
package: github.com/fastygo/templ/utils
related:
  - label: "Теги"
    href: /docs/utils/tags/
  - label: "Кнопка"
    href: /docs/primitives/button/
  - label: "Input"
    href: /docs/primitives/input/
  - label: "Карточка"
    href: /docs/components/card/
  - label: "Alert"
    href: /docs/components/alert/
api:
  - name: "Cn"
    type: "function"
    description: "Объединяет строки CSS-классов"
  - name: "Compose"
    type: "function"
    description: "Собирает классы из Variants и map выбора ключей"
  - name: "Variants"
    type: "struct"
    description: "Рецепт cva-стиля: Base и ByKey"
  - name: "MergeAttrs"
    type: "function"
    description: "Сливает templ.Attributes без дублирующих ключей"
  - name: "DOMAttrs"
    type: "function"
    description: "id, role, tabindex для DOM-узла"
  - name: "ControlAttrs"
    type: "function"
    description: "Атрибуты form-control"
  - name: "AriaLabel"
    type: "function"
    description: "Статический aria-label"
  - name: "InputChrome"
    type: "recipe"
    description: "Варианты поверхности Input, Textarea, Select"
  - name: "ControlChrome"
    type: "recipe"
    description: "Варианты поверхности Checkbox, Radio, Switch"
  - name: "CardVariants"
    type: "recipe"
    description: "Варианты Card"
  - name: "AlertVariants"
    type: "recipe"
    description: "Варианты Alert"
  - name: "TitleTag"
    type: "function"
    description: "Разрешает уровень заголовка h1–h6"
---

Общие Go-хелперы для registry-кирпичей: классы, атрибуты, ARIA и рецепты вариантов.

## Обзор

Пакет utils содержит общие Go-хелперы для registry-кирпичей.
Utils предоставляет рецепты классов, слияние атрибутов и ARIA-хелперы.

## Сценарии использования

- Скопировать utils один раз в модуль consumer-приложения
- Собирать variant-классы в кирпичах button и input
- Сливать templ-атрибуты без дублирующих ключей

## Семантика

- Имя пакета в consumer-коде — uiutils
- Структура Variants следует cva-стилю: Base плюс карты ByKey
- Рецепты объявлены как vars; кирпичи вызывают Compose или обёртки-хелперы

## Реестр allow-list-source

Значения `api.*.allow-list-source` в spec кирпичей разрешаются здесь:

| Идентификатор | Расположение в коде |
|---------------|---------------------|
| `utils.tags.TagGroupLayout` | `utils/tags.go` — Block, Box |
| `utils.tags.TagGroupStack` | `utils/tags.go` — Stack |
| `utils.tags.TagGroupGroup` | `utils/tags.go` — Group |
| `utils.tags.TagGroupContainer` | `utils/tags.go` — Container |
| `utils.tags.TagGroupList` | `utils/tags.go` — List |
| `utils.tags.TagGroupText` | `utils/tags.go` — Text |
| `utils.recipes.InputChrome` | `utils/utils.go` — InputChrome.ByKey.variant |
| `utils.recipes.ControlChrome` | `utils/utils.go` — ControlChrome.ByKey.variant |
| `utils.recipes.CardVariants` | `utils/utils.go` — CardVariants.ByKey.variant |
| `utils.recipes.AlertVariants` | `utils/utils.go` — AlertVariants.ByKey.variant |
| `utils.helpers.TitleTag` | `utils/utils.go` — функция TitleTag |
| `ui.button.ButtonVariants` | `ui/button/button.templ` — локальный Variants кирпича |
| `ui.badge.BadgeVariants` | `ui/badge/badge.templ` — локальный Variants кирпича |

## Пример compose.variant

```go
import uiutils "github.com/fastygo/templ/utils"

classes := uiutils.Compose(uiutils.InputChrome, map[string]string{
	"variant": "outline",
}, "mt-2")
```

## Пример merge.attrs

```go
attrs := uiutils.MergeAttrs(
	uiutils.AriaLabel("Search"),
	templ.Attributes{"name": "q"},
)
```
