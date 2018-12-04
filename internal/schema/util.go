package schema

import (
	"fmt"
	"strings"

	"github.com/matthewmueller/go-gen"
)

func coerce(schema *Schema, dt string) (kind string, err error) {
	// handle SETOF
	if strings.HasPrefix(dt, "SETOF ") {
		t, err := coerce(schema, dt[len("SETOF "):])
		if err != nil {
			return "", err
		}
		return "[]" + t, nil
	}

	// determine if it's a slice
	if strings.HasSuffix(dt, "[]") {
		dt = dt[:len(dt)-2]
		t, err := coerce(schema, dt)
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
	idx := strings.Index(dt, "(")
	if idx >= 0 {
		dt = dt[:idx]
	}

	switch dt {
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
		for _, enum := range schema.Enums {
			name := enum.Name

			if schema.Name != "" && schema.Name != "public" {
				name = fmt.Sprintf(`%s.%s`, schema.Name, name)
			}

			// remove quotes
			kind := strings.Replace(dt, "\"", "", -1)

			if name == kind {
				return "enum." + gen.Pascal(enum.Name), nil
			}
		}

		return "", fmt.Errorf(`pogo/coerce: don't understand the data type: %s`, dt)
	}
}

func coerceFilter(schema *Schema, dt string) (kind string, err error) {
	if strings.HasPrefix(dt, "SETOF ") || strings.HasSuffix(dt, "[]") {
		return "List", nil
	}

	switch dt {
	case "text", "citext":
		return "String", nil
	case "boolean":
		return "Boolean", nil
	case "integer", "smallint", "bigint":
		return "Int", nil
	case "real":
		return "Float", nil
	case "double", "float":
		return "Float", nil
	case "time with time zone", "time without time zone":
		return "String", nil
	case "date", "timestamp with time zone", "timestamp without time zone":
		return "DateTime", nil
	case "json", "jsonb":
		return "JSON", nil
	}

	// maybe an enum?
	for _, enum := range schema.Enums {
		name := enum.Name

		if schema.Name != "" && schema.Name != "public" {
			name = fmt.Sprintf(`%s.%s`, schema.Name, name)
		}

		// remove quotes
		kind := strings.Replace(dt, "\"", "", -1)
		if name == kind {
			return "Enum", nil
		}
	}

	return "", fmt.Errorf(`pogo/coerceFilter: don't understand the data type: %s`, dt)
}
