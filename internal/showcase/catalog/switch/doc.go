package formswitch

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "switch",
		Title:       "Switch",
		Section:     "components",
		Description: "Toggle switch (formswitch / ui.Switch).",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@ui.Switch(ui.SwitchProps{Name: "airplane"})`, previewDefault),
			showcaseutil.Variant("checked", "On", "", `@ui.Switch(ui.SwitchProps{Checked: true})`, previewChecked),
		},
		API: []registry.APIField{
			{Name: "Name", Type: "string", Description: "Form field name"},
			{Name: "Checked", Type: "bool", Description: "On state"},
			{Name: "AriaLabel", Type: "string", Description: "Accessible name"},
		},
		Related: []registry.RelatedLink{
			{Label: "Checkbox", Href: "/docs/components/checkbox"},
			{Label: "Toggle", Href: "/docs/components/toggle"},
		},
	})
}
