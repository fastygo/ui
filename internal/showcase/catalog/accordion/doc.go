package accordion

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "accordion",
		Title:       "Accordion",
		Section:     "components",
		Description: "Vertically stacked expandable sections (wireframe; data-ui8kit accordion).",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.Stack(templ.Attributes{"data-ui8kit": "accordion"}) { @ui.Button[data-ui8kit-trigger] … }`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Attrs", Type: "templ.Attributes", Description: "data-ui8kit on root"},
			{Name: "Trigger", Type: "Button", Description: "data-ui8kit-trigger"},
			{Name: "Panel", Type: "Box", Description: "data-ui8kit-panel"},
		},
		Related: []registry.RelatedLink{
			{Label: "Collapsible", Href: "/docs/components/collapsible"},
			{Label: "Tabs", Href: "/docs/components/tabs"},
		},
	})
}
