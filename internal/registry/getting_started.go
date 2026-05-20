package registry

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/fastygo/templ/ui"
)

func init() {
	registerGettingStarted()
}

func registerGettingStarted() {
	Register(Page{
		Slug:        "introduction",
		Title:       "Introduction",
		Section:     "getting-started",
		Description: "FastyGo UI is a wireframe-first component gallery on Go and templ. Pages show live previews, copy-pasteable templ snippets, and prop tables—similar to shadcn/ui, without shipping a monolithic UI package.",
		Source:      "github.com/fastygo/ui",
		Package:     "internal/site",
		Variants: []Variant{
			VariantFromFunc("overview", "Wireframe scope", "Structure, semantics, and accessibility come first; visual brand polish is a later phase.", introCode, introPreview),
		},
		Related: []RelatedLink{
			{Label: "Installation", Href: "/docs/installation"},
			{Label: "Components index", Href: "/docs"},
		},
	})
	Register(Page{
		Slug:        "installation",
		Title:       "Installation",
		Section:     "getting-started",
		Description: "Use this repository as the reference app, or copy individual components from the gallery into your project (future fastygo add workflow).",
		Source:      "github.com/fastygo/ui",
		Package:     "cmd/server",
		Variants: []Variant{
			VariantFromFunc("clone", "From git", "Clone the app, install Bun for CSS, run templ generate, then start the server.", installCode, installPreview),
		},
		Related: []RelatedLink{
			{Label: "Theming", Href: "/docs/theming"},
			{Label: "Button", Href: "/docs/components/button"},
		},
	})
	Register(Page{
		Slug:        "theming",
		Title:       "Theming",
		Section:     "getting-started",
		Description: "Semantic tokens (background, foreground, primary, …) come from tweakcn-style CSS variables. The header theme toggle switches light/dark via theme.js.",
		Source:      "web/static/css/tweakcn.css",
		Package:     "web/static/js/theme.js",
		Variants: []Variant{
			VariantFromFunc("tokens", "Semantic tokens", "Use token utilities such as bg-background and text-foreground in templ class strings.", themingCode, themingPreview),
		},
		Related: []RelatedLink{
			{Label: "Introduction", Href: "/docs/introduction"},
		},
	})
}

const introCode = `@ui.Stack(ui.StackProps{}) {
  @ui.Title(ui.TitleProps{Order: 1}, "Component gallery")
  @ui.Text(ui.TextProps{}, "Wireframe previews + API tables.")
}`

const installCode = `bun install
go mod download
bun run build:css
go tool templ generate ./...
bun run go`

const themingCode = `<body class="bg-background text-foreground">
  <!-- theme.js toggles .dark on html -->
</body>`

func introPreview(ctx context.Context, w io.Writer) error {
	if err := ui.Title(ui.TitleProps{Order: 2}, "Component gallery").Render(ctx, w); err != nil {
		return err
	}
	return ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "Wireframe previews and API tables.").Render(ctx, w)
}

func installPreview(ctx context.Context, w io.Writer) error {
	return ui.Text(ui.TextProps{Class: "text-sm font-mono text-muted-foreground"}, "bun run go → http://127.0.0.1:8080/docs").Render(ctx, w)
}

func themingPreview(ctx context.Context, w io.Writer) error {
	boxes := []templ.Component{
		ui.Box(ui.BoxProps{Class: "inline-block h-10 w-16 rounded-md border border-border bg-background"}),
		ui.Box(ui.BoxProps{Class: "inline-block h-10 w-16 rounded-md bg-primary"}),
		ui.Box(ui.BoxProps{Class: "inline-block h-10 w-16 rounded-md bg-muted"}),
	}
	for _, b := range boxes {
		if err := b.Render(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
