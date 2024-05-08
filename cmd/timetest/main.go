package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// Point represents a point that may be null.
type Time string

var _ pgtype.TimeScanner = (*Time)(nil)

const (
	microsecondsPerSecond = 1000000
	microsecondsPerMinute = 60 * microsecondsPerSecond
	microsecondsPerHour   = 60 * microsecondsPerMinute
	microsecondsPerDay    = 24 * microsecondsPerHour
	microsecondsPerMonth  = 30 * microsecondsPerDay
)

func (s *Time) ScanTime(t pgtype.Time) error {
	fmt.Println(t.Valid)
	usec := t.Microseconds
	hours := usec / microsecondsPerHour
	usec -= hours * microsecondsPerHour
	minutes := usec / microsecondsPerMinute
	usec -= minutes * microsecondsPerMinute
	seconds := usec / microsecondsPerSecond
	usec -= seconds * microsecondsPerSecond
	if usec > 0 {
		*s = Time(fmt.Sprintf("%02d:%02d:%02d.%06d", hours, minutes, seconds, usec))
		return nil
	}
	*s = Time(fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds))
	return nil
}

func run() error {
	conn, err := pgx.Connect(context.TODO(), "postgres://localhost:5432/pogo?sslmode=disable")
	if err != nil {
		return err
	}
	defer conn.Close(context.TODO())
	var result Time = ""
	if err := conn.QueryRow(context.TODO(), "SELECT '08:00:00'::time without time zone").Scan(&result); err != nil {
		return err
	}
	fmt.Println(Test(string(result)))
	return nil
}

func Test(t string) string {
	return t
}
