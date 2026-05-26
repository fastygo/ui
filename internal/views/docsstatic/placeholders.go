package docsstatic

import (
	"encoding/base64"
	"strings"
)

const (
	indexPlaceholderBadge  = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 240 160"><g fill="#000"><rect x="58" y="52" width="138" height="56" rx="18" opacity=".28"/><rect x="82" y="70" width="86" height="22" rx="11"/><circle cx="180" cy="81" r="7"/></g></svg>`
	indexPlaceholderBox    = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 240 160"><g fill="#000"><rect x="44" y="40" width="152" height="92" rx="18" opacity=".28"/><rect x="68" y="62" width="104" height="48" rx="12"/><rect x="82" y="118" width="76" height="8" rx="4" opacity=".55"/></g></svg>`
	indexPlaceholderButton = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 240 160"><g fill="#000"><rect x="52" y="54" width="140" height="52" rx="18" opacity=".28"/><rect x="76" y="72" width="92" height="20" rx="10"/></g></svg>`
	indexPlaceholderRadio  = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 240 160"><g fill="#000"><rect x="54" y="50" width="132" height="60" rx="18" opacity=".26"/><circle cx="82" cy="80" r="12"/><circle cx="82" cy="80" r="5" fill="#fff"/><rect x="104" y="72" width="66" height="16" rx="8"/></g></svg>`
	indexPlaceholderAlert  = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 240 160"><g fill="#000"><rect x="42" y="44" width="156" height="72" rx="18" opacity=".28"/><circle cx="70" cy="80" r="9"/><rect x="90" y="66" width="76" height="14" rx="7"/><rect x="90" y="88" width="104" height="10" rx="5" opacity=".7"/></g></svg>`
	indexPlaceholderCard   = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 240 160"><g fill="#000"><rect x="52" y="36" width="136" height="94" rx="18" opacity=".30"/><rect x="72" y="54" width="96" height="28" rx="10"/><rect x="72" y="94" width="78" height="10" rx="5"/><rect x="72" y="110" width="52" height="8" rx="4" opacity=".7"/></g></svg>`
	indexPlaceholderForm   = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 240 160"><g fill="#000"><rect x="46" y="34" width="148" height="96" rx="18" opacity=".28"/><rect x="68" y="54" width="104" height="14" rx="7"/><rect x="68" y="78" width="104" height="14" rx="7"/><rect x="68" y="102" width="68" height="14" rx="7"/><rect x="144" y="102" width="28" height="14" rx="7" opacity=".72"/></g></svg>`
)

var indexPlaceholderSVGBySuffix = map[string]string{
	"/components/alert/":  indexPlaceholderAlert,
	"/components/card/":   indexPlaceholderCard,
	"/components/form/":   indexPlaceholderForm,
	"/primitives/badge/":  indexPlaceholderBadge,
	"/primitives/box/":    indexPlaceholderBox,
	"/primitives/button/": indexPlaceholderButton,
	"/primitives/radio/":  indexPlaceholderRadio,
}

func indexCardPlaceholderStyle(href string) string {
	svg := indexCardPlaceholderSVG(href)
	if svg == "" {
		return ""
	}
	return `--docs-index-card-mask: url("data:image/svg+xml;base64,` + base64.StdEncoding.EncodeToString([]byte(svg)) + `");`
}

func indexCardPlaceholderSVG(href string) string {
	for suffix, svg := range indexPlaceholderSVGBySuffix {
		if strings.HasSuffix(href, suffix) {
			return svg
		}
	}
	return ""
}
