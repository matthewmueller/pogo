package main

import (
	"github.com/apex/log"
	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/_examples/fluent/pogo"
	"github.com/matthewmueller/pogo/_examples/fluent/pogo/team"
)

func main() {
	config, err := pgx.ParseURI("postgres://localhost:5432/pogo?sslmode=disable")
	if err != nil {
		log.WithError(err).Fatal("postgres uri invalid")
	}

	db, err := pgx.Connect(config)
	if err != nil {
		log.WithError(err).Fatal("unable to connect to postgres")
	}

	pogo.Team(db).Create()
	pogo.Team(db).Create(
		team.TeamName("hi"),
	)

}
