package previews

import "github.com/fastygo/ui/internal/registry"

// RegisterFromRegistry registers all variant previews from the runtime docs registry.
// Used during migration so catalog showcase.go previews remain the single render source.
func RegisterFromRegistry() error {
	for _, page := range registry.AllPages() {
		for _, v := range page.Variants {
			id := page.Slug + "." + v.ID
			if err := Register(Demo{ID: id, Preview: v.Preview}); err != nil {
				return err
			}
		}
	}
	return nil
}
