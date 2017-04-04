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

func init() {
	log.SetHandler(text.New(os.Stderr))
	model.XOLog = log.Infof

	var ev environment
	err := env.Parse(&ev)
	if err != nil {
		log.WithError(err).Fatal("unable to parse env variables")
	}

	config, err := pgx.ParseURI(ev.PostgresURL)
	if err != nil {
		log.WithError(err).Fatal("postgres uri invalid")
	}

	db, err = pgx.Connect(config)
	if err != nil {
		log.WithError(err).Fatal("unable to connect to postgres")
	}

	_, err = db.Exec(`
		delete from jack.teams where slack_team_id = 'T123123'
	`)
	if err != nil {
		log.WithError(err).Fatal("unable to delete team")
	}
}

func TestTeamCreate(t *testing.T) {
	bytes := []byte(`
  {
		"id": "54a93908-08d8-4eb6-8cff-7a232aace285",
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

	teams := model.NewTeam(db)
	s, err := teams.Insert(&team)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, *s.SlackTeamID, "T123123")
	assert.Equal(t, *s.SlackTeamAccessToken, "T123123")
	assert.Equal(t, *s.SlackBotAccessToken, "T123123")
	assert.Equal(t, *s.SlackBotID, "U123123")
	assert.Equal(t, *s.TeamName, "Test Team")
	assert.Equal(t, *s.Scope, []string{"email"})
	assert.Equal(t, *s.Email, "testteam@gmail.com")
	assert.Equal(t, *s.Active, true)
	assert.Equal(t, *s.FreeTeammates, 4)
	assert.Equal(t, *s.CostPerUser, 1)
}

func TestTeamUpdate(t *testing.T) {
	bytes := []byte(`
  {
		"id": "54a93908-08d8-4eb6-8cff-7a232aace285",
    "slack_bot_id": "U123123",
    "scope": ["email", "another"],
    "email": "matt@gmail.com",
		"stripe_id": "abc123"
  }
  `)

	var team model.Team
	err := json.Unmarshal(bytes, &team)
	if err != nil {
		t.Fatal(err)
	}

	teams := model.NewTeam(db)
	s, err := teams.Update(team.ID, &team)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, *s.SlackTeamID, "T123123")
	assert.Equal(t, *s.SlackTeamAccessToken, "T123123")
	assert.Equal(t, *s.SlackBotAccessToken, "T123123")
	assert.Equal(t, *s.SlackBotID, "U123123")
	assert.Equal(t, *s.TeamName, "Test Team")
	assert.Equal(t, *s.Scope, []string{"email", "another"})
	assert.Equal(t, *s.Email, "matt@gmail.com")
	assert.Equal(t, *s.StripeID, "abc123")
}

func TestTeamFind(t *testing.T) {
	bytes := []byte(`
  {
		"id": "54a93908-08d8-4eb6-8cff-7a232aace285"
  }
  `)

	var team model.Team
	err := json.Unmarshal(bytes, &team)
	if err != nil {
		t.Fatal(err)
	}

	teams := model.NewTeam(db)
	s, err := teams.Find(team.ID)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, *s.SlackTeamID, "T123123")
	assert.Equal(t, *s.SlackTeamAccessToken, "T123123")
	assert.Equal(t, *s.SlackBotAccessToken, "T123123")
	assert.Equal(t, *s.SlackBotID, "U123123")
	assert.Equal(t, *s.TeamName, "Test Team")
	assert.Equal(t, *s.Scope, []string{"email", "another"})
	assert.Equal(t, *s.Email, "matt@gmail.com")
	assert.Equal(t, *s.StripeID, "abc123")
}

func TestTeamDelete(t *testing.T) {
	bytes := []byte(`
  {
		"id": "54a93908-08d8-4eb6-8cff-7a232aace285"
  }
  `)

	var team model.Team
	err := json.Unmarshal(bytes, &team)
	if err != nil {
		t.Fatal(err)
	}

	teams := model.NewTeam(db)
	err = teams.Delete(team.ID)
	if err != nil {
		t.Fatal(err)
	}
}

// func TestTeamDelete(t *testing.T) {
// 	bytes := []byte(`
//   {
// 		"id": "54a93908-08d8-4eb6-8cff-7a232aace285",
//     "slack_bot_id": "U123123",
//     "scope": ["email", "another"],
//     "email": "matt@gmail.com"
//   }
//   `)
//
// 	var team model.Team
// 	err := json.Unmarshal(bytes, &team)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	teams := model.NewTeam(db)
// 	s, err := teams.Update(team.ID, &team)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	assert.Equal(t, *s.SlackTeamID, "T123123")
// 	assert.Equal(t, *s.SlackTeamAccessToken, "T123123")
// 	assert.Equal(t, *s.SlackBotAccessToken, "T123123")
// 	assert.Equal(t, *s.SlackBotID, "U123123")
// 	assert.Equal(t, *s.TeamName, "Test Team")
// 	assert.Equal(t, *s.Scope, []string{"email", "another"})
// 	assert.Equal(t, *s.Email, "matt@gmail.com")
// }

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
