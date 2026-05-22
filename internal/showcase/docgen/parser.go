package docgen

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"gopkg.in/yaml.v3"
)

var demoDirectiveRe = regexp.MustCompile(`^\{\{demo\s+id="([^"]+)"\s*\}\}\s*$`)

// ParseFile parses a markdown file with YAML front matter into a DocPage.
func ParseFile(locale, sourceFile string, raw []byte) (DocPage, error) {
	meta, body, err := splitFrontMatter(raw)
	if err != nil {
		return DocPage{}, fmt.Errorf("%s: %w", sourceFile, err)
	}
	if err := ValidateMeta(meta); err != nil {
		return DocPage{}, fmt.Errorf("%s: %w", sourceFile, err)
	}
	if !isGettingStarted(meta.Section) {
		if err := ValidateBodyRules(sourceFile, body); err != nil {
			return DocPage{}, err
		}
	}
	blocks, demoIDs, err := parseBody(sourceFile, body)
	if err != nil {
		return DocPage{}, err
	}
	headings := extractHeadings(blocks)
	page := DocPage{
		Locale:     locale,
		Meta:       meta,
		OutputPath: OutputRelPath(locale, meta),
		PublicPath: PublicPath(locale, meta),
		Blocks:     blocks,
		Headings:   headings,
		DemoIDs:    demoIDs,
		SourceFile: sourceFile,
	}
	if err := ValidatePage(page); err != nil {
		return DocPage{}, err
	}
	return page, nil
}

func isGettingStarted(section string) bool {
	s := strings.ToLower(strings.TrimSpace(section))
	return s == "getting-started" || s == "getting_started" || s == "start"
}

func splitFrontMatter(raw []byte) (PageMeta, string, error) {
	s := string(raw)
	if !strings.HasPrefix(s, "---\n") && !strings.HasPrefix(s, "---\r\n") {
		return PageMeta{}, "", fmt.Errorf("missing YAML front matter")
	}
	s = strings.TrimPrefix(s, "---")
	s = strings.TrimPrefix(s, "\r\n")
	s = strings.TrimPrefix(s, "\n")
	end := strings.Index(s, "\n---")
	if end < 0 {
		return PageMeta{}, "", fmt.Errorf("unterminated front matter")
	}
	fm := s[:end]
	body := strings.TrimPrefix(s[end+4:], "\n")
	body = strings.TrimPrefix(body, "\r\n")
	var meta PageMeta
	if err := yaml.Unmarshal([]byte(fm), &meta); err != nil {
		return PageMeta{}, "", fmt.Errorf("front matter: %w", err)
	}
	return meta, body, nil
}

func parseBody(sourceFile, body string) ([]Block, []string, error) {
	segments := splitDemoSegments(body)
	var blocks []Block
	var demoIDs []string
	var templFenceIndex int
	for i, seg := range segments {
		if seg.demoID != "" {
			if seg.code == nil {
				return nil, nil, fmt.Errorf("%s: demo %q requires a fenced code block immediately after", sourceFile, seg.demoID)
			}
			if err := validateFenceLang(sourceFile, seg.code.Language); err != nil {
				return nil, nil, err
			}
			blocks = append(blocks, DemoBlock{ID: seg.demoID, Code: *seg.code})
			demoIDs = append(demoIDs, seg.demoID)
			continue
		}
		if strings.TrimSpace(seg.text) == "" {
			continue
		}
		chunkBlocks, err := parseMarkdownChunk(sourceFile, seg.text, &templFenceIndex)
		if err != nil {
			return nil, nil, err
		}
		blocks = append(blocks, chunkBlocks...)
		if i == 0 && len(chunkBlocks) > 0 {
			// First segment may duplicate description; allowed.
		}
	}
	return blocks, demoIDs, nil
}

type bodySegment struct {
	text   string
	demoID string
	code   *CodeBlock
}

func splitDemoSegments(body string) []bodySegment {
	lines := strings.Split(body, "\n")
	var segments []bodySegment
	var textBuf strings.Builder
	var pendingDemo string

	flushText := func() {
		if textBuf.Len() > 0 {
			segments = append(segments, bodySegment{text: textBuf.String()})
			textBuf.Reset()
		}
	}

	i := 0
	for i < len(lines) {
		line := lines[i]
		trim := strings.TrimSpace(line)
		if m := demoDirectiveRe.FindStringSubmatch(trim); m != nil {
			flushText()
			pendingDemo = m[1]
			i++
			for i < len(lines) && strings.TrimSpace(lines[i]) == "" {
				i++
			}
			if i >= len(lines) || !strings.HasPrefix(strings.TrimSpace(lines[i]), "```") {
				segments = append(segments, bodySegment{demoID: pendingDemo})
				pendingDemo = ""
				continue
			}
			fence := strings.TrimSpace(lines[i])
			lang := strings.TrimPrefix(fence, "```")
			lang = strings.TrimSpace(lang)
			i++
			var codeLines []string
			for i < len(lines) {
				if strings.TrimSpace(lines[i]) == "```" {
					i++
					break
				}
				codeLines = append(codeLines, lines[i])
				i++
			}
			segments = append(segments, bodySegment{
				demoID: pendingDemo,
				code:   &CodeBlock{Language: lang, Source: strings.TrimRight(strings.Join(codeLines, "\n"), "\n")},
			})
			pendingDemo = ""
			continue
		}
		textBuf.WriteString(line)
		textBuf.WriteByte('\n')
		i++
	}
	flushText()
	return segments
}

