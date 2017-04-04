{{ $shortClass := shortname .Table.TableName }}
{{ $class := classname .Table.TableName }}
{{ $shortModel := shortname .Table.TableName $shortClass }}
{{ $model := modelname .Table.TableName }}
{{ $return := modelreturn .Table.TableName }}
package model

// Delete the {{ $model }} from the database.
func ({{ $shortClass }} *{{ $class }}) Delete({{ primaryname .Columns }} {{ primarytype .Columns }}) (err error) {
	// sql query
	const sqlstr = `
    DELETE FROM {{ schema .Schema .Table.TableName }}
    WHERE {{ primaryname .Columns }} = $1
  `

	// run query
	XOLog(sqlstr, {{ primaryname .Columns }})
	_, err = {{ $shortClass }}.DB.Exec(sqlstr, {{ primaryname .Columns }})
	if err != nil {
		return err
	}

	return nil
}
