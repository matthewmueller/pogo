package pogo

import (
	"fmt"
	"net/url"

	gen "github.com/matthewmueller/go-gen"
	"github.com/matthewmueller/pogo/internal/postgres"
	"github.com/matthewmueller/pogo/internal/schema"
	"github.com/matthewmueller/pogo/internal/sqlite"
	"github.com/matthewmueller/pogo/internal/vfs"
)

// Driver interface
type Driver interface {
	Introspector
	Generator
}

// enforce existing drivers
var _ Driver = (*sqlite.DB)(nil)
var _ Driver = (*postgres.DB)(nil)

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

	var ss []string
	for _, schema := range schemas {
		if schema == "" {
			continue
		}
		ss = append(ss, schema)
	}

	var generator Generator

	// open the database
	switch u.Scheme {
	case "postgres":
		dr, err := postgres.Open(u.String())
		if err != nil {
			return err
		}
		generator = dr
		// add a schema if we don't have one
		if len(ss) == 0 {
			ss = append(ss, "public")
		}
	case "", "sqlite", "sqlite3":
		// no schema is just a filepath
		dr, err := sqlite.Open(u.String())
		if err != nil {
			return err
		}
		generator = dr
	default:
		return fmt.Errorf("unsupported scheme: %s", u.Scheme)
	}

	// generate the virtual filesystem
	fs, err := generator.Generate(ss)
	if err != nil {
		return err
	}

	// write to the filesystem
	if err := vfs.Write(fs, outdir); err != nil {
		return err
	}

	// format the code
	if err := gen.FormatAll(outdir); err != nil {
		return err
	}

	return nil
}
