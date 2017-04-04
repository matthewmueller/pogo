{{ $shortClass := shortname .Table.TableName }}
{{ $class := classnameMM .Table.TableName }}
{{ $shortModel := shortname .Table.TableName $shortClass }}
{{ $model := modelnameMM .Table.TableName }}
{{ $return := modelreturnMM .Table.TableName }}
package model

import (
	"errors"
	"strings"
)

// Update the {{ $model }} by the Primary Key
func ({{ $shortClass }} *{{ $class }}) Update({{ fkparams .ForeignKeys .Columns }}, {{ $shortModel }} *{{ $model }}) ({{ $return }} {{ $model }}, err error) {
	fields := {{ $shortClass }}.getFields({{ $shortModel }})

	// first check if we have the foreign keys
	{{ range .ForeignKeys }}
		if {{ .ColumnName }} == nil {
			return {{ $return }}, errors.New(`"{{ .ColumnName }}" must be non-nil`)
		}
	{{ end }}

	// don't update the foreign keys
	{{ range .ForeignKeys }}
		delete(fields, "{{ .ColumnName }}")
	{{ end }}

	// prepare the slices
	c, i, v := querySlices(fields, 1)

	// sql query
	sqlstr := `UPDATE {{ schema .Schema .Table.TableName }} SET (` +
		strings.Join(c, ", ") + `) = (` +
		strings.Join(i, ", ") + `)
		WHERE {{ fkwhere .ForeignKeys }}
		RETURNING {{ fields .Columns }}`

	// run query
	values := []interface{}{}
	{{ range .ForeignKeys }}
		values = append(values, {{ .ColumnName }})
	{{ end }}
	values = append(values, v...)
	XOLog(sqlstr, values...)

	row := {{ $shortClass }}.DB.QueryRow(sqlstr, values...)
	err = row.Scan({{ gofields .Columns $return }})
	if err != nil {
		return {{ $return }}, err
	}

	return {{ $return }}, nil
}
