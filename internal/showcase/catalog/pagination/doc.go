package pagination

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "pagination",
		Title:       "Pagination",
		Section:     "components",
		Description: "Page navigation controls.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.Group { prev / numbers / next }`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Buttons", Type: "Button", Description: "Prev / page / Next"},
		},
		Related: []registry.RelatedLink{
			{Label: "Data Table", Href: "/docs/components/data-table"},
			{Label: "Breadcrumb", Href: "/docs/components/breadcrumb"},
		},
	})
}
