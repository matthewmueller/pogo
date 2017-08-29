{{/*************************************************************************/}}
{{/* Variables */}}
{{/*************************************************************************/}}

{{ $pkg := .Settings.Package }}
{{ $t := tablename .Schema .Table }}
{{ $tn := .Table.Name }}
{{ $c := map (split "_" $tn) mcapitalize | join "" }}
{{ $cv := $c | lower }}
{{ $m := map (split "_" $tn) msingular | join "" }}
{{ $mv := $m | lower }}
{{ $co := colnames .Table.Columns }}
{{ $idxs := indexes .Table.Indexes }}
{{ $cof := map $co (mprintf "\"%s\"") | join ", " }}
{{ $cog := map $co mcapitalize (mprefix "&cols.") | join ", " }}
{{ $fkparams := fkparams .Schema .Table }}
{{ $fkwhere := fkwhere .Table.ForeignKeys }}
{{ $fks := fknames .Table.ForeignKeys }}
{{ $fklist := map $fks mcamelize | join ", " }}

{{/*************************************************************************/}}
{{/* Our Package */}}
{{/*************************************************************************/}}

package {{ $cv }}

{{/*************************************************************************/}}
{{/* Pogo marker */}}
{{/*************************************************************************/}}

// GENERATED BY POGO. DO NOT EDIT.

{{/*************************************************************************/}}
{{/* Errors */}}
{{/*************************************************************************/}}

// Err{{ $m }}NotFound returned if the {{ $mv }} is not found
var Err{{ $m }}NotFound = errors.New("{{ $mv }} not found")

{{/*************************************************************************/}}
{{/* All the columns in our table */}}
{{/*************************************************************************/}}

// columns in `{{ $t }}`
type columns struct {
  {{ range .Table.Columns }}{{ $t := coerce $.Schema .DataType }}
  {{ .Name | capitalize }} *{{ $t }} `json:"{{ .Name }},omitempty"` {{ if .Comment }}// {{ .Comment }}{{ end }}{{ end }}
}

{{/*************************************************************************/}}
{{/* This contains our fluent parameter container */}}
{{/*************************************************************************/}}

// {{ $m }} fluent API
type {{ $m }} struct {
	columns *columns
}

{{/*************************************************************************/}}
{{/* Helper to create the fluent API */}}
{{/*************************************************************************/}}

// New `{{ $t }}` API
func New() *{{ $m }} {
	return &{{ $m }}{&columns{}}
}

{{/*************************************************************************/}}
{{/* Generate each of the fluent methods for the fluent parameter API      */}}
{{/* This is very messy right now because we have some custom accessors    */}}
{{/* for uuid.                                                             */}}
{{/*                                                                       */}}
{{/* NOTE: This would probably be better solved at pgx level, but I spent  */}}
{{/* far too long trying to get that working with scanning nil *uuid.UUID  */}}
{{/* Fortunately, this isn't a big deal, because the API can remain stable */}}
{{/*************************************************************************/}}

{{ range .Table.Columns }}
{{- $nu := .Name | capitalize -}}
{{- $nc := .Name | camelize -}}
{{- $dt := coerceaccessor $.Schema .DataType -}}
// {{ $nu }} sets the `{{ .Name }}`
func ({{ $mv }} *{{ $m }}) {{ $nu }}({{ $nc }} {{ $dt }}) *{{ $m }} {
	{{ $mv }}.columns.{{ $nu }} = {{ decode $pkg $nc $dt }}
	return {{ $mv }}
}

// Get{{ $nu }} returns the `{{ .Name }}` if set
func ({{ $mv }} *{{ $m }}) Get{{ $nu }}() ({{ $nc }} *{{ $dt }}) {
	return {{ encode $pkg $mv $nu $dt }}
}
{{ end }}

{{/*************************************************************************/}}
{{/* Implement the Marshaler & Unmarshaler interfaces */}}
{{/*************************************************************************/}}

// MarshalJSON marshals the `{{ $mv }}` into JSON
func ({{ $mv }} *{{ $m }}) MarshalJSON() ([]byte, error) {
	return json.Marshal({{ $mv }}.columns)
}

// UnmarshalJSON unmarshals json to a `{{ $mv }}`
func ({{ $mv }} *{{ $m }}) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, {{ $mv }}.columns)
}

{{/*************************************************************************/}}
{{/* Implement the Stringer interface */}}
{{/*************************************************************************/}}

func ({{ $mv }} *{{ $m }}) String() string {
	return "{{ $mv }} TODO"
}

{{/*************************************************************************/}}
{{/* Private helper to get all the non-nil columns in our table */}}
{{/*************************************************************************/}}

