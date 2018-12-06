package postgres

import (
	"errors"
	"path/filepath"

	"github.com/matthewmueller/pogo/internal/template"
	"github.com/matthewmueller/pogo/internal/templates"
	"github.com/matthewmueller/pogo/internal/vfs"
)

var pogoT = template.MustCompile("pogo", templates.MustAssetString("internal/templates/go_pg_pogo.gotext"))
var modelT = template.MustCompile("model", templates.MustAssetString("internal/templates/go_pg_model.gotext"))
var enumT = template.MustCompile("enum", templates.MustAssetString("internal/templates/go_pg_enum.gotext"))

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
	files["pogo.go"], err = pogoT(template.Map{
		"Package": "pogo",
		"Schema":  schema,
	})
	if err != nil {
		return nil, err
	}

	for _, table := range schema.Tables {
		slug := table.Slug()
		path := filepath.Join(slug, slug+".go")
		files[path], err = modelT(template.Map{
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
		files[path], err = enumT(template.Map{
			"Package": "enum",
			"Schema":  schema,
			"Enum":    enum,
		})
		if err != nil {
			return nil, err
		}
	}

	return vfs.Map(files), nil
}
