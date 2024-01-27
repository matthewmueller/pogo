package sqlite

import (
	"path/filepath"

	"github.com/matthewmueller/pogo/internal/importer"
	"github.com/matthewmueller/pogo/internal/template"
	"github.com/matthewmueller/pogo/internal/templates"
	"github.com/matthewmueller/pogo/internal/vfs"
)

var pogoT = template.MustCompile("pogo", templates.MustAssetString("go_sq_pogo.gotext"))
var modelT = template.MustCompile("model", templates.MustAssetString("go_sq_model.gotext"))

// Generate the filesystem
func (s *DB) Generate(imp *importer.Importer, schemas []string) (vfs.FileSystem, error) {
	// TODO: remove this
	schema, err := s.Introspect("")
	if err != nil {
		return nil, err
	}

	files := make(map[string]string)
	files["pogo.go"], err = pogoT(template.Map{
		"Importer": imp,
		"Package":  "pogo",
		"Schema":   schema,
	})
	if err != nil {
		return nil, err
	}

	for _, table := range schema.Tables {
		slug := table.Slug()
		path := filepath.Join(slug, slug+".go")
		files[path], err = modelT(template.Map{
			"Importer": imp,
			"Package":  slug,
			"Schema":   schema,
			"Table":    table,
		})
		if err != nil {
			return nil, err
		}
	}

	return vfs.Map(files), nil
}
