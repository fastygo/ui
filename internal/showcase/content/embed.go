package content

import "embed"

// FS holds localized documentation markdown sources.
//
//go:embed all:en all:ru
var FS embed.FS
