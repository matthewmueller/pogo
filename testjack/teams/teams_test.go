package teams_test

import (
	"testing"

	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/testjack"
	"github.com/matthewmueller/pogo/testjack/teams"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func DB(t *testing.T) (testjack.DB, func()) {
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

	teamname := uuid.NewV4().String()
	email := uuid.NewV4().String()
	teamID := uuid.NewV4().String()
	teamAccessToken := uuid.NewV4().String()
	botAccessToken := uuid.NewV4().String()
	botID := uuid.NewV4().String()

	team := teams.New().
		TeamName(teamname).
		Email(email).
		SlackTeamID(teamID).
		SlackTeamAccessToken(teamAccessToken).
		SlackBotAccessToken(botAccessToken).
		SlackBotID(botID).
		Active(true).
		CostPerUser(1)

	tm, err := teams.Insert(db, team)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, teamname, *tm.GetTeamName())
	assert.Equal(t, email, *tm.GetEmail())
	assert.Equal(t, teamID, *tm.GetSlackTeamID())
	assert.Equal(t, teamAccessToken, *tm.GetSlackTeamAccessToken())
	assert.Equal(t, botAccessToken, *tm.GetSlackBotAccessToken())
	assert.Equal(t, botID, *tm.GetSlackBotID())
}

func TestInsertWithID(t *testing.T) {
	db, close := DB(t)
	defer close()

	id := uuid.NewV4()
	teamname := uuid.NewV4().String()
	email := uuid.NewV4().String()
	teamID := uuid.NewV4().String()
	teamAccessToken := uuid.NewV4().String()
	botAccessToken := uuid.NewV4().String()
	botID := uuid.NewV4().String()

	team := teams.New().
		ID(id).
		TeamName(teamname).
		Email(email).
		SlackTeamID(teamID).
		SlackTeamAccessToken(teamAccessToken).
		SlackBotAccessToken(botAccessToken).
		SlackBotID(botID).
		Active(true).
		CostPerUser(1)

	tm, err := teams.Insert(db, team)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, id, *tm.GetID())
	assert.Equal(t, teamname, *tm.GetTeamName())
	assert.Equal(t, email, *tm.GetEmail())
	assert.Equal(t, teamID, *tm.GetSlackTeamID())
	assert.Equal(t, teamAccessToken, *tm.GetSlackTeamAccessToken())
	assert.Equal(t, botAccessToken, *tm.GetSlackBotAccessToken())
	assert.Equal(t, botID, *tm.GetSlackBotID())
	assert.Equal(t, true, *tm.GetActive())
	assert.Equal(t, 1, *tm.GetCostPerUser())
}

func TestFind(t *testing.T) {
	db, close := DB(t)
	defer close()

	id := uuid.NewV4()
	teamname := uuid.NewV4().String()
	email := uuid.NewV4().String()
	teamID := uuid.NewV4().String()
	teamAccessToken := uuid.NewV4().String()
	botAccessToken := uuid.NewV4().String()
	botID := uuid.NewV4().String()

	team := teams.New().
		ID(id).
		TeamName(teamname).
		Email(email).
		SlackTeamID(teamID).
		SlackTeamAccessToken(teamAccessToken).
		SlackBotAccessToken(botAccessToken).
		SlackBotID(botID).
		Active(true).
		CostPerUser(1)

	_, err := teams.Insert(db, team)
	if err != nil {
		t.Fatal(err)
	}

	tm, err := teams.Find(db, id)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, id, *tm.GetID())
	assert.Equal(t, teamname, *tm.GetTeamName())
	assert.Equal(t, email, *tm.GetEmail())
	assert.Equal(t, teamID, *tm.GetSlackTeamID())
	assert.Equal(t, teamAccessToken, *tm.GetSlackTeamAccessToken())
	assert.Equal(t, botAccessToken, *tm.GetSlackBotAccessToken())
	assert.Equal(t, botID, *tm.GetSlackBotID())
	assert.Equal(t, true, *tm.GetActive())
	assert.Equal(t, 1, *tm.GetCostPerUser())
}

