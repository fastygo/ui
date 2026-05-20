package docs

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/fastygo/ui/internal/registry"
	_ "github.com/fastygo/ui/internal/showcase"
)

func TestComponentDoc_buttonRenders(t *testing.T) {
	page, ok := registry.PageBySlug("button")
	if !ok {
		t.Fatal("expected button page in registry")
	}
	var buf bytes.Buffer
	if err := ComponentDoc(page).Render(context.Background(), &buf); err != nil {
		t.Fatal(err)
	}
	html := buf.String()
	if !strings.Contains(html, "Button") {
		t.Fatal("expected button title in output")
	}
	if !strings.Contains(html, "API") {
		t.Fatal("expected API section")
	}
}
