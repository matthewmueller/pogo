package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jackc/pgx"
	teammate "github.com/matthewmueller/pogo/test/jack/pogo/teammate"
)

func main() {
	cfg, err := pgx.ParseConnectionString("postgres://localhost:5432/pogo-jack?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	db, err := pgx.Connect(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}
	defer db.Close()

	actual, err := teammate.Find(db, teammate.NewFilter().SlackIDIn("b", "c"))
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	buf, err := json.Marshal(actual)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	fmt.Fprintf(os.Stdout, "%s", string(buf))
}