// get all the non-nil columns
func getColumns({{ $mv }} *{{ $m }}) map[string]interface{} {
  columns := make(map[string]interface{})
  {{ range .Table.Columns }}{{ $col := .Name | capitalize }}
  if {{ $mv }}.columns.{{ $col }} != nil {
    columns["{{ .Name }}"] = {{ $mv }}.columns.{{ $col }}
  }{{ end }}
  
  return columns
}

{{/*************************************************************************/}}
{{/* pogo.$TABLE.Find(): find one row by it's primary key */}}
{{/*************************************************************************/}}

// Find a `{{ $mv }}` by its {{ $fks | join "`, `" | printf "`%s`"}}
func Find(db {{ $pkg }}.DB, {{ $fkparams }}) (*{{ $m }}, error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT {{ $cof }}
	FROM {{ $t }}
	WHERE {{ $fkwhere }}
	`
	{{ $pkg }}.Log(sqlstr, {{ $fklist }})

	// run the query
	var cols *columns
	row := db.QueryRow(sqlstr, {{ $fklist }})
	if e := row.Scan({{ $cog }}); e != nil {
    if e == pgx.ErrNoRows {
      return nil,  Err{{ $m }}NotFound
    }
    return nil, e
  }

	return &{{ $m }}{cols}, nil
}

{{/*************************************************************************/}}
{{/* pogo.$TABLE.Insert(): insert a new row into the table */}}
{{/*************************************************************************/}}

// Insert a `{{ $mv }}` into `{{ $t }}`
func Insert(db {{ $pkg }}.DB, {{ $mv }} *{{ $m }}) (*{{ $m }}, error) {
	// get all the non-nil fields and prepare them for the query
	_c, _i, _v := {{ $pkg }}.Slice(getColumns({{ $mv }}), 0)

	// sql insert query, primary key provided by sequence
	sqlstr := `
	INSERT INTO {{ $t }} (` + strings.Join(_c, ", ") + `)
	VALUES (` + strings.Join(_i, ", ") + `)
	RETURNING {{ $cof }}
  `
	{{ $pkg }}.Log(sqlstr, _v...)

	// run the query
	cols := &columns{}
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan({{ $cog }}); e != nil {
    return nil, e
  }

	return &{{ $m }}{cols}, nil
}

{{/*************************************************************************/}}
{{/* pogo.$TABLE.Update(): update an existing row in the table */}}
{{/*************************************************************************/}}

// Update a `{{ $m }}` by its {{ $fks | join "`, `" | printf "`%s`"}}
func Update(db {{ $pkg }}.DB, {{ $fkparams }}, {{ $mv }} *{{ $m }}) (*{{ $m }}, error) {
	fields := getColumns({{ $mv }})

	// first check if we have the foreign keys
  {{ range .Table.ForeignKeys }}
	if {{ .Name | camelize }} == nil {
    return nil, errors.New(`"{{ .Name | camelize }}" must be non-nil`)
  }
  {{- end }}

	// don't update the foreign keys
	{{ range .Table.ForeignKeys }}
	delete(fields, "{{ .Name }}")
  {{- end }}

	// prepare the slices
	_c, _i, _v := {{ $pkg }}.Slice(fields, {{ len .Table.ForeignKeys }})

	// sql query
	sqlstr := `UPDATE {{ $t }} SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `)
		WHERE {{ $fkwhere }}
		RETURNING {{ $cof }}`

	// setup the query
	values := []interface{}{}
	{{ range .Table.ForeignKeys }}values = append(values, {{ .Name | camelize }})
  {{ end }}
	values = append(values, _v...)
	{{ $pkg }}.Log(sqlstr, values...)

	// run the query
	var cols *columns
	row := db.QueryRow(sqlstr, values...)
	if e := row.Scan({{ $cog }}); e != nil {
    if e == pgx.ErrNoRows {
      return nil, Err{{ $m }}NotFound
    }
    return nil, e
  }

	return &{{ $m }}{cols}, nil
}

{{/*****************************************************************************/}}
{{/* pogo.$TABLE.Delete(): delete a row using its foreign keys */}}
{{/*****************************************************************************/}}

// Delete a `{{ $m }}` by its {{ $fks | join "`, `" | printf "`%s`"}}
func Delete(db {{ $pkg }}.DB, {{ $fkparams }}) error {
	// sql query
	const sqlstr = `
	DELETE FROM {{ $t }}
	WHERE {{ $fkwhere }}
	`
	{{ $pkg }}.Log(sqlstr, {{ $fklist }})

	// run query
	if _, e := db.Exec(sqlstr, {{ $fklist }}); e != nil {
    if e == pgx.ErrNoRows {
      return Err{{ $m }}NotFound
    }
    return e
  }

	return nil
}
