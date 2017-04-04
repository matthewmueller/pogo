{{ $shortClass := shortname .Table.TableName }}
{{ $class := classnameMM .Table.TableName }}
{{ $shortModel := shortname .Table.TableName $shortClass }}
{{ $model := modelnameMM .Table.TableName }}
{{ $return := modelreturnMM .Table.TableName }}
package model

// Delete the {{ $model }} from the database.
func ({{ $shortClass }} *{{ $class }}) Delete({{ fkparams .ForeignKeys .Columns }}) (err error) {
	// sql query
	const sqlstr = `
    DELETE FROM {{ schema .Schema .Table.TableName }}
    WHERE {{ fkwhere .ForeignKeys }}
  `

	// run query
	XOLog(sqlstr, {{ fklist .ForeignKeys }})
	_, err = {{ $shortClass }}.DB.Exec(sqlstr, {{ fklist .ForeignKeys }})
	if err != nil {
		return err
	}

	return nil
}
