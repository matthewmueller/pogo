package main

import (
	"log"
	"os"

	"github.com/matthewmueller/commander"
	"github.com/matthewmueller/pogo"
)

func main() {
	cmd := commander.New("pogo", "ORM code generator for postgres")

	// flags
	url := cmd.Flag("db", "database connection string").Required().String()
	schema := cmd.Flag("schema", "database schema").Default("public").String()
	out := cmd.Flag("dir", "output directory to write to").Default("pogo").String()

	// run the generator
	cmd.Run(func() error {
		return pogo.Generate(*url, *out, *schema)
	})

	if err := cmd.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
