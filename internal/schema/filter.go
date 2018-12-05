package schema

import (
	"fmt"

	gen "github.com/matthewmueller/go-gen"
)

func newFilter(
	name string,
	dataType DataType,
	notNull bool,
) *Filter {
	return &Filter{
		name,
		dataType,
		notNull,
	}
}

// Filter struct
type Filter struct {
	name     string   // column name
	dataType DataType // column type (only for columns)
	notNull  bool     // column can't be null
}

// // FilterType string
// type FilterType string

// // Field Types
// var (
// 	Null       FilterType = "NULL"
// 	ID         FilterType = "ID"
// 	String     FilterType = "String"
// 	Int        FilterType = "Int"
// 	Float      FilterType = "Float"
// 	Enumerable FilterType = "Enum"
// 	Boolean    FilterType = "Boolean"
// 	DateTime   FilterType = "DateTime"
// 	List       FilterType = "List"
// 	JSON       FilterType = "JSON"
// )

// Fields gets the filter fields based on the type
func (f *Filter) Fields() (fields []*FilterField, err error) {
	switch f.dataType.(type) {
	case *Null:
		// add no filters
	case *JSON:
		// TODO: add filters
	case *String:
		// field equals
		fields = append(fields, &FilterField{
			name:        f.name,
			dataType:    f.dataType,
			description: f.name + " equals",
			format:      fmt.Sprintf(`"%s" = %%s`, f.name),
		})

		// field doesn't equal
		fields = append(fields, &FilterField{
			name:        f.name + "Not",
			dataType:    f.dataType,
			description: f.name + " doesn't equal",
			format:      fmt.Sprintf(`"%s" != %%s`, f.name),
		})

		// field contains
		fields = append(fields, &FilterField{
			name:        f.name + "Contains",
			dataType:    f.dataType,
			description: f.name + " contains",
			format:      fmt.Sprintf(`"%s" LIKE '%%%%' || %%s || '%%%%'`, f.name),
		})

		// field doesn't contain
		fields = append(fields, &FilterField{
			name:        f.name + "NotContains",
			dataType:    f.dataType,
			description: f.name + " doesn't contain",
			format:      fmt.Sprintf(`"%s" NOT LIKE '%%%%' || %%s || '%%%%'`, f.name),
		})

		// field starts with
		fields = append(fields, &FilterField{
			name:        f.name + "StartsWith",
			dataType:    f.dataType,
			description: f.name + " starts with",
			format:      fmt.Sprintf(`"%s" LIKE %%s || '%%%%'`, f.name),
		})

		// field doesn't start with
		fields = append(fields, &FilterField{
			name:        f.name + "NotStartsWith",
			dataType:    f.dataType,
			description: f.name + " doesn't start with",
			format:      fmt.Sprintf(`"%s" NOT LIKE %%s || '%%%%'`, f.name),
		})

		// field ends with
		fields = append(fields, &FilterField{
			name:        f.name + "EndsWith",
			dataType:    f.dataType,
			description: f.name + " ends with",
			format:      fmt.Sprintf(`"%s" LIKE '%%%%' || %%s`, f.name),
		})

		// field doesn't end with
		fields = append(fields, &FilterField{
			name:        f.name + "NotEndsWith",
			dataType:    f.dataType,
			description: f.name + " doesn't end with",
			format:      fmt.Sprintf(`"%s" NOT LIKE '%%%%' || %%s`, f.name),
		})

		// field is less than
		fields = append(fields, &FilterField{
			name:        f.name + "Lt",
			dataType:    f.dataType,
			description: f.name + " is less than",
			format:      fmt.Sprintf(`"%s" < %%s`, f.name),
		})

		// field is less than or equal
		fields = append(fields, &FilterField{
			name:        f.name + "Lte",
			dataType:    f.dataType,
			description: f.name + " is less than or equal",
			format:      fmt.Sprintf(`"%s" <= %%s`, f.name),
		})

		// field is greater than
		fields = append(fields, &FilterField{
			name:        f.name + "Gt",
			dataType:    f.dataType,
			description: f.name + " is greater than",
			format:      fmt.Sprintf(`"%s" > %%s`, f.name),
		})

		// field is greater than or equal
		fields = append(fields, &FilterField{
			name:        f.name + "Gte",
			dataType:    f.dataType,
			description: f.name + " is greater than or equal",
			format:      fmt.Sprintf(`"%s" >= %%s`, f.name),
		})

		// field is in list
		fields = append(fields, &FilterField{
			name:        f.name + "In",
			dataType:    f.dataType,
			description: f.name + " is in",
			format:      fmt.Sprintf(`"%s" IN (%%s)`, f.name),
			spread:      `, `,
		})

		// field is not in list
		fields = append(fields, &FilterField{
			name:        f.name + "NotIn",
			dataType:    f.dataType,
			description: f.name + " is not in",
			format:      fmt.Sprintf(`"%s" NOT IN (%%s)`, f.name),
			spread:      `, `,
		})

		// if nullable
		if !f.notNull {
			fields = append(fields, &FilterField{
				name:        "nullable_" + f.name,
				dataType:    f.dataType,
				description: "nullable " + f.name + " equals",
				nullable:    true,
				format:      fmt.Sprintf(`"%s" = %%s`, f.name),
				nullformat:  fmt.Sprintf(`"%s" IS NULL`, f.name),
			})

			fields = append(fields, &FilterField{
				name:        "nullable_" + f.name + "_not",
				dataType:    f.dataType,
				description: "nullable " + f.name + " is not equal",
				nullable:    true,
				format:      fmt.Sprintf(`"%s" != %%s`, f.name),
				nullformat:  fmt.Sprintf(`"%s" IS NOT NULL`, f.name),
			})
		}

	case *Integer:
		// field equals
		fields = append(fields, &FilterField{
			name:        f.name,
			dataType:    f.dataType,
			description: f.name + " equals",
			format:      fmt.Sprintf(`"%s" = %%s`, f.name),
		})

		// field doesn't equal
		fields = append(fields, &FilterField{
			name:        f.name + "Not",
			dataType:    f.dataType,
			description: f.name + " doesn't equal",
			format:      fmt.Sprintf(`"%s" != %%s`, f.name),
		})

		// field is less than
		fields = append(fields, &FilterField{
			name:        f.name + "Lt",
			dataType:    f.dataType,
			description: f.name + " is less than",
			format:      fmt.Sprintf(`"%s" < %%s`, f.name),
		})

		// field is less than or equal
		fields = append(fields, &FilterField{
			name:        f.name + "Lte",
			dataType:    f.dataType,
			description: f.name + " is less than or equal",
			format:      fmt.Sprintf(`"%s" <= %%s`, f.name),
		})

		// field is greater than
		fields = append(fields, &FilterField{
			name:        f.name + "Gt",
			dataType:    f.dataType,
			description: f.name + " is greater than",
			format:      fmt.Sprintf(`"%s" > %%s`, f.name),
		})

		// field is greater than or equal
		fields = append(fields, &FilterField{
			name:        f.name + "Gte",
			dataType:    f.dataType,
			description: f.name + " is greater than or equal",
			format:      fmt.Sprintf(`"%s" >= %%s`, f.name),
		})

		// field is in list
		fields = append(fields, &FilterField{
			name:        f.name + "In",
			dataType:    f.dataType,
			description: f.name + " is in",
			format:      fmt.Sprintf(`"%s" IN (%%s)`, f.name),
			spread:      `, `,
		})

		// field is not in list
		fields = append(fields, &FilterField{
			name:        f.name + "NotIn",
			dataType:    f.dataType,
			description: f.name + " is not in",
			format:      fmt.Sprintf(`"%s" NOT IN (%%s)`, f.name),
			spread:      `, `,
		})

		// if nullable
		if !f.notNull {
			fields = append(fields, &FilterField{
				name:        "nullable_" + f.name,
				dataType:    f.dataType,
				description: "nullable " + f.name + " equals",
				nullable:    true,
				format:      fmt.Sprintf(`"%s" = %%s`, f.name),
				nullformat:  fmt.Sprintf(`"%s" IS NULL`, f.name),
			})

			fields = append(fields, &FilterField{
				name:        "nullable_" + f.name + "_not",
				dataType:    f.dataType,
				description: "nullable " + f.name + " is not equal",
				nullable:    true,
				format:      fmt.Sprintf(`"%s" != %%s`, f.name),
				nullformat:  fmt.Sprintf(`"%s" IS NOT NULL`, f.name),
			})
		}

	case *Float:
		// field equals
		fields = append(fields, &FilterField{
			name:     f.name,
			dataType: f.dataType,
			format:   fmt.Sprintf(`"%s" = %%s`, f.name),
		})

		// field doesn't equal
		fields = append(fields, &FilterField{
			name:     f.name + "Not",
			dataType: f.dataType,
			format:   fmt.Sprintf(`"%s" != %%s`, f.name),
		})

		// field is less than
		fields = append(fields, &FilterField{
			name:     f.name + "Lt",
			dataType: f.dataType,
			format:   fmt.Sprintf(`"%s" < %%s`, f.name),
		})

		// field is less than or equal
		fields = append(fields, &FilterField{
			name:     f.name + "Lte",
			dataType: f.dataType,
			format:   fmt.Sprintf(`"%s" <= %%s`, f.name),
		})

		// field is greater than
		fields = append(fields, &FilterField{
			name:     f.name + "Gt",
			dataType: f.dataType,
			format:   fmt.Sprintf(`"%s" > %%s`, f.name),
		})

		// field is greater than or equal
		fields = append(fields, &FilterField{
			name:     f.name + "Gte",
			dataType: f.dataType,
			format:   fmt.Sprintf(`"%s" >= %%s`, f.name),
		})

		// field is in list
		fields = append(fields, &FilterField{
			name:     f.name + "In",
			dataType: f.dataType,
			spread:   `, `,
			format:   fmt.Sprintf(`"%s" IN (%%s)`, f.name),
		})

		// field is not in list
		fields = append(fields, &FilterField{
			name:     f.name + "NotIn",
			dataType: f.dataType,
			spread:   `, `,
			format:   fmt.Sprintf(`"%s" NOT IN (%%s)`, f.name),
		})

		// if nullable
		if !f.notNull {
			fields = append(fields, &FilterField{
				name:        "nullable_" + f.name,
				description: "nullable " + f.name + " equals",
				nullable:    true,
				dataType:    f.dataType,
				format:      fmt.Sprintf(`"%s" = %%s`, f.name),
				nullformat:  fmt.Sprintf(`"%s" IS NULL`, f.name),
			})

			fields = append(fields, &FilterField{
				name:        "nullable_" + f.name + "_not",
				description: "nullable " + f.name + " is not equal",
				nullable:    true,
				dataType:    f.dataType,
				format:      fmt.Sprintf(`"%s" != %%s`, f.name),
				nullformat:  fmt.Sprintf(`"%s" IS NOT NULL`, f.name),
			})
		}

	case *Boolean:
		// field equals
		fields = append(fields, &FilterField{
			name:        f.name,
			description: f.name + " is equal to",
			dataType:    f.dataType,
			format:      fmt.Sprintf(`"%s" = %%s`, f.name),
		})

		// field doesn't equal
		fields = append(fields, &FilterField{
			name:        f.name + "Not",
			description: f.name + " is not equal to",
			dataType:    f.dataType,
			format:      fmt.Sprintf(`"%s" != %%s`, f.name),
		})

		// if nullable
		if !f.notNull {
			fields = append(fields, &FilterField{
				name:        "nullable_" + f.name,
				description: "nullable " + f.name + " equals",
				nullable:    true,
				dataType:    f.dataType,
				format:      fmt.Sprintf(`"%s" = %%s`, f.name),
				nullformat:  fmt.Sprintf(`"%s" IS NULL`, f.name),
			})

			fields = append(fields, &FilterField{
				name:        "nullable_" + f.name + "_not",
				description: "nullable " + f.name + " is not equal",
				nullable:    true,
				dataType:    f.dataType,
				format:      fmt.Sprintf(`"%s" != %%s`, f.name),
				nullformat:  fmt.Sprintf(`"%s" IS NOT NULL`, f.name),
			})
		}

	case *DateTime:
		// field equals
		fields = append(fields, &FilterField{
			name:        f.name,
			description: f.name + " is equal to",
			dataType:    f.dataType,
			format:      fmt.Sprintf(`"%s" = %%s`, f.name),
			value:       `%s.Format("2006-01-02 15:04:05.999999999Z07:00")`,
		})

		// field doesn't equal
		fields = append(fields, &FilterField{
			name:        f.name + "Not",
			description: f.name + " is not equal to",
			dataType:    f.dataType,
			format:      fmt.Sprintf(`"%s" != %%s`, f.name),
			value:       `%s.Format("2006-01-02 15:04:05.999999999Z07:00")`,
		})

		// field is in list
		fields = append(fields, &FilterField{
			name:        f.name + "In",
			description: f.name + " is in",
			dataType:    f.dataType,
			spread:      `, `,
			format:      fmt.Sprintf(`"%s" IN (%%s)`, f.name),
			value:       `%s.Format("2006-01-02 15:04:05.999999999Z07:00")`,
		})

		// field is not in list
		fields = append(fields, &FilterField{
			name:        f.name + "NotIn",
			description: f.name + " is not in",
			dataType:    f.dataType,
			spread:      `, `,
			format:      fmt.Sprintf(`"%s" NOT IN (%%s)`, f.name),
			value:       `%s.Format("2006-01-02 15:04:05.999999999Z07:00")`,
		})

		// field is less than
		fields = append(fields, &FilterField{
			name:        f.name + "Lt",
			description: f.name + " is less than",
			dataType:    f.dataType,
			format:      fmt.Sprintf(`"%s" < %%s`, f.name),
			value:       `%s.Format("2006-01-02 15:04:05.999999999Z07:00")`,
		})

		// field is less than or equal
		fields = append(fields, &FilterField{
			name:        f.name + "Lte",
			description: f.name + " is less than or equal",
			dataType:    f.dataType,
			format:      fmt.Sprintf(`"%s" <= %%s`, f.name),
			value:       `%s.Format("2006-01-02 15:04:05.999999999Z07:00")`,
		})

		// field is greater than
		fields = append(fields, &FilterField{
			name:        f.name + "Gt",
			description: f.name + " is greater than",
			dataType:    f.dataType,
			format:      fmt.Sprintf(`"%s" > %%s`, f.name),
			value:       `%s.Format("2006-01-02 15:04:05.999999999Z07:00")`,
		})

		// field is greater than or equal
		fields = append(fields, &FilterField{
			name:        f.name + "Gte",
			description: f.name + " is greater than or equal",
			dataType:    f.dataType,
			format:      fmt.Sprintf(`"%s" >= %%s`, f.name),
			value:       `%s.Format("2006-01-02 15:04:05.999999999Z07:00")`,
		})

		// if nullable
		if !f.notNull {
			fields = append(fields, &FilterField{
				name:        "nullable_" + f.name,
				description: "nullable " + f.name + " equals",
				nullable:    true,
				dataType:    f.dataType,
				format:      fmt.Sprintf(`"%s" = %%s`, f.name),
				nullformat:  fmt.Sprintf(`"%s" IS NULL`, f.name),
				value:       `%s.Format("2006-01-02 15:04:05.999999999Z07:00")`,
			})

			fields = append(fields, &FilterField{
				name:        "nullable_" + f.name + "_not",
				description: "nullable " + f.name + " is not equal",
				nullable:    true,
				dataType:    f.dataType,
				format:      fmt.Sprintf(`"%s" != %%s`, f.name),
				nullformat:  fmt.Sprintf(`"%s" IS NOT NULL`, f.name),
				value:       `%s.Format("2006-01-02 15:04:05.999999999Z07:00")`,
			})
		}

	case *Enumerable:
		// field equals
		fields = append(fields, &FilterField{
			name:        f.name,
			description: f.name + " is equal to",
			dataType:    f.dataType,
			format:      fmt.Sprintf(`"%s" = %%s`, f.name),
		})

		// field doesn't equal
		fields = append(fields, &FilterField{
			name:        f.name + "Not",
			description: f.name + " is not equal to",
			dataType:    f.dataType,
			format:      fmt.Sprintf(`"%s" != %%s`, f.name),
		})

		// field is in list
		fields = append(fields, &FilterField{
			name:        f.name + "In",
			description: f.name + " is in",
			dataType:    f.dataType,
			spread:      `, `,
			format:      fmt.Sprintf(`"%s" IN (%%s)`, f.name),
		})

		// field is not in list
		fields = append(fields, &FilterField{
			name:        f.name + "NotIn",
			description: f.name + " is not in",
			dataType:    f.dataType,
			spread:      `, `,
			format:      fmt.Sprintf(`"%s" NOT IN (%%s)`, f.name),
		})

		// if nullable
		if !f.notNull {
			fields = append(fields, &FilterField{
				name:        "nullable_" + f.name,
				description: "nullable " + f.name + " equals",
				nullable:    true,
				dataType:    f.dataType,
				format:      fmt.Sprintf(`"%s" = %%s`, f.name),
				nullformat:  fmt.Sprintf(`"%s" IS NULL`, f.name),
			})

			fields = append(fields, &FilterField{
				name:        "nullable_" + f.name + "_not",
				description: "nullable " + f.name + " is not equal",
				nullable:    true,
				dataType:    f.dataType,
				format:      fmt.Sprintf(`"%s" != %%s`, f.name),
				nullformat:  fmt.Sprintf(`"%s" IS NOT NULL`, f.name),
			})
		}

	case *List:
		// field equals
		fields = append(fields, &FilterField{
			name:        f.name + "Contains",
			description: f.name + " contains",
			dataType:    f.dataType,
		})

		// if nullable
		if !f.notNull {
			fields = append(fields, &FilterField{
				name:        "nullable_" + f.name,
				description: "nullable " + f.name + " equals",
				nullable:    true,
				dataType:    f.dataType,
				format:      fmt.Sprintf(`"%s" = %%s`, f.name),
				nullformat:  fmt.Sprintf(`"%s" IS NULL`, f.name),
			})

			fields = append(fields, &FilterField{
				name:        "nullable_" + f.name + "_not",
				description: "nullable " + f.name + " is not equal",
				nullable:    true,
				dataType:    f.dataType,
				format:      fmt.Sprintf(`"%s" != %%s`, f.name),
				nullformat:  fmt.Sprintf(`"%s" IS NOT NULL`, f.name),
			})
		}

		// // field equals
		// fields = append(fields, &FilterField{
		// 	name:     f.name + "ContainsEvery",
		// description: f.name + " contains every",
		// 	DType: f.Type,
		// 	spread:   `, `,
		// 	// format:   fmt.Sprintf(`%s NOT IN (%%s)`, f.name),
		// })

		// // field equals
		// fields = append(fields, &FilterField{
		// 	name:     f.name + "ContainsSome",
		// description: f.name + " contains some",
		// 	DType: f.Type,
		// 	spread:   `, `,
		// 	// format:   fmt.Sprintf(`%s NOT IN (%%s)`, f.name),
		// })

	default:
		return fields, fmt.Errorf("filter fields: unknown type %q", f.dataType.String())
	}

	return fields, nil
}

