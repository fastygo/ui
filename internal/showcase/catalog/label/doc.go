package label

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "label",
		Title:       "Label",
		Section:     "components",
		Description: "Accessible label for form controls.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@showcaseutil.RenderLabel(ctx, w, ui.LabelProps{HTMLFor: "email"}, "Email")`, previewDefault),
			showcaseutil.Variant("required", "Required hint", "", `Pair with aria-required on control`, previewRequired),
		},
		API: []registry.APIField{
			{Name: "HTMLFor", Type: "string", Description: "id of associated control"},
			{Name: "Class", Type: "string", Description: "Tailwind utilities"},
		},
		Related: []registry.RelatedLink{
			{Label: "Input", Href: "/docs/components/input"},
			{Label: "Checkbox", Href: "/docs/components/checkbox"},
		},
	})
}
