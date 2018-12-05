package sqlite

import (
	"fmt"
	"strings"

	"github.com/matthewmueller/pogo/internal/schema"
)

// Coercer struct
type Coercer struct{}

var _ schema.Coercer = (*Coercer)(nil)

// Type returns a Go type from an SQL type
func (c *Coercer) Type(sqlType string) (goType string, err error) {
	switch strings.ToUpper(sqlType) {
	case "NULL":
		return `nil`, nil
	case "INTEGER":
		// TODO: handle int64 gracefully
		return `int`, nil
	case "REAL":
		// TODO: handle float64 gracefully
		return `float`, nil
	case "TEXT":
		return `string`, nil
	case "BLOB":
		return `[]byte`, nil
	default:
		return "", fmt.Errorf(`sqlite coerce: unrecognized format %q`, sqlType)
	}
}

// FilterType coerces an SQL type into a filter type
func (c *Coercer) FilterType(sqlType string) (filterType schema.FilterType, err error) {
	switch strings.ToUpper(sqlType) {
	case "NULL":
		return schema.Null, nil
	case "INTEGER":
		return schema.Int, nil
	case "REAL":
		return schema.Float, nil
	case "TEXT":
		return schema.String, nil
	case "BLOB":
		return schema.JSON, nil
	default:
		return "", fmt.Errorf(`sqlite filter type: unrecognized format %q`, sqlType)
	}
}
