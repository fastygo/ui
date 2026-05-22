package docgen

import (
	"fmt"
	"strings"

	"github.com/fastygo/ui/internal/doclocale"
)

var knownSections = map[string]struct{}{
	"getting-started": {},
	"components":      {},
	"blocks":          {},
}

func (o LoadOptions) routing() doclocale.Routing {
	return doclocale.Routing{
		Default: o.DefaultLocale,
		Locales: o.Locales,
	}.Normalize()
}

func applyPagePaths(r doclocale.Routing, page *DocPage) {
	page.PublicPath = r.PublicPath(page.Locale, page.Meta.Section, page.Meta.Slug)
	page.OutputPath = r.OutputRelPath(page.Locale, page.Meta.Section, page.Meta.Slug)
}

// PublicPath returns the canonical URL path for a page (trailing slash).
func PublicPath(locale string, meta PageMeta) string {
	return defaultRouting().PublicPath(locale, meta.Section, meta.Slug)
}

// OutputRelPath returns the relative file path under the docs output root.
func OutputRelPath(locale string, meta PageMeta) string {
	return defaultRouting().OutputRelPath(locale, meta.Section, meta.Slug)
}

func defaultRouting() doclocale.Routing {
	return doclocale.Routing{Default: "en", Locales: []string{"en", "ru"}}.Normalize()
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
