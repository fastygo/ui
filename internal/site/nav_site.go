package site

import (
	"github.com/fastygo/ui/internal/fixtures"
	"github.com/fastygo/ui/internal/ui/layout"
)

func siteNav(fix fixtures.Locale) []layout.NavItem {
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
			items = append(items, layout.NavItem{
				Label: link.Label,
				Path:  link.Path,
				Icon:  icon,
			})
		}
	}
	items = append(items, layout.NavItem{Label: fix.Nav.Sample, Path: "/sample", Icon: "layout-dashboard"})
	return items
}
