package pogo

import (
	"fmt"
	"net/url"

	"github.com/matthewmueller/pogo/internal/schema"
	"github.com/matthewmueller/pogo/internal/vfs"
)

// Driver interface
type Driver interface {
	Introspector
	Generator
}

// Introspector interface
type Introspector interface {
	Introspect(schemaName string) (*schema.Schema, error)
}

// Generator interface
type Generator interface {
	Generate(schemas []string) (vfs.FileSystem, error)
}

// Generate function
func Generate(uri string, outdir string, schemas ...string) error {
	u, err := url.Parse(uri)
	if err != nil {
		return err
	}
	switch u.Scheme {
	case "postgres":
		return fmt.Errorf("FINSIH ME")
	case "sqlite", "sqlite3":
		return fmt.Errorf("FINSIH ME")
	default:
		return fmt.Errorf("unsupported scheme: %s", u.Scheme)
	}
}
