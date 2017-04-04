{{ $shortClass := shortname .Table.TableName }}
{{ $class := classnameMM .Table.TableName }}
{{ $shortModel := shortname .Table.TableName $shortClass }}
{{ $model := modelnameMM .Table.TableName }}
package model

// {{ $class }} class
type {{ $class }} struct {
	DB DB
}

// {{ $model }} model
type {{ $model }} struct {
  {{ range .Columns }}
  {{ field .ColumnName }} {{ fieldtype .DataType }} `json:"{{ .ColumnName }},omitempty"`{{ end }}
}

// New{{ $model }} model
func New{{ $model }}(db DB) {{ $class }} {
	return {{ $class }}{
		DB: db,
	}
}

// getFields fetch the non-nil fields
func ({{ $shortClass }} *{{ $class }}) getFields({{ $shortModel }} *{{ $model }}) map[string]interface{} {
	fields := map[string]interface{}{}

  {{ range .Columns }}
  if {{ $shortModel }}.{{ field .ColumnName }} != nil {
    fields["{{ .ColumnName }}"] = {{ $shortModel }}.{{ field .ColumnName }}
  }
  {{ end }}

	return fields
}
