package pogo

import (
	"bytes"
	"io"
	"io/ioutil"
	"os/exec"
	"text/template"

	"github.com/apex/log"
	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/templates"
	"github.com/pkg/errors"
)

// Field contains field information.
// type Field struct {
// 	Name    string
// 	Type    string
// 	NilType string
// 	Len     int
// 	Col     *postgres.Column
// 	Comment string
// }

// RelType represents the different types of relational storage (table/view).
// type RelType uint

// ModelData struct
// type ModelData struct {
// 	Package string
// 	Schema  string
// }

// // TemplateData is a template item for a table
// type TemplateData struct {
// 	Package     string
// 	Schema      string
// 	Table       *postgres.Table
// 	Columns     []*postgres.Column
// 	ForeignKeys []*postgres.ForeignKey
// 	Indexes     []*Index
// }

// Settings struct
type Settings struct {
	Address string
	Schema  string
	Package string
}

type templateData struct {
	Settings *Settings
	Schema   *Schema
	Table    *Table
	Enum     *Enum
}

// Generate the models
func Generate(db *pgx.Conn, settings *Settings) (files map[string]string, err error) {
	if settings.Schema == "" {
		settings.Schema = "public"
	}

	if settings.Package == "" {
		settings.Package = "pogo"
	}

	schema, err := introspect(db, settings.Schema)
	if err != nil {
		return files, err
	}

	files = make(map[string]string)

	// generate the entry file
	code, err := generate(settings.Package, templates.MustAsset("templates/pogo.go.tpl"), &templateData{
		Settings: settings,
		Schema:   schema,
	})
	if err != nil {
		return files, err
	}
	formatted, err := format(code)
	if err != nil {
		return files, err
	}
	files[settings.Package+".go"] = formatted

	// generate the codec file
	code, err = generate(settings.Package, templates.MustAsset("templates/codec.go.tpl"), &templateData{
		Settings: settings,
		Schema:   schema,
	})
	if err != nil {
		return files, err
	}
	formatted, err = format(code)
	if err != nil {
		return files, err
	}
	files["codec.go"] = formatted

	// build each model file from the tables
	for _, table := range schema.Tables {
		// if table.Name != "teams" {
		// 	continue
		// }

		// pick the template based on the type of relationship
		template := templates.MustAsset("templates/model.go.tpl")
		if isManyToMany(table) {
			template = templates.MustAsset("templates/model-many-to-many.go.tpl")
		}

		// generate a model for each table
		code, err := generate(table.Name, template, &templateData{
			Settings: settings,
			Schema:   schema,
			Table:    table,
		})
		if err != nil {
			return files, err
		}

		formatted, err := format(code)
		if err != nil {
			return files, err
		}

		files[table.Name+"/"+table.Name+".go"] = formatted
	}

	// build a test file for each model
	// for _, table := range schema.Tables {
	// 	// pick the template based on the type of relationship
	// 	template := templates.MustAsset("templates/model_test.go.tpl")
	// 	if isManyToMany(table) {
	// 		continue
	// 		// template = templates.MustAsset("templates/model-many-to-many.go.tpl")
	// 	}

	// 	// generate a test file for each table
	// 	code, err := generate(table.Name, template, &templateData{
	// 		Settings: settings,
	// 		Schema:   schema,
	// 		Table:    table,
	// 	})
	// 	if err != nil {
	// 		return files, err
	// 	}

	// 	formatted, err := format(code)
	// 	if err != nil {
	// 		return files, err
	// 	}

	// 	// fmt.Println(formatted)
	// 	files[table.Name+"_test.go"] = formatted
	// }

	// build each enum file
	for _, enum := range schema.Enums {
		// generate each enum enum
		code, err := generate(enum.Name, templates.MustAsset("templates/enum.go.tpl"), &templateData{
			Settings: settings,
			Schema:   schema,
			Enum:     enum,
		})
		if err != nil {
			return files, err
		}

		formatted, err := format(code)
		if err != nil {
			return files, err
		}

		files["enum/"+enum.Name+".go"] = formatted
	}

	return files, nil
}

// generate the model from a table
func generate(name string, raw []byte, data *templateData) (string, error) {
	tpl, err := template.New(name).Funcs(templateMap).Parse(string(raw))
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	if e := tpl.Execute(&b, data); e != nil {
		return "", e
	}

	return string(b.Bytes()), nil
}

// format the output using goimports
func format(input string) (output string, err error) {
	cmd := exec.Command("goimports")
	stdin, err := cmd.StdinPipe()

	if err != nil {
		return output, err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return output, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return output, err
	}

	reader := bytes.NewBufferString(input)

	if e := cmd.Start(); e != nil {
		return output, e
	}

	io.Copy(stdin, reader)
	stdin.Close()

	formatted, err := ioutil.ReadAll(stdout)
	if err != nil {
		return output, err
	}

	formattingError, err := ioutil.ReadAll(stderr)
	if err != nil {
		return output, err
	}

	stderr.Close()
	stdout.Close()

	if e := cmd.Wait(); e != nil {
		log.Infof("input %s", input)
		return output, errors.New(string(formattingError))
	}

	return string(formatted), nil
}

// Check if the relationship is many-to-many
func isManyToMany(table *Table) bool {
	hasPrimary := false
	for _, c := range table.Columns {
		if c.IsPrimaryKey {
			hasPrimary = true
			break
		}
	}
	if hasPrimary {
		return false
	}

	if len(table.ForeignKeys) == 2 {
		return true
	}

	return false
}
