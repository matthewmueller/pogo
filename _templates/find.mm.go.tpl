{{ $shortClass := shortname .Table.TableName }}
{{ $class := classnameMM .Table.TableName }}
{{ $shortModel := shortname .Table.TableName $shortClass }}
{{ $model := modelnameMM .Table.TableName }}
{{ $return := modelreturnMM .Table.TableName }}
package {{ .Package }}

// GENERATED BY POGO. DO NOT EDIT.

// Find a team by id
func ({{ $shortClass }} *{{ $class }}) Find({{ fkparams .ForeignKeys .Columns }}) ({{ $return }} {{ $model }}, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT {{ fields .Columns }}
    FROM {{ schema .Schema .Table.TableName }}
    WHERE {{ fkwhere .ForeignKeys }}`

	DBLog(sqlstr, {{ fklist .ForeignKeys }})
	row := {{ $shortClass }}.DB.QueryRow(sqlstr, {{ fklist .ForeignKeys }})
	err = row.Scan({{ gofields .Columns $return }})
	if err != nil {
		return {{ $return }}, err
	}

	return {{ $return }}, nil
}