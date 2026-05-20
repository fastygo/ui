package textarea

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "textarea",
		Title:       "Textarea",
		Section:     "components",
		Description: "Multi-line text input.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@ui.Textarea(ui.TextareaProps{Placeholder: "Message"})`, previewDefault),
			showcaseutil.Variant("disabled", "Disabled", "", `@ui.Textarea(ui.TextareaProps{Disabled: true})`, previewDisabled),
		},
		API: []registry.APIField{
			{Name: "Placeholder", Type: "string", Description: "Placeholder text"},
			{Name: "Rows", Type: "int", Description: "Visible row count"},
			{Name: "Disabled", Type: "bool", Description: "Disables control"},
		},
		Related: []registry.RelatedLink{
			{Label: "Input", Href: "/docs/components/input"},
			{Label: "Form", Href: "/docs/components/form"},
		},
	})
}
