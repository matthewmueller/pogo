package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/pogo"
)

var dburl = flag.String("db", "", "database")
var schema = flag.String("schema", "public", "schema name")
var pathdir = flag.String("path", "model", "path to output")

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

	if *schema == "" {
		*schema = "public"
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.WithError(err).Fatal("unable to get the current working directory")
	}
	*pathdir = path.Join(cwd, *pathdir)

	output, err := pogo.Generate(db, *schema, path.Base(*pathdir))
	if err != nil {
		log.WithError(err).Fatal("unable to generate models")
	}

	err = pogo.Write(output, *pathdir)
	if err != nil {
		log.WithError(err).Fatal("unable to write out models")
	}

	fmt.Println(output)
}
