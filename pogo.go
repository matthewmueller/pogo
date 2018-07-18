package pogo

import (
	"github.com/matthewmueller/go-gen"

	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/paths"
	"github.com/matthewmueller/pogo/postgres"
	"github.com/matthewmueller/pogo/templates"
	"github.com/matthewmueller/pogo/util"
)

var pogoTemplate = gen.MustCompile("pogo", templates.MustAssetString("templates/pogo.gotmpl"))
var modelTemplate = gen.MustCompile("model", templates.MustAssetString("templates/model.gotmpl"))
var enumTemplate = gen.MustCompile("enum", templates.MustAssetString("templates/enum.gotmpl"))

// template data
type data map[string]interface{}

// Pogo struct
type Pogo struct {
	URL    string // database connection string
	Schema string // Schema to use
	Output string // Output path
}

// Run pogo's generator and write to the output
func (p *Pogo) Run() error {
	files, err := p.Generate()
	if err != nil {
		return err
	}

	if err := util.WriteFiles(files); err != nil {
		return err
	}

	return gen.FormatAll(p.Output)
}

// Generate pogo files
func (p *Pogo) Generate() (files map[string]string, err error) {
	files = make(map[string]string)

	cfg, err := pgx.ParseURI(p.URL)
	if err != nil {
		return files, err
	}

	conn, err := pgx.Connect(cfg)
	if err != nil {
		return files, err
	}
	defer conn.Close()

	path, err := paths.Load(p.Output)
	if err != nil {
		return files, err
	}

	schema, err := postgres.Introspect(conn, p.Schema)
	if err != nil {
		return files, err
	}

	relpath := path.Rel("pogo.go")
	files[relpath], err = pogoTemplate(data{
		"Path":   path,
		"Schema": schema,
	})
	if err != nil {
		return files, err
	}

	for _, table := range schema.Tables {
		path, err := path.New("./" + table.Slug())
		if err != nil {
			return files, err
		}

		relpath := path.Rel(table.Slug() + ".go")
		files[relpath], err = modelTemplate(data{
			"Path":   path,
			"Schema": schema,
			"Table":  table,
		})
		if err != nil {
			return files, err
		}
	}

	enumpath, err := path.New("./enum")
	if err != nil {
		return files, err
	}
	for _, enum := range schema.Enums {
		relpath := enumpath.Rel(enum.Slug() + ".go")
		files[relpath], err = enumTemplate(data{
			"Path":   enumpath,
			"Schema": schema,
			"Enum":   enum,
		})
		if err != nil {
			return files, err
		}
	}

	return files, nil
}

// // Config struct
// type Config struct {
// 	DB     database.Database
// 	Schema string
// 	Dir    string
// }

// // New pogo
// func New(db database.Database, schema, dir string) *Pogo {
// 	return &Pogo{&Config{
// 		DB:     db,
// 		Schema: schema,
// 		Dir:    dir,
// 	}}
// }

// // Pogo struct
// type Pogo struct {
// 	cfg *Config
// }

// // pogo templates
// var template = struct {
// 	Pogo  string
// 	Model string
// 	Many  string
// 	Enum  string
// }{
// 	Pogo:  string(templates.MustAsset("templates/pogo.gotmpl")),
// 	Model: string(templates.MustAsset("templates/model.gotmpl")),
// 	Many:  string(templates.MustAsset("templates/many.gotmpl")),
// 	Enum:  string(templates.MustAsset("templates/enum.gotmpl")),
// }

// // Run pogo
// func (p *Pogo) Run(ctx context.Context) (err error) {
// 	// files map
// 	files := map[string]string{}

// 	pkgname := text.Lower(text.Camel(filepath.Base(p.cfg.Dir)))

// 	// introspect the schema
// 	schema, err := p.cfg.DB.Introspect(p.cfg.Schema)
// 	if err != nil {
// 		return err
// 	}

// 	abspath, err := filepath.Abs(p.cfg.Dir)
// 	if err != nil {
// 		return err
// 	}

// 	importer, err := importer(abspath)
// 	if err != nil {
// 		return err
// 	}

// 	fns := gen.Functions{
// 		"import": importer,
// 	}

// 	// base file
// 	path := pkgname + ".go"
// 	files[path], err = gen.Compile("pogo.gotmpl", template.Pogo, fns, gen.Data{
// 		"Package": pkgname,
// 		"Schema":  schema,
// 	})
// 	if err != nil {
// 		return fmt.Errorf("error generating %s: %v", path, err)
// 	}

// 	// generate models for each table
// 	for _, table := range schema.Tables {
// 		tpl := template.Model
// 		if isManyToMany(table) {
// 			tpl = template.Many
// 		}

// 		data := gen.Data{
// 			"Package": pkgname,
// 			"Schema":  schema,
// 			"Table":   table,
// 		}

// 		// generate the model
// 		model := text.Lower(text.Camel(table.Model()))
// 		path := filepath.Join(model, model+".go")
// 		files[path], err = gen.Compile(path, tpl, data, fns)
// 		if err != nil {
// 			return fmt.Errorf("error generating %s: %v", path, err)
// 		}
// 	}

// 	// generate each enum
// 	for _, en := range schema.Enums {
// 		name := en.Name
// 		path := filepath.Join("enum", name+".go")
// 		files[path], err = gen.Compile("pogo.gotmpl", template.Enum, gen.Data{
// 			"Package": pkgname,
// 			"Schema":  schema,
// 			"Enum":    en,
// 		})
// 		if err != nil {
// 			return fmt.Errorf("error generating %s: %v", path, err)
// 		}
// 	}

// 	for path, code := range files {
// 		outpath := filepath.Join(p.cfg.Dir, path)
// 		outdir := filepath.Dir(outpath)

// 		if err := os.MkdirAll(outdir, 0755); err != nil {
// 			return err
// 		}

// 		if err := ioutil.WriteFile(outpath, []byte(code), 0644); err != nil {
// 			return err
// 		}
// 	}

// 	if err := gen.FormatAll(p.cfg.Dir); err != nil {
// 		return err
// 	}

// 	return nil
// }

// // Check if the relationship is many-to-many
// func isManyToMany(table *database.Table) bool {
// 	var pks []string

// 	for _, c := range table.Columns {
// 		if c.IsPrimaryKey {
// 			pks = append(pks, c.Name)
// 		}
// 	}
// 	if len(pks) > 1 {
// 		return true
// 	}

// 	// no primary keys but at least one unique foreign key pair
// 	for _, idx := range table.Indexes {
// 		if idx.IsPrimary || !idx.IsUnique {
// 			continue
// 		}

// 		if len(idx.Columns) >= 2 {
// 			return true
// 		}
// 	}

// 	return false
// }

// // importer fn
// func importer(abspath string) (func(...string) string, error) {
// 	gopath := build.Default.GOPATH
// 	importBase, err := filepath.Rel(filepath.Join(gopath, "src"), abspath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return func(s ...string) string {
// 		path := importBase
// 		for _, p := range s {
// 			path = filepath.Join(path, p)
// 		}
// 		return path
// 	}, nil
// }
