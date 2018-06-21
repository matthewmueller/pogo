package pogo

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	gen "github.com/matthewmueller/go-gen"
	text "github.com/matthewmueller/go-text"
	"github.com/matthewmueller/pogo/database"
	"github.com/matthewmueller/pogo/templates"
)

// Config struct
type Config struct {
	DB     database.Database
	Schema string
	Dir    string
}

// New pogo
func New(db database.Database, schema, dir string) *Pogo {
	return &Pogo{&Config{
		DB:     db,
		Schema: schema,
		Dir:    dir,
	}}
}

// Pogo struct
type Pogo struct {
	cfg *Config
}

// pogo templates
var template = struct {
	Pogo  string
	Model string
	Many  string
	Enum  string
}{
	Pogo:  string(templates.MustAsset("templates/pogo.gotmpl")),
	Model: string(templates.MustAsset("templates/model.gotmpl")),
	Many:  string(templates.MustAsset("templates/many.gotmpl")),
	Enum:  string(templates.MustAsset("templates/enum.gotmpl")),
}

// Run pogo
func (p *Pogo) Run(ctx context.Context) (err error) {
	pkgname := text.Lower(text.Camel(filepath.Base(p.cfg.Dir)))

	// introspect the schema
	schema, err := p.cfg.DB.Introspect(p.cfg.Schema)
	if err != nil {
		return err
	}

	// files map
	files := map[string]string{}

	// base file
	path := pkgname + ".go"
	files[path], err = gen.Compile("pogo.gotmpl", template.Pogo, gen.Data{
		"Package": pkgname,
		"Schema":  schema,
	})
	if err != nil {
		return fmt.Errorf("error generating %s: %v", path, err)
	}

	// generate models for each table
	for _, table := range schema.Tables {
		if isManyToMany(table) {
			continue
		}

		// generate the model
		path := filepath.Join(table.Name, table.Name+".go")
		files[path], err = gen.Compile("pogo.gotmpl", template.Model, gen.Data{
			"Package": pkgname,
			"Schema":  schema,
			"Table":   table,
		})
		if err != nil {
			return fmt.Errorf("error generating %s: %v", path, err)
		}
	}

	// generate models for many-to-many tables
	for _, table := range schema.Tables {
		if !isManyToMany(table) {
			continue
		}

		// generate join model
		path := filepath.Join(table.Name, table.Name+".go")
		files[path], err = gen.Compile("pogo.gotmpl", template.Many, gen.Data{
			"Package": pkgname,
			"Schema":  schema,
			"Table":   table,
		})
		if err != nil {
			return fmt.Errorf("error generating %s: %v", path, err)
		}
	}

	// generate each enum
	for _, en := range schema.Enums {
		name := en.Name
		path := filepath.Join("enum", name+".go")
		files[path], err = gen.Compile("pogo.gotmpl", template.Enum, gen.Data{
			"Package": pkgname,
			"Schema":  schema,
			"Enum":    en,
		})
		if err != nil {
			return fmt.Errorf("error generating %s: %v", path, err)
		}
	}

	for path, code := range files {
		outpath := filepath.Join(p.cfg.Dir, path)
		outdir := filepath.Dir(outpath)

		if err := os.MkdirAll(outdir, 0755); err != nil {
			return err
		}

		if err := ioutil.WriteFile(outpath, []byte(code), 0644); err != nil {
			return err
		}
	}

	if err := gen.FormatAll(p.cfg.Dir); err != nil {
		return err
	}

	return nil
}

// Check if the relationship is many-to-many
func isManyToMany(table *database.Table) bool {
	var pks []string

	for _, c := range table.Columns {
		if c.IsPrimaryKey {
			pks = append(pks, c.Name)
		}
	}
	if len(pks) > 1 {
		return true
	}

	// no primary keys but at least one unique foreign key pair
	for _, idx := range table.Indexes {
		if idx.IsPrimary || !idx.IsUnique {
			continue
		}

		if len(idx.Columns) >= 2 {
			return true
		}
	}

	return false
}
