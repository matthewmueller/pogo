package pogo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"text/template"

	"github.com/matthewmueller/pogo/db"
	"github.com/matthewmueller/pogo/postgres"
	"github.com/pkg/errors"
	"github.com/visualfc/gotools/goimports"
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
func init() {

}

// TemplateData is a template item for a table
type TemplateData struct {
	Schema      string
	Table       *postgres.Table
	Columns     []*postgres.Column
	ForeignKeys []*postgres.ForeignKey
}

// Generate the database
func Generate(db db.DB, schema string, outpath string) (err error) {
	// files, err := ioutil.ReadDir("./templates")
	// if err != nil {
	// 	return errors.Wrap(err, "unable to read the template directory")
	// }

	methods := []string{"table", "find", "insert", "delete", "update"}
	// templates := map[string][]byte{}
	// for _, typ := range types {
	// 	name := typ
	// 	// TODO: read a template with fallbacks
	// 	tpl, err1 := ioutil.ReadFile(path.Join("templates", name+".go.tpl"))
	// 	if err1 != nil {
	// 		return errors.Wrap(err1, "unable to read the template file: "+name)
	// 	}
	// 	base := strings.TrimSuffix(path.Base(name), ".go.tpl")
	// 	templates[base] = tpl
	// }

	tables, err := postgres.Tables(db, schema)
	if err != nil {
		return errors.Wrap(err, "unable to lookup tables")
	}

	output := map[string]string{}
	for _, table := range tables {
		// if table.TableName != "standups_teammates" {
		// 	continue
		// }

		columns, err1 := postgres.Columns(db, schema, table.TableName)
		if err1 != nil {
			return errors.Wrap(err1, "unable to lookup columns")
		}

		fks, err1 := postgres.ForeignKeys(db, schema, table.TableName)
		if err1 != nil {
			return errors.Wrap(err1, "unable to lookup foreign keys")
		}

		data := TemplateData{
			Schema:      schema,
			Table:       table,
			Columns:     columns,
			ForeignKeys: fks,
		}

		tableType := TableType(columns, fks)
		for _, method := range methods {
			// if method != "update" {
			// 	continue
			// }

			outputFile := path.Join(outpath, table.TableName+"."+method+".go")
			if method == "table" {
				outputFile = path.Join(outpath, table.TableName+".go")
			}

			buf, err := loadTemplate(templatePath(method, tableType), templatePath(method, ""))
			if err != nil {
				return errors.Wrap(err, "unable to load a template")
			}

			tpl, err1 := template.New(outputFile).Funcs(TemplateFunctions()).Parse(string(buf))
			if err1 != nil {
				return errors.Wrap(err1, "could not parse the template for: "+outputFile)
			}

			var b bytes.Buffer
			tpl.Execute(&b, data)
			bytes, err1 := ioutil.ReadAll(&b)
			if err1 != nil {
				return errors.Wrap(err1, "could not read in buffer for: "+outputFile)
			}

			// TODO: swap with: exec.Command("goimports", params...).Run()
			//       move to the end
			bytes, err = goimports.Process(outputFile, bytes, &goimports.Options{
				Comments: true,
			})
			if err != nil {
				return err
			}

			output[outputFile] = string(bytes)
		}
	}

	fmt.Println(output)
	return nil
}

// TableType
func TableType(columns []*postgres.Column, fks []*postgres.ForeignKey) string {
	hasPrimary := false
	for _, c := range columns {
		if c.IsPrimaryKey {
			hasPrimary = true
			break
		}
	}
	if hasPrimary {
		return ""
	}

	if len(fks) == 2 {
		return "mm"
	}

	return ""
}

func templatePath(basename string, typ string) string {
	if typ != "" {
		basename += "." + typ
	}
	return path.Join("templates", basename+".go.tpl")
}

func loadTemplate(paths ...string) (bytes []byte, err error) {
	for _, path := range paths {
		fmt.Println(path)
		if _, err := os.Stat(path); err == nil {
			return ioutil.ReadFile(path)
		}
	}
	return bytes, errors.New("no template exists")
}
