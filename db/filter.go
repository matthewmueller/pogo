package db

import (
	"fmt"

	gen "github.com/matthewmueller/go-gen"
)

// Filter struct
type Filter struct {
	Name        string // column name
	DataType    string // column type (only for columns)
	FKReference string // fk name (only for fk references)
}

// Fields gets the filter fields
func (f *Filter) Fields(schema *Schema) (fields []*FilterField, err error) {
	// TODO: finish
	if f.FKReference != "" {
		return fields, nil
	}

	kind, err := coerceFilter(schema, f.DataType)
	if err != nil {
		return fields, err
	}

	switch kind {
	case "ID", "String":
		// field equals
		fields = append(fields, &FilterField{
			Name:        f.Name,
			Description: f.Name + " equals",
			DataType:    f.DataType,
			format:      fmt.Sprintf("%s = %%s", f.Name),
		})

		// field doesn't equal
		fields = append(fields, &FilterField{
			Name:        f.Name + "Not",
			Description: f.Name + " doesn't equal",
			DataType:    f.DataType,
			format:      fmt.Sprintf("%s != %%s", f.Name),
		})

		// field contains
		fields = append(fields, &FilterField{
			Name:        f.Name + "Contains",
			Description: f.Name + " contains",
			DataType:    f.DataType,
			format:      fmt.Sprintf("%s LIKE %%%%s%%", f.Name),
		})

		// field doesn't contain
		fields = append(fields, &FilterField{
			Name:        f.Name + "NotContains",
			Description: f.Name + " doesn't contain",
			DataType:    f.DataType,
			format:      fmt.Sprintf("%s NOT LIKE %%%%s%%", f.Name),
		})

		// field starts with
		fields = append(fields, &FilterField{
			Name:        f.Name + "StartsWith",
			Description: f.Name + " starts with",
			DataType:    f.DataType,
			format:      fmt.Sprintf("%s LIKE %%s%%", f.Name),
		})

		// field doesn't start with
		fields = append(fields, &FilterField{
			Name:        f.Name + "NotStartsWith",
			Description: f.Name + " doesn't start with",
			DataType:    f.DataType,
			format:      fmt.Sprintf("%s NOT LIKE %%s%%", f.Name),
		})

		// field ends with
		fields = append(fields, &FilterField{
			Name:        f.Name + "EndsWith",
			Description: f.Name + " ends with",
			DataType:    f.DataType,
			format:      fmt.Sprintf("%s LIKE %%%%s", f.Name),
		})

		// field doesn't end with
		fields = append(fields, &FilterField{
			Name:        f.Name + "NotEndsWith",
			Description: f.Name + " doesn't end with",
			DataType:    f.DataType,
			format:      fmt.Sprintf("%s NOT LIKE %%%%s", f.Name),
		})

		// field is less than
		fields = append(fields, &FilterField{
			Name:        f.Name + "Lt",
			Description: f.Name + " is less than",
			DataType:    f.DataType,
			format:      fmt.Sprintf("%s < %%s", f.Name),
		})

		// field is less than or equal
		fields = append(fields, &FilterField{
			Name:        f.Name + "Lte",
			Description: f.Name + " is less than or equal",
			DataType:    f.DataType,
			format:      fmt.Sprintf("%s <= %%s", f.Name),
		})

		// field is greater than
		fields = append(fields, &FilterField{
			Name:        f.Name + "Gt",
			Description: f.Name + " is greater than",
			DataType:    f.DataType,
			format:      fmt.Sprintf("%s > %%s", f.Name),
		})

		// field is greater than or equal
		fields = append(fields, &FilterField{
			Name:        f.Name + "Gte",
			Description: f.Name + " is greater than or equal",
			DataType:    f.DataType,
			format:      fmt.Sprintf("%s >= %%s", f.Name),
		})

		// field is in list
		fields = append(fields, &FilterField{
			Name:        f.Name + "In",
			Description: f.Name + " is in",
			DataType:    f.DataType + "[]",
			format:      fmt.Sprintf("%s IN (%%s)", f.Name),
		})

		// field is not in list
		fields = append(fields, &FilterField{
			Name:        f.Name + "NotIn",
			Description: f.Name + " is not in",
			DataType:    f.DataType + "[]",
			format:      fmt.Sprintf("%s NOT IN (%%s)", f.Name),
		})

	case "Int":
		// field equals
		fields = append(fields, &FilterField{
			Name:        f.Name,
			Description: f.Name + " equals",
			DataType:    f.DataType,
			format:      fmt.Sprintf("%s = %%s", f.Name),
		})

		// field doesn't equal
		fields = append(fields, &FilterField{
			Name:        f.Name + "Not",
			Description: f.Name + " doesn't equal",
			DataType:    f.DataType,
			format:      fmt.Sprintf("%s != %%s", f.Name),
		})

		// field is less than
		fields = append(fields, &FilterField{
			Name:        f.Name + "Lt",
			Description: f.Name + " is less than",
			DataType:    f.DataType,
			format:      fmt.Sprintf("%s < %%s", f.Name),
		})

		// field is less than or equal
		fields = append(fields, &FilterField{
			Name:        f.Name + "Lte",
			Description: f.Name + " is less than or equal",
			DataType:    f.DataType,
			format:      fmt.Sprintf("%s <= %%s", f.Name),
		})

		// field is greater than
		fields = append(fields, &FilterField{
			Name:        f.Name + "Gt",
			Description: f.Name + " is greater than",
			DataType:    f.DataType,
			format:      fmt.Sprintf("%s > %%s", f.Name),
		})

		// field is greater than or equal
		fields = append(fields, &FilterField{
			Name:        f.Name + "Gte",
			Description: f.Name + " is greater than or equal",
			DataType:    f.DataType,
			format:      fmt.Sprintf("%s >= %%s", f.Name),
		})

		// field is in list
		fields = append(fields, &FilterField{
			Name:        f.Name + "In",
			Description: f.Name + " is in",
			DataType:    f.DataType + "[]",
			format:      fmt.Sprintf("%s IN (%%s)", f.Name),
		})

		// field is not in list
		fields = append(fields, &FilterField{
			Name:        f.Name + "NotIn",
			Description: f.Name + " is not in",
			DataType:    f.DataType + "[]",
			format:      fmt.Sprintf("%s NOT IN (%%s)", f.Name),
		})

	case "Float":
		// field equals
		fields = append(fields, &FilterField{
			Name:     f.Name,
			DataType: f.DataType,
		})

		// field doesn't equal
		fields = append(fields, &FilterField{
			Name:     f.Name + "Not",
			DataType: f.DataType,
		})

		// field is less than
		fields = append(fields, &FilterField{
			Name:     f.Name + "Lt",
			DataType: f.DataType,
		})

		// field is less than or equal
		fields = append(fields, &FilterField{
			Name:     f.Name + "Lte",
			DataType: f.DataType,
		})

		// field is greater than
		fields = append(fields, &FilterField{
			Name:     f.Name + "Gt",
			DataType: f.DataType,
		})

		// field is greater than or equal
		fields = append(fields, &FilterField{
			Name:     f.Name + "Gte",
			DataType: f.DataType,
		})

		// field is in list
		fields = append(fields, &FilterField{
			Name:     f.Name + "In",
			DataType: f.DataType + "[]",
		})

		// field is not in list
		fields = append(fields, &FilterField{
			Name:     f.Name + "NotIn",
			DataType: f.DataType + "[]",
		})

	case "Boolean":
		// field equals
		fields = append(fields, &FilterField{
			Name:     f.Name,
			DataType: f.DataType,
		})

		// field doesn't equal
		fields = append(fields, &FilterField{
			Name:     f.Name + "Not",
			DataType: f.DataType,
		})

	case "DateTime":
		// field equals
		fields = append(fields, &FilterField{
			Name:     f.Name,
			DataType: f.DataType,
		})

		// field doesn't equal
		fields = append(fields, &FilterField{
			Name:     f.Name + "Not",
			DataType: f.DataType,
		})

		// field is in list
		fields = append(fields, &FilterField{
			Name:     f.Name + "In",
			DataType: f.DataType + "[]",
		})

		// field is not in list
		fields = append(fields, &FilterField{
			Name:     f.Name + "NotIn",
			DataType: f.DataType + "[]",
		})

		// field is less than
		fields = append(fields, &FilterField{
			Name:     f.Name + "Lt",
			DataType: f.DataType,
		})

		// field is less than or equal
		fields = append(fields, &FilterField{
			Name:     f.Name + "Lte",
			DataType: f.DataType,
		})

		// field is greater than
		fields = append(fields, &FilterField{
			Name:     f.Name + "Gt",
			DataType: f.DataType,
		})

		// field is greater than or equal
		fields = append(fields, &FilterField{
			Name:     f.Name + "Gte",
			DataType: f.DataType,
		})

	case "Enum":
		// field equals
		fields = append(fields, &FilterField{
			Name:     f.Name,
			DataType: f.DataType,
		})

		// field doesn't equal
		fields = append(fields, &FilterField{
			Name:     f.Name + "Not",
			DataType: f.DataType,
		})

		// field is in list
		fields = append(fields, &FilterField{
			Name:     f.Name + "In",
			DataType: f.DataType + "[]",
		})

		// field is not in list
		fields = append(fields, &FilterField{
			Name:     f.Name + "NotIn",
			DataType: f.DataType + "[]",
		})

	case "List":
		// field equals
		fields = append(fields, &FilterField{
			Name:     f.Name + "Contains",
			DataType: f.DataType,
		})

		// field equals
		fields = append(fields, &FilterField{
			Name:     f.Name + "ContainsEvery",
			DataType: f.DataType + "[]",
		})

		// field equals
		fields = append(fields, &FilterField{
			Name:     f.Name + "ContainsSome",
			DataType: f.DataType + "[]",
		})

	default:
		return fields, fmt.Errorf("filter fields: unknown type %s", kind)
	}

	return fields, nil
}

// FilterField struct
type FilterField struct {
	Name        string
	Description string
	DataType    string
	format      string
}

// Pascal case
func (f *FilterField) Pascal() string {
	return gen.Pascal(f.Name)
}

// Camel case
func (f *FilterField) Camel() string {
	return gen.Camel(f.Name)
}

// Type of filter
func (f *FilterField) Type(schema *Schema) (string, error) {
	return coerce(schema, f.DataType)
}

// Format returns the filter's condition sprintf format
func (f *FilterField) Format() string {
	return f.format
}
