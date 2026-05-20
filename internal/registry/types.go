package registry

import "github.com/a-h/templ"

// Section groups documentation pages in the sidebar (shadcn-style).
type Section struct {
	ID    string
	Label string
}

// Page is one showcase route (component, block, or guide).
type Page struct {
	Slug        string
	Title       string
	Description string
	Section     string
	Path        string
	Source      string
	Package     string
	Variants    []Variant
	API         []APIField
	Related     []RelatedLink
}

// Variant is one stacked example on a page (preview + code snippet).
type Variant struct {
	ID          string
	Title       string
	Description string
	Preview     templ.Component
	Code        string
}

// APIField documents a prop or exported type field.
type APIField struct {
	Name        string
	Type        string
	Required    bool
	Description string
}

// RelatedLink points to sibling docs or external references.
type RelatedLink struct {
	Label string
	Href  string
}
