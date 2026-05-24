package doclocale_test

import (
	"path/filepath"
	"testing"

	"github.com/fastygo/ui/internal/doclocale"
)

func TestRouting_publicAndStaticPaths(t *testing.T) {
	r := doclocale.Routing{Default: "en", Locales: []string{"en", "ru"}}.Normalize()

	cases := []struct {
		name       string
		public     func() string
		url        string
		staticWant string
	}{
		{
			name:   "en home",
			public: func() string { return r.DocsHomePath("en") },
			url:    "/docs/",
			staticWant: "en",
		},
		{
			name:   "ru home",
			public: func() string { return r.DocsHomePath("ru") },
			url:    "/ru/docs/",
			staticWant: "ru",
		},
		{
			name:   "en button",
			public: func() string { return r.PublicPath("en", "primitives", "button") },
			url:    "/docs/primitives/button/",
			staticWant: "en/primitives/button",
		},
		{
			name:   "ru button",
			public: func() string { return r.PublicPath("ru", "primitives", "button") },
			url:    "/ru/docs/primitives/button/",
			staticWant: "ru/primitives/button",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := filepath.ToSlash(r.StaticFileRelPath(tc.url))
			if got != tc.staticWant {
				t.Fatalf("StaticFileRelPath(%q) = %q, want %q", tc.url, got, tc.staticWant)
			}
			if tc.public() != tc.url {
				t.Fatalf("PublicPath = %q, want %q", tc.public(), tc.url)
			}
		})
	}
}

func TestRouting_defaultRu(t *testing.T) {
	r := doclocale.Routing{Default: "ru", Locales: []string{"en", "ru"}}.Normalize()

	if got, want := r.PublicPath("ru", "primitives", "button"), "/docs/primitives/button/"; got != want {
		t.Fatalf("ru public = %q, want %q", got, want)
	}
	if got, want := r.PublicPath("en", "primitives", "button"), "/en/docs/primitives/button/"; got != want {
		t.Fatalf("en public = %q, want %q", got, want)
	}
	if got, want := filepath.ToSlash(r.StaticFileRelPath("/docs/primitives/button/")), "ru/primitives/button"; got != want {
		t.Fatalf("static ru default = %q, want %q", got, want)
	}
}

func TestRouting_alternatePublicPath(t *testing.T) {
	r := doclocale.Routing{Default: "en", Locales: []string{"en", "ru"}}.Normalize()
	if got, want := r.AlternatePublicPath("/docs/primitives/button/", "ru"), "/ru/docs/primitives/button/"; got != want {
		t.Fatalf("alternate = %q, want %q", got, want)
	}
}
