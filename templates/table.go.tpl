{{ $shortClass := shortname .Table.TableName }}
{{ $class := classname .Table.TableName }}
{{ $shortModel := shortname .Table.TableName $shortClass }}
{{ $model := modelname .Table.TableName }}
package model

// {{ classname .Table.TableName }} class
type {{ classname .Table.TableName }} struct {
	DB DB
}

// {{ modelname .Table.TableName }} model
type {{ modelname .Table.TableName }} struct {
  {{ range .Columns }}
  {{ field .ColumnName }} {{ fieldtype .DataType }} `json:"{{ .ColumnName }},omitempty"`{{ end }}
}

// New{{ modelname .Table.TableName }} model
func New{{ modelname .Table.TableName }}(db DB) {{ classname .Table.TableName }} {
	return {{ classname .Table.TableName }}{
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
