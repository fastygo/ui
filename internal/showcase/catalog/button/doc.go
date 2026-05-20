package button

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "button",
		Title:       "Button",
		Section:     "components",
		Description: "Triggers an action or navigates when rendered as a link. Built on github.com/fastygo/templ/ui.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@showcaseutil.Button(ui.ButtonProps{}, "Button").Render(ctx, w)`, previewDefault),
			showcaseutil.Variant("secondary", "Secondary", "", `@showcaseutil.Button(ui.ButtonProps{Variant: "secondary"}, "Secondary").Render(ctx, w)`, previewSecondary),
			showcaseutil.Variant("outline", "Outline", "", `@showcaseutil.Button(ui.ButtonProps{Variant: "outline"}, "Outline").Render(ctx, w)`, previewOutline),
			showcaseutil.Variant("destructive", "Destructive", "", `@showcaseutil.Button(ui.ButtonProps{Variant: "destructive"}, "Destructive").Render(ctx, w)`, previewDestructive),
			showcaseutil.Variant("ghost", "Ghost", "", `@showcaseutil.Button(ui.ButtonProps{Variant: "ghost"}, "Ghost").Render(ctx, w)`, previewGhost),
			showcaseutil.Variant("link", "Link", "", `@showcaseutil.Button(ui.ButtonProps{Variant: "link"}, "Link").Render(ctx, w)`, previewLink),
			showcaseutil.Variant("sizes", "Sizes", "", `sm / default / lg via ButtonProps.Size`, previewSizes),
		},
		API: []registry.APIField{
			{Name: "Variant", Type: "string", Description: "default | secondary | destructive | outline | ghost | link | unstyled"},
			{Name: "Size", Type: "string", Description: "default | sm | lg | icon"},
			{Name: "Type", Type: "string", Description: "button | submit | reset"},
			{Name: "Href", Type: "string", Description: "When set, renders an anchor instead of button"},
			{Name: "Disabled", Type: "bool", Description: "Disables interaction"},
			{Name: "Class", Type: "string", Description: "Additional Tailwind utilities"},
			{Name: "AriaLabel", Type: "string", Description: "Accessible name when visible text is insufficient"},
		},
		Related: []registry.RelatedLink{
			{Label: "Toggle", Href: "/docs/components/toggle"},
			{Label: "Form", Href: "/docs/components/form"},
		},
	})
}
