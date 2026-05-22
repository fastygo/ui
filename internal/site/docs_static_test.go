package site_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/fastygo/ui/internal/site"
)

func TestStaticDocsRoutes_defaultLocale(t *testing.T) {
	root := filepath.Join("..", "..", "web", "static", "docs")
	if _, err := os.Stat(filepath.Join(root, "en", "index.html")); err != nil {
		t.Skip("docs not built; run go run ./cmd/docgen")
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		t.Fatal(err)
	}

	mux := http.NewServeMux()
	feat := site.NewFeature([]string{"en", "ru"}, "en", abs)
	feat.Routes(mux)

	cases := []struct {
		path       string
		wantStatus int
	}{
		{"/docs/", 200},
		{"/docs/components/button/", 200},
		{"/ru/docs/", 200},
		{"/en/docs/", 301},
		{"/en/docs/components/button/", 301},
	}
	for _, tc := range cases {
		t.Run(tc.path, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tc.path, nil)
			rec := httptest.NewRecorder()
			mux.ServeHTTP(rec, req)
			if rec.Code != tc.wantStatus {
				t.Fatalf("GET %s: status %d, want %d", tc.path, rec.Code, tc.wantStatus)
			}
			if tc.wantStatus == 301 {
				loc := rec.Header().Get("Location")
				if loc == "" {
					t.Fatal("expected redirect location")
				}
			}
		})
	}
}
