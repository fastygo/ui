package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fastygo/ui/internal/showcase/specvalidate"
)

func main() {
	root := "internal/showcase/content/en"
	if len(os.Args) > 1 {
		root = os.Args[1]
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "validate-spec: %v\n", err)
		os.Exit(1)
	}
	paths, err := specvalidate.Discover(abs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "validate-spec: discover: %v\n", err)
		os.Exit(1)
	}
	if len(paths) == 0 {
		fmt.Println("validate-spec: no spec docs found")
		os.Exit(0)
	}
	var all []specvalidate.Error
	for _, path := range paths {
		doc, err := specvalidate.Load(path)
		if err != nil {
			rel, _ := filepath.Rel(abs, path)
			all = append(all, specvalidate.Error{File: rel, Message: err.Error()})
			continue
		}
		all = append(all, specvalidate.Validate(doc, abs)...)
	}
	if len(all) > 0 {
		fmt.Fprintf(os.Stderr, "validate-spec: %d error(s)\n", len(all))
		for _, e := range all {
			fmt.Fprintf(os.Stderr, "  %s\n", e.String())
		}
		os.Exit(1)
	}
	fmt.Printf("validate-spec: OK (%d docs under %s)\n", len(paths), abs)
}
