package docs

import (
	"github.com/a-h/templ"
	"github.com/fastygo/ui/internal/registry"
)

// ComponentPageData drives the component documentation layout.
type ComponentPageData struct {
	Page registry.Page
}

// IndexPageData is the /docs landing content.
type IndexPageData struct {
	Title       string
	Description string
	Sections    []IndexSection
}

// IndexSection groups links on the docs home page.
type IndexSection struct {
	Label string
	Links []IndexLink
}

// IndexLink is one card/link on the docs index.
type IndexLink struct {
	Title       string
	Description string
	Href        string
}

// VariantBlockData is one example block (preview + code + headings).
type VariantBlockData struct {
	Title       string
	Description string
	Preview     templ.Component
	Code        string
}
