package site

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/fastygo/ui/internal/doclocale"
)

// registerStaticDocsRoutes serves prebuilt documentation HTML when present.
func (f *Feature) registerStaticDocsRoutes(mux *http.ServeMux) {
	root := resolveDocsRoot(f.staticDocsRoot)
	logDocsRoot(root)
	if !docsRootValid(root) {
		return
	}
	f.staticDocsRoot = root

	routing := f.docsRouting()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f.serveStaticDoc(w, r)
	})
	for _, loc := range routing.Locales {
		if loc == routing.Default {
			mux.HandleFunc("GET /docs/", handler)
			mux.HandleFunc("GET /docs", func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, "/docs/", http.StatusFound)
			})
			registerDefaultLocalePrefixRedirects(mux, routing, loc)
			continue
		}
		prefix := "/" + loc + "/docs"
		mux.HandleFunc("GET "+prefix+"/", handler)
		mux.HandleFunc("GET "+prefix, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, prefix+"/", http.StatusFound)
		})
	}
}

// registerDefaultLocalePrefixRedirects sends /en/docs/… to /docs/… when en is default.
func registerDefaultLocalePrefixRedirects(mux *http.ServeMux, routing doclocale.Routing, defaultLocale string) {
	prefix := "/" + defaultLocale + "/docs"
	mux.HandleFunc("GET "+prefix, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/docs/", http.StatusMovedPermanently)
	})
	mux.HandleFunc("GET "+prefix+"/", func(w http.ResponseWriter, r *http.Request) {
		target := strings.TrimPrefix(r.URL.Path, prefix)
		if target == "" || target == "/" {
			target = "/"
		}
		http.Redirect(w, r, "/docs"+target, http.StatusMovedPermanently)
	})
}

func (f *Feature) serveStaticDoc(w http.ResponseWriter, r *http.Request) {
	if f.staticDocsRoot == "" {
		http.NotFound(w, r)
		return
	}
	path := r.URL.Path
	if !strings.HasSuffix(path, "/") {
		http.Redirect(w, r, path+"/", http.StatusMovedPermanently)
		return
	}

	routing := f.docsRouting()
	for _, candidate := range f.docFileCandidates(routing, path) {
		if fi, err := os.Stat(candidate); err == nil && !fi.IsDir() {
			http.ServeFile(w, r, candidate)
			return
		}
	}
	http.NotFound(w, r)
}

func (f *Feature) docFileCandidates(routing doclocale.Routing, urlPath string) []string {
	rel := routing.StaticFileRelPath(urlPath)
	if rel == "" {
		return nil
	}
	out := []string{filepath.Join(f.staticDocsRoot, rel, "index.html")}

	_, suffix, ok := routing.ParseDocsURL(urlPath)
	if !ok {
		return out
	}
	if suffix == "" {
		return append(out, filepath.Join(f.staticDocsRoot, "index.html"))
	}
	return append(out, filepath.Join(f.staticDocsRoot, filepath.FromSlash(suffix), "index.html"))
}
