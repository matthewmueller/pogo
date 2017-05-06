package pogo

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/knq/snaker"
	"github.com/matthewmueller/pogo/postgres"
)

// Coerce struct
type Coerce struct {
	Schema string
	Enums  []*postgres.Enum
}

// NewCoerce struct
func NewCoerce(schema string, enums []*postgres.Enum) Coerce {
	return Coerce{
		Schema: schema,
		Enums:  enums,
	}
}

// Coerce a postgres type into a Go type based on the column definition.
func (c *Coerce) Coerce(dt string) (kind string) {
	// handle SETOF
	if strings.HasPrefix(dt, "SETOF ") {
		t := c.Coerce(dt[len("SETOF "):])
		return "[]" + t
	}

	// determine if it's a slice
	if strings.HasSuffix(dt, "[]") {
		dt = dt[:len(dt)-2]
		t := c.Coerce(dt)
		return "*[]" + strings.TrimPrefix(t, "*")
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
		return "*uuid.UUID"
	case "text":
		return "*string"
	case "boolean":
		return "*bool"
	case "integer", "smallint", "bigint":
		return "*int"
	case "real":
		return "*float32"
	case "double", "float":
		return "*float64"
	case "date", "timestamp with time zone", "time with time zone", "time without time zone", "timestamp without time zone":
		return "*time.Time"
	case "json":
		return "*map[string]interface{}"
	case "numeric", "decimal":
		return "*decimal.Decimal"
	default:
		for _, enum := range c.Enums {
			name := enum.Name
			if c.Schema != "" && c.Schema != "public" {
				name = c.Schema + "." + name
			}

			if name == dt {
				return "*" + snaker.SnakeToCamelIdentifier(enum.Name)
			}
		}
		panic("don't understand the data type `" + dt + "`.\nPlease open an issue: https://github.com/matthewmueller/pogo/issues")
	}

	// extract precision
	// dt, precision, _ = parsePrecision(dt)

	// switch dt {
	// case "boolean":
	// 	nilVal = "nil"
	// 	typ = "*bool"
	// 	if nullable {
	// 		nilVal = "nil"
	// 		typ = "*bool"
	// 	}
	//
	// case "character", "character varying", "text", "money", "inet":
	// 	nilVal = `""`
	// 	typ = "string"
	// 	if nullable {
	// 		nilVal = "nil"
	// 		typ = "*string"
	// 	}

	// case "smallint":
	// 	nilVal = "0"
	// 	typ = "int16"
	// 	if nullable {
	// 		nilVal = "sql.NullInt64{}"
	// 		typ = "sql.NullInt64"
	// 	}
	// case "integer":
	// 	nilVal = "0"
	// 	typ = args.Int32Type
	// 	if nullable {
	// 		nilVal = "sql.NullInt64{}"
	// 		typ = "sql.NullInt64"
	// 	}
	// case "bigint":
	// 	nilVal = "0"
	// 	typ = "int64"
	// 	if nullable {
	// 		nilVal = "sql.NullInt64{}"
	// 		typ = "sql.NullInt64"
	// 	}
	//
	// case "smallserial":
	// 	nilVal = "0"
	// 	typ = "uint16"
	// 	if nullable {
	// 		nilVal = "sql.NullInt64{}"
	// 		typ = "sql.NullInt64"
	// 	}
	// case "serial":
	// 	nilVal = "0"
	// 	typ = args.Uint32Type
	// 	if nullable {
	// 		nilVal = "sql.NullInt64{}"
	// 		typ = "sql.NullInt64"
	// 	}
	// case "bigserial":
	// 	nilVal = "0"
	// 	typ = "uint64"
	// 	if nullable {
	// 		nilVal = "sql.NullInt64{}"
	// 		typ = "sql.NullInt64"
	// 	}
	//
	// case "real":
	// 	nilVal = "0.0"
	// 	typ = "float32"
	// 	if nullable {
	// 		nilVal = "sql.NullFloat64{}"
	// 		typ = "sql.NullFloat64"
	// 	}
	// case "numeric", "double precision":
	// 	nilVal = "0.0"
	// 	typ = "float64"
	// 	if nullable {
	// 		nilVal = "sql.NullFloat64{}"
	// 		typ = "sql.NullFloat64"
	// 	}
	//
	// case "bytea":
	// 	asSlice = true
	// 	typ = "byte"
	//
	// case "date", "timestamp with time zone", "time with time zone", "time without time zone", "timestamp without time zone":
	// 	nilVal = "time.Time{}"
	// 	typ = "time.Time"
	// 	if nullable {
	// 		nilVal = "nil"
	// 		typ = "*time.Time"
	// 	}
	//
	// case "interval":
	// 	typ = "*time.Duration"
	//
	// case "json":
	// 	typ = "map[string]interface{}"
	// 	nilVal = "map[string]interface{}{}"
	// 	if nullable {
	// 		nilVal = "nil"
	// 		typ = "*map[string]interface{}"
	// 	}
	//
	// case `"char"`, "bit":
	// 	// FIXME: this needs to actually be tested ...
	// 	// i think this should be 'rune' but I don't think database/sql
	// 	// supports 'rune' as a type?
	// 	//
	// 	// this is mainly here because postgres's pg_catalog.* meta tables have
	// 	// this as a type.
	// 	//typ = "rune"
	// 	nilVal = `uint8(0)`
	// 	typ = "uint8"
	//
	// case `"any"`, "bit varying":
	// 	asSlice = true
	// 	typ = "byte"
	//
	// case "hstore":
	// 	typ = "hstore.Hstore"
	//
	// case "uuid":
	// 	nilVal = "uuid.New()"
	// 	typ = "uuid.UUID"

	// default:
	// if strings.HasPrefix(dt, args.Schema+".") {
	// 	// in the same schema, so chop off
	// 	typ = snaker.SnakeToCamelIdentifier(dt[len(args.Schema)+1:])
	// 	nilVal = typ + "(0)"
	// } else {
	// 	typ = snaker.SnakeToCamelIdentifier(dt)
	// 	nilVal = typ + "{}"
	// }
	// }

	// special case for []slice
	// if typ == "string" && asSlice {
	// 	return precision, "[]string{}", "[]string"
	// }
	//
	// // correct type if slice
	// if asSlice {
	// 	typ = "[]" + typ
	// 	nilVal = "nil"
	// }

	// return kind
}

// precScaleRE is the regexp that matches "(precision[,scale])" definitions in a
// database.
var precScaleRE = regexp.MustCompile(`\(([0-9]+)(\s*,[0-9]+)?\)$`)

// ParsePrecision extracts (precision[,scale]) strings from a data type and
// returns the data type without the string.
func parsePrecision(dt string) (string, int, int) {
	var err error

	precision := -1
	scale := -1

	m := precScaleRE.FindStringSubmatchIndex(dt)
	if m != nil {
		// extract precision
		precision, err = strconv.Atoi(dt[m[2]:m[3]])
		if err != nil {
			panic("could not convert precision")
		}

		// extract scale
		if m[4] != -1 {
			scale, err = strconv.Atoi(dt[m[4]+1 : m[5]])
			if err != nil {
				panic("could not convert scale")
			}
		}

		// change dt
		dt = dt[:m[0]] + dt[m[1]:]
	}

	return dt, precision, scale
}
