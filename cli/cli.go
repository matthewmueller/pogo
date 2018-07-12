package cli

import (
	"github.com/matthewmueller/commander"
	"github.com/matthewmueller/pogo"
)

// Run the CLI
func Run(args []string) error {
	cmd := commander.New("pogo", "ORM code generator for postgres")

	// flags
	url := cmd.Flag("db", "database connection string").Required().String()
	schema := cmd.Flag("schema", "database schema").Default("public").String()
	out := cmd.Flag("dir", "output directory to write to").Default("pogo").String()

	// run the generator
	cmd.Run(func() error { return run(*url, *schema, *out) })

	return cmd.Parse(args)
}

func run(url, schema, output string) error {
	pogo := pogo.Pogo{
		URL:    url,
		Schema: schema,
		Output: output,
	}

	return pogo.Run()
}
