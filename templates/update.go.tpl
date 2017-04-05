{{ $shortClass := shortname .Table.TableName }}
{{ $class := classname .Table.TableName }}
{{ $shortModel := shortname .Table.TableName $shortClass }}
{{ $model := modelname .Table.TableName }}
{{ $return := modelreturn .Table.TableName }}
package {{ .Package }}

// GENERATED BY POGO. DO NOT EDIT.

import (
	"errors"
	"strings"
)

// Update the {{ $model }} by the Primary Key
func ({{ $shortClass }} *{{ $class }}) Update({{ primaryname .Columns }} {{ primarytype .Columns }}, {{ $shortModel }} *{{ $model }}) ({{ $return }} {{ $model }}, err error) {
	fields := {{ $shortClass }}.getFields({{ $shortModel }})

	// first check if we have the primary key
	if {{ primaryname .Columns }} == nil {
		return {{ $return }}, errors.New(`primary key "{{ primaryname .Columns }}" must be non-nil`)
	}

	// don't update the primary key
	delete(fields, "{{ primaryname .Columns }}")

	// prepare the slices
	c, i, v := querySlices(fields, 1)

	// sql query
	sqlstr := `UPDATE {{ schema .Schema .Table.TableName }} SET (` +
		strings.Join(c, ", ") + `) = (` +
		strings.Join(i, ", ") + `)
		WHERE {{ primaryname .Columns }} = $1
		RETURNING {{ fields .Columns }}`

	// run query
	values := append([]interface{}{ {{ $shortModel }}.{{ primaryid .Columns }} }, v...)
	XOLog(sqlstr, values...)

	row := {{ $shortClass }}.DB.QueryRow(sqlstr, values...)
	err = row.Scan({{ gofields .Columns $return }})
	if err != nil {
		return {{ $return }}, err
	}

	return {{ $return }}, nil
}
