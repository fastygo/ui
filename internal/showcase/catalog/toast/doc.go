package toast

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "toast",
		Title:       "Toast",
		Section:     "components",
		Description: "Transient notification wireframe.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.Box { message + action }`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Class", Type: "string", Description: "Card-like surface"},
		},
		Related: []registry.RelatedLink{
			{Label: "Alert", Href: "/docs/components/alert"},
			{Label: "Dialog", Href: "/docs/components/dialog"},
		},
	})
}
