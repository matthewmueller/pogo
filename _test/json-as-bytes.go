package main

import (
	"github.com/apex/log"
	"github.com/matthewmueller/pgx"
	"github.com/pkg/errors"
)

type User struct {
	Data *[]byte `json:"data"`
}

func main() {
	data := []byte(`{ "name": "matt", "age": 27 }`)
	user := User{
		Data: &data,
	}

	db, err := Connect("postgres://localhost:5432/pogo?sslmode=disable")
	if err != nil {
		log.WithError(err).Fatal("unable to connect to pg")
	}

	var u User
	err = db.QueryRow("insert into users (data) values ($1) returning data", &user.Data).Scan(&u.Data)
	if err != nil {
		log.WithError(err).Fatal("unable to insert")
	}
}

// Connect to postgres
func Connect(conn string) (db *pgx.Conn, err error) {
	config, err := pgx.ParseURI(conn)
	if err != nil {
		return db, errors.Wrap(err, "postgres uri invalid")
	}

	db, err = pgx.Connect(config)
	if err != nil {
		return db, errors.Wrap(err, "unable to connect to postgres")
	}

	return db, nil
}
