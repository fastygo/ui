package container

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "container",
		Title:       "Container",
		Section:     "components",
		Description: "Centers content with a max-width constraint.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@ui.Container(ui.ContainerProps{Class: "mx-auto max-w-3xl px-4"}) { … }`, previewDefault),
			showcaseutil.Variant("section", "Section", "", `@ui.Container(ui.ContainerProps{Tag: "section"}) { … }`, previewSection),
		},
		API: []registry.APIField{
			{Name: "Class", Type: "string", Description: "Tailwind utilities"},
			{Name: "Tag", Type: "string", Description: "motion.div | section | main"},
			{Name: "Attrs", Type: "templ.Attributes", Description: "Extra attributes"},
		},
		Related: []registry.RelatedLink{
			{Label: "Stack", Href: "/docs/components/stack"},
			{Label: "Box", Href: "/docs/components/box"},
		},
	})
}
