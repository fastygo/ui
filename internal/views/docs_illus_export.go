package views

import "strings"

func docsIllusExportTitle(section string) string {
	switch section {
	case "primitives":
		return "Docs index illustration export — primitives"
	case "components":
		return "Docs index illustration export — components"
	default:
		return "Docs index illustration export"
	}
}

func docsIllusExportSlug(href string) string {
	href = strings.TrimSuffix(href, "/")
	if i := strings.LastIndex(href, "/"); i >= 0 {
		return href[i+1:]
	}
	return href
}
