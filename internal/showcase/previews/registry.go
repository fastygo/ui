package previews

import (
	"fmt"
	"sync"

	"github.com/a-h/templ"
)

// Demo is a registered live preview for static documentation.
type Demo struct {
	ID       string
	Preview  templ.Component
	Patterns []string
}

var (
	mu    sync.RWMutex
	demos = map[string]Demo{}
)

// Register adds a demo preview. Panics on duplicate ID when using MustRegister.
func Register(d Demo) error {
	if d.ID == "" {
		return fmt.Errorf("previews: demo id is required")
	}
	if d.Preview == nil {
		return fmt.Errorf("previews: demo %q has nil preview", d.ID)
	}
	mu.Lock()
	defer mu.Unlock()
	if _, exists := demos[d.ID]; exists {
		return fmt.Errorf("previews: duplicate demo id %q", d.ID)
	}
	demos[d.ID] = d
	return nil
}

// MustRegister registers a demo or panics.
func MustRegister(d Demo) {
	if err := Register(d); err != nil {
		panic(err)
	}
}

// Get returns a registered demo.
func Get(id string) (Demo, bool) {
	mu.RLock()
	defer mu.RUnlock()
	d, ok := demos[id]
	return d, ok
}

// All returns all registered demos sorted by ID.
func All() []Demo {
	mu.RLock()
	defer mu.RUnlock()
	out := make([]Demo, 0, len(demos))
	for _, d := range demos {
		out = append(out, d)
	}
	return out
}

// Reset clears the registry (for tests).
func Reset() {
	mu.Lock()
	defer mu.Unlock()
	demos = map[string]Demo{}
}
