package site

import (
	"context"
	"strings"

	"github.com/fastygo/framework/pkg/web/locale"
	"github.com/fastygo/ui/internal/doclocale"
	"github.com/fastygo/ui/internal/fixtures"
	"github.com/fastygo/ui/internal/ui/layout"
)

func (f *Feature) siteNav(ctx context.Context, fix fixtures.Locale) []layout.NavItem {
	routing := f.docsRouting()
	current := strings.ToLower(strings.TrimSpace(locale.From(ctx)))
	if current == "" {
		current = routing.Default
	}

	items := []layout.NavItem{
		{Label: fix.Nav.Home, Path: "/", Icon: "home"},
	}
	if len(fix.DocNav) > 0 {
		section := fix.Nav.DocsSection
		if section == "" {
			section = "Documentation"
		}
		items = append(items, layout.NavItem{Section: true, Label: section})
		for _, link := range fix.DocNav {
			icon := link.Icon
			if icon == "" {
				icon = "book-open"
			}
			path := link.Path
			if strings.HasPrefix(path, "/docs") {
				path = routing.AlternatePublicPath(path, current)
			}
			items = append(items, layout.NavItem{
				Label: link.Label,
				Path:  path,
				Icon:  icon,
			})
		}
	}
	items = append(items, layout.NavItem{Label: fix.Nav.Sample, Path: "/sample", Icon: "layout-dashboard"})
	return items
}

func (f *Feature) docsRouting() doclocale.Routing {
	return doclocale.Routing{
		Default: f.defaultLocale,
		Locales: f.available,
	}.Normalize()
}
