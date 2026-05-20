package docgen

import (
	"fmt"
	"strings"
)

var knownSections = map[string]struct{}{
	"getting-started": {},
	"components":      {},
	"blocks":          {},
}

// PublicPath returns the canonical URL path for a page (trailing slash).
func PublicPath(locale string, meta PageMeta) string {
	base := registryPagePath(meta.Section, meta.Slug)
	if locale == "" || locale == "en" {
		return ensureTrailingSlash(base)
	}
	return ensureTrailingSlash("/" + locale + base)
}

func registryPagePath(section, slug string) string {
	section = strings.Trim(strings.ToLower(section), "/")
	slug = strings.Trim(strings.ToLower(slug), "/")
	switch section {
	case "getting-started", "getting_started", "start":
		return "/docs/" + slug
	case "blocks":
		return "/docs/blocks/" + slug
	default:
		return "/docs/components/" + slug
	}
}

// OutputRelPath returns the relative file path under the docs output root.
func OutputRelPath(locale string, meta PageMeta) string {
	slug := strings.Trim(strings.ToLower(meta.Slug), "/")
	section := strings.Trim(strings.ToLower(meta.Section), "/")
	var parts []string
	if locale != "" && locale != "en" {
		parts = append(parts, locale)
	}
	switch section {
	case "getting-started", "getting_started", "start":
		parts = append(parts, slug)
	case "blocks":
		parts = append(parts, "blocks", slug)
	default:
		parts = append(parts, "components", slug)
	}
	parts = append(parts, "index.html")
	return strings.Join(parts, "/")
}

func ensureTrailingSlash(p string) string {
	if p == "" {
		return "/"
	}
	if !strings.HasSuffix(p, "/") {
		return p + "/"
	}
	return p
}

// NormalizeHref normalizes internal doc links to trailing-slash form.
func NormalizeHref(href string) string {
	href = strings.TrimSpace(href)
	if href == "" || strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") || strings.HasPrefix(href, "#") {
		return href
	}
	if !strings.HasPrefix(href, "/") {
		href = "/" + href
	}
	if strings.HasPrefix(href, "/docs") && !strings.HasSuffix(href, "/") && !strings.Contains(strings.TrimPrefix(href, "/docs"), ".") {
		return href + "/"
	}
	return href
}

// DemoID builds the canonical preview registry id for a page slug and variant id.
func DemoID(pageSlug, variantID string) string {
	return pageSlug + "." + variantID
}

// ValidateMeta checks required front matter fields.
func ValidateMeta(meta PageMeta) error {
	if strings.TrimSpace(meta.Slug) == "" {
		return fmt.Errorf("front matter: slug is required")
	}
	if strings.TrimSpace(meta.Section) == "" {
		return fmt.Errorf("front matter: section is required")
	}
	if _, ok := knownSections[strings.ToLower(meta.Section)]; !ok {
		return fmt.Errorf("front matter: unknown section %q", meta.Section)
	}
	if strings.TrimSpace(meta.Title) == "" {
		return fmt.Errorf("front matter: title is required")
	}
	if strings.TrimSpace(meta.Description) == "" {
		return fmt.Errorf("front matter: description is required")
	}
	return nil
}
