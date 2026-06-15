package docsstatic

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestIllustrationPatterns_allEntriesCovered(t *testing.T) {
	entries := IndexIllustrationEntries()
	if len(entries) != 52 {
		t.Fatalf("expected 52 registry entries, got %d", len(entries))
	}
	for _, entry := range entries {
		if _, ok := indexCardIllustration(entry.Href); !ok {
			t.Fatalf("href %q not found in suffix map", entry.Href)
		}
	}
}

func TestIllustrationPatterns_rendersWithoutError(t *testing.T) {
	ctx := context.Background()
	for _, entry := range IndexIllustrationEntries() {
		var buf bytes.Buffer
		if err := IndexIllustrationComponent(entry.Href, IllustrationEmbedded).Render(ctx, &buf); err != nil {
			t.Fatalf("render %q: %v", entry.Href, err)
		}
		if buf.Len() == 0 {
			t.Fatalf("render %q: empty output", entry.Href)
		}
	}
}

func TestIndexCardIllustrationClass(t *testing.T) {
	if got := IndexCardIllustrationClass("/docs/primitives/button/"); got != indexCardIllustratedClass {
		t.Fatalf("button class = %q, want %q", got, indexCardIllustratedClass)
	}
	if got := IndexCardIllustrationClass("/docs/introduction/"); got != "" {
		t.Fatalf("introduction class = %q, want empty", got)
	}
}

func TestIllustrationRootClass_layoutModifiers(t *testing.T) {
	if got := illustrationRootClass(IllustrationEmbedded); !strings.Contains(got, "absolute right-4 top-4") {
		t.Fatalf("embedded class missing positioning: %q", got)
	}
	if got := illustrationRootClass(IllustrationStandalone); !strings.Contains(got, "relative") {
		t.Fatalf("standalone class missing modifier: %q", got)
	}
}

func TestIllusOpacityToneClass(t *testing.T) {
	for tone, want := range map[string]string{
		"soft":   illusBgSoft,
		"border": illusBgBorder,
		"ink":    illusBgInk,
		"accent": illusBgAccent,
	} {
		if got := IllusOpacityToneClass(tone); got != want {
			t.Fatalf("tone %q = %q, want %q", tone, got, want)
		}
	}
}
