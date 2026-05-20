---
slug: "blog-card"
section: "components"
title: "Blog Card"
description: "Переиспользуемые карточки блога с плейсхолдером медиа — вертикальная и горизонтальная раскладки."
source: "github.com/fastygo/ui/internal/ui/components/blogcard"
package: "github.com/fastygo/ui/internal/ui/components/blogcard"
related:
  - label: "Card"
    href: "/docs/components/card/"
  - label: "Stack"
    href: "/docs/components/stack/"
api:
  - name: "Title"
    type: "string"
    description: "Заголовок статьи"
  - name: "Excerpt"
    type: "string"
    description: "Краткое описание"
  - name: "Href"
    type: "string"
    description: "Ссылка «Читать далее»"
  - name: "MediaURL"
    type: "string"
    description: "Зарезервировано для будущего медиа-примитива; в wireframe используется muted-плейсхолдер"
  - name: "MediaAlt"
    type: "string"
    description: "Доступная подпись для плейсхолдера медиа"
  - name: "DateLabel"
    type: "string"
    description: "Дата публикации (уже отформатированная)"
---

Переиспользуемые карточки блога с плейсхолдером медиа.

## Вертикальная

Медиа и текст друг под другом для сеток и лент.

{{demo id="blog-card.vertical"}}

```templ
@blogcard.VerticalBlogCard(blogcard.DefaultVertical())
```

## Горизонтальная

Медиа и текст рядом для плотных списков.

{{demo id="blog-card.horizontal"}}

```templ
@blogcard.HorizontalBlogCard(blogcard.DefaultHorizontal())
```
