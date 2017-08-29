package teams_test

import (
	"encoding/json"
	"testing"

	"github.com/apex/log"
	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/jack"
	"github.com/matthewmueller/pogo/jack/teams"
	uuid "github.com/satori/go.uuid"
)

func DB(t *testing.T) (jack.DB, func()) {
	config, err := pgx.ParseURI("postgres://localhost:5432/pogo?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	db, err := pgx.Connect(config)
	if err != nil {
		t.Fatal(err)
	}

	return db, func() {
		if e := db.Close(); e != nil {
			t.Fatal(e)
		}
	}
}

func TestInsert(t *testing.T) {
	db, close := DB(t)
	defer close()

	team := teams.New().
		TeamName(uuid.NewV4().String()).
		Email(uuid.NewV4().String()).
		SlackTeamID(uuid.NewV4().String()).
		SlackTeamAccessToken(uuid.NewV4().String()).
		SlackBotAccessToken(uuid.NewV4().String()).
		SlackBotID(uuid.NewV4().String()).
		Active(true).
		CostPerUser(1)

	tm, err := teams.Insert(db, team)
	if err != nil {
		t.Fatal(err)
	}

	buf, err := json.Marshal(tm)
	if err != nil {
		t.Fatal(err)
	}

	log.Infof("buf %s", buf)
}
