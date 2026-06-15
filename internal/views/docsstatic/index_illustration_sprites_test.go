package docsstatic

import "testing"

func TestIndexIllustrationSpriteMeta_allRegistryEntries(t *testing.T) {
	primCount := 0
	compCount := 0
	for i, entry := range IndexIllustrationEntries() {
		meta, ok := indexIllustrationSpriteMeta(entry.Href)
		if !ok {
			t.Fatalf("entry %d href %q: no sprite meta", i, entry.Href)
		}
		if meta.section != entry.Section {
			t.Fatalf("href %q section = %q, want %q", entry.Href, meta.section, entry.Section)
		}
		switch entry.Section {
		case IllusSectionPrimitives:
			if meta.index != primCount {
				t.Fatalf("href %q primitive index = %d, want %d", entry.Href, meta.index, primCount)
			}
			primCount++
		case IllusSectionComponents:
			if meta.index != compCount {
				t.Fatalf("href %q component index = %d, want %d", entry.Href, meta.index, compCount)
			}
			compCount++
		default:
			t.Fatalf("unexpected section %q", entry.Section)
		}
	}
	if primCount != 26 || compCount != 26 {
		t.Fatalf("sprite counts: primitives=%d components=%d, want 26 each", primCount, compCount)
	}
}

func TestIndexIllustrationSpriteMeta_backgroundPosition(t *testing.T) {
	cases := []struct {
		href string
		want string
	}{
		{"/docs/primitives/badge/", "background-position:0px 0px"},
		{"/docs/primitives/input/", "background-position:-144px -160px"},
		{"/docs/components/tabs/", "background-position:-144px -320px"},
	}
	for _, tc := range cases {
		meta, ok := indexIllustrationSpriteMeta(tc.href)
		if !ok {
			t.Fatalf("href %q: not found", tc.href)
		}
		if got := meta.backgroundPositionStyle(); got != tc.want {
			t.Fatalf("href %q position = %q, want %q", tc.href, got, tc.want)
		}
	}
}

func TestIndexIllustrationSpriteMeta_unknownHref(t *testing.T) {
	if _, ok := indexIllustrationSpriteMeta("/docs/introduction/"); ok {
		t.Fatal("introduction should not have sprite meta")
	}
}
