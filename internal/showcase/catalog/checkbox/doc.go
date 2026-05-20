package checkbox

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "checkbox",
		Title:       "Checkbox",
		Section:     "components",
		Description: "Boolean checkbox input.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@ui.Checkbox(ui.CheckboxProps{Name: "terms"})`, previewDefault),
			showcaseutil.Variant("checked", "Checked", "", `@ui.Checkbox(ui.CheckboxProps{Checked: true})`, previewChecked),
		},
		API: []registry.APIField{
			{Name: "Name", Type: "string", Description: "Form field name"},
			{Name: "Checked", Type: "bool", Description: "Initial checked state"},
			{Name: "Disabled", Type: "bool", Description: "Disables control"},
		},
		Related: []registry.RelatedLink{
			{Label: "Radio", Href: "/docs/components/radio"},
			{Label: "Switch", Href: "/docs/components/switch"},
		},
	})
}
