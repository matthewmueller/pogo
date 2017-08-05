{{/*************************************************************************/}}
{{/* Variables */}}
{{/*************************************************************************/}}

{{ $t := tablename $.Schema .Table }}
{{ $tn := .Table.Name }}
{{ $c := $tn | capitalize }}
{{ $cv := $tn | capitalize | lower }}
{{ $m := $tn | singular }}
{{ $mv := $tn | singular | lower }}
{{ $mvg := print "&" $mv "." }}
{{ $f := $tn | singular | lower }}
{{ $p := primary .Table.Columns }}
{{ $pt := coerce $.Schema $p.DataType }}
{{ $co := colnames .Table.Columns }}
{{ $idxs := idxnames .Table.Indexes }}
{{ $cof := map $co (mprintf "\"%s\"") | join ", " }}
{{ $cog := map $co mcapitalize (mprefix $mvg) | join ", " }}

{{/*************************************************************************/}}
{{/* Our Package */}}
{{/*************************************************************************/}}

package {{ .Settings.Package }}

{{/*************************************************************************/}}
{{/* The table we'll attach our CRUD methods onto */}}
{{/*************************************************************************/}}

// {{ $c }} class
type {{ $c }} struct {
  db *DB
}

{{/*************************************************************************/}}
{{/* The model that contains all our database fields */}}
{{/*************************************************************************/}}

// {{ $m }} model
type {{ $m }} struct {
  {{ range .Table.Columns }}{{ $t := coerce $.Schema .DataType }}
  {{ .Name | capitalize }} {{ $t }} `json:"{{ .Name }},omitempty"` {{ if .Comment }}// {{ .Comment }}{{ end }}{{ end }}
}

{{/*************************************************************************/}}
{{/* Private class constructor, accessed via pogo.$TABLE */}}
{{/*************************************************************************/}}

// {{ $f }} constructor
func {{ $f }}(db *pgx.Conn) *{{ $c }} {
  return &{{ $c }}{db}
}

{{/*************************************************************************/}}
{{/* Private helper to get all the non-nil fields on our model */}}
{{/*************************************************************************/}}

// get all the non-nil fields
func fields({{ $mv }} *{{ $m }}) map[string]interface{} {
  fields := make(map[string]interface{})
  {{ range .Table.Columns }}{{ $field := .Name | capitalize }}
  if {{ $mv }}.{{ $field }} != nil {
    fields["{{ .Name }}"] = {{ $mv }}.{{ $field }}
  }{{ end }}
  
  return fields
}

{{/*************************************************************************/}}
{{/* pogo.$TABLE.Find(): find one row by it's primary key */}}
{{/*************************************************************************/}}

{{ if $p }}
// Find a {{ $mv }} by "{{ $p.Name }}"
func ({{ $cv }} *{{ $c }}) Find({{ $p.Name }} {{ $pt }}) ({{ $mv }} *{{ $m }}, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT {{ $cof }}
    FROM {{ $t }}
    WHERE "{{ $p.Name }}" = $1`

	Log(sqlstr, {{ $p.Name }})
	row := {{ $cv }}.db.QueryRow(sqlstr, {{ $p.Name }})
  err = row.Scan({{ $cog }})
	if err != nil {
		if err == pgx.ErrNoRows {
			return {{ $mv }},  Err{{ $m }}NotFound
		}
		return {{ $mv }}, err
	}

	return &{{ $mv }}, nil
}
{{ end }}

{{/*************************************************************************/}}
{{/* pogo.$TABLE.FindBy...(): find a row by its unique non-primary indexes */}}
{{/*************************************************************************/}}

{{ range $idx := .Table.Indexes }}
{{ $cols := idxcolnames $idx }}
{{ $idxmethod := map $cols mcapitalize | join "And" }}
{{ $idxparams := idxparams $.Schema $idx }}
{{ $indexvars := map $cols mcamelize | join ", " }}
// FindBy{{ $idxmethod }} find a {{ $mv }} by {{ $cols | join "` and `" | printf "`%s`"}}
func ({{ $cv }} *{{ $c }}) FindBy{{ $idxmethod }}({{ $idxparams }}) ({{ $mv }} {{ $m }}, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
		SELECT {{ $cof }}
		FROM {{ $t }}
		WHERE {{ idxwhere $idx }}`

	Log(sqlstr, {{ $indexvars }})
	row := {{ $cv }}.db.QueryRow(sqlstr, {{ $indexvars }})
	err = row.Scan({{ $cog }})
	if err != nil {
		if err == pgx.ErrNoRows {
			return {{ $mv }},  Err{{ $m }}NotFound
		}
		return {{ $mv }}, err
	}

	return {{ $mv }}, nil
}
{{ end }}

{{/*************************************************************************/}}
{{/* pogo.$TABLE.Insert(): insert a new row into the table */}}
{{/*************************************************************************/}}

// Insert a `{{ $mv }}` into the `{{ $t }}` table.
func ({{ $cv }} *{{ $c }}) Insert({{ $mv }} {{ $m }}) (*{{ $m }}, error) {
	// get all the non-nil fields and prepare them for the query
	_c, _i, _v := pogo.slice(fields({{ $mv }}), 0)

	// sql insert query, primary key provided by sequence
	sqlstr := `
	INSERT INTO {{ $t }} (` + strings.Join(_c, ", ") + `)
	VALUES (` + strings.Join(_i, ", ") + `)
	RETURNING {{ $cof }}`

	Log(sqlstr, _v...)
	row := {{ $cv }}.db.QueryRow(sqlstr, _v...)
	if e := row.Scan({{ $cog }}); e != nil {
    return nil, e
  }

	return &{{ $mv }}, nil
}

