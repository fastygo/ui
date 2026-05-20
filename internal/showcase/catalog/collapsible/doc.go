package collapsible

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "collapsible",
		Title:       "Collapsible",
		Section:     "components",
		Description: "Single expand/collapse region.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@ui.Button + @ui.Box { data-ui8kit=disclosure }`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Trigger", Type: "Button", Description: "Expands panel"},
			{Name: "Panel", Type: "Box", Description: "Collapsible content"},
		},
		Related: []registry.RelatedLink{
			{Label: "Accordion", Href: "/docs/components/accordion"},
			{Label: "Sheet", Href: "/docs/components/sheet"},
		},
	})
}
