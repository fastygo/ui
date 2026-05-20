package tooltip

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "tooltip",
		Title:       "Tooltip",
		Section:     "components",
		Description: "Hint on hover/focus (wireframe).",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.Button + @ui.Box[role=tooltip]`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Role", Type: "string", Description: "tooltip on hint box"},
		},
		Related: []registry.RelatedLink{
			{Label: "Popover", Href: "/docs/components/popover"},
			{Label: "Button", Href: "/docs/components/button"},
		},
	})
}
