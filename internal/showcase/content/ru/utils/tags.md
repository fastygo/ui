---
slug: tags
section: utils
title: "Теги"
description: "Ограничение имён HTML-тегов по группам примитивов и безопасный выбор через ResolveTag."
source: github.com/fastygo/templ/utils
package: github.com/fastygo/templ/utils
related:
  - label: "Утилиты"
    href: /docs/utils/utils/
  - label: "Block"
    href: /docs/primitives/block/
  - label: "Box"
    href: /docs/primitives/box/
  - label: "Text"
    href: /docs/primitives/text/
  - label: "Title"
    href: /docs/primitives/title/
api:
  - name: "TagGroup"
    type: "тип"
    description: "Набор разрешённых семантических тегов"
  - name: "ResolveTag"
    type: "function"
    description: "Возвращает tag или fallback, если tag не разрешён"
  - name: "IsAllowedTag"
    type: "function"
    description: "Проверяет, входит ли tag в группу"
  - name: "TagGroupLayout"
    type: "константа"
    description: "div, section, article, aside, header, footer, main, nav, figure — Block, Box"
  - name: "TagGroupStack"
    type: "константа"
    description: "layout-теги, ul, ol — Stack"
  - name: "TagGroupGroup"
    type: "константа"
    description: "layout-теги, fieldset — Group"
  - name: "TagGroupList"
    type: "константа"
    description: "ul, ol, dl, menu — List"
  - name: "TagGroupText"
    type: "константа"
    description: "p, span, small, code, blockquote, em, strong, time, cite, abbr, mark, address, pre, figcaption — Text"
  - name: "TagGroupContainer"
    type: "константа"
    description: "div, main, section — Container"
  - name: "TagGroupHeading"
    type: "константа"
    description: "h1–h6 — Title"
---

Ограничение имён HTML-тегов по группам примитивов и безопасный выбор через ResolveTag.

## Обзор

Модуль tags ограничивает имена HTML-тегов для каждой группы примитивов.
ResolveTag возвращает fallback, если запрошенный tag не разрешён.

## Сценарии использования

- Block и Box безопасно выбирают landmark-теги
- Text и Title безопасно выбирают семантические prose-теги

## Семантика

- TagGroup — перечисление int в том же Go-пакете
- ResolveTag приводит запрошенный tag к нижнему регистру и обрезает пробелы
- Недопустимые tag никогда не попадают в разметку; fallback всегда побеждает

## Пример resolve.layout

```go
tag := uiutils.ResolveTag("main", "div", uiutils.TagGroupLayout)
// tag == "main"

tag := uiutils.ResolveTag("table", "div", uiutils.TagGroupLayout)
// tag == "div"
```
