package docgen

import (
	"sort"

	"github.com/fastygo/ui/internal/doclocale"
	"github.com/fastygo/ui/internal/ui/layout"
)

// BuildNavItems constructs sidebar navigation for static docs pages.
func BuildNavItems(pages []DocPage, routing doclocale.Routing, locale, activePath string) []layout.NavItem {
	routing = routing.Normalize()
	var items []layout.NavItem
	items = append(items,
		layout.NavItem{Label: "Overview", Path: "/", Icon: "home"},
		layout.NavItem{Label: "Docs home", Path: routing.DocsHomePath(locale), Icon: "book-open"},
	)
	bySection := map[string][]DocPage{}
	for _, p := range pages {
		if p.Locale != locale {
			continue
		}
		bySection[p.Meta.Section] = append(bySection[p.Meta.Section], p)
	}
	order := []string{"getting-started", "components", "blocks"}
	for _, sec := range order {
		secPages := bySection[sec]
		if len(secPages) == 0 {
			continue
		}
		sort.Slice(secPages, func(i, j int) bool {
			return secPages[i].Meta.Title < secPages[j].Meta.Title
		})
		items = append(items, layout.NavItem{Section: true, Label: sectionLabel(sec)})
		for _, p := range secPages {
			items = append(items, layout.NavItem{
				Label: p.Meta.Title,
				Path:  p.PublicPath,
				Icon:  navIconForSlug(p.Meta.Slug),
			})
		}
	}
	return items
}

func sectionLabel(id string) string {
	switch id {
	case "getting-started":
		return "Getting Started"
	case "components":
		return "Components"
	case "blocks":
		return "Blocks"
	default:
		return id
	}
}

func navIconForSlug(slug string) string {
	switch slug {
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
	case "card", "alert", "badge", "avatar", "blog-card":
		return "layout-grid"
	default:
		return "layout-grid"
	}
}
