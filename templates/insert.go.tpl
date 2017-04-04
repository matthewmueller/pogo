{{ $shortClass := shortname .Table.TableName }}
{{ $class := classname .Table.TableName }}
{{ $shortModel := shortname .Table.TableName $shortClass }}
{{ $model := modelname .Table.TableName }}
{{ $return := modelreturn .Table.TableName }}
package model

// Insert the {{ $model }} to the database.
func ({{ $shortClass }} *{{ $class }}) Insert({{ $shortModel }} *{{ $model }}) ({{ $return }} {{ $model }}, err error) {
	// get all the non-nil fields and prepare them for the query
	c, i, v := querySlices({{ $shortClass }}.getFields({{ $shortModel }}), 0)

	// sql insert query, primary key provided by sequence
	sqlstr := `
	INSERT INTO {{ schema .Schema .Table.TableName }} (` + strings.Join(c, ", ") + `)
	VALUES (` + strings.Join(i, ", ") + `)
	RETURNING {{ fields .Columns }}`

	XOLog(sqlstr, v...)
	row := m.DB.QueryRow(sqlstr, v...)
	err = row.Scan({{ gofields .Columns $return }})
	if err != nil {
	  return {{ $return }}, errors.Wrap(err, "could not insert into '{{ .Table.TableName }}'")
	}

	return {{ $return }}, nil
}
