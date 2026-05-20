package docsarticle

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "docs-article",
		Title:       "Docs Article",
		Section:     "blocks",
		Description: "Block scaffold — section wireframe with placeholder copy for future github.com/fastygo/blocks extraction.",
		Source:      "internal/ui/blocks",
		Package:     "internal/ui/blocks",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Wireframe", "", `@ui.Stack { @ui.Title … "Docs Article" }`, previewDefault),
			showcaseutil.Variant("compact", "Compact", "", `Denser spacing variant`, previewCompact),
		},
		API: []registry.APIField{
			{Name: "Title", Type: "string", Description: "Section heading"},
			{Name: "Body", Type: "string", Description: "Supporting copy"},
		},
		Related: []registry.RelatedLink{
			{Label: "Card", Href: "/docs/components/card"},
			{Label: "Stack", Href: "/docs/components/stack"},
		},
	})
}
