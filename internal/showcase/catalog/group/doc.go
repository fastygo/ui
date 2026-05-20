package group

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "group",
		Title:       "Group",
		Section:     "components",
		Description: "Horizontal flex row for grouping controls.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@ui.Group(ui.GroupProps{Class: "flex items-center gap-2"}) { … }`, previewDefault),
			showcaseutil.Variant("wrap", "Wrap", "", `@ui.Group(ui.GroupProps{Class: "flex flex-wrap gap-2"}) { … }`, previewWrap),
		},
		API: []registry.APIField{
			{Name: "Class", Type: "string", Description: "Tailwind utilities"},
			{Name: "Tag", Type: "string", Description: "div | span"},
			{Name: "Attrs", Type: "templ.Attributes", Description: "Extra attributes"},
		},
		Related: []registry.RelatedLink{
			{Label: "Stack", Href: "/docs/components/stack"},
			{Label: "Box", Href: "/docs/components/box"},
		},
	})
}
