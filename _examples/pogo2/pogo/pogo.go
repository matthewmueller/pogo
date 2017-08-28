package pogo

import (
	"strconv"

	"github.com/jackc/pgx"
)

// DB interface
type DB interface {
	Exec(string, ...interface{}) (pgx.CommandTag, error)
	Query(string, ...interface{}) (*pgx.Rows, error)
	QueryRow(string, ...interface{}) *pgx.Row
}

// Log provides the log func used by generated queries.
var Log = func(string, ...interface{}) {}

const (
	// UpsertDoNothing Do nothing if there's a conflict
	UpsertDoNothing = "DO NOTHING"
	// UpsertDoUpdate Perform an update when there's a conflict
	UpsertDoUpdate = "DO UPDATE"
)

// Slice is a runtime helper
func Slice(fields map[string]interface{}, offset int) (c []string, i []string, v []interface{}) {
	n := offset + 1
	for col, val := range fields {
		c = append(c, `"`+col+`"`)
		i = append(i, "$"+strconv.Itoa(n))
		v = append(v, val)
		n++
	}
	return c, i, v
}
