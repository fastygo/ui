package docgen

import (
	"sort"

	"github.com/fastygo/ui/internal/doclocale"
	"github.com/fastygo/ui/internal/fixtures"
	"github.com/fastygo/ui/internal/ui/layout"
)

// BuildNavItems constructs sidebar navigation for static docs pages.
func BuildNavItems(pages []DocPage, routing doclocale.Routing, locale, activePath string, fix fixtures.Locale) []layout.NavItem {
	routing = routing.Normalize()
	overview := fix.Docs.NavOverview
	if overview == "" {
		overview = "Overview"
	}
	docsHome := fix.Docs.NavDocsHome
	if docsHome == "" {
		docsHome = "Docs home"
	}
	var items []layout.NavItem
	items = append(items,
		layout.NavItem{Label: overview, Path: "/", Icon: "home"},
		layout.NavItem{Label: docsHome, Path: routing.DocsHomePath(locale), Icon: "book-open"},
	)
	bySection := map[string][]DocPage{}
	for _, p := range pages {
		if p.Locale != locale {
			continue
		}
		bySection[p.Meta.Section] = append(bySection[p.Meta.Section], p)
	}
	order := []string{"getting-started", "primitives", "utils", "components", "blocks"}
	for _, sec := range order {
		secPages := bySection[sec]
		if len(secPages) == 0 {
			continue
		}
		sort.Slice(secPages, func(i, j int) bool {
			return secPages[i].Meta.Title < secPages[j].Meta.Title
		})
		items = append(items, layout.NavItem{Section: true, Label: sectionLabel(fix, sec)})
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

func sectionLabel(fix fixtures.Locale, id string) string {
	switch id {
	case "getting-started":
		if fix.Docs.SectionGettingStarted != "" {
			return fix.Docs.SectionGettingStarted
		}
		return "Getting Started"
	case "primitives":
		if fix.Docs.SectionPrimitives != "" {
			return fix.Docs.SectionPrimitives
		}
		return "Primitives"
	case "utils":
		if fix.Docs.SectionUtils != "" {
			return fix.Docs.SectionUtils
		}
		return "Utils"
	case "components":
		if fix.Docs.SectionComponents != "" {
			return fix.Docs.SectionComponents
		}
		return "Components"
	case "blocks":
		if fix.Docs.SectionBlocks != "" {
			return fix.Docs.SectionBlocks
		}
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
