package stack

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug: "stack", Title: "Stack", Section: "components",
		Description: "Vertical flex column for stacking children with gap utilities.",
		Source:      "github.com/fastygo/templ/ui", Package: "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@ui.Stack(ui.StackProps{Class: "gap-2"}) { … }`, previewDefault),
			showcaseutil.Variant("horizontal", "Row", "", `@ui.Stack(ui.StackProps{Class: "flex-row gap-2"}) { … }`, previewHorizontal),
			showcaseutil.Variant("nav", "Nav tag", "", `@ui.Stack(ui.StackProps{Tag: "nav"}) { … }`, previewNav),
		},
		API: []registry.APIField{
			{Name: "Class", Type: "string", Description: "Tailwind utilities including gap"},
			{Name: "Tag", Type: "string", Description: "div | nav | section | ul | …"},
			{Name: "Attrs", Type: "templ.Attributes", Description: "Extra HTML attributes"},
		},
		Related: []registry.RelatedLink{
			{Label: "Group", Href: "/docs/components/group"},
			{Label: "Box", Href: "/docs/components/box"},
		},
	})
}
