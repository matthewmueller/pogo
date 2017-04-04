{{ $shortClass := shortname .Table.TableName }}
{{ $class := classname .Table.TableName }}
{{ $shortModel := shortname .Table.TableName $shortClass }}
{{ $model := modelname .Table.TableName }}
{{ $return := modelreturn .Table.TableName }}
package model

// Find a team by id
func ({{ $shortClass }} *{{ $class }}) Find({{ primaryname .Columns }} {{ primarytype .Columns }}) ({{ $return }} {{ $model }}, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT {{ fields .Columns }}
    FROM {{ schema .Schema .Table.TableName }}
    WHERE {{ primaryname .Columns }} = $1`

	XOLog(sqlstr, {{ primaryname .Columns }})
	row := {{ $shortClass }}.DB.QueryRow(sqlstr, {{ primaryname .Columns }})
	err = row.Scan({{ gofields .Columns $return }})
	if err != nil {
		return {{ $return }}, err
	}

	return {{ $return }}, nil
}
