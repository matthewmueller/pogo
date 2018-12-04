package pogo

import (
	"fmt"
	"net/url"

	"github.com/matthewmueller/pogo/internal/schema"
	"golang.org/x/tools/godoc/vfs"
)

// var pogoTemplate = mustCompile("pogo", templates.MustAssetString("templates/pogo.gotmpl"))
// var pgModelTemplate = mustCompile("model", templates.MustAssetString("templates/pg_model.gotmpl"))
// var pgEnumTemplate = mustCompile("enum", templates.MustAssetString("templates/pg_enum.gotmpl"))
// var sqModelTemplate = mustCompile("model", templates.MustAssetString("templates/sq_model.gotmpl"))

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
func Generate(uri string, schemas []string) (vfs.FileSystem, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	switch u.Scheme {

	default:
		return nil, fmt.Errorf("unsupported scheme: %s", u.Scheme)
	}
}

func generatePG(u *url.URL, schemas ...string) (vfs.FileSystem, error) {
	return nil, nil
}

func generateSQLite(u *url.URL, schemas ...string) (vfs.FileSystem, error) {
	return nil, nil
}

// // template data
// type data map[string]interface{}

// // Pogo struct
// type Pogo struct {
// 	URL    string // database connection string
// 	Schema string // Schema to use
// 	Output string // Output path
// }

// // Run pogo's generator and write to the output
// func (p *Pogo) Run() error {
// 	files, err := p.Generate()
// 	if err != nil {
// 		return err
// 	}

// 	if err := util.WriteFiles(files); err != nil {
// 		return err
// 	}

// 	return gen.FormatAll(p.Output)
// }

// // Generate pogo files
// func (p *Pogo) Generate() (files map[string]string, err error) {
// 	files = make(map[string]string)

// 	cfg, err := pgx.ParseURI(p.URL)
// 	if err != nil {
// 		return files, err
// 	}

// 	conn, err := pgx.Connect(cfg)
// 	if err != nil {
// 		return files, err
// 	}
// 	defer conn.Close()

// 	path, err := paths.Load(p.Output)
// 	if err != nil {
// 		return files, err
// 	}

// 	schema, err := postgres.Introspect(conn, p.Schema)
// 	if err != nil {
// 		return files, err
// 	}

// 	relpath := path.Rel("pogo.go")
// 	files[relpath], err = pogoTemplate(data{
// 		"Path":   path,
// 		"Schema": schema,
// 	})
// 	if err != nil {
// 		return files, err
// 	}

// 	for _, table := range schema.Tables {
// 		path, err := path.New("./" + table.Slug())
// 		if err != nil {
// 			return files, err
// 		}

// 		relpath := path.Rel(table.Slug() + ".go")
// 		files[relpath], err = pgModelTemplate(data{
// 			"Path":   path,
// 			"Schema": schema,
// 			"Table":  table,
// 		})
// 		if err != nil {
// 			return files, err
// 		}
// 	}

// 	enumpath, err := path.New("./enum")
// 	if err != nil {
// 		return files, err
// 	}
// 	for _, enum := range schema.Enums {
// 		relpath := enumpath.Rel(enum.Slug() + ".go")
// 		files[relpath], err = pgEnumTemplate(data{
// 			"Path":   enumpath,
// 			"Schema": schema,
// 			"Enum":   enum,
// 		})
// 		if err != nil {
// 			return files, err
// 		}
// 	}

// 	return files, nil
// }
