package tabs

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "tabs",
		Title:       "Tabs",
		Section:     "components",
		Description: "Tabbed interface (wireframe; data-ui8kit tabs).",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.Stack { tablist + panels }`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Attrs", Type: "templ.Attributes", Description: "data-ui8kit=tabs on root"},
			{Name: "Tab", Type: "Button", Description: "data-ui8kit-tab"},
			{Name: "Panel", Type: "Box", Description: "data-ui8kit-panel"},
		},
		Related: []registry.RelatedLink{
			{Label: "Accordion", Href: "/docs/components/accordion"},
			{Label: "Navigation Menu", Href: "/docs/components/navigation-menu"},
		},
	})
}
