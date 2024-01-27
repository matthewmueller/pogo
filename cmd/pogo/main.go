package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/matthewmueller/commander"
	"github.com/matthewmueller/pogo"
)

func main() {
	log := &log.Logger{
		Handler: cli.Default,
		Level:   log.InfoLevel,
	}

	cmd := commander.New("pogo", "ORM code generator for postgres")

	// flags
	dbUrl := cmd.Flag("db", "database connection string").Envar("DATABASE_URL").String()
	schema := cmd.Flag("schema", "database schema").Default("").String()
	out := cmd.Flag("dir", "output directory to write to").Default(filepath.Join("internal", "pogo")).String()

	// run the generator
	cmd.Run(func() error {
		if *dbUrl == "" {
			return fmt.Errorf("missing --db flag or DATABASE_URL environment variable")
		}
		return pogo.Generate(*dbUrl, *out, *schema)
	})

	if err := cmd.Parse(os.Args[1:]); err != nil {
		log.Fatal(err.Error())
	}
}
