package docsstatic

// PageData is the view model for a static documentation page.
type PageData struct {
	Title       string
	Description string
	Source      string
	Blocks      []Block
	API         []APIField
	Related     []RelatedLink
	TOC                 []TOCHeading
	TOCLabel            string
	APISectionTitle     string
	RelatedSectionTitle string
}

// TOCHeading is one in-page table-of-contents entry.
type TOCHeading struct {
	Level int
	Text  string
	ID    string
}

// Block is one documentation segment.
type Block interface {
	blockKind() string
}

func (ParagraphBlock) blockKind() string   { return "paragraph" }
func (HeadingBlock) blockKind() string     { return "heading" }
func (ListBlock) blockKind() string        { return "list" }
func (CodeBlock) blockKind() string        { return "code" }
func (PreviewCodeBlock) blockKind() string { return "preview" }

// ParagraphBlock is plain prose.
type ParagraphBlock struct {
	Text string
}

// HeadingBlock is a section heading.
type HeadingBlock struct {
	Level int
	Text  string
	ID    string
}

// ListBlock is a bullet list.
type ListBlock struct {
	Items []string
}

// CodeBlock is a standalone fenced code block.
type CodeBlock struct {
	Language        string
	Source          string
	HighlightedHTML string
}

// PreviewCodeBlock is a templ fence with live preview HTML and collapsible source.
type PreviewCodeBlock struct {
	ID              string
	Source          string
	HTML            string
	HighlightedHTML string
	SourceFile      string
	FenceIndex      int
}

// APIField documents a prop.
type APIField struct {
	Name        string
	Type        string
	Description string
}

// RelatedLink points to another page.
type RelatedLink struct {
	Label string
	Href  string
}
