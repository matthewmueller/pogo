package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// Open a URL
func Open(uri string) (*DB, error) {
	conn, err := pgx.Connect(context.TODO(), uri)
	if err != nil {
		return nil, err
	}
	return &DB{
		conn,
	}, nil
}

// DB is a connection
type DB struct {
	*pgx.Conn
}
