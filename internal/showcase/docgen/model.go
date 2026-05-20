package docgen

// PageMeta is YAML front matter for a documentation page.
type PageMeta struct {
	Slug        string        `yaml:"slug"`
	Section     string        `yaml:"section"`
	Title       string        `yaml:"title"`
	Description string        `yaml:"description"`
	Source      string        `yaml:"source,omitempty"`
	Package     string        `yaml:"package,omitempty"`
	Related     []RelatedLink `yaml:"related,omitempty"`
	API         []APIField    `yaml:"api,omitempty"`
}

// RelatedLink points to another docs page or external resource.
type RelatedLink struct {
	Label    string `yaml:"label"`
	Href     string `yaml:"href"`
	External bool   `yaml:"external,omitempty"`
}

// APIField documents a prop or exported type field.
type APIField struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	Required    bool   `yaml:"required,omitempty"`
	Description string `yaml:"description"`
}

// DocPage is a parsed, locale-resolved documentation page.
type DocPage struct {
	Locale      string
	Meta        PageMeta
	OutputPath  string // relative path under output root, e.g. components/blog-card/index.html
	PublicPath  string // URL path, e.g. /docs/components/blog-card/
	Blocks      []Block
	Headings    []Heading
	DemoIDs     []string
	SourceFile  string // path within embed FS for debugging
	FallbackEN  bool   // true when locale content fell back to English
}

// Heading is a table-of-contents entry extracted from the page.
type Heading struct {
	Level int
	Text  string
	ID    string
}

// Block is one rendered documentation segment.
type Block interface {
	blockKind() string
}

func (ParagraphBlock) blockKind() string  { return "paragraph" }
func (HeadingBlock) blockKind() string    { return "heading" }
func (ListBlock) blockKind() string       { return "list" }
func (DemoBlock) blockKind() string       { return "demo" }
func (CodeBlock) blockKind() string       { return "code" }

// ParagraphBlock is plain prose (may contain inline markdown formatting as plain text initially).
type ParagraphBlock struct {
	Text string
}

// HeadingBlock is a section heading from markdown.
type HeadingBlock struct {
	Level int
	Text  string
	ID    string
}

// ListBlock is a bullet list.
type ListBlock struct {
	Items []string
}

// DemoBlock binds a preview registry ID to a code snippet.
type DemoBlock struct {
	ID   string
	Code CodeBlock
}

// CodeBlock is a fenced code snippet (standalone or attached to a demo).
type CodeBlock struct {
	Language string
	Source   string
}

// SearchEntry is one row in search-index.json.
type SearchEntry struct {
	Locale      string   `json:"locale"`
	Section     string   `json:"section"`
	Slug        string   `json:"slug"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Href        string   `json:"href"`
	Headings    []string `json:"headings,omitempty"`
}

// RegistryItem is one row in registry-manifest.json.
type RegistryItem struct {
	Slug    string   `json:"slug"`
	Section string   `json:"section"`
	Title   string   `json:"title"`
	Source  string   `json:"source,omitempty"`
	Package string   `json:"package,omitempty"`
	Demos   []string `json:"demos"`
}
