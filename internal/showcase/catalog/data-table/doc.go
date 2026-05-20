package datatable

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "data-table",
		Title:       "Data Table",
		Section:     "components",
		Description: "Table with toolbar (wireframe).",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.Table inside @ui.Stack`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Table", Type: "Table", Description: "Semantic table"},
			{Name: "Input", Type: "Input", Description: "Filter field"},
		},
		Related: []registry.RelatedLink{
			{Label: "Table", Href: "/docs/components/table"},
			{Label: "Pagination", Href: "/docs/components/pagination"},
		},
	})
}
