package sqlite

import (
	"errors"
	"path/filepath"

	"github.com/matthewmueller/pogo/internal/template"
	"github.com/matthewmueller/pogo/internal/templates"
	"github.com/matthewmueller/pogo/internal/vfs"
)

var pogoT = template.MustCompile("pogo", templates.MustAssetString("internal/templates/pogo.gotmpl"))
var modelT = template.MustCompile("model", templates.MustAssetString("internal/templates/sq_model.gotmpl"))

// Generate the filesystem
func (s *DB) Generate(schemas []string) (vfs.FileSystem, error) {
	if len(schemas) == 0 {
		return nil, errors.New("schema not specified")
	}

	schema, err := s.Introspect(schemas[0])
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

	return vfs.Map(files), nil
}
