package docsstatic

// PageData is the view model for a static documentation page.
type PageData struct {
	Title       string
	Description string
	Source      string
	Blocks      []Block
	API         []APIField
	Related     []RelatedLink
}

// Block is one documentation segment.
type Block interface {
	blockKind() string
}

func (ParagraphBlock) blockKind() string   { return "paragraph" }
func (HeadingBlock) blockKind() string     { return "heading" }
func (ListBlock) blockKind() string        { return "list" }
func (DemoBlock) blockKind() string        { return "demo" }
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
}

// ListBlock is a bullet list.
type ListBlock struct {
	Items []string
}

// DemoBlock binds a preview registry ID to a code snippet.
type DemoBlock struct {
	ID         string
	CodeSource string
}

// CodeBlock is a standalone fenced code block.
type CodeBlock struct {
	Source string
}

// PreviewCodeBlock is a templ fence with live preview HTML and collapsible source.
type PreviewCodeBlock struct {
	ID         string
	Source     string
	HTML       string
	SourceFile string
	FenceIndex int
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
