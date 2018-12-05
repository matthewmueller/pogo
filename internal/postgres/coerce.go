package postgres

import (
	"fmt"
	"strings"

	gen "github.com/matthewmueller/go-gen"
	"github.com/matthewmueller/pogo/internal/schema"
)

// Coercer struct
type Coercer struct {
	SchemaName string
	Enums      []*schema.Enum
}

var _ schema.Coercer = (*Coercer)(nil)

// Type returns a Go type from an SQL type
func (c *Coercer) Type(sqlType string) (goType string, err error) {
	// handle SETOF
	if strings.HasPrefix(sqlType, "SETOF ") {
		t, err := c.Type(sqlType[len("SETOF "):])
		if err != nil {
			return "", err
		}
		return "[]" + t, nil
	}

	// determine if it's a slice
	if strings.HasSuffix(sqlType, "[]") {
		sqlType = sqlType[:len(sqlType)-2]
		t, err := c.Type(sqlType)
		if err != nil {
			return "", err
		}
		return "[]" + strings.TrimPrefix(t, "*"), nil
	}

	// ignore the content of functions
	// TODO: not sure if this is a good idea or not
	// was primarily added for numeric(10, 10)
	// NOTE: this won't work for float(4) vs float(64)
	// where float(4) should be a float32 and float(64)
	// should be a float64
	idx := strings.Index(sqlType, "(")
	if idx >= 0 {
		sqlType = sqlType[:idx]
	}

	switch sqlType {
	case "uuid", "citext":
		return "string", nil
	case "text":
		return "string", nil
	case "boolean":
		return "bool", nil
	case "integer", "smallint", "bigint":
		return "int", nil
	case "real":
		return "float32", nil
	case "double", "float":
		return "float64", nil
	case "time with time zone", "time without time zone":
		return "string", nil
	case "date", "timestamp with time zone", "timestamp without time zone":
		return "time.Time", nil
	case "json", "jsonb":
		return "json.RawMessage", nil
	case "numeric", "decimal":
		return "decimal.Decimal", nil
	default:
		for _, enum := range c.Enums {
			name := enum.Name

			if c.SchemaName != "" && c.SchemaName != "public" {
				name = fmt.Sprintf(`%s.%s`, c.SchemaName, name)
			}

			// remove quotes
			kind := strings.Replace(sqlType, "\"", "", -1)

			if name == kind {
				return "enum." + gen.Pascal(enum.Name), nil
			}
		}

		return "", fmt.Errorf(`pogo/coerce: don't understand the data type: %s`, sqlType)
	}
}

// FilterType coerces an SQL type into a filter type
func (c *Coercer) FilterType(sqlType string) (filterType schema.FilterType, err error) {
	if strings.HasPrefix(sqlType, "SETOF ") || strings.HasSuffix(sqlType, "[]") {
		return schema.List, nil
	}

	switch sqlType {
	case "text", "citext":
		return schema.String, nil
	case "boolean":
		return schema.Boolean, nil
	case "integer", "smallint", "bigint":
		return schema.Int, nil
	case "real":
		return schema.Float, nil
	case "double", "float":
		return schema.Float, nil
	case "time with time zone", "time without time zone":
		return schema.String, nil
	case "date", "timestamp with time zone", "timestamp without time zone":
		return schema.DateTime, nil
	case "json", "jsonb":
		return schema.JSON, nil
	}

	// maybe an enum?
	for _, enum := range c.Enums {
		name := enum.Name

		if c.SchemaName != "" && c.SchemaName != "public" {
			name = fmt.Sprintf(`%s.%s`, c.SchemaName, name)
		}

		// remove quotes
		kind := strings.Replace(sqlType, "\"", "", -1)
		if name == kind {
			return schema.Enumerable, nil
		}
	}

	return "", fmt.Errorf(`pogo/coerceFilter: don't understand the data type: %s`, sqlType)
}
