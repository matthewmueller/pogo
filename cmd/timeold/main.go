package main

import (
	"fmt"
	"log"

	"github.com/jackc/pgx"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg, err := pgx.ParseConnectionString("postgres://localhost:5432/pogo?sslmode=disable")
	if err != nil {
		return err
	}
	conn, err := pgx.Connect(cfg)
	if err != nil {
		return err
	}
	defer conn.Close()
	result := ""
	if err := conn.QueryRow("SELECT '08:00'::time without time zone").Scan(&result); err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
