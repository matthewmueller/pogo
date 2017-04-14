package pogo

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"

	"github.com/matthewmueller/pogo/bin"
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

// ModelData struct
type ModelData struct {
	Package string
	Schema  string
}

// TemplateData is a template item for a table
type TemplateData struct {
	Package     string
	Schema      string
	Table       *postgres.Table
	Columns     []*postgres.Column
	ForeignKeys []*postgres.ForeignKey
	Indexes     []*Index
}

// EnumData is a template item for enums
type EnumData struct {
	Package string
	Schema  string
	Enum    *postgres.Enum
}

// Index data
type Index struct {
	Name      string
	Type      string
	IsUnique  bool
	IsPrimary bool
	Columns   []*postgres.IndexColumn
}

// Generate the database
func Generate(db db.DB, schema string, pkg string) (output map[string]string, err error) {
	// files, err := ioutil.ReadDir("./templates")
	// if err != nil {
	// 	return errors.Wrap(err, "unable to read the template directory")
	// }

	methods := []string{"table", "find", "insert", "delete", "update", "find-many", "upsert"}
	output = map[string]string{}
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
		return output, errors.Wrap(err, "unable to lookup tables")
	}

	enums, err := postgres.Enums(db, schema)
	if err != nil {
		return output, errors.Wrap(err, "unable to lookup enums")
	}

	coerce := NewCoerce(schema, enums)

	// First generate the model file
	buf, err := loadTemplate(templatePath("model", ""))
	if err != nil {
		return output, errors.Wrap(err, "unable to load a model template")
	}
	tpl, err := template.New(pkg + ".go").Funcs(TemplateFunctions(&coerce)).Parse(string(buf))
	if err != nil {
		return output, errors.Wrap(err, "could not parse the template for: "+pkg+".go")
	}
	data := ModelData{
		Schema:  schema,
		Package: pkg,
	}
	var b bytes.Buffer
	tpl.Execute(&b, data)
	by, err := ioutil.ReadAll(&b)
	if err != nil {
		return output, errors.Wrap(err, "could not read in buffer for: "+pkg+".go")
	}
	output[pkg+".go"] = string(by)

	for _, table := range tables {
		// if table.TableName != "reports" {
		// 	continue
		// }

		columns, err1 := postgres.Columns(db, schema, table.TableName)
		if err1 != nil {
			return output, errors.Wrap(err1, "unable to lookup columns")
		}

		fks, err1 := postgres.ForeignKeys(db, schema, table.TableName)
		if err1 != nil {
			return output, errors.Wrap(err1, "unable to lookup foreign keys")
		}

		indexes, err1 := postgres.Indexes(db, schema, table.TableName)
		if err != nil {
			return output, errors.Wrap(err1, "unable to lookup the indexes")
		}

		var indices []*Index
		for _, index := range indexes {

			cols, err2 := postgres.IndexColumns(db, schema, table.TableName, index.IndexName)
			if err2 != nil {
				return output, errors.Wrap(err2, "unable to get the columns")
			}

			indices = append(indices, &Index{
				Name:      index.IndexName,
				IsUnique:  index.IsUnique,
				IsPrimary: index.IsPrimary,
				Columns:   cols,
			})
		}

		data := TemplateData{
			Schema:      schema,
			Table:       table,
			Columns:     columns,
			ForeignKeys: fks,
			Package:     pkg,
			Indexes:     indices,
		}

		tableType := TableType(columns, fks)
		for _, method := range methods {
			// if method != "table" {
			// 	continue
			// }

			outputFile := path.Join(table.TableName + "." + method + ".go")
			if method == "table" {
				outputFile = path.Join(table.TableName + ".go")
			}

			buf, err1 := loadTemplate(templatePath(method, tableType), templatePath(method, ""))
			if err1 != nil {
				return output, errors.Wrap(err1, "unable to load a template")
			}

			tpl, err1 := template.New(outputFile).Funcs(TemplateFunctions(&coerce)).Parse(string(buf))
			if err1 != nil {
				return output, errors.Wrap(err1, "could not parse the template for: "+outputFile)
			}

			var b bytes.Buffer
			tpl.Execute(&b, data)
			bytes, err1 := ioutil.ReadAll(&b)
			if err1 != nil {
				return output, errors.Wrap(err1, "could not read in buffer for: "+outputFile)
			}

			// TODO: swap with: exec.Command("goimports", params...).Run()
			//       move to the end
			bytes, err = goimports.Process(outputFile, bytes, &goimports.Options{
				Comments: true,
			})
			if err != nil {
				return output, err
			}

			output[outputFile] = string(bytes)
		}
	}

	enumTpl, err := loadTemplate(templatePath("enum", ""))
	if err != nil {
		return output, errors.Wrap(err, "unable to load enum template")
	}

	for _, enum := range enums {
		outputFile := path.Join("enum." + strings.Replace(enum.Name, "_", "-", -1) + ".go")
		// enum.Name
		data := EnumData{
			Package: pkg,
			Schema:  schema,
			Enum:    enum,
		}

		tpl, err1 := template.New(outputFile).Funcs(TemplateFunctions(&coerce)).Parse(string(enumTpl))
		if err1 != nil {
			return output, errors.Wrap(err1, "could not parse the template for: "+outputFile)
		}

		var b bytes.Buffer
		tpl.Execute(&b, data)
		bytes, err1 := ioutil.ReadAll(&b)
		if err1 != nil {
			return output, errors.Wrap(err1, "could not read in buffer for: "+outputFile)
		}

		// TODO: swap with: exec.Command("goimports", params...).Run()
		//       move to the end
		// bytes, err = goimports.Process(outputFile, bytes, &goimports.Options{
		// 	Comments: true,
		// })
		// if err != nil {
		// 	return output, err
		// }

		output[outputFile] = string(bytes)
	}

	return output, nil
}

// Write out the models
func Write(models map[string]string, outpath string) (err error) {
	if _, err := os.Stat(outpath); os.IsNotExist(err) {
		err1 := os.MkdirAll(outpath, os.ModePerm)
		if err1 != nil {
			return err1
		}
	}

	for basepath, model := range models {
		filepath := path.Join(outpath, basepath)

		if _, err := os.Stat(filepath); err == nil {
			buf, err1 := ioutil.ReadFile(filepath)
			if err1 != nil {
				return err1
			}
			if !strings.Contains(string(buf), "// GENERATED BY POGO") {
				return errors.New("refusing to write over " + filepath + ". please rename this file")
			}
		}

		err := ioutil.WriteFile(filepath, []byte(model), os.ModePerm)
		if err != nil {
			return errors.Wrap(err, "unable to write out: "+filepath)
		}
	}

	// build goimports parameters, closing files
	params := []string{"-w"}
	for basepath := range models {
		params = append(params, path.Join(outpath, basepath))
	}

	// process written files with goimports
	return exec.Command("goimports", params...).Run()
}

// TableType get the table type
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
		bytes, err = bin.Asset(path)
		if err == nil {
			return bytes, nil
		}
	}
	return bytes, errors.New("no template exists")
}
