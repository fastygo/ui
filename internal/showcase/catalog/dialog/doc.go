package dialog

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "dialog",
		Title:       "Dialog",
		Section:     "components",
		Description: "Modal dialog wireframe (data-ui8kit dialog).",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.Box[data-ui8kit=dialog] { title + actions }`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Attrs", Type: "templ.Attributes", Description: "data-ui8kit=dialog"},
			{Name: "Title", Type: "Title", Description: "aria-labelledby target"},
		},
		Related: []registry.RelatedLink{
			{Label: "Alert Dialog", Href: "/docs/components/alert-dialog"},
			{Label: "Sheet", Href: "/docs/components/sheet"},
		},
	})
}
