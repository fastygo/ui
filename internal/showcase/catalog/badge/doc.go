package badge

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "badge",
		Title:       "Badge",
		Section:     "components",
		Description: "Small status label chip.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@showcaseutil.RenderBadge(ctx, w, ui.BadgeProps{}, "Badge")`, previewDefault),
			showcaseutil.Variant("secondary", "Secondary", "", `@showcaseutil.RenderBadge(ctx, w, ui.BadgeProps{Variant: "secondary"}, "Secondary")`, previewSecondary),
			showcaseutil.Variant("outline", "Outline", "", `@showcaseutil.RenderBadge(ctx, w, ui.BadgeProps{Variant: "outline"}, "Outline")`, previewOutline),
			showcaseutil.Variant("destructive", "Destructive", "", `@showcaseutil.RenderBadge(ctx, w, ui.BadgeProps{Variant: "destructive"}, "Alert")`, previewDestructive),
		},
		API: []registry.APIField{
			{Name: "Variant", Type: "string", Description: "default | secondary | destructive | outline"},
			{Name: "Size", Type: "string", Description: "default | sm | lg"},
			{Name: "Class", Type: "string", Description: "Extra utilities"},
		},
		Related: []registry.RelatedLink{
			{Label: "Button", Href: "/docs/components/button"},
			{Label: "Alert", Href: "/docs/components/alert"},
		},
	})
}
