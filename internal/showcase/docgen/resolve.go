package docgen

import (
	"fmt"

	"github.com/fastygo/ui/internal/showcase/previews"
)

// ResolveDemos ensures every demo block on pages has a registered preview.
func ResolveDemos(pages []DocPage) error {
	for _, page := range pages {
		for _, id := range page.DemoIDs {
			if _, ok := previews.Get(id); !ok {
				return fmt.Errorf("%s: unknown demo id %q (not in preview registry)", page.SourceFile, id)
			}
		}
	}
	return nil
}
