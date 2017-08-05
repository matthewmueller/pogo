{{ $shortClass := shortname .Table.TableName }}
{{ $class := classname .Table.TableName }}
{{ $shortModel := shortname .Table.TableName $shortClass }}
{{ $model := modelname .Table.TableName }}
{{ $singular := modelreturn .Table.TableName }}
{{ $return := pluralize $singular }}
package {{ .Package }}

import "github.com/matthewmueller/pgx"

// GENERATED BY POGO. DO NOT EDIT.

// FindMany find many {{ $model }}s by a condition
func ({{ $shortClass }} *{{ $class }}) FindMany(condition string, params... interface{}) ({{ $return }} []{{ $model }}, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT {{ fields .Columns }}
    FROM {{ schema .Schema .Table.TableName }}
    WHERE ` + condition

	DBLog(sqlstr, params...)
  rows, err := {{ $shortClass }}.DB.Query(sqlstr, params...)
  if err != nil {
    return {{ $return }}, err
  }
  defer rows.Close()

  for rows.Next() {
    {{ $singular }} := {{ $model }}{}
    err = rows.Scan({{ gofields .Columns $singular }})
    if err != nil {
      return {{ $return }}, err
    }
    {{ $return }} = append({{ $return }}, {{ $singular }})
  }

  if rows.Err() != nil {
    return {{ $return }}, rows.Err()
  }

	// ensure we return an empty array
	// rather than nil when we marshal
	if len({{ $return }}) == 0 {
		return make([]{{ $model }}, 0), nil
	}

  return {{ $return }}, nil
}