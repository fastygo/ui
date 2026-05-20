package progress

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "progress",
		Title:       "Progress",
		Section:     "components",
		Description: "Progress indicator wireframe.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.Box { track + fill }`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Role", Type: "string", Description: "progressbar"},
			{Name: "AriaValuenow", Type: "string", Description: "Current value"},
		},
		Related: []registry.RelatedLink{
			{Label: "Skeleton", Href: "/docs/components/skeleton"},
			{Label: "Slider", Href: "/docs/components/slider"},
		},
	})
}
