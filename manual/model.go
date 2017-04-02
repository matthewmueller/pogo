package model

import (
	"strconv"

	"github.com/jackc/pgx"
)

// DB is the common interface for database operations that can be used with
// types from schema 'jack'.
//
// This should work with database/sql.DB and database/sql.Tx.
type DB interface {
	Exec(string, ...interface{}) (pgx.CommandTag, error)
	Query(string, ...interface{}) (*pgx.Rows, error)
	QueryRow(string, ...interface{}) *pgx.Row
}

// XOLog provides the log func used by generated queries.
var XOLog = func(string, ...interface{}) {}

func querySlices(fields map[string]interface{}, offset int) (c []string, i []string, v []interface{}) {
	n := offset + 1
	for col, val := range fields {
		c = append(c, col)
		i = append(i, "$"+strconv.Itoa(n))
		v = append(v, val)
		n++
	}
	return c, i, v
}
