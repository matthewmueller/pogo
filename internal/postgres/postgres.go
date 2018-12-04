package postgres

import "github.com/jackc/pgx"

// Open a URL
func Open(uri string) (*DB, error) {
	cfg, err := pgx.ParseConnectionString(uri)
	if err != nil {
		return nil, err
	}
	conn, err := pgx.Connect(cfg)
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
