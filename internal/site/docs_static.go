package site

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const staticDocsRoot = "web/static/docs"

// registerStaticDocsRoutes serves prebuilt documentation HTML when present.
func (f *Feature) registerStaticDocsRoutes(mux *http.ServeMux) {
	if _, err := os.Stat(staticDocsRoot); err != nil {
		return
	}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveStaticDoc(w, r)
	})
	mux.HandleFunc("GET /docs/", handler)
	mux.HandleFunc("GET /docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/docs/", http.StatusFound)
	})
	mux.HandleFunc("GET /ru/docs/", handler)
	mux.HandleFunc("GET /ru/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/ru/docs/", http.StatusFound)
	})
}

func serveStaticDoc(w http.ResponseWriter, r *http.Request) {
	rel := staticDocsRelPath(r.URL.Path)
	if rel == "" && !strings.HasSuffix(r.URL.Path, "/") {
		http.NotFound(w, r)
		return
	}
	candidates := []string{
		filepath.Join(staticDocsRoot, rel, "index.html"),
	}
	if rel == "" {
		candidates = []string{filepath.Join(staticDocsRoot, "index.html")}
	}
	for _, c := range candidates {
		if fi, err := os.Stat(c); err == nil && !fi.IsDir() {
			http.ServeFile(w, r, c)
			return
		}
	}
	http.NotFound(w, r)
}

// staticDocsRelPath maps a request path to a file path under web/static/docs.
func staticDocsRelPath(urlPath string) string {
	rel := strings.TrimPrefix(urlPath, "/")
	rel = strings.TrimSuffix(rel, "/")
	if strings.HasPrefix(rel, "ru/docs") {
		suffix := strings.TrimPrefix(rel, "ru/docs")
		suffix = strings.TrimPrefix(suffix, "/")
		if suffix == "" {
			return "ru"
		}
		return filepath.Join("ru", filepath.FromSlash(suffix))
	}
	if strings.HasPrefix(rel, "docs") {
		suffix := strings.TrimPrefix(rel, "docs")
		suffix = strings.TrimPrefix(suffix, "/")
		return filepath.FromSlash(suffix)
	}
	return ""
}
