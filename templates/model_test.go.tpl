{{/*************************************************************************/}}
{{/* Variables */}}
{{/*************************************************************************/}}

{{ $pkg := .Settings.Package }}
{{ $t := tablename $.Schema .Table }}
{{ $tn := .Table.Name }}
{{ $c := $tn | capitalize }}
{{ $cv := $c | lower }}
{{ $m := $tn | singular }}
{{ $mv := $m | lower }}
{{ $mvg := print "&" $mv "." }}
{{ $p := primary .Table.Columns }}
{{ $pt := coerce $.Schema $p.DataType }}
{{ $co := colnames .Table.Columns }}
{{ $idxs := indexes .Table.Indexes }}
{{ $cof := map $co (mprintf "\"%s\"") | join ", " }}
{{ $cog := map $co mcapitalize (mprefix $mvg) | join ", " }}

{{/*************************************************************************/}}
{{/* Our Package */}}
{{/*************************************************************************/}}

package {{ $pkg }}_test

{{/*************************************************************************/}}
{{/* Pogo marker */}}
{{/*************************************************************************/}}

// GENERATED BY POGO. DO NOT EDIT.

{{/*************************************************************************/}}
{{/* Client factory */}}
{{/*************************************************************************/}}

func {{ $tn }}DB(t *testing.T) {{ $pkg }}.DB {
	config, err := pgx.ParseURI("{{ .Settings.Address }}")
	if err != nil {
		t.Fatal(err)
	}

	db, err := pgx.Connect(config)
	if err != nil {
		t.Fatal(err)
	}

	return db
}

{{/*************************************************************************/}}
{{/* Test Insert */}}
{{/*************************************************************************/}}

func Test{{ $c }}Insert(t *testing.T) {
	// setup the model
	model := {{ $pkg }}.New({{ $tn }}DB(t))

	// random values
	{{ range .Table.Columns }}_{{ .Name | capitalize | lower }} := {{ fake $.Settings $.Schema .Name .DataType }}
	{{ end }}

	// struct
	{{ $mv }}1 := {{ $pkg }}.{{ $m }}{
		{{ range .Table.Columns }}{{ .Name | capitalize }}: &_{{ .Name | capitalize | lower }},
    {{ end }}
	}

	{{ $mv }}2, err := model.{{ $m }}.Insert({{ $mv }}1)
	if err != nil {
		t.Fatal(err)
	}

	// assertions
  {{ range .Table.Columns }}assert.Equal(t, _{{ .Name | capitalize | lower }}, *{{ $mv }}2.{{ .Name | capitalize }})
  {{ end }}

	// cleanup
	if e := model.{{ $m }}.Delete(&_{{ $p.Name | capitalize | lower }}); e != nil {
		t.Fatal(e)
	}
}

{{/*************************************************************************/}}
{{/* Test Update */}}
{{/*************************************************************************/}}

func Test{{ $c }}Update(t *testing.T) {
	// setup the model
	model := {{ $pkg }}.New({{ $tn }}DB(t))

	// random values
	{{ range .Table.Columns }}_{{ .Name | capitalize | lower }} := {{ fake $.Settings $.Schema .Name .DataType }}
	{{ end }}

	// struct
	{{ $mv }}1 := {{ $pkg }}.{{ $m }}{
		{{ range .Table.Columns }}{{ .Name | capitalize }}: &_{{ .Name | capitalize | lower }},
    {{ end }}
	}

	{{ $mv }}2, err := model.{{ $m }}.Insert({{ $mv }}1)
	if err != nil {
		t.Fatal(err)
	}

  // random values
	{{ range .Table.Columns }}_{{ .Name | capitalize | lower }}2 := {{ fake $.Settings $.Schema .Name .DataType }}
	{{ end }}

  // random values
	{{ range .Table.Columns }}{{ $mv }}2.{{ .Name | capitalize }} = &_{{ .Name | capitalize | lower }}2
	{{ end }}

  {{ $mv }}3, err := model.{{ $m }}.Update(*{{ $mv }}2, &_{{ $p.Name | capitalize | lower }})
	if err != nil {
		t.Fatal(err)
	}

	// assertions
  {{ range .Table.Columns }}assert.Equal(t, _{{ .Name | capitalize | lower }}2, *{{ $mv }}3.{{ .Name | capitalize }})
  {{ end }}

	// cleanup
	if e := model.{{ $m }}.Delete(&_{{ $p.Name | capitalize | lower }}); e != nil {
		t.Fatal(e)
	}
}