package site

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/ui/layout"
)

func docsNavItems(active string) []layout.NavItem {
	var items []layout.NavItem
	items = append(items,
		layout.NavItem{Label: "Overview", Path: "/", Icon: "home"},
		layout.NavItem{Label: "Docs home", Path: "/docs", Icon: "book-open"},
	)
	for _, sec := range registry.Sections() {
		items = append(items, layout.NavItem{Section: true, Label: sec.Label})
		for _, p := range registry.PagesInSection(sec.ID) {
			items = append(items, layout.NavItem{
				Label: p.Title,
				Path:  p.Path,
				Icon:  navIconForPage(p),
			})
		}
	}
	return items
}

func navIconForPage(p registry.Page) string {
	switch p.Slug {
	case "button", "toggle", "toggle-group":
		return "mouse-pointer-click"
	case "input", "textarea", "select", "checkbox", "radio", "form", "switch", "slider", "combobox":
		return "text-cursor-input"
	case "table", "data-table":
		return "table"
	case "dialog", "sheet", "drawer", "alert-dialog":
		return "panel-right"
	case "tabs", "accordion", "navigation-menu", "menubar", "breadcrumb":
		return "layout-list"
	case "card", "alert", "badge", "avatar":
		return "layout-grid"
	default:
		return "layout-grid"
	}
}
