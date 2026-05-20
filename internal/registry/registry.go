package registry

import (
	"sort"
	"strings"
	"sync"
)

var (
	mu    sync.RWMutex
	pages = map[string]Page{}
)

// Register adds or replaces a documentation page. Slug must be unique.
func Register(p Page) {
	if strings.TrimSpace(p.Slug) == "" {
		return
	}
	if p.Path == "" {
		p.Path = PagePath(p.Section, p.Slug)
	}
	mu.Lock()
	pages[p.Slug] = p
	mu.Unlock()
}

// PagePath builds the canonical docs URL for a section and slug.
func PagePath(section, slug string) string {
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

// PageByPath returns a page for an exact docs path, or false.
func PageByPath(path string) (Page, bool) {
	path = normalizePath(path)
	mu.RLock()
	defer mu.RUnlock()
	for _, p := range pages {
		if normalizePath(p.Path) == path {
			return p, true
		}
	}
	return Page{}, false
}

// PageBySlug returns a page by slug.
func PageBySlug(slug string) (Page, bool) {
	mu.RLock()
	defer mu.RUnlock()
	p, ok := pages[strings.ToLower(slug)]
	return p, ok
}

// AllPages returns pages sorted by section then title.
func AllPages() []Page {
	mu.RLock()
	defer mu.RUnlock()
	out := make([]Page, 0, len(pages))
	for _, p := range pages {
		out = append(out, p)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Section != out[j].Section {
			return sectionOrder(out[i].Section) < sectionOrder(out[j].Section)
		}
		return out[i].Title < out[j].Title
	})
	return out
}

// Sections returns sidebar sections that have at least one page.
func Sections() []Section {
	seen := map[string]struct{}{}
	var ids []string
	for _, p := range AllPages() {
		if p.Section == "" {
			continue
		}
		if _, ok := seen[p.Section]; ok {
			continue
		}
		seen[p.Section] = struct{}{}
		ids = append(ids, p.Section)
	}
	sort.Slice(ids, func(i, j int) bool {
		return sectionOrder(ids[i]) < sectionOrder(ids[j])
	})
	out := make([]Section, 0, len(ids))
	for _, id := range ids {
		out = append(out, Section{ID: id, Label: sectionLabel(id)})
	}
	return out
}

// PagesInSection returns pages for one section ID.
func PagesInSection(sectionID string) []Page {
	var out []Page
	for _, p := range AllPages() {
		if p.Section == sectionID {
			out = append(out, p)
		}
	}
	return out
}

func normalizePath(p string) string {
	p = strings.TrimSpace(p)
	if p == "" {
		return "/"
	}
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}
	if len(p) > 1 && strings.HasSuffix(p, "/") {
		p = strings.TrimSuffix(p, "/")
	}
	return p
}

func sectionOrder(id string) int {
	switch id {
	case "getting-started":
		return 0
	case "components":
		return 1
	case "blocks":
		return 2
	default:
		return 99
	}
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
