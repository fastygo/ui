---
slug: "blog-card"
section: "components"
title: "Карточка блога"
description: "Переиспользуемые карточки блога с плейсхолдером медиа — вертикальная и горизонтальная раскладки."
source: "github.com/fastygo/ui/internal/ui/components/blogcard"
package: "github.com/fastygo/ui/internal/ui/components/blogcard"
related:
  - label: "Карточка"
    href: "/docs/components/card/"
  - label: "Стек"
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

```templ
import "github.com/fastygo/ui/internal/ui/components/blogcard"

templ Example() {
	@blogcard.VerticalBlogCard(blogcard.DefaultVertical())
}
```

## Горизонтальная

Медиа и текст рядом — для плотных списков.

```templ
import "github.com/fastygo/ui/internal/ui/components/blogcard"

templ Example() {
	@blogcard.HorizontalBlogCard(blogcard.DefaultHorizontal())
}
```