{{/*************************************************************************/}}
{{/* pogo.$TABLE.Update(): update an existing row in the table */}}
{{/*************************************************************************/}}

// Update a {{ $mv }} by its `{{ $p.Name }}`
func ({{ $cv }} *{{ $c }}) Update({{ $mv }} {{ $m }}, {{ $p.Name }} {{ $pt }}) (*{{ $m }}, error) {
	fieldset := fields({{ $mv }})

	// first check if we have the primary key
	if {{ $p.Name }} == nil {
		return nil, errors.New(`primary key "{{ $p.Name }}" must be non-nil`)
	}

	// don't update the primary key
	delete(fieldset, "{{ $p.Name }}")

	// prepare the slices
	_c, _i, _v := pogo.slice(fieldset, 1)

	// sql query
	sqlstr := `UPDATE {{ $t }} SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `)
		WHERE "{{ $p.Name }}" = $1
		RETURNING {{ $cof }}`

	// run query
	values := append([]interface{}{ {{ $p.Name }} }, _v...)
	Log(sqlstr, values...)

	row := {{ $cv }}.db.QueryRow(sqlstr, values...)
	if e := row.Scan({{ $cog }}); e != nil {
    return nil, e
  }

	return {{ $mv }}, nil
}

{{/*****************************************************************************/}}
{{/* pogo.$TABLE.UpdateBy...(): update a row by its unique non-primary indexes */}}
{{/*****************************************************************************/}}

{{ range $idx := .Table.Indexes }}
{{ $cols := idxcolnames $idx }}
{{ $idxmethod := map $cols mcapitalize | join "And" }}
{{ $idxparams := idxparams $.Schema $idx }}
// UpdateBy{{ $idxmethod }} find a {{ $m }}
func ({{ $cv }} *{{ $c }}) UpdateBy{{ $idxmethod }}({{ $mv }} {{ $m }}, {{ $idxparams }}) (*{{ $m }}, error) {
	fieldset := fields({{ $mv }})

	// first check if we have all the keys we need
	{{ range $idx.Columns }}if {{ .Name | camelize }} == nil {
		return nil, errors.New(`{{ .Name | camelize }} must be non-nil`)
	}
	{{ end }}

	// don't update the keys
	{{ range $idx.Columns }}delete(fieldset, "{{ .Name | camelize }}")
	{{ end }}

	// prepare the slices
	_c, _i, _v := pogo.slice(fieldset, {{ len $cols }})

	// sql query
	sqlstr := `UPDATE {{ $t }} SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `) ` +
		`WHERE {{ idxwhere $idx }} ` +
		`RETURNING {{ $cof }}`

	// run query
	values := []interface{}{}
	{{ range .Columns }}values = append(values, {{ .Name | camelize }})
	{{ end }}
	values = append(values, _v...)
	Log(sqlstr, values...)

	row := {{ $cv }}.DB.QueryRow(sqlstr, values...)
	err = row.Scan({{ $cog }})
	if err != nil {
		return {{ $mv }}, err
	}

	return {{ $mv }}, nil
}
{{ end }}

{{/*****************************************************************************/}}
{{/* pogo.$TABLE.Delete(): delete a row using its primary index */}}
{{/*****************************************************************************/}}

// Delete a `{{ $mv }}` from the `{{ $t }}` table
func ({{ $cv }} *{{ $c }}) Delete({{ $p.Name }} {{ $pt }}) error {
	// sql query
	sqlstr := `DELETE FROM {{ $t }} WHERE "{{ $p.Name }}" = $1`

	// run query
	Log(sqlstr, {{ $p.Name }})
	if _, e := {{ $cv }}.db.Exec(sqlstr, {{ $p.Name }}); e != nil {
    if e == pgx.ErrNoRows {
      return Err{{ $m }}NotFound
    }
    return e
  }

	return nil
}

{{/*****************************************************************************/}}
{{/* pogo.$TABLE.DeleteBy...(): delete a row by its unique non-primary indexes */}}
{{/*****************************************************************************/}}

{{ range $idx := .Table.Indexes }}
{{ $cols := idxcolnames $idx }}
{{ $idxmethod := map $cols mcapitalize | join "And" }}
{{ $idxparams := idxparams $.Schema $idx }}
{{ $indexvars := map $cols mcamelize | join ", " }}
// DeleteBy{{ $idxmethod }} find a {{ $m }}
func ({{ $cv }} *{{ $c }}) DeleteBy{{ $idxmethod }}({{ $idxparams }}) error {
	// sql delete query
	sqlstr := `DELETE FROM {{ $t }} WHERE {{ idxwhere $idx }}`

	Log(sqlstr, {{ $indexvars }})
	if _, err := {{ $cv }}.DB.Exec(sqlstr, {{ $indexvars }}); e != nil {
    if e == pgx.ErrNoRows {
      return Err{{ $m }}NotFound
    }
    return e
  }

	return nil
}
{{ end }}
