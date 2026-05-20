package fixtures

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed locale/*.json
var localeFS embed.FS

// DocNavLink is a sidebar or home-page link to static documentation.
type DocNavLink struct {
	Label string `json:"label"`
	Path  string `json:"path"`
	Icon  string `json:"icon,omitempty"`
}

// Locale holds embedded copy for one language.
type Locale struct {
	Brand string `json:"brand"`
	Nav   struct {
		Home        string `json:"home"`
		Sample      string `json:"sample"`
		DocsSection string `json:"docs_section"`
	} `json:"nav"`
	DocNav []DocNavLink `json:"doc_nav"`
	Theme struct {
		Label             string `json:"label"`
		SwitchToDarkLabel string `json:"switch_to_dark"`
		SwitchToLight     string `json:"switch_to_light"`
	} `json:"theme"`
	LanguageToggleLabel string `json:"language_toggle_label"`
	Dashboard           struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Body        string `json:"body"`
	} `json:"dashboard"`
	SampleStub struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Body        string `json:"body"`
	} `json:"sample_stub"`
	Docs struct {
		IndexTitle       string `json:"index_title"`
		IndexDescription string `json:"index_description"`
	} `json:"docs"`
}

// LoadLocale reads locale/{code}.json (e.g. en, ru).
func LoadLocale(code string) (Locale, error) {
	raw, err := localeFS.ReadFile("locale/" + code + ".json")
	if err != nil {
		return Locale{}, fmt.Errorf("fixtures: read locale %q: %w", code, err)
	}
	var out Locale
	if err := json.Unmarshal(raw, &out); err != nil {
		return Locale{}, fmt.Errorf("fixtures: parse locale %q: %w", code, err)
	}
	return out, nil
}
