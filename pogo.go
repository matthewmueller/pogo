package pogo

import (
	"fmt"

	"github.com/matthewmueller/pogo/internal/gofmt"
	"github.com/matthewmueller/pogo/internal/importer"
	"github.com/matthewmueller/pogo/internal/postgres"
	"github.com/matthewmueller/pogo/internal/schema"
	"github.com/matthewmueller/pogo/internal/sqlite"
	"github.com/matthewmueller/pogo/internal/vfs"
	"github.com/xo/dburl"
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
	Generate(imp *importer.Importer, schemas []string) (vfs.FileSystem, error)
}

// Generate function
func Generate(uri string, outdir string, schemas ...string) error {
	u, err := dburl.Parse(uri)
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
		dr, err := postgres.Open(u.DSN)
		if err != nil {
			return err
		}
		generator = dr
		// add a schema if we don't have one
		if len(ss) == 0 {
			ss = append(ss, "public")
		}
	case "", "sqlite", "sqlite3":
		dr, err := sqlite.Open(u.DSN)
		if err != nil {
			return err
		}
		generator = dr
	default:
		return fmt.Errorf("unsupported scheme: %s", u.Scheme)
	}

	// TODO: this is overkill, we just need to pass the path in
	imp := importer.New(outdir)

	// generate the virtual filesystem
	fs, err := generator.Generate(imp, ss)
	if err != nil {
		return err
	}

	// write to the filesystem
	if err := vfs.Write(fs, outdir); err != nil {
		return err
	}

	// format the code
	if err := gofmt.FormatAll(outdir); err != nil {
		return err
	}

	return nil
}
