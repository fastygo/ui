package popover

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "popover",
		Title:       "Popover",
		Section:     "components",
		Description: "Floating content panel.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.Button + floating @ui.Box`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Class", Type: "string", Description: "Floating panel utilities"},
		},
		Related: []registry.RelatedLink{
			{Label: "Dropdown Menu", Href: "/docs/components/dropdown-menu"},
			{Label: "Tooltip", Href: "/docs/components/tooltip"},
		},
	})
}
