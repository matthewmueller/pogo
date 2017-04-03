package model_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/caarlos0/env"
	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/manual"
	"github.com/stretchr/testify/assert"
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

func TestTeamCreate(t *testing.T) {
	bytes := []byte(`
  {
    "slack_team_id": "T123123",
    "slack_team_access_token": "T123123",
    "slack_bot_access_token": "T123123",
    "slack_bot_id": "U123123",
    "team_name": "Test Team",
    "scope": ["email"],
    "email": "testteam@gmail.com"
  }
  `)

	var team model.Team
	err := json.Unmarshal(bytes, &team)
	if err != nil {
		t.Fatal(err)
	}

	s, err := team.Insert(db)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, s.SlackTeamID, "T123123")
	assert.Equal(t, s.SlackTeamAccessToken, "T123123")
	assert.Equal(t, s.SlackBotAccessToken, "T123123")
	assert.Equal(t, s.SlackBotID, "T123123")
	assert.Equal(t, s.TeamName, "Test Team")
	assert.Equal(t, s.Scope, []string{"email"})
	assert.Equal(t, s.Email, "email")
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
