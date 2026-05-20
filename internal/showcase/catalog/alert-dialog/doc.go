package alertdialog

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "alert-dialog",
		Title:       "Alert Dialog",
		Section:     "components",
		Description: "Modal that interrupts flow.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.Box { alert copy + confirm }`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Title", Type: "Title", Description: "Alert heading"},
		},
		Related: []registry.RelatedLink{
			{Label: "Dialog", Href: "/docs/components/dialog"},
			{Label: "Alert", Href: "/docs/components/alert"},
		},
	})
}
