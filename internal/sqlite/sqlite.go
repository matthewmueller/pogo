package sqlite

import (
	"database/sql"
	"net/url"
	"path/filepath"

	"github.com/matthewmueller/pogo"
	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// New DB
func New(db *sql.DB) *DB {
	return &DB{db}
}

// Open a URL
func Open(uri string) (*DB, error) {
	// parse the URL
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	// make path relative to the current working directory
	path := filepath.Join(".", u.Path)
	url := path + "?" + u.Query().Encode()

	// open the database
	conn, err := sql.Open("sqlite3", url)
	if err != nil {
		return nil, err
	}

	return &DB{conn}, nil
}

// DB is a connection
type DB struct {
	*sql.DB
}

var _ pogo.Driver = (*DB)(nil)

// var _ pogo.Driver = (*DB)(nil)
