package model_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/caarlos0/env"
	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/manual"
	uuid "github.com/satori/go.uuid"
)

var db *pgx.Conn

type environment struct {
	PostgresURL string `env:"POSTGRES_URL,required"`
}

func TestSetup(t *testing.T) {
	log.SetHandler(text.New(os.Stderr))
	model.XOLog = log.Infof

	var ev environment
	err := env.Parse(&ev)
	if err != nil {
		t.Fatal("unable to parse env variables")
	}

	config, err := pgx.ParseURI(ev.PostgresURL)
	if err != nil {
		t.Fatal("postgres uri invalid")
	}

	db, err = pgx.Connect(config)
	if err != nil {
		t.Fatal("unable to connect to postgres")
	}
}

func TestStandupCreate(t *testing.T) {
	name := "standup"
	slackChannelID := "C123145"
	teamID := uuid.NewV4()
	now := "11:00:00"
	timezone := "America/Los_Angeles"
	standup := model.Standup{
		Name:           &name,
		SlackChannelID: &slackChannelID,
		Time:           &now,
		TeamID:         &teamID,
		Timezone:       &timezone,
	}

	s, err := standup.Insert(db)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(s)
}

// func TestStandupCreate(t *testing.T) {
// 	name := "standup"
// 	slackChannelID := "C123145"
// 	teamID := uuid.NewV4()
// 	now := "11:00:00"
// 	timezone := "America/Los_Angeles"
// 	standup := model.Standup{
// 		Name:           &name,
// 		SlackChannelID: &slackChannelID,
// 		Time:           &now,
// 		TeamID:         &teamID,
// 		Timezone:       &timezone,
// 	}
//
// 	s, err := standup.Insert(db)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	fmt.Println(s)
// }
