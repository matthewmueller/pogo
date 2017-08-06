package main

import (
	"os"
	"path"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/pogo"
)

var (
	cli = kingpin.New("pogo", "postgres model generator")

	dburl   = cli.Flag("db", "database connection string").String()
	schema  = cli.Flag("schema", "database schema").Default("public").String()
	outpath = cli.Flag("path", "output path to write to").Default("pogo").String()
)

func main() {
	log.SetHandler(text.New(os.Stderr))
	kingpin.MustParse(cli.Parse(os.Args[1:]))

	config, err := pgx.ParseURI(*dburl)
	if err != nil {
		log.WithError(err).Fatal("postgres uri invalid")
	}

	db, err := pgx.Connect(config)
	if err != nil {
		log.WithError(err).Fatal("unable to connect to postgres")
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.WithError(err).Fatal("unable to get the current working directory")
	}
	*outpath = path.Join(cwd, *outpath)

	output, err := pogo.Generate(db, &pogo.Settings{
		Schema:  *schema,
		Package: path.Base(*outpath),
	})
	if err != nil {
		log.WithError(err).Fatal("unable to generate models")
		return
	}

	if e := pogo.Write(output, *outpath); e != nil {
		log.WithError(err).Fatal("unable to write out models")
	}
}
