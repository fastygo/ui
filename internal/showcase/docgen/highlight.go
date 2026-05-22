package docgen

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/alecthomas/chroma/v2"
	htmlfmt "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

var htmlFormatter = htmlfmt.New(
	htmlfmt.WithClasses(true),
	htmlfmt.WithCSSComments(false),
)

// HighlightCodeBlocks fills HighlightedHTML on every code-bearing block.
func HighlightCodeBlocks(pages []DocPage) error {
	for pi := range pages {
		for bi, b := range pages[pi].Blocks {
			switch block := b.(type) {
			case CodeBlock:
				html, err := HighlightSource(block.Language, block.Source)
				if err != nil {
					return fmt.Errorf("%s: highlight code: %w", pages[pi].SourceFile, err)
				}
				block.HighlightedHTML = html
				pages[pi].Blocks[bi] = block
			case PreviewCodeBlock:
				html, err := HighlightSource(block.Language, block.Source)
				if err != nil {
					return fmt.Errorf("%s: highlight preview source: %w", pages[pi].SourceFile, err)
				}
				block.HighlightedHTML = html
				pages[pi].Blocks[bi] = block
			}
		}
	}
	return nil
}

// HighlightSource renders syntax-highlighted HTML for a fenced block.
func HighlightSource(language, source string) (string, error) {
	lexer := resolveLexer(language, source)
	tokens, err := chroma.Tokenise(lexer, nil, source)
	if err != nil {
		return "", err
	}
	tokens = normalizeTemplAtTokens(tokens)
	var buf bytes.Buffer
	if err := htmlFormatter.Format(&buf, styles.Get("modus-operandi"), chroma.Literator(tokens...)); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// normalizeTemplAtTokens reclassifies @ from Chroma's Go Error token (invalid in Go, valid in templ).
// chroma.Other maps to class "x" (no error background) in the HTML formatter.
func normalizeTemplAtTokens(tokens []chroma.Token) []chroma.Token {
	for i := range tokens {
		if tokens[i].Type == chroma.Error && tokens[i].Value == "@" {
			tokens[i].Type = chroma.Other
		}
	}
	return tokens
}

func resolveLexer(language, source string) chroma.Lexer {
	lang := highlightLanguage(language)
	if lang != "" {
		if l := lexers.Get(lang); l != nil {
			return chroma.Coalesce(l)
		}
	}
	if l := lexers.Analyse(source); l != nil {
		return chroma.Coalesce(l)
	}
	return chroma.Coalesce(lexers.Fallback)
}

func highlightLanguage(language string) string {
	switch strings.ToLower(strings.TrimSpace(language)) {
	case "templ", "gotemplate":
		return "go"
	case "shell", "sh":
		return "bash"
	case "ts", "typescript":
		return "typescript"
	case "js", "javascript":
		return "javascript"
	default:
		return strings.ToLower(strings.TrimSpace(language))
	}
}

// WriteChromaCSS writes Modus-role Chroma styles scoped to docs code blocks.
func WriteChromaCSS(w *bytes.Buffer) error {
	w.WriteString(modusTokenCSS)
	return nil
}

// Modus token roles mapped to --code-* variables in tweakcn.css (modus-operandi / modus-vivendi).
const modusTokenCSS = `/* Docs code blocks: Modus Operandi / Vivendi via --code-* tokens. */
.docs-code .chroma {
  margin: 0;
  background: transparent !important;
  color: var(--code-default);
  font-family: var(--font-mono);
  font-variant-ligatures: none;
  font-feature-settings: normal;
  -webkit-text-size-adjust: none;
}

.docs-code .chroma code {
  font-family: inherit;
  color: inherit;
}

.docs-code .chroma .line {
  display: flex;
}

.docs-code .chroma .hl {
  background-color: var(--code-highlight);
}

.docs-code .chroma .err {
  color: var(--code-error-fg);
  background-color: var(--code-error-bg);
}

.docs-code .chroma .x {
  color: var(--code-at);
  background-color: transparent;
}

.docs-code .chroma .c,
.docs-code .chroma .ch,
.docs-code .chroma .cm,
.docs-code .chroma .c1,
.docs-code .chroma .cs,
.docs-code .chroma .cp,
.docs-code .chroma .cpf,
.docs-code .chroma .gu {
  color: var(--code-comment);
}

.docs-code .chroma .k,
.docs-code .chroma .kd,
.docs-code .chroma .kr {
  color: var(--code-keyword);
}

.docs-code .chroma .kc,
.docs-code .chroma .kp {
  color: var(--code-keyword-const);
}

.docs-code .chroma .kt {
  color: var(--code-type);
}

.docs-code .chroma .kn {
  color: var(--code-keyword-op);
}

.docs-code .chroma .s,
.docs-code .chroma .sa,
.docs-code .chroma .sb,
.docs-code .chroma .sc,
.docs-code .chroma .dl,
.docs-code .chroma .sd,
.docs-code .chroma .s2,
.docs-code .chroma .sh,
.docs-code .chroma .si,
.docs-code .chroma .sx,
.docs-code .chroma .sr,
.docs-code .chroma .s1,
.docs-code .chroma .ss,
.docs-code .chroma .ld {
  color: var(--code-string);
}

.docs-code .chroma .m,
.docs-code .chroma .mb,
.docs-code .chroma .mf,
.docs-code .chroma .mh,
.docs-code .chroma .mi,
.docs-code .chroma .il,
.docs-code .chroma .mo,
.docs-code .chroma .l,
.docs-code .chroma .se {
  color: var(--code-number);
}

.docs-code .chroma .nf,
.docs-code .chroma .fm {
  color: var(--code-function);
}

.docs-code .chroma .na,
.docs-code .chroma .nc,
.docs-code .chroma .nd,
.docs-code .chroma .ne,
.docs-code .chroma .nx {
  color: var(--code-name);
}

.docs-code .chroma .nv,
.docs-code .chroma .vc,
.docs-code .chroma .vg,
.docs-code .chroma .vi,
.docs-code .chroma .vm {
  color: var(--code-variable);
}

.docs-code .chroma .no,
.docs-code .chroma .nb,
.docs-code .chroma .ni,
.docs-code .chroma .bp {
  color: var(--code-builtin);
}

.docs-code .chroma .nt {
  color: var(--code-tag);
}

.docs-code .chroma .o,
.docs-code .chroma .ow,
.docs-code .chroma .or {
  color: var(--code-operator);
}

.docs-code .chroma .p {
  color: var(--code-default);
}

.docs-code .chroma .gd {
  color: var(--code-diff-del);
}

.docs-code .chroma .gi {
  color: var(--code-diff-add);
}

.docs-code .chroma .ge {
  font-style: italic;
}

.docs-code .chroma .gs {
  font-weight: 600;
}
`
