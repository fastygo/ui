package doclocale

import (
	"path/filepath"
	"strings"
)

// Routing maps documentation locales to public URLs and on-disk paths.
type Routing struct {
	Default  string
	Locales  []string
	Labels   map[string]string // optional display labels, e.g. "en" -> "En"
}

// Normalize fills defaults and normalizes locale codes.
func (r Routing) Normalize() Routing {
	out := r
	if out.Default == "" {
		out.Default = "en"
	}
	out.Default = strings.ToLower(strings.TrimSpace(out.Default))
	if len(out.Locales) == 0 {
		out.Locales = []string{out.Default}
	}
	seen := map[string]struct{}{}
	var locales []string
	for _, raw := range out.Locales {
		code := strings.ToLower(strings.TrimSpace(raw))
		if code == "" {
			continue
		}
		if _, ok := seen[code]; ok {
			continue
		}
		seen[code] = struct{}{}
		locales = append(locales, code)
	}
	if len(locales) == 0 {
		locales = []string{out.Default}
	}
	out.Locales = locales
	if out.Labels == nil {
		out.Labels = map[string]string{
			"en": "En",
			"ru": "Ru",
		}
	}
	return out
}

// Label returns the short display label for a locale code.
func (r Routing) Label(locale string) string {
	r = r.Normalize()
	code := strings.ToLower(strings.TrimSpace(locale))
	if label, ok := r.Labels[code]; ok && strings.TrimSpace(label) != "" {
		return label
	}
	if code == "" {
		return ""
	}
	return strings.ToUpper(code)
}

// PublicPath returns the canonical URL path for a docs page (trailing slash).
func (r Routing) PublicPath(locale, section, slug string) string {
	r = r.Normalize()
	base := registryPagePath(section, slug)
	code := strings.ToLower(strings.TrimSpace(locale))
	if code == "" {
		code = r.Default
	}
	if code == r.Default {
		return ensureTrailingSlash(base)
	}
	return ensureTrailingSlash("/" + code + base)
}

// OutputRelPath returns the relative file path under the docs output root.
func (r Routing) OutputRelPath(locale, section, slug string) string {
	r = r.Normalize()
	code := strings.ToLower(strings.TrimSpace(locale))
	if code == "" {
		code = r.Default
	}
	slug = strings.Trim(strings.ToLower(slug), "/")
	section = strings.Trim(strings.ToLower(section), "/")
	parts := []string{code}
	switch section {
	case "getting-started", "getting_started", "start":
		parts = append(parts, slug)
	case "blocks":
		parts = append(parts, "blocks", slug)
	case "primitives":
		parts = append(parts, "primitives", slug)
	case "utils":
		parts = append(parts, "utils", slug)
	default:
		parts = append(parts, "components", slug)
	}
	parts = append(parts, "index.html")
	return strings.Join(parts, "/")
}

// DocsHomePath returns the docs index URL for a locale.
func (r Routing) DocsHomePath(locale string) string {
	return r.PublicPath(locale, "getting-started", "")
}

// IndexOutputRelPath returns the relative path for a locale docs index file.
func (r Routing) IndexOutputRelPath(locale string) string {
	r = r.Normalize()
	code := strings.ToLower(strings.TrimSpace(locale))
	if code == "" {
		code = r.Default
	}
	return filepath.Join(code, "index.html")
}

// ParseDocsURL extracts locale and page suffix from a docs request path.
// suffix uses forward slashes, e.g. "components/button".
func (r Routing) ParseDocsURL(urlPath string) (locale, suffix string, ok bool) {
	r = r.Normalize()
	rel := strings.TrimPrefix(strings.TrimSpace(urlPath), "/")
	rel = strings.TrimSuffix(rel, "/")
	if rel == "docs" {
		return r.Default, "", true
	}
	for _, loc := range r.Locales {
		prefix := loc + "/docs"
		if rel == prefix {
			if loc == r.Default {
				return r.Default, "", true
			}
			return loc, "", true
		}
		if strings.HasPrefix(rel, prefix+"/") {
			return loc, strings.TrimPrefix(rel, prefix+"/"), true
		}
	}
	if strings.HasPrefix(rel, "docs/") {
		return r.Default, strings.TrimPrefix(rel, "docs/"), true
	}
	return "", "", false
}

// StaticFileRelPath maps a request path to a path under web/static/docs.
func (r Routing) StaticFileRelPath(urlPath string) string {
	locale, suffix, ok := r.ParseDocsURL(urlPath)
	if !ok {
		return ""
	}
	if suffix == "" {
		return locale
	}
	return filepath.Join(locale, filepath.FromSlash(suffix))
}

// AlternatePublicPath returns the same logical page in another locale.
func (r Routing) AlternatePublicPath(currentPath, targetLocale string) string {
	r = r.Normalize()
	currentPath = ensureTrailingSlash(currentPath)
	target := strings.ToLower(strings.TrimSpace(targetLocale))
	currentLocale, suffix, ok := r.ParseDocsURL(currentPath)
	if !ok {
		return r.DocsHomePath(target)
	}
	if currentLocale == target {
		return currentPath
	}
	section, slug := suffixToSectionSlug(suffix)
	if slug == "" && suffix == "" {
		return r.DocsHomePath(target)
	}
	return r.PublicPath(target, section, slug)
}

func suffixToSectionSlug(suffix string) (section, slug string) {
	suffix = strings.Trim(suffix, "/")
	if suffix == "" {
		return "getting-started", ""
	}
	parts := strings.Split(suffix, "/")
	switch parts[0] {
	case "blocks":
		if len(parts) >= 2 {
			return "blocks", parts[1]
		}
	case "components":
		if len(parts) >= 2 {
			return "components", parts[1]
		}
	case "primitives":
		if len(parts) >= 2 {
			return "primitives", parts[1]
		}
	case "utils":
		if len(parts) >= 2 {
			return "utils", parts[1]
		}
	default:
		return "getting-started", parts[0]
	}
	return "components", suffix
}

func registryPagePath(section, slug string) string {
	section = strings.Trim(strings.ToLower(section), "/")
	slug = strings.Trim(strings.ToLower(slug), "/")
	switch section {
	case "getting-started", "getting_started", "start":
		if slug == "" {
			return "/docs"
		}
		return "/docs/" + slug
	case "blocks":
		return "/docs/blocks/" + slug
	case "primitives":
		return "/docs/primitives/" + slug
	case "utils":
		return "/docs/utils/" + slug
	default:
		return "/docs/components/" + slug
	}
}

func ensureTrailingSlash(p string) string {
	if p == "" || p == "/" {
		return "/"
	}
	if !strings.HasSuffix(p, "/") {
		return p + "/"
	}
	return p
}
