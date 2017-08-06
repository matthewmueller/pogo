package main

import (
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/jackc/pgx"
)

type Team struct {
	ID *uuid.UUID
	SlackTeamID *string
	SlackTeamAccessToken *string
	SlackBotAccessToken *string
	SlackBotID *string
	TeamName *string
	Scope []string
	email
	stripe_id
	active
	free_teammates
	cost_per_user
	trial_ends
	created_at
	updated_at
}

func main() {
	log.SetHandler(text.New(os.Stderr))

	config, err := pgx.ParseURI("postgres://localhost:5432/pogo?sslmode=disable")
	if err != nil {
		log.WithError(err).Fatal("postgres uri invalid")
	}

	db, err := pgx.Connect(config)
	if err != nil {
		log.WithError(err).Fatal("unable to connect to postgres")
	}


	// sql insert query, primary key provided by sequence
	sqlstr := `
	insert into jack.teams
	values ()
	INSERT INTO {{ schema .Schema .Table.TableName }} (` + strings.Join(_c, ", ") + `)
	VALUES (` + strings.Join(_i, ", ") + `)
	RETURNING {{ fields .Columns }}`
	db.QueryRow(`
		insert into 
	`)

}
