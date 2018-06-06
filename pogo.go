package pogo

import (
	"context"
	"fmt"
	"path/filepath"

	text "github.com/matthewmueller/go-text"
	"github.com/matthewmueller/pogo/database"
	"github.com/matthewmueller/pogo/template"
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

	// generate ${pkgname}.go
	path := pkgname + ".go"
	files[path], err = template.Generate(&template.Base{
		Package: pkgname,
		Schema:  schema,
	})
	if err != nil {
		return fmt.Errorf("error generating %s: %v", path, err)
	}

	// generate each table
	for _, table := range schema.Tables {
		name := table.Name
		path := filepath.Join(name, name+".go")

		switch {
		case isManyToMany(table):
			files[path], err = template.Generate(&template.ManyToMany{
				Package: pkgname,
				Schema:  schema,
				Table:   table,
			})
		default:
			files[path], err = template.Generate(&template.Model{
				Package: pkgname,
				Schema:  schema,
				Table:   table,
			})
		}

		if err != nil {
			return fmt.Errorf("error generating %s: %v", path, err)
		}
	}

	// generate each enum
	for _, enum := range schema.Enums {
		name := enum.Name
		path := filepath.Join("name", name+".go")

		files[path], err = template.Generate(&template.Enum{
			Package: pkgname,
			Schema:  schema,
			Enum:    enum,
		})
		if err != nil {
			return fmt.Errorf("error generating %s: %v", path, err)
		}
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

	return len(pks) > 1
}

// // Generate the models
// func (p *Pogo) Generate(ctx context.Context) (files []*File, err error) {

// 	var f File
// 	f.Name = p.cfg.PkgName + ".go"
// 	f.Data, err = gen.Generate("core", `

// 	`,
// 		gen.Field("Schema", p.cfg.Schema),
// 		gen.Field("Package", p.cfg.PkgName),
// 	)
// 	if err != nil {
// 		return files, err
// 	}
// 	files = append(files, &f)

// 	return files, nil
// }
