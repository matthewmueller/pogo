package postgres

import (
	"errors"
	"path/filepath"

	gen "github.com/matthewmueller/go-gen"
	"github.com/matthewmueller/pogo/internal/template"
	"github.com/matthewmueller/pogo/internal/templates"
	"github.com/matthewmueller/pogo/internal/vfs"
)

var pogot = template.MustCompile("pogo", templates.MustAssetString("internal/templates/pogo.gotmpl"))
var modelt = template.MustCompile("model", templates.MustAssetString("internal/templates/pg_model.gotmpl"))
var enumt = template.MustCompile("enum", templates.MustAssetString("internal/templates/pg_enum.gotmpl"))

// Generate the database binding
func (d *DB) Generate(schemas []string) (vfs.FileSystem, error) {
	if len(schemas) == 0 {
		return nil, errors.New("schema not specified")
	}

	schema, err := d.Introspect(schemas[0])
	if err != nil {
		return nil, err
	}

	files := make(map[string]string)
	files["pogo.go"], err = pogot(template.Map{
		"Package": "pogo",
		"Schema":  schema,
	})
	if err != nil {
		return nil, err
	}

	for _, table := range schema.Tables {
		slug := table.Slug()
		path := filepath.Join(slug, slug+".go")
		files[path], err = modelt(template.Map{
			"Package": slug,
			"Schema":  schema,
			"Table":   table,
		})
		if err != nil {
			return nil, err
		}
	}

	for _, enum := range schema.Enums {
		slug := enum.Slug()
		path := filepath.Join("enum/", slug+".go")
		files[path], err = enumt(template.Map{
			"Package": slug,
			"Schema":  schema,
			"Enum":    enum,
		})
		if err != nil {
			return nil, err
		}
	}

	// format all the code
	for path, code := range files {
		formatted, err := gen.Format(code)
		if err != nil {
			return nil, err
		}
		files[path] = formatted
	}

	return vfs.Map(files), nil
}
