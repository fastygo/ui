package form

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "form",
		Title:       "Form",
		Section:     "components",
		Description: "Form landmark with item helpers.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Login", "", `@ui.Form(ui.FormProps{}) { @ui.FormItem … }`, previewDefault),
			showcaseutil.Variant("inline", "Inline", "", `Compact horizontal FormItem layout`, previewInline),
		},
		API: []registry.APIField{
			{Name: "Action", Type: "string", Description: "Form action URL"},
			{Name: "Method", Type: "string", Description: "GET | POST"},
			{Name: "FormItem", Type: "component", Description: "Label + control group"},
		},
		Related: []registry.RelatedLink{
			{Label: "Input", Href: "/docs/components/input"},
			{Label: "Button", Href: "/docs/components/button"},
		},
	})
}
