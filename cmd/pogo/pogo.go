package main

import (
	"flag"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/pogo"
)

var dburl = flag.String("db", "", "database")

func main() {
	log.SetHandler(text.New(os.Stderr))
	flag.Parse()

	if *dburl == "" {
		log.Fatal("pogo needs a database connection string")
	}

	config, err := pgx.ParseURI(*dburl)
	if err != nil {
		log.WithError(err).Fatal("postgres uri invalid")
	}

	db, err := pgx.Connect(config)
	if err != nil {
		log.WithError(err).Fatal("unable to connect to postgres")
	}

	err = pogo.Generate(db)
	if err != nil {
		log.WithError(err).Fatal("unable to generate models")
	}
}
