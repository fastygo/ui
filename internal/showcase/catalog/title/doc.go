package title

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "title",
		Title:       "Title",
		Section:     "components",
		Description: "Semantic heading with order 1–6.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("h1", "Heading 1", "", `@ui.Title(ui.TitleProps{Order: 1}, "Page title")`, previewH1),
			showcaseutil.Variant("h2", "Heading 2", "", `@ui.Title(ui.TitleProps{Order: 2}, "Section")`, previewH2),
			showcaseutil.Variant("h3", "Heading 3", "", `@ui.Title(ui.TitleProps{Order: 3}, "Subsection")`, previewH3),
		},
		API: []registry.APIField{
			{Name: "Order", Type: "int", Description: "1–6 maps to h1–h6"},
			{Name: "Class", Type: "string", Description: "Tailwind utilities"},
		},
		Related: []registry.RelatedLink{
			{Label: "Text", Href: "/docs/components/text"},
			{Label: "Stack", Href: "/docs/components/stack"},
		},
	})
}
