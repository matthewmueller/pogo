package pogo

import (
	"strings"

	"github.com/knq/snaker"
)

// Coerce a postgres type into a Go type based on the column definition.
func Coerce(schema *Schema, dt string) (kind string) {
	// handle SETOF
	if strings.HasPrefix(dt, "SETOF ") {
		t := Coerce(schema, dt[len("SETOF "):])
		return "[]" + t
	}

	// determine if it's a slice
	if strings.HasSuffix(dt, "[]") {
		dt = dt[:len(dt)-2]
		t := Coerce(schema, dt)
		return "[]" + strings.TrimPrefix(t, "*")
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
	case "uuid":
		return "uuid.UUID"
	case "text":
		return "string"
	case "boolean":
		return "bool"
	case "integer", "smallint", "bigint":
		return "int"
	case "real":
		return "float32"
	case "double", "float":
		return "float64"
	case "date", "timestamp with time zone", "time with time zone", "time without time zone", "timestamp without time zone":
		return "time.Time"
	case "json":
		return "map[string]interface{}"
	case "numeric", "decimal":
		return "decimal.Decimal"
	default:
		for _, enum := range schema.Enums {
			name := enum.Name
			if schema.Name != "" && schema.Name != "public" {
				name = schema.Name + "." + name
			}

			if name == dt {
				return "enum." + snaker.SnakeToCamelIdentifier(enum.Name)
			}
		}
		panic("don't understand the data type `" + dt + "`.\nPlease open an issue: https://github.com/matthewmueller/pogo/issues")
	}
}

// // precScaleRE is the regexp that matches "(precision[,scale])" definitions in a
// // database.
// var precScaleRE = regexp.MustCompile(`\(([0-9]+)(\s*,[0-9]+)?\)$`)

// // ParsePrecision extracts (precision[,scale]) strings from a data type and
// // returns the data type without the string.
// func parsePrecision(dt string) (string, int, int) {
// 	var err error

// 	precision := -1
// 	scale := -1

// 	m := precScaleRE.FindStringSubmatchIndex(dt)
// 	if m != nil {
// 		// extract precision
// 		precision, err = strconv.Atoi(dt[m[2]:m[3]])
// 		if err != nil {
// 			panic("could not convert precision")
// 		}

// 		// extract scale
// 		if m[4] != -1 {
// 			scale, err = strconv.Atoi(dt[m[4]+1 : m[5]])
// 			if err != nil {
// 				panic("could not convert scale")
// 			}
// 		}

// 		// change dt
// 		dt = dt[:m[0]] + dt[m[1]:]
// 	}

// 	return dt, precision, scale
// }
