package main

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

func main() {
	url := os.Getenv("POSTGRES_URL")

	cfg, err := pgx.ParseURI(url)
	if err != nil {
		panic(err)
	}

	conn, err := pgx.Connect(cfg)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	rows, err := conn.Query(`select * from jack.events where "time" is null`, nil)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		fmt.Println("row!")
	}
	if rows.Err() != nil {
		panic(rows.Err())
	}
}
