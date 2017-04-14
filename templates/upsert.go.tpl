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

  "github.com/matthewmueller/pgx"
)

// Upsert the {{ $model }} by the Primary Key
func ({{ $shortClass }} *{{ $class }}) Upsert({{ $shortModel }} *{{ $model }}, action string) ({{ $return }} {{ $model }}, err error) {
	fields := {{ $shortClass }}.getFields({{ $shortModel }})

	// prepare the slices
	c, i, v := querySlices(fields, 0)

  // determine on conflict action
  var upsertAction string
  if action == UpsertDoUpdate {
    upsertAction = `DO UPDATE SET (` + strings.Join(c, ", ") + `) = ( EXCLUDED.` + strings.Join(c, ", EXCLUDED.") + `)`
  } else if action == UpsertDoNothing {
    upsertAction = UpsertDoNothing
  } else {
    return {{ $return }}, errors.New("invalid upsert action")
  }

	// sql query
  sqlstr := `INSERT INTO {{ schema .Schema .Table.TableName }} (` + strings.Join(c, ", ") + `) ` +
	`VALUES (` + strings.Join(i, ", ") + `) ` +
  `ON CONFLICT ("{{ primaryname .Columns }}") ` +
  upsertAction + ` ` +
  `RETURNING {{ fields .Columns }}`

	// run query
  DBLog(sqlstr, v...)
	row := {{ $shortClass }}.DB.QueryRow(sqlstr, v...)
	err = row.Scan({{ gofields .Columns $return }})
	if err != nil && err != pgx.ErrNoRows {
	  return {{ $return }}, err
	}

	return {{ $return }}, nil
}

{{ range $idx := .Indexes }}
{{ if .IsUnique }}{{ if not .IsPrimary }}
// UpsertBy{{ indexmethod $idx }} find a {{ $model }}
func ({{ $shortClass }} *{{ $class }}) UpsertBy{{ indexmethod $idx }}({{ $shortModel }} *{{ $model }}, action string) ({{ $return }} {{ $model }}, err error) {
  fields := {{ $shortClass }}.getFields({{ $shortModel }})

	// prepare the slices
	c, i, v := querySlices(fields, 0)

  // determine on conflict action
  var upsertAction string
  if action == UpsertDoUpdate {
    upsertAction = `DO UPDATE SET (` + strings.Join(c, ", ") + `) = ( EXCLUDED.` + strings.Join(c, ", EXCLUDED.") + `)`
  } else if action == UpsertDoNothing {
    upsertAction = UpsertDoNothing
  } else {
    return {{ $return }}, errors.New("invalid upsert action")
  }

  // sql query
  sqlstr := `INSERT INTO {{ schema $.Schema $.Table.TableName }} (` + strings.Join(c, ", ") + `) ` +
	`VALUES (` + strings.Join(i, ", ") + `) ` +
  `ON CONFLICT ({{ indexparamlist $idx }}) ` +
  upsertAction + ` ` +
  `RETURNING {{ fields $.Columns }}`

	// run query
  DBLog(sqlstr, v...)
	row := {{ $shortClass }}.DB.QueryRow(sqlstr, v...)
	err = row.Scan({{ gofields $.Columns $return }})
	if err != nil && err != pgx.ErrNoRows {
	  return {{ $return }}, err
	}

	return {{ $return }}, nil
}
{{ end }}{{ end }}
{{ end }}
