package pogo

import (
	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/_examples/fluent/pogo/team"
)

// // Pogo st
// type Pogo struct{}

// Team fn
func Team(db *pgx.Conn) *team.Team {
	return &team.Team{
		db: db,
	}
}
