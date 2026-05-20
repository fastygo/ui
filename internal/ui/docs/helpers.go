package docs

import "github.com/fastygo/ui/internal/registry"

func apiFieldName(f registry.APIField) string {
	if f.Required {
		return f.Name + " *"
	}
	return f.Name
}
