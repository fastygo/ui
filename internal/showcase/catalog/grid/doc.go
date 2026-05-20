package grid

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "grid",
		Title:       "Grid",
		Section:     "components",
		Description: "CSS grid layout with columns.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Two columns", "", `@ui.Grid(ui.GridProps{Class: "grid-cols-2 gap-4"}) { … }`, previewDefault),
			showcaseutil.Variant("three", "Three columns", "", `@ui.Grid(ui.GridProps{Class: "grid-cols-3 gap-2"}) { … }`, previewThree),
		},
		API: []registry.APIField{
			{Name: "Class", Type: "string", Description: "grid-cols-* and gap utilities"},
			{Name: "GridCol", Type: "component", Description: "Column cell wrapper"},
		},
		Related: []registry.RelatedLink{
			{Label: "Stack", Href: "/docs/components/stack"},
			{Label: "Container", Href: "/docs/components/container"},
		},
	})
}
