package views

import (
	"github.com/a-h/templ"
	"github.com/fastygo/templ/ui"
	templutils "github.com/fastygo/templ/utils"
	"github.com/fastygo/ui/internal/ui/layout"
	"github.com/fastygo/ui/internal/views/docsstatic"
)

type docsIllusLabIndexPreview struct {
	Label string
	Title string
	Desc  string
	Href  string
}

func docsIllusLabIndexPreviews() []docsIllusLabIndexPreview {
	return []docsIllusLabIndexPreview{
		{
			Label: "Short title",
			Title: "Button",
			Desc:  "Primary actions.",
			Href:  "/docs/primitives/button/",
		},
		{
			Label: "Long title",
			Title: "Alert Dialog",
			Desc:  "Blocking modal confirmation pattern.",
			Href:  "/docs/components/alert-dialog/",
		},
		{
			Label: "With description",
			Title: "Form",
			Desc:  "Field layout and actions.",
			Href:  "/docs/components/form/",
		},
	}
}

func docsIllusLabProductionIllusStandalone(href string) templ.Component {
	return docsstatic.IndexIllustrationComponent(href, docsstatic.IllustrationStandalone)
}

func docsIllusLabProductionIllusEmbedded(href string) templ.Component {
	return docsstatic.IndexIllustrationSpriteComponent(href, docsstatic.IllustrationEmbedded)
}

func docsIllusLabIndexCardClass(href string) string {
	class := "prose-docs docs-illus-lab-prose-card docs-index-card p-3"
	if illusClass := docsstatic.IndexCardIllustrationClass(href); illusClass != "" {
		class += " relative overflow-hidden " + illusClass
	}
	return class
}

func docsIllusLabThemeButtonProps(props layout.ThemeToggleProps) ui.ButtonProps {
	label := props.Label
	if label == "" {
		label = "Theme"
	}
	switchToDark := props.SwitchToDarkLabel
	if switchToDark == "" {
		switchToDark = "Switch to dark theme"
	}
	switchToLight := props.SwitchToLightLabel
	if switchToLight == "" {
		switchToLight = "Switch to light theme"
	}
	return ui.ButtonProps{
		ID:      "ui8kit-theme-toggle",
		Type:    "button",
		Variant: "unstyled",
		Class: templutils.Cn(
			"inline-flex h-8 w-8 items-center justify-center rounded-md bg-transparent p-0 text-muted-foreground transition-colors hover:bg-accent hover:text-accent-foreground",
		),
		Attrs: templutils.MergeAttrs(
			templ.Attributes{
				"data-switch-to-dark-label":  switchToDark,
				"data-switch-to-light-label": switchToLight,
				"title":                      label,
			},
			templutils.AriaLabel(label),
			templutils.AriaPressed(false),
		),
	}
}