func parseMarkdownChunk(sourceFile, chunk string, templFenceIndex *int) ([]Block, error) {
	md := goldmark.New(
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
	)
	src := []byte(strings.TrimSpace(chunk))
	if len(src) == 0 {
		return nil, nil
	}
	root := md.Parser().Parse(text.NewReader(src))
	var blocks []Block
	if err := ast.Walk(root, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		switch nn := n.(type) {
		case *ast.Heading:
			text := headingText(nn, src)
			idStr := slugHeading(text)
			if v, ok := nn.AttributeString("id"); ok {
				if s, ok := v.([]byte); ok {
					idStr = string(s)
				} else if s, ok := v.(string); ok {
					idStr = s
				}
			}
			blocks = append(blocks, HeadingBlock{Level: nn.Level, Text: text, ID: idStr})
			return ast.WalkSkipChildren, nil
		case *ast.Paragraph:
			if _, ok := n.Parent().(*ast.ListItem); ok {
				return ast.WalkContinue, nil
			}
			t := textFromChildren(nn, src)
			if strings.TrimSpace(t) != "" {
				blocks = append(blocks, ParagraphBlock{Text: t})
			}
			return ast.WalkSkipChildren, nil
		case *ast.List:
			if nn.Start == 0 {
				var items []string
				for child := nn.FirstChild(); child != nil; child = child.NextSibling() {
					if li, ok := child.(*ast.ListItem); ok {
						items = append(items, textFromChildren(li, src))
					}
				}
				if len(items) > 0 {
					blocks = append(blocks, ListBlock{Items: items})
				}
				return ast.WalkSkipChildren, nil
			}
		case *ast.FencedCodeBlock:
			lang := string(nn.Language(src))
			source := strings.TrimRight(string(nn.Text(src)), "\n")
			if err := validateFenceLang(sourceFile, lang); err != nil {
				return ast.WalkStop, err
			}
			lang = strings.ToLower(strings.TrimSpace(lang))
			if lang == "templ" {
				if err := ValidateTemplExample(sourceFile, source); err != nil {
					return ast.WalkStop, err
				}
				*templFenceIndex++
				idx := *templFenceIndex
				blocks = append(blocks, PreviewCodeBlock{
					ID:         previewCacheID(sourceFile, idx, source),
					Language:   lang,
					Source:     source,
					SourceFile: sourceFile,
					FenceIndex: idx,
				})
			} else {
				blocks = append(blocks, CodeBlock{Language: lang, Source: source})
			}
			return ast.WalkSkipChildren, nil
		case *ast.HTMLBlock, *ast.RawHTML:
			return ast.WalkStop, fmt.Errorf("%s: raw HTML is not allowed", sourceFile)
		}
		return ast.WalkContinue, nil
	}); err != nil {
		return nil, err
	}
	return blocks, nil
}

func headingText(h *ast.Heading, src []byte) string {
	var buf bytes.Buffer
	for child := h.FirstChild(); child != nil; child = child.NextSibling() {
		if t, ok := child.(*ast.Text); ok {
			buf.Write(t.Segment.Value(src))
		}
	}
	return strings.TrimSpace(buf.String())
}

func textFromChildren(n ast.Node, src []byte) string {
	var buf bytes.Buffer
	_ = ast.Walk(n, func(child ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			if t, ok := child.(*ast.Text); ok {
				buf.Write(t.Segment.Value(src))
			}
		}
		return ast.WalkContinue, nil
	})
	return strings.TrimSpace(buf.String())
}

func extractHeadings(blocks []Block) []Heading {
	var out []Heading
	for _, b := range blocks {
		if h, ok := b.(HeadingBlock); ok && h.Level >= 2 {
			out = append(out, Heading{Level: h.Level, Text: h.Text, ID: h.ID})
		}
	}
	return out
}

func slugHeading(text string) string {
	text = strings.ToLower(text)
	var b strings.Builder
	prevDash := false
	for _, r := range text {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
			prevDash = false
			continue
		}
		if !prevDash {
			b.WriteByte('-')
			prevDash = true
		}
	}
	return strings.Trim(b.String(), "-")
}
