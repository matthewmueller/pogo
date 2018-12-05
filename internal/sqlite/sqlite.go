package sqlite

import (
	"database/sql"

	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// New DB
func New(db *sql.DB) *DB {
	return &DB{db}
}

// Open the database
func Open(path string) (*DB, error) {
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return &DB{conn}, nil
}

// DB is a connection
type DB struct {
	*sql.DB
}
