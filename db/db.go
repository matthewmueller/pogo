package db

import (
	"github.com/matthewmueller/pgx"
)

// DB is the common interface for database operations that can be used with
// types from schema 'public'.
//
// This should work with database/sql.DB and database/sql.Tx.
type DB interface {
	Exec(string, ...interface{}) (pgx.CommandTag, error)
	Query(string, ...interface{}) (*pgx.Rows, error)
	QueryRow(string, ...interface{}) *pgx.Row
}
