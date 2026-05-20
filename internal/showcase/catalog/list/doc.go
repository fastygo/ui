package list

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "list",
		Title:       "List",
		Section:     "components",
		Description: "Semantic ul/ol/dl list containers.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("unordered", "Unordered", "", `@ui.List(ui.ListProps{Class: "list-disc pl-5"}) { … }`, previewUnordered),
			showcaseutil.Variant("ordered", "Ordered", "", `@ui.List(ui.ListProps{Tag: "ol", Class: "list-decimal pl-5"}) { … }`, previewOrdered),
		},
		API: []registry.APIField{
			{Name: "Tag", Type: "string", Description: "ul | ol | dl | menu"},
			{Name: "Class", Type: "string", Description: "Tailwind utilities"},
			{Name: "Attrs", Type: "templ.Attributes", Description: "Extra attributes"},
		},
		Related: []registry.RelatedLink{
			{Label: "Table", Href: "/docs/components/table"},
			{Label: "Breadcrumb", Href: "/docs/components/breadcrumb"},
		},
	})
}
