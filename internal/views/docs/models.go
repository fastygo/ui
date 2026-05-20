package docs

// IndexData is the docs home page body.
type IndexData struct {
	Title       string
	Description string
	Sections    []IndexSection
}

// IndexSection groups cards on the docs index.
type IndexSection struct {
	Label string
	Links []IndexLink
}

// IndexLink is one entry on the docs home page.
type IndexLink struct {
	Title       string
	Description string
	Href        string
}
