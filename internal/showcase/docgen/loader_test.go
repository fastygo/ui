package docgen_test

import (
	"testing"

	"github.com/fastygo/ui/internal/showcase/docgen"
)

func TestLoadAll_skipsDraftContent(t *testing.T) {
	pages, err := docgen.LoadAll(docgen.LoadOptions{Locales: []string{"en"}})
	if err != nil {
		t.Fatal(err)
	}
	draftSlugs := map[string]struct{}{
		"calendar": {}, "collapsible": {}, "combobox": {}, "command": {},
		"context-menu": {}, "data-table": {}, "hover-card": {}, "popover": {},
		"sheet": {}, "skeleton": {}, "toast": {}, "toggle": {}, "toggle-group": {},
		"tooltip": {},
	}
	for _, p := range pages {
		if p.Locale != "en" || p.Meta.Section != "components" {
			continue
		}
		if _, draft := draftSlugs[p.Meta.Slug]; draft {
			t.Fatalf("draft component %q should not be published", p.Meta.Slug)
		}
	}
}

func TestLoadAll_ruSourcePagesMatchEnPublicKeys(t *testing.T) {
	t.Helper()
	pages, err := docgen.LoadAll(docgen.LoadOptions{Locales: []string{"en", "ru"}})
	if err != nil {
		t.Fatal(err)
	}
	enKeys := map[string]struct{}{}
	for _, p := range pages {
		if p.Locale != "en" {
			continue
		}
		enKeys[pageKey(p)] = struct{}{}
	}
	for _, p := range pages {
		if p.Locale != "ru" || p.FallbackEN {
			continue
		}
		key := pageKey(p)
		if _, ok := enKeys[key]; !ok {
			t.Fatalf("ru source page %q is not in en public catalog", key)
		}
	}
}

func pageKey(p docgen.DocPage) string {
	return p.Meta.Section + "/" + p.Meta.Slug
}