func TestFindBy(t *testing.T) {
	db, close := DB(t)
	defer close()

	id := uuid.NewV4()
	teamname := uuid.NewV4().String()
	email := uuid.NewV4().String()
	teamID := uuid.NewV4().String()
	teamAccessToken := uuid.NewV4().String()
	botAccessToken := uuid.NewV4().String()
	botID := uuid.NewV4().String()

	team := teams.New().
		ID(id).
		TeamName(teamname).
		Email(email).
		SlackTeamID(teamID).
		SlackTeamAccessToken(teamAccessToken).
		SlackBotAccessToken(botAccessToken).
		SlackBotID(botID).
		Active(true).
		CostPerUser(1)

	_, err := teams.Insert(db, team)
	if err != nil {
		t.Fatal(err)
	}

	tm, err := teams.FindBySlackTeamID(db, teamID)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, id, *tm.GetID())
	assert.Equal(t, teamname, *tm.GetTeamName())
	assert.Equal(t, email, *tm.GetEmail())
	assert.Equal(t, teamID, *tm.GetSlackTeamID())
	assert.Equal(t, teamAccessToken, *tm.GetSlackTeamAccessToken())
	assert.Equal(t, botAccessToken, *tm.GetSlackBotAccessToken())
	assert.Equal(t, botID, *tm.GetSlackBotID())
	assert.Equal(t, true, *tm.GetActive())
	assert.Equal(t, 1, *tm.GetCostPerUser())
}

func TestUpdate(t *testing.T) {
	db, close := DB(t)
	defer close()

	id := uuid.NewV4()
	teamname := uuid.NewV4().String()
	email := uuid.NewV4().String()
	teamID := uuid.NewV4().String()
	teamAccessToken := uuid.NewV4().String()
	botAccessToken := uuid.NewV4().String()
	botID := uuid.NewV4().String()

	team := teams.New().
		ID(id).
		TeamName(teamname).
		Email(email).
		SlackTeamID(teamID).
		SlackTeamAccessToken(teamAccessToken).
		SlackBotAccessToken(botAccessToken).
		SlackBotID(botID).
		Active(true).
		CostPerUser(1)

	_, err := teams.Insert(db, team)
	if err != nil {
		t.Fatal(err)
	}

	id2 := uuid.NewV4()
	teamname2 := uuid.NewV4().String()
	email2 := uuid.NewV4().String()
	teamID2 := uuid.NewV4().String()
	teamAccessToken2 := uuid.NewV4().String()
	botAccessToken2 := uuid.NewV4().String()
	botID2 := uuid.NewV4().String()

	team2 := teams.New().
		ID(id2).
		TeamName(teamname2).
		Email(email2).
		SlackTeamID(teamID2).
		SlackTeamAccessToken(teamAccessToken2).
		SlackBotAccessToken(botAccessToken2).
		SlackBotID(botID2).
		Active(false).
		CostPerUser(5)

	tm, err := teams.Update(db, id, team2)

	assert.Equal(t, id, *tm.GetID())
	assert.Equal(t, teamname2, *tm.GetTeamName())
	assert.Equal(t, email2, *tm.GetEmail())
	assert.Equal(t, teamID2, *tm.GetSlackTeamID())
	assert.Equal(t, teamAccessToken2, *tm.GetSlackTeamAccessToken())
	assert.Equal(t, botAccessToken2, *tm.GetSlackBotAccessToken())
	assert.Equal(t, botID2, *tm.GetSlackBotID())
	assert.Equal(t, false, *tm.GetActive())
	assert.Equal(t, 5, *tm.GetCostPerUser())
}

