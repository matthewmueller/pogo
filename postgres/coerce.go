package postgres

//
// import (
// 	"strings"
//
// 	"github.com/knq/snaker"
// )
//
// // Coerce a postgres type into a Go type based on the column definition.
// func Coerce(dt string, nullable bool) (precision int, nilVal string, typ string) {
// 	asSlice := false
//
// 	// handle SETOF
// 	if strings.HasPrefix(dt, "SETOF ") {
// 		_, _, t := Coerce(dt[len("SETOF "):], false)
// 		return 0, "nil", "[]" + t
// 	}
//
// 	// determine if it's a slice
// 	if strings.HasSuffix(dt, "[]") {
// 		dt = dt[:len(dt)-2]
// 		asSlice = true
// 	}
//
// 	// extract precision
// 	dt, precision, _ = args.ParsePrecision(dt)
//
// 	switch dt {
// 	case "boolean":
// 		nilVal = "false"
// 		typ = "bool"
// 		if nullable {
// 			nilVal = "sql.NullBool{}"
// 			typ = "sql.NullBool"
// 		}
//
// 	case "character", "character varying", "text", "money", "inet":
// 		nilVal = `""`
// 		typ = "string"
// 		if nullable {
// 			nilVal = "nil"
// 			typ = "*string"
// 		}
//
// 	case "smallint":
// 		nilVal = "0"
// 		typ = "int16"
// 		if nullable {
// 			nilVal = "sql.NullInt64{}"
// 			typ = "sql.NullInt64"
// 		}
// 	case "integer":
// 		nilVal = "0"
// 		typ = args.Int32Type
// 		if nullable {
// 			nilVal = "sql.NullInt64{}"
// 			typ = "sql.NullInt64"
// 		}
// 	case "bigint":
// 		nilVal = "0"
// 		typ = "int64"
// 		if nullable {
// 			nilVal = "sql.NullInt64{}"
// 			typ = "sql.NullInt64"
// 		}
//
// 	case "smallserial":
// 		nilVal = "0"
// 		typ = "uint16"
// 		if nullable {
// 			nilVal = "sql.NullInt64{}"
// 			typ = "sql.NullInt64"
// 		}
// 	case "serial":
// 		nilVal = "0"
// 		typ = args.Uint32Type
// 		if nullable {
// 			nilVal = "sql.NullInt64{}"
// 			typ = "sql.NullInt64"
// 		}
// 	case "bigserial":
// 		nilVal = "0"
// 		typ = "uint64"
// 		if nullable {
// 			nilVal = "sql.NullInt64{}"
// 			typ = "sql.NullInt64"
// 		}
//
// 	case "real":
// 		nilVal = "0.0"
// 		typ = "float32"
// 		if nullable {
// 			nilVal = "sql.NullFloat64{}"
// 			typ = "sql.NullFloat64"
// 		}
// 	case "numeric", "double precision":
// 		nilVal = "0.0"
// 		typ = "float64"
// 		if nullable {
// 			nilVal = "sql.NullFloat64{}"
// 			typ = "sql.NullFloat64"
// 		}
//
// 	case "bytea":
// 		asSlice = true
// 		typ = "byte"
//
// 	case "date", "timestamp with time zone", "time with time zone", "time without time zone", "timestamp without time zone":
// 		nilVal = "time.Time{}"
// 		typ = "time.Time"
// 		if nullable {
// 			nilVal = "nil"
// 			typ = "*time.Time"
// 		}
//
// 	case "interval":
// 		typ = "*time.Duration"
//
// 	case "json":
// 		typ = "map[string]interface{}"
// 		nilVal = "map[string]interface{}{}"
// 		if nullable {
// 			nilVal = "nil"
// 			typ = "*map[string]interface{}"
// 		}
//
// 	case `"char"`, "bit":
// 		// FIXME: this needs to actually be tested ...
// 		// i think this should be 'rune' but I don't think database/sql
// 		// supports 'rune' as a type?
// 		//
// 		// this is mainly here because postgres's pg_catalog.* meta tables have
// 		// this as a type.
// 		//typ = "rune"
// 		nilVal = `uint8(0)`
// 		typ = "uint8"
//
// 	case `"any"`, "bit varying":
// 		asSlice = true
// 		typ = "byte"
//
// 	case "hstore":
// 		typ = "hstore.Hstore"
//
// 	case "uuid":
// 		nilVal = "uuid.New()"
// 		typ = "uuid.UUID"
//
// 	default:
// 		if strings.HasPrefix(dt, args.Schema+".") {
// 			// in the same schema, so chop off
// 			typ = snaker.SnakeToCamelIdentifier(dt[len(args.Schema)+1:])
// 			nilVal = typ + "(0)"
// 		} else {
// 			typ = snaker.SnakeToCamelIdentifier(dt)
// 			nilVal = typ + "{}"
// 		}
// 	}
//
// 	// special case for []slice
// 	if typ == "string" && asSlice {
// 		return precision, "[]string{}", "[]string"
// 	}
//
// 	// correct type if slice
// 	if asSlice {
// 		typ = "[]" + typ
// 		nilVal = "nil"
// 	}
//
// 	return precision, nilVal, typ
// }
