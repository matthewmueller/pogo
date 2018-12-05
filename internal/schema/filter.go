package schema

import (
	"fmt"

	gen "github.com/matthewmueller/go-gen"
)

// Filter struct
type Filter struct {
	Name     string   // column name
	DataType DataType // column type (only for columns)
	NotNull  bool     // column can't be null
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
	switch f.DataType.(type) {
	case *Null:
		// add no filters
	case *JSON:
		// TODO: add filters
	case *String:
		// field equals
		fields = append(fields, &FilterField{
			Name:        f.Name,
			DataType:    f.DataType,
			Description: f.Name + " equals",
			Format:      fmt.Sprintf(`"%s" = %%s`, f.Name),
		})

		// field doesn't equal
		fields = append(fields, &FilterField{
			Name:        f.Name + "Not",
			DataType:    f.DataType,
			Description: f.Name + " doesn't equal",
			Format:      fmt.Sprintf(`"%s" != %%s`, f.Name),
		})

		// field contains
		fields = append(fields, &FilterField{
			Name:        f.Name + "Contains",
			DataType:    f.DataType,
			Description: f.Name + " contains",
			Format:      fmt.Sprintf(`"%s" LIKE '%%%%' || %%s || '%%%%'`, f.Name),
		})

		// field doesn't contain
		fields = append(fields, &FilterField{
			Name:        f.Name + "NotContains",
			DataType:    f.DataType,
			Description: f.Name + " doesn't contain",
			Format:      fmt.Sprintf(`"%s" NOT LIKE '%%%%' || %%s || '%%%%'`, f.Name),
		})

		// field starts with
		fields = append(fields, &FilterField{
			Name:        f.Name + "StartsWith",
			DataType:    f.DataType,
			Description: f.Name + " starts with",
			Format:      fmt.Sprintf(`"%s" LIKE %%s || '%%%%'`, f.Name),
		})

		// field doesn't start with
		fields = append(fields, &FilterField{
			Name:        f.Name + "NotStartsWith",
			DataType:    f.DataType,
			Description: f.Name + " doesn't start with",
			Format:      fmt.Sprintf(`"%s" NOT LIKE %%s || '%%%%'`, f.Name),
		})

		// field ends with
		fields = append(fields, &FilterField{
			Name:        f.Name + "EndsWith",
			DataType:    f.DataType,
			Description: f.Name + " ends with",
			Format:      fmt.Sprintf(`"%s" LIKE '%%%%' || %%s`, f.Name),
		})

		// field doesn't end with
		fields = append(fields, &FilterField{
			Name:        f.Name + "NotEndsWith",
			DataType:    f.DataType,
			Description: f.Name + " doesn't end with",
			Format:      fmt.Sprintf(`"%s" NOT LIKE '%%%%' || %%s`, f.Name),
		})

		// field is less than
		fields = append(fields, &FilterField{
			Name:        f.Name + "Lt",
			DataType:    f.DataType,
			Description: f.Name + " is less than",
			Format:      fmt.Sprintf(`"%s" < %%s`, f.Name),
		})

		// field is less than or equal
		fields = append(fields, &FilterField{
			Name:        f.Name + "Lte",
			DataType:    f.DataType,
			Description: f.Name + " is less than or equal",
			Format:      fmt.Sprintf(`"%s" <= %%s`, f.Name),
		})

		// field is greater than
		fields = append(fields, &FilterField{
			Name:        f.Name + "Gt",
			DataType:    f.DataType,
			Description: f.Name + " is greater than",
			Format:      fmt.Sprintf(`"%s" > %%s`, f.Name),
		})

		// field is greater than or equal
		fields = append(fields, &FilterField{
			Name:        f.Name + "Gte",
			DataType:    f.DataType,
			Description: f.Name + " is greater than or equal",
			Format:      fmt.Sprintf(`"%s" >= %%s`, f.Name),
		})

		// field is in list
		fields = append(fields, &FilterField{
			Name:        f.Name + "In",
			DataType:    f.DataType,
			Description: f.Name + " is in",
			Format:      fmt.Sprintf(`"%s" IN (%%s)`, f.Name),
			Spread:      `, `,
		})

		// field is not in list
		fields = append(fields, &FilterField{
			Name:        f.Name + "NotIn",
			DataType:    f.DataType,
			Description: f.Name + " is not in",
			Format:      fmt.Sprintf(`"%s" NOT IN (%%s)`, f.Name),
			Spread:      `, `,
		})

		// if nullable
		if !f.NotNull {
			fields = append(fields, &FilterField{
				Name:        "nullable_" + f.Name,
				DataType:    f.DataType,
				Description: "nullable " + f.Name + " equals",
				Nullable:    true,
				Format:      fmt.Sprintf(`"%s" = %%s`, f.Name),
				NullFormat:  fmt.Sprintf(`"%s" IS NULL`, f.Name),
			})

			fields = append(fields, &FilterField{
				Name:        "nullable_" + f.Name + "_not",
				DataType:    f.DataType,
				Description: "nullable " + f.Name + " is not equal",
				Nullable:    true,
				Format:      fmt.Sprintf(`"%s" != %%s`, f.Name),
				NullFormat:  fmt.Sprintf(`"%s" IS NOT NULL`, f.Name),
			})
		}

	case *Integer:
		// field equals
		fields = append(fields, &FilterField{
			Name:        f.Name,
			DataType:    f.DataType,
			Description: f.Name + " equals",
			Format:      fmt.Sprintf(`"%s" = %%s`, f.Name),
		})

		// field doesn't equal
		fields = append(fields, &FilterField{
			Name:        f.Name + "Not",
			DataType:    f.DataType,
			Description: f.Name + " doesn't equal",
			Format:      fmt.Sprintf(`"%s" != %%s`, f.Name),
		})

		// field is less than
		fields = append(fields, &FilterField{
			Name:        f.Name + "Lt",
			DataType:    f.DataType,
			Description: f.Name + " is less than",
			Format:      fmt.Sprintf(`"%s" < %%s`, f.Name),
		})

		// field is less than or equal
		fields = append(fields, &FilterField{
			Name:        f.Name + "Lte",
			DataType:    f.DataType,
			Description: f.Name + " is less than or equal",
			Format:      fmt.Sprintf(`"%s" <= %%s`, f.Name),
		})

		// field is greater than
		fields = append(fields, &FilterField{
			Name:        f.Name + "Gt",
			DataType:    f.DataType,
			Description: f.Name + " is greater than",
			Format:      fmt.Sprintf(`"%s" > %%s`, f.Name),
		})

		// field is greater than or equal
		fields = append(fields, &FilterField{
			Name:        f.Name + "Gte",
			DataType:    f.DataType,
			Description: f.Name + " is greater than or equal",
			Format:      fmt.Sprintf(`"%s" >= %%s`, f.Name),
		})

		// field is in list
		fields = append(fields, &FilterField{
			Name:        f.Name + "In",
			DataType:    f.DataType,
			Description: f.Name + " is in",
			Format:      fmt.Sprintf(`"%s" IN (%%s)`, f.Name),
			Spread:      `, `,
		})

		// field is not in list
		fields = append(fields, &FilterField{
			Name:        f.Name + "NotIn",
			DataType:    f.DataType,
			Description: f.Name + " is not in",
			Format:      fmt.Sprintf(`"%s" NOT IN (%%s)`, f.Name),
			Spread:      `, `,
		})

		// if nullable
		if !f.NotNull {
			fields = append(fields, &FilterField{
				Name:        "nullable_" + f.Name,
				DataType:    f.DataType,
				Description: "nullable " + f.Name + " equals",
				Nullable:    true,
				Format:      fmt.Sprintf(`"%s" = %%s`, f.Name),
				NullFormat:  fmt.Sprintf(`"%s" IS NULL`, f.Name),
			})

			fields = append(fields, &FilterField{
				Name:        "nullable_" + f.Name + "_not",
				DataType:    f.DataType,
				Description: "nullable " + f.Name + " is not equal",
				Nullable:    true,
				Format:      fmt.Sprintf(`"%s" != %%s`, f.Name),
				NullFormat:  fmt.Sprintf(`"%s" IS NOT NULL`, f.Name),
			})
		}

	case *Float:
		// field equals
		fields = append(fields, &FilterField{
			Name:     f.Name,
			DataType: f.DataType,
			Format:   fmt.Sprintf(`"%s" = %%s`, f.Name),
		})

		// field doesn't equal
		fields = append(fields, &FilterField{
			Name:     f.Name + "Not",
			DataType: f.DataType,
			Format:   fmt.Sprintf(`"%s" != %%s`, f.Name),
		})

		// field is less than
		fields = append(fields, &FilterField{
			Name:     f.Name + "Lt",
			DataType: f.DataType,
			Format:   fmt.Sprintf(`"%s" < %%s`, f.Name),
		})

		// field is less than or equal
		fields = append(fields, &FilterField{
			Name:     f.Name + "Lte",
			DataType: f.DataType,
			Format:   fmt.Sprintf(`"%s" <= %%s`, f.Name),
		})

		// field is greater than
		fields = append(fields, &FilterField{
			Name:     f.Name + "Gt",
			DataType: f.DataType,
			Format:   fmt.Sprintf(`"%s" > %%s`, f.Name),
		})

		// field is greater than or equal
		fields = append(fields, &FilterField{
			Name:     f.Name + "Gte",
			DataType: f.DataType,
			Format:   fmt.Sprintf(`"%s" >= %%s`, f.Name),
		})

		// field is in list
		fields = append(fields, &FilterField{
			Name:     f.Name + "In",
			DataType: f.DataType,
			Spread:   `, `,
			Format:   fmt.Sprintf(`"%s" IN (%%s)`, f.Name),
		})

		// field is not in list
		fields = append(fields, &FilterField{
			Name:     f.Name + "NotIn",
			DataType: f.DataType,
			Spread:   `, `,
			Format:   fmt.Sprintf(`"%s" NOT IN (%%s)`, f.Name),
		})

		// if nullable
		if !f.NotNull {
			fields = append(fields, &FilterField{
				Name:        "nullable_" + f.Name,
				Description: "nullable " + f.Name + " equals",
				Nullable:    true,
				DataType:    f.DataType,
				Format:      fmt.Sprintf(`"%s" = %%s`, f.Name),
				NullFormat:  fmt.Sprintf(`"%s" IS NULL`, f.Name),
			})

			fields = append(fields, &FilterField{
				Name:        "nullable_" + f.Name + "_not",
				Description: "nullable " + f.Name + " is not equal",
				Nullable:    true,
				DataType:    f.DataType,
				Format:      fmt.Sprintf(`"%s" != %%s`, f.Name),
				NullFormat:  fmt.Sprintf(`"%s" IS NOT NULL`, f.Name),
			})
		}

	case *Boolean:
		// field equals
		fields = append(fields, &FilterField{
			Name:        f.Name,
			Description: f.Name + " is equal to",
			DataType:    f.DataType,
			Format:      fmt.Sprintf(`"%s" = %%s`, f.Name),
		})

		// field doesn't equal
		fields = append(fields, &FilterField{
			Name:        f.Name + "Not",
			Description: f.Name + " is not equal to",
			DataType:    f.DataType,
			Format:      fmt.Sprintf(`"%s" != %%s`, f.Name),
		})

		// if nullable
		if !f.NotNull {
			fields = append(fields, &FilterField{
				Name:        "nullable_" + f.Name,
				Description: "nullable " + f.Name + " equals",
				Nullable:    true,
				DataType:    f.DataType,
				Format:      fmt.Sprintf(`"%s" = %%s`, f.Name),
				NullFormat:  fmt.Sprintf(`"%s" IS NULL`, f.Name),
			})

			fields = append(fields, &FilterField{
				Name:        "nullable_" + f.Name + "_not",
				Description: "nullable " + f.Name + " is not equal",
				Nullable:    true,
				DataType:    f.DataType,
				Format:      fmt.Sprintf(`"%s" != %%s`, f.Name),
				NullFormat:  fmt.Sprintf(`"%s" IS NOT NULL`, f.Name),
			})
		}

	case *DateTime:
		// field equals
		fields = append(fields, &FilterField{
			Name:        f.Name,
			Description: f.Name + " is equal to",
			DataType:    f.DataType,
			Format:      fmt.Sprintf(`"%s" = %%s`, f.Name),
			Value:       `%s.Format("2006-01-02 15:04:05.999999999Z07:00")`,
		})

		// field doesn't equal
		fields = append(fields, &FilterField{
			Name:        f.Name + "Not",
			Description: f.Name + " is not equal to",
			DataType:    f.DataType,
			Format:      fmt.Sprintf(`"%s" != %%s`, f.Name),
			Value:       `%s.Format("2006-01-02 15:04:05.999999999Z07:00")`,
		})

		// field is in list
		fields = append(fields, &FilterField{
			Name:        f.Name + "In",
			Description: f.Name + " is in",
			DataType:    f.DataType,
			Spread:      `, `,
			Format:      fmt.Sprintf(`"%s" IN (%%s)`, f.Name),
			Value:       `%s.Format("2006-01-02 15:04:05.999999999Z07:00")`,
		})

		// field is not in list
		fields = append(fields, &FilterField{
			Name:        f.Name + "NotIn",
			Description: f.Name + " is not in",
			DataType:    f.DataType,
			Spread:      `, `,
			Format:      fmt.Sprintf(`"%s" NOT IN (%%s)`, f.Name),
			Value:       `%s.Format("2006-01-02 15:04:05.999999999Z07:00")`,
		})

		// field is less than
		fields = append(fields, &FilterField{
			Name:        f.Name + "Lt",
			Description: f.Name + " is less than",
			DataType:    f.DataType,
			Format:      fmt.Sprintf(`"%s" < %%s`, f.Name),
			Value:       `%s.Format("2006-01-02 15:04:05.999999999Z07:00")`,
		})

		// field is less than or equal
		fields = append(fields, &FilterField{
			Name:        f.Name + "Lte",
			Description: f.Name + " is less than or equal",
			DataType:    f.DataType,
			Format:      fmt.Sprintf(`"%s" <= %%s`, f.Name),
			Value:       `%s.Format("2006-01-02 15:04:05.999999999Z07:00")`,
		})

		// field is greater than
		fields = append(fields, &FilterField{
			Name:        f.Name + "Gt",
			Description: f.Name + " is greater than",
			DataType:    f.DataType,
			Format:      fmt.Sprintf(`"%s" > %%s`, f.Name),
			Value:       `%s.Format("2006-01-02 15:04:05.999999999Z07:00")`,
		})

		// field is greater than or equal
		fields = append(fields, &FilterField{
			Name:        f.Name + "Gte",
			Description: f.Name + " is greater than or equal",
			DataType:    f.DataType,
			Format:      fmt.Sprintf(`"%s" >= %%s`, f.Name),
			Value:       `%s.Format("2006-01-02 15:04:05.999999999Z07:00")`,
		})

		// if nullable
		if !f.NotNull {
			fields = append(fields, &FilterField{
				Name:        "nullable_" + f.Name,
				Description: "nullable " + f.Name + " equals",
				Nullable:    true,
				DataType:    f.DataType,
				Format:      fmt.Sprintf(`"%s" = %%s`, f.Name),
				NullFormat:  fmt.Sprintf(`"%s" IS NULL`, f.Name),
				Value:       `%s.Format("2006-01-02 15:04:05.999999999Z07:00")`,
			})

			fields = append(fields, &FilterField{
				Name:        "nullable_" + f.Name + "_not",
				Description: "nullable " + f.Name + " is not equal",
				Nullable:    true,
				DataType:    f.DataType,
				Format:      fmt.Sprintf(`"%s" != %%s`, f.Name),
				NullFormat:  fmt.Sprintf(`"%s" IS NOT NULL`, f.Name),
				Value:       `%s.Format("2006-01-02 15:04:05.999999999Z07:00")`,
			})
		}

	case *Enumerable:
		// field equals
		fields = append(fields, &FilterField{
			Name:        f.Name,
			Description: f.Name + " is equal to",
			DataType:    f.DataType,
			Format:      fmt.Sprintf(`"%s" = %%s`, f.Name),
		})

		// field doesn't equal
		fields = append(fields, &FilterField{
			Name:        f.Name + "Not",
			Description: f.Name + " is not equal to",
			DataType:    f.DataType,
			Format:      fmt.Sprintf(`"%s" != %%s`, f.Name),
		})

		// field is in list
		fields = append(fields, &FilterField{
			Name:        f.Name + "In",
			Description: f.Name + " is in",
			DataType:    f.DataType,
			Spread:      `, `,
			Format:      fmt.Sprintf(`"%s" IN (%%s)`, f.Name),
		})

		// field is not in list
		fields = append(fields, &FilterField{
			Name:        f.Name + "NotIn",
			Description: f.Name + " is not in",
			DataType:    f.DataType,
			Spread:      `, `,
			Format:      fmt.Sprintf(`"%s" NOT IN (%%s)`, f.Name),
		})

		// if nullable
		if !f.NotNull {
			fields = append(fields, &FilterField{
				Name:        "nullable_" + f.Name,
				Description: "nullable " + f.Name + " equals",
				Nullable:    true,
				DataType:    f.DataType,
				Format:      fmt.Sprintf(`"%s" = %%s`, f.Name),
				NullFormat:  fmt.Sprintf(`"%s" IS NULL`, f.Name),
			})

			fields = append(fields, &FilterField{
				Name:        "nullable_" + f.Name + "_not",
				Description: "nullable " + f.Name + " is not equal",
				Nullable:    true,
				DataType:    f.DataType,
				Format:      fmt.Sprintf(`"%s" != %%s`, f.Name),
				NullFormat:  fmt.Sprintf(`"%s" IS NOT NULL`, f.Name),
			})
		}

	case *List:
		// field equals
		fields = append(fields, &FilterField{
			Name:        f.Name + "Contains",
			Description: f.Name + " contains",
			DataType:    f.DataType,
		})

		// if nullable
		if !f.NotNull {
			fields = append(fields, &FilterField{
				Name:        "nullable_" + f.Name,
				Description: "nullable " + f.Name + " equals",
				Nullable:    true,
				DataType:    f.DataType,
				Format:      fmt.Sprintf(`"%s" = %%s`, f.Name),
				NullFormat:  fmt.Sprintf(`"%s" IS NULL`, f.Name),
			})

			fields = append(fields, &FilterField{
				Name:        "nullable_" + f.Name + "_not",
				Description: "nullable " + f.Name + " is not equal",
				Nullable:    true,
				DataType:    f.DataType,
				Format:      fmt.Sprintf(`"%s" != %%s`, f.Name),
				NullFormat:  fmt.Sprintf(`"%s" IS NOT NULL`, f.Name),
			})
		}

		// // field equals
		// fields = append(fields, &FilterField{
		// 	Name:     f.Name + "ContainsEvery",
		// Description: f.Name + " contains every",
		// 	DType: f.Type,
		// 	Spread:   `, `,
		// 	// Format:   fmt.Sprintf(`%s NOT IN (%%s)`, f.Name),
		// })

		// // field equals
		// fields = append(fields, &FilterField{
		// 	Name:     f.Name + "ContainsSome",
		// Description: f.Name + " contains some",
		// 	DType: f.Type,
		// 	Spread:   `, `,
		// 	// Format:   fmt.Sprintf(`%s NOT IN (%%s)`, f.Name),
		// })

	default:
		return fields, fmt.Errorf("filter fields: unknown type %q", f.DataType.String())
	}

	return fields, nil
}

// FilterField struct
type FilterField struct {
	Name        string
	DataType    DataType
	Description string
	Nullable    bool
	Format      string
	NullFormat  string
	Spread      string
	Value       string
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
func (f *FilterField) Type() string {
	ptr := ""
	if f.Nullable {
		ptr = "*"
	}
	if f.Spread != "" {
		return "..." + ptr + f.DataType.String()
	}
	return ptr + f.DataType.String()
}

// Coerce the value of the filter field
func (f *FilterField) Coerce(value string) (string, error) {
	if f.Value == "" {
		return value, nil
	}
	return fmt.Sprintf(f.Value, value), nil
}
