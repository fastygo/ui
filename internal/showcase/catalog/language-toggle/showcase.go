package languagetoggle

import (
	"context"
	"io"

	"github.com/fastygo/framework/pkg/web/view"
	"github.com/fastygo/ui/internal/ui/components/toggles"
)

func previewDefault(ctx context.Context, w io.Writer) error {
	return toggles.LanguageToggle(view.LanguageToggleData{
		CurrentLabel: "EN", CurrentLocale: "en", NextLocale: "ru", NextHref: "/?lang=ru",
		DefaultLocale: "en", AvailableLocales: []string{"en", "ru"}, Label: "Switch language",
	}).Render(ctx, w)
}
