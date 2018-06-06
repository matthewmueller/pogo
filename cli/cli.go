package cli

import (
	"context"

	"github.com/jackc/pgx"
	"github.com/matthewmueller/commander"
	"github.com/matthewmueller/pogo"
	"github.com/matthewmueller/pogo/database"
)

// Run the CLI
func Run(args []string) error {
	cmd := commander.New("pogo", "ORM code generator for postgres")

	// flags
	db := cmd.Flag("db", "database connection string").Required().String()
	schema := cmd.Flag("schema", "database schema").Default("public").String()
	dir := cmd.Flag("dir", "output directory to write to").Default("pogo").String()

	// run the generator
	cmd.Run(func() error { return run(*db, *schema, *dir) })

	return cmd.Parse(args)
}

func run(url, schema, dir string) error {
	connCfg, err := pgx.ParseURI(url)
	if err != nil {
		return err
	}

	conn, err := pgx.Connect(connCfg)
	if err != nil {
		return err
	}

	db := &database.Postgres{Conn: conn}

	return pogo.
		New(db, schema, dir).
		Run(context.TODO())
}
