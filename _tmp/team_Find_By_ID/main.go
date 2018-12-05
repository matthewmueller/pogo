package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx"
	team "github.com/matthewmueller/pogo/pogo/team"
)

func main() {
	now := time.Date(2018, 9, 5, 0, 0, 0, 0, time.UTC)
	_ = now

	cfg, err := pgx.ParseConnectionString("postgres://localhost:5432/pogo?sslmode=disable")
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

	actual, err := team.FindByID(db, 2)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	buf, err := json.Marshal(actual)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	fmt.Fprintf(os.Stdout, "%s", string(buf))
}
