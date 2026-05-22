package site

import (
	"net/http"
)

func (f *Feature) registerDocsRoutes(mux *http.ServeMux) {
	f.registerStaticDocsRoutes(mux)
}
