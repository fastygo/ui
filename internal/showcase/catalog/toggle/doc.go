package toggle

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "toggle",
		Title:       "Toggle",
		Section:     "components",
		Description: "Pressable toggle button (wireframe).",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@ui.Button(ui.ButtonProps{Variant: "outline", Attrs: templ.Attributes{"aria-pressed": "false"}})`, previewDefault),
			showcaseutil.Variant("pressed", "Pressed", "", `aria-pressed="true"`, previewPressed),
		},
		API: []registry.APIField{
			{Name: "AriaPressed", Type: "string", Description: "true | false"},
			{Name: "Variant", Type: "string", Description: "Button variant"},
		},
		Related: []registry.RelatedLink{
			{Label: "Toggle Group", Href: "/docs/components/toggle-group"},
			{Label: "Switch", Href: "/docs/components/switch"},
		},
	})
}
