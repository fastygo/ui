package input

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "input",
		Title:       "Input",
		Section:     "components",
		Description: "Single-line text input control.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@ui.Input(ui.InputProps{Placeholder: "Email"})`, previewDefault),
			showcaseutil.Variant("disabled", "Disabled", "", `@ui.Input(ui.InputProps{Disabled: true})`, previewDisabled),
			showcaseutil.Variant("file", "File", "", `@ui.Input(ui.InputProps{Type: "file"})`, previewFile),
		},
		API: []registry.APIField{
			{Name: "Type", Type: "string", Description: "text | email | password | file | range | …"},
			{Name: "Placeholder", Type: "string", Description: "Placeholder text"},
			{Name: "Disabled", Type: "bool", Description: "Disables input"},
			{Name: "Class", Type: "string", Description: "Tailwind utilities"},
		},
		Related: []registry.RelatedLink{
			{Label: "Textarea", Href: "/docs/components/textarea"},
			{Label: "Form", Href: "/docs/components/form"},
		},
	})
}