func newFilterField(
	name string,
	dataType DataType,
	description string,
	nullable bool,
	format string,
	nullformat string,
	spread string,
	value string,
) *FilterField {
	return &FilterField{
		name,
		dataType,
		description,
		nullable,
		format,
		nullformat,
		spread,
		value,
	}
}

// FilterField struct
type FilterField struct {
	name        string
	dataType    DataType
	description string
	nullable    bool
	format      string
	nullformat  string
	spread      string
	value       string
}

// Name fn
func (f *FilterField) Name() string {
	return fmt.Sprintf("%q", f.name)
}

// Pascal case
func (f *FilterField) Pascal() string {
	return gen.Pascal(f.name)
}

// Camel case
func (f *FilterField) Camel() string {
	return gen.Camel(f.name)
}

// Description of the field
func (f *FilterField) Description() string {
	return f.description
}

// Spread fn
func (f *FilterField) Spread() string {
	return f.spread
}

// Nullable fn
func (f *FilterField) Nullable() bool {
	return f.nullable
}

// Format fn
func (f *FilterField) Format() string {
	return f.format
}

// NullFormat fn
func (f *FilterField) NullFormat() string {
	return f.nullformat
}

// Type of filter
func (f *FilterField) Type() string {
	ptr := ""
	if f.nullable {
		ptr = "*"
	}
	if f.spread != "" {
		return "..." + ptr + f.dataType.String()
	}
	return ptr + f.dataType.String()
}

// Coerce the value of the filter field
func (f *FilterField) Coerce(value string) (string, error) {
	if f.value == "" {
		return value, nil
	}
	return fmt.Sprintf(f.value, value), nil
}