func TestUpdateBy(t *testing.T) {
	db, close := DB(t)
	defer close()

	id := uuid.NewV4()
	teamname := uuid.NewV4().String()
	email := uuid.NewV4().String()
	teamID := uuid.NewV4().String()
	teamAccessToken := uuid.NewV4().String()
	botAccessToken := uuid.NewV4().String()
	botID := uuid.NewV4().String()

	team := teams.New().
		ID(id).
		TeamName(teamname).
		Email(email).
		SlackTeamID(teamID).
		SlackTeamAccessToken(teamAccessToken).
		SlackBotAccessToken(botAccessToken).
		SlackBotID(botID).
		Active(true).
		CostPerUser(1)

	_, err := teams.Insert(db, team)
	if err != nil {
		t.Fatal(err)
	}

	id2 := uuid.NewV4()
	teamname2 := uuid.NewV4().String()
	email2 := uuid.NewV4().String()
	teamID2 := uuid.NewV4().String()
	teamAccessToken2 := uuid.NewV4().String()
	botAccessToken2 := uuid.NewV4().String()
	botID2 := uuid.NewV4().String()

	team2 := teams.New().
		ID(id2).
		TeamName(teamname2).
		Email(email2).
		SlackTeamID(teamID2).
		SlackTeamAccessToken(teamAccessToken2).
		SlackBotAccessToken(botAccessToken2).
		SlackBotID(botID2).
		Active(false).
		CostPerUser(5)

	tm, err := teams.UpdateBySlackBotAccessToken(db, botAccessToken, team2)

	assert.Equal(t, id2, *tm.GetID())
	assert.Equal(t, teamname2, *tm.GetTeamName())
	assert.Equal(t, email2, *tm.GetEmail())
	assert.Equal(t, teamID2, *tm.GetSlackTeamID())
	assert.Equal(t, teamAccessToken2, *tm.GetSlackTeamAccessToken())
	assert.Equal(t, botAccessToken, *tm.GetSlackBotAccessToken())
	assert.Equal(t, botID2, *tm.GetSlackBotID())
	assert.Equal(t, false, *tm.GetActive())
	assert.Equal(t, 5, *tm.GetCostPerUser())
}

func TestDelete(t *testing.T) {
	db, close := DB(t)
	defer close()

	id := uuid.NewV4()
	teamname := uuid.NewV4().String()
	email := uuid.NewV4().String()
	teamID := uuid.NewV4().String()
	teamAccessToken := uuid.NewV4().String()
	botAccessToken := uuid.NewV4().String()
	botID := uuid.NewV4().String()

	team := teams.New().
		ID(id).
		TeamName(teamname).
		Email(email).
		SlackTeamID(teamID).
		SlackTeamAccessToken(teamAccessToken).
		SlackBotAccessToken(botAccessToken).
		SlackBotID(botID).
		Active(true).
		CostPerUser(1)

	_, err := teams.Insert(db, team)
	if err != nil {
		t.Fatal(err)
	}

	if e := teams.Delete(db, id); e != nil {
		t.Fatal(e)
	}

	_, err = teams.Find(db, id)
	assert.Equal(t, teams.ErrTeamNotFound, err)
}

func TestDeleteBy(t *testing.T) {
	db, close := DB(t)
	defer close()

	id := uuid.NewV4()
	teamname := uuid.NewV4().String()
	email := uuid.NewV4().String()
	teamID := uuid.NewV4().String()
	teamAccessToken := uuid.NewV4().String()
	botAccessToken := uuid.NewV4().String()
	botID := uuid.NewV4().String()

	team := teams.New().
		ID(id).
		TeamName(teamname).
		Email(email).
		SlackTeamID(teamID).
		SlackTeamAccessToken(teamAccessToken).
		SlackBotAccessToken(botAccessToken).
		SlackBotID(botID).
		Active(true).
		CostPerUser(1)

	_, err := teams.Insert(db, team)
	if err != nil {
		t.Fatal(err)
	}

	if e := teams.DeleteBySlackBotAccessToken(db, botAccessToken); e != nil {
		t.Fatal(e)
	}

	_, err = teams.Find(db, id)
	assert.Equal(t, teams.ErrTeamNotFound, err)
}
