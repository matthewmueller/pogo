package pogo

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/knq/snaker"
	uuid "github.com/satori/go.uuid"
)

// Fake a postgres value given a type
func Fake(settings *Settings, schema *Schema, name string, dt string) (kind string) {
	// determine if it's a slice
	if strings.HasSuffix(dt, "[]") {
		kind = coerce(schema, dt[:len(dt)-2])
		// TODO: random number
		t := Fake(settings, schema, name, dt[:len(dt)-2])
		return "[]" + kind + "{" + t + "}"
	}

	switch dt {
	case "uuid":
		return "uuid.NewV4()"
	case "text":
		return fmt.Sprintf(`"%s"`, uuid.NewV4())
	case "boolean":
		return "false"
	case "integer":
		return strconv.Itoa(rand.Intn(1000))
	case "timestamp with time zone", "timestamp":
		return "time.Now()"
	case "json":
		return "map[string]interface{}{}"
	default:
		for _, enum := range schema.Enums {
			name := enum.Name
			if schema.Name != "" && schema.Name != "public" {
				name = schema.Name + "." + name
			}

			if name == dt {
				return fmt.Sprintf(`%s.%s%s`, settings.Package, coerce(schema, dt), snaker.SnakeToCamelIdentifier(enum.Values[0].Label))
			}
		}
		panic("don't understand the data type `" + dt + "` for `" + name + "`.\nPlease open an issue: https://github.com/matthewmueller/pogo/issues")
	}
}
