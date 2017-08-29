package standupsteammates_test

import (
	"testing"

	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/testjack"
	"github.com/matthewmueller/pogo/testjack/standups"
	"github.com/matthewmueller/pogo/testjack/standupsteammates"
	"github.com/matthewmueller/pogo/testjack/teammates"
	"github.com/matthewmueller/pogo/testjack/teams"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func id() string {
	return uuid.NewV4().String()
}

func DB(t *testing.T) (testjack.DB, func()) {
	config, err := pgx.ParseURI("postgres://localhost:5432/pogo?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	db, err := pgx.Connect(config)
	if err != nil {
		t.Fatal(err)
	}

	// clear out the DB before starting
	if _, e := db.Exec("DELETE FROM jack.standups_teammates WHERE true"); e != nil {
		t.Fatal(e)
	}
	if _, e := db.Exec("DELETE FROM jack.standups WHERE true"); e != nil {
		t.Fatal(e)
	}
	if _, e := db.Exec("DELETE FROM jack.teammates WHERE true"); e != nil {
		t.Fatal(e)
	}
	if _, e := db.Exec("DELETE FROM jack.teams WHERE true"); e != nil {
		t.Fatal(e)
	}

	return db, func() {
		if e := db.Close(); e != nil {
			t.Fatal(e)
		}
	}
}

func standup(t *testing.T, db testjack.DB) *standups.Standup {
	tm, err := teams.Insert(db, teams.New().
		TeamName(id()).
		Email(id()).
		SlackTeamID(id()).
		SlackTeamAccessToken(id()).
		SlackBotAccessToken(id()).
		SlackBotID(id()))
	if err != nil {
		t.Fatal(err)
	}

	s, err := standups.Insert(db, standups.New().
		Name(id()).
		SlackChannelID(id()).
		TeamID(*tm.GetID()))
	if err != nil {
		t.Fatal(err)
	}
	return s
}

func teammate(t *testing.T, db testjack.DB) *teammates.Teammate {
	te, err := teammates.Insert(db, teammates.New().SlackID(id()).Username(id()))
	if err != nil {
		t.Fatal(err)
	}
	return te
}

func TestInsert(t *testing.T) {
	db, close := DB(t)
	defer close()

	s := standup(t, db)
	te := teammate(t, db)

	ts := standupsteammates.New().StandupID(*s.GetID()).TeammateID(*te.GetID())
	st, err := standupsteammates.Insert(db, ts)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, *s.GetID(), *st.GetStandupID())
	assert.Equal(t, *te.GetID(), *st.GetTeammateID())
}

// func TestInsertWithID(t *testing.T) {
// 	db, close := DB(t)
// 	defer close()

// 	id := uuid.NewV4()
// 	teamname := uuid.NewV4().String()
// 	email := uuid.NewV4().String()
// 	teamID := uuid.NewV4().String()
// 	teamAccessToken := uuid.NewV4().String()
// 	botAccessToken := uuid.NewV4().String()
// 	botID := uuid.NewV4().String()

// 	team := teams.New().
// 		ID(id).
// 		TeamName(teamname).
// 		Email(email).
// 		SlackTeamID(teamID).
// 		SlackTeamAccessToken(teamAccessToken).
// 		SlackBotAccessToken(botAccessToken).
// 		SlackBotID(botID).
// 		Active(true).
// 		CostPerUser(1)

// 	tm, err := teams.Insert(db, team)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.Equal(t, id, *tm.GetID())
// 	assert.Equal(t, teamname, *tm.GetTeamName())
// 	assert.Equal(t, email, *tm.GetEmail())
// 	assert.Equal(t, teamID, *tm.GetSlackTeamID())
// 	assert.Equal(t, teamAccessToken, *tm.GetSlackTeamAccessToken())
// 	assert.Equal(t, botAccessToken, *tm.GetSlackBotAccessToken())
// 	assert.Equal(t, botID, *tm.GetSlackBotID())
// 	assert.Equal(t, true, *tm.GetActive())
// 	assert.Equal(t, 1, *tm.GetCostPerUser())
// }

// func TestFind(t *testing.T) {
// 	db, close := DB(t)
// 	defer close()

// 	id := uuid.NewV4()
// 	teamname := uuid.NewV4().String()
// 	email := uuid.NewV4().String()
// 	teamID := uuid.NewV4().String()
// 	teamAccessToken := uuid.NewV4().String()
// 	botAccessToken := uuid.NewV4().String()
// 	botID := uuid.NewV4().String()

// 	team := teams.New().
// 		ID(id).
// 		TeamName(teamname).
// 		Email(email).
// 		SlackTeamID(teamID).
// 		SlackTeamAccessToken(teamAccessToken).
// 		SlackBotAccessToken(botAccessToken).
// 		SlackBotID(botID).
// 		Active(true).
// 		CostPerUser(1)

// 	_, err := teams.Insert(db, team)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	tm, err := teams.Find(db, id)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.Equal(t, id, *tm.GetID())
// 	assert.Equal(t, teamname, *tm.GetTeamName())
// 	assert.Equal(t, email, *tm.GetEmail())
// 	assert.Equal(t, teamID, *tm.GetSlackTeamID())
// 	assert.Equal(t, teamAccessToken, *tm.GetSlackTeamAccessToken())
// 	assert.Equal(t, botAccessToken, *tm.GetSlackBotAccessToken())
// 	assert.Equal(t, botID, *tm.GetSlackBotID())
// 	assert.Equal(t, true, *tm.GetActive())
// 	assert.Equal(t, 1, *tm.GetCostPerUser())
// }

// func TestFindBy(t *testing.T) {
// 	db, close := DB(t)
// 	defer close()

// 	id := uuid.NewV4()
// 	teamname := uuid.NewV4().String()
// 	email := uuid.NewV4().String()
// 	teamID := uuid.NewV4().String()
// 	teamAccessToken := uuid.NewV4().String()
// 	botAccessToken := uuid.NewV4().String()
// 	botID := uuid.NewV4().String()

// 	team := teams.New().
// 		ID(id).
// 		TeamName(teamname).
// 		Email(email).
// 		SlackTeamID(teamID).
// 		SlackTeamAccessToken(teamAccessToken).
// 		SlackBotAccessToken(botAccessToken).
// 		SlackBotID(botID).
// 		Active(true).
// 		CostPerUser(1)

// 	_, err := teams.Insert(db, team)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	tm, err := teams.FindBySlackTeamID(db, teamID)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.Equal(t, id, *tm.GetID())
// 	assert.Equal(t, teamname, *tm.GetTeamName())
// 	assert.Equal(t, email, *tm.GetEmail())
// 	assert.Equal(t, teamID, *tm.GetSlackTeamID())
// 	assert.Equal(t, teamAccessToken, *tm.GetSlackTeamAccessToken())
// 	assert.Equal(t, botAccessToken, *tm.GetSlackBotAccessToken())
// 	assert.Equal(t, botID, *tm.GetSlackBotID())
// 	assert.Equal(t, true, *tm.GetActive())
// 	assert.Equal(t, 1, *tm.GetCostPerUser())
// }

// func TestUpdate(t *testing.T) {
// 	db, close := DB(t)
// 	defer close()

// 	id := uuid.NewV4()
// 	teamname := uuid.NewV4().String()
// 	email := uuid.NewV4().String()
// 	teamID := uuid.NewV4().String()
// 	teamAccessToken := uuid.NewV4().String()
// 	botAccessToken := uuid.NewV4().String()
// 	botID := uuid.NewV4().String()

// 	team := teams.New().
// 		ID(id).
// 		TeamName(teamname).
// 		Email(email).
// 		SlackTeamID(teamID).
// 		SlackTeamAccessToken(teamAccessToken).
// 		SlackBotAccessToken(botAccessToken).
// 		SlackBotID(botID).
// 		Active(true).
// 		CostPerUser(1)

// 	_, err := teams.Insert(db, team)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	id2 := uuid.NewV4()
// 	teamname2 := uuid.NewV4().String()
// 	email2 := uuid.NewV4().String()
// 	teamID2 := uuid.NewV4().String()
// 	teamAccessToken2 := uuid.NewV4().String()
// 	botAccessToken2 := uuid.NewV4().String()
// 	botID2 := uuid.NewV4().String()

// 	team2 := teams.New().
// 		ID(id2).
// 		TeamName(teamname2).
// 		Email(email2).
// 		SlackTeamID(teamID2).
// 		SlackTeamAccessToken(teamAccessToken2).
// 		SlackBotAccessToken(botAccessToken2).
// 		SlackBotID(botID2).
// 		Active(false).
// 		CostPerUser(5)

// 	tm, err := teams.Update(db, id, team2)

// 	assert.Equal(t, id, *tm.GetID())
// 	assert.Equal(t, teamname2, *tm.GetTeamName())
// 	assert.Equal(t, email2, *tm.GetEmail())
// 	assert.Equal(t, teamID2, *tm.GetSlackTeamID())
// 	assert.Equal(t, teamAccessToken2, *tm.GetSlackTeamAccessToken())
// 	assert.Equal(t, botAccessToken2, *tm.GetSlackBotAccessToken())
// 	assert.Equal(t, botID2, *tm.GetSlackBotID())
// 	assert.Equal(t, false, *tm.GetActive())
// 	assert.Equal(t, 5, *tm.GetCostPerUser())
// }

// func TestUpdateBy(t *testing.T) {
// 	db, close := DB(t)
// 	defer close()

// 	id := uuid.NewV4()
// 	teamname := uuid.NewV4().String()
// 	email := uuid.NewV4().String()
// 	teamID := uuid.NewV4().String()
// 	teamAccessToken := uuid.NewV4().String()
// 	botAccessToken := uuid.NewV4().String()
// 	botID := uuid.NewV4().String()

// 	team := teams.New().
// 		ID(id).
// 		TeamName(teamname).
// 		Email(email).
// 		SlackTeamID(teamID).
// 		SlackTeamAccessToken(teamAccessToken).
// 		SlackBotAccessToken(botAccessToken).
// 		SlackBotID(botID).
// 		Active(true).
// 		CostPerUser(1)

// 	_, err := teams.Insert(db, team)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	id2 := uuid.NewV4()
// 	teamname2 := uuid.NewV4().String()
// 	email2 := uuid.NewV4().String()
// 	teamID2 := uuid.NewV4().String()
// 	teamAccessToken2 := uuid.NewV4().String()
// 	botAccessToken2 := uuid.NewV4().String()
// 	botID2 := uuid.NewV4().String()

// 	team2 := teams.New().
// 		ID(id2).
// 		TeamName(teamname2).
// 		Email(email2).
// 		SlackTeamID(teamID2).
// 		SlackTeamAccessToken(teamAccessToken2).
// 		SlackBotAccessToken(botAccessToken2).
// 		SlackBotID(botID2).
// 		Active(false).
// 		CostPerUser(5)

// 	tm, err := teams.UpdateBySlackBotAccessToken(db, botAccessToken, team2)

// 	assert.Equal(t, id2, *tm.GetID())
// 	assert.Equal(t, teamname2, *tm.GetTeamName())
// 	assert.Equal(t, email2, *tm.GetEmail())
// 	assert.Equal(t, teamID2, *tm.GetSlackTeamID())
// 	assert.Equal(t, teamAccessToken2, *tm.GetSlackTeamAccessToken())
// 	assert.Equal(t, botAccessToken, *tm.GetSlackBotAccessToken())
// 	assert.Equal(t, botID2, *tm.GetSlackBotID())
// 	assert.Equal(t, false, *tm.GetActive())
// 	assert.Equal(t, 5, *tm.GetCostPerUser())
// }

// func TestDelete(t *testing.T) {
// 	db, close := DB(t)
// 	defer close()

// 	id := uuid.NewV4()
// 	teamname := uuid.NewV4().String()
// 	email := uuid.NewV4().String()
// 	teamID := uuid.NewV4().String()
// 	teamAccessToken := uuid.NewV4().String()
// 	botAccessToken := uuid.NewV4().String()
// 	botID := uuid.NewV4().String()

// 	team := teams.New().
// 		ID(id).
// 		TeamName(teamname).
// 		Email(email).
// 		SlackTeamID(teamID).
// 		SlackTeamAccessToken(teamAccessToken).
// 		SlackBotAccessToken(botAccessToken).
// 		SlackBotID(botID).
// 		Active(true).
// 		CostPerUser(1)

// 	_, err := teams.Insert(db, team)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if e := teams.Delete(db, id); e != nil {
// 		t.Fatal(e)
// 	}

// 	_, err = teams.Find(db, id)
// 	assert.Equal(t, teams.ErrTeamNotFound, err)
// }

// func TestDeleteBy(t *testing.T) {
// 	db, close := DB(t)
// 	defer close()

// 	id := uuid.NewV4()
// 	teamname := uuid.NewV4().String()
// 	email := uuid.NewV4().String()
// 	teamID := uuid.NewV4().String()
// 	teamAccessToken := uuid.NewV4().String()
// 	botAccessToken := uuid.NewV4().String()
// 	botID := uuid.NewV4().String()

// 	team := teams.New().
// 		ID(id).
// 		TeamName(teamname).
// 		Email(email).
// 		SlackTeamID(teamID).
// 		SlackTeamAccessToken(teamAccessToken).
// 		SlackBotAccessToken(botAccessToken).
// 		SlackBotID(botID).
// 		Active(true).
// 		CostPerUser(1)

// 	_, err := teams.Insert(db, team)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if e := teams.DeleteBySlackBotAccessToken(db, botAccessToken); e != nil {
// 		t.Fatal(e)
// 	}

// 	_, err = teams.Find(db, id)
// 	assert.Equal(t, teams.ErrTeamNotFound, err)
// }

// func TestFindMany(t *testing.T) {
// 	db, close := DB(t)
// 	defer close()

// 	id := uuid.NewV4()
// 	teamname := uuid.NewV4().String()
// 	email := uuid.NewV4().String()
// 	teamID := uuid.NewV4().String()
// 	teamAccessToken := uuid.NewV4().String()
// 	botAccessToken := uuid.NewV4().String()
// 	botID := uuid.NewV4().String()

// 	team := teams.New().
// 		ID(id).
// 		TeamName(teamname).
// 		Email(email).
// 		SlackTeamID(teamID).
// 		SlackTeamAccessToken(teamAccessToken).
// 		SlackBotAccessToken(botAccessToken).
// 		SlackBotID(botID).
// 		Active(true).
// 		CostPerUser(9)

// 	_, err := teams.Insert(db, team)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	id2 := uuid.NewV4()
// 	teamname2 := uuid.NewV4().String()
// 	email2 := uuid.NewV4().String()
// 	teamID2 := uuid.NewV4().String()
// 	teamAccessToken2 := uuid.NewV4().String()
// 	botAccessToken2 := uuid.NewV4().String()
// 	botID2 := uuid.NewV4().String()

// 	team2 := teams.New().
// 		ID(id2).
// 		TeamName(teamname2).
// 		Email(email2).
// 		SlackTeamID(teamID2).
// 		SlackTeamAccessToken(teamAccessToken2).
// 		SlackBotAccessToken(botAccessToken2).
// 		SlackBotID(botID2).
// 		Active(true).
// 		CostPerUser(9)

// 	_, err = teams.Insert(db, team2)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	tms, err := teams.FindMany(db, teams.Where("cost_per_user = $1", 9))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.Equal(t, 2, len(tms))
// }

// func FindManyNoneFound(t *testing.T) {
// 	db, close := DB(t)
// 	defer close()

// 	tms, err := teams.FindMany(db, teams.Where("cost_per_user = $1", 910))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	buf, err := json.Marshal(tms)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.Equal(t, "[]", string(buf))
// }

// func TestFindOne(t *testing.T) {
// 	db, close := DB(t)
// 	defer close()

// 	id := uuid.NewV4()
// 	teamname := uuid.NewV4().String()
// 	email := uuid.NewV4().String()
// 	teamID := uuid.NewV4().String()
// 	teamAccessToken := uuid.NewV4().String()
// 	botAccessToken := uuid.NewV4().String()
// 	botID := uuid.NewV4().String()

// 	team := teams.New().
// 		ID(id).
// 		TeamName(teamname).
// 		Email(email).
// 		SlackTeamID(teamID).
// 		SlackTeamAccessToken(teamAccessToken).
// 		SlackBotAccessToken(botAccessToken).
// 		SlackBotID(botID).
// 		Active(true).
// 		CostPerUser(11)

// 	_, err := teams.Insert(db, team)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	id2 := uuid.NewV4()
// 	teamname2 := uuid.NewV4().String()
// 	email2 := uuid.NewV4().String()
// 	teamID2 := uuid.NewV4().String()
// 	teamAccessToken2 := uuid.NewV4().String()
// 	botAccessToken2 := uuid.NewV4().String()
// 	botID2 := uuid.NewV4().String()

// 	team2 := teams.New().
// 		ID(id2).
// 		TeamName(teamname2).
// 		Email(email2).
// 		SlackTeamID(teamID2).
// 		SlackTeamAccessToken(teamAccessToken2).
// 		SlackBotAccessToken(botAccessToken2).
// 		SlackBotID(botID2).
// 		Active(true).
// 		CostPerUser(11)

// 	_, err = teams.Insert(db, team2)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	tm, err := teams.FindOne(db, teams.Where("cost_per_user = $1", 11))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.Equal(t, id, *tm.GetID())
// 	assert.Equal(t, teamname, *tm.GetTeamName())
// 	assert.Equal(t, email, *tm.GetEmail())
// 	assert.Equal(t, teamID, *tm.GetSlackTeamID())
// 	assert.Equal(t, teamAccessToken, *tm.GetSlackTeamAccessToken())
// 	assert.Equal(t, botAccessToken, *tm.GetSlackBotAccessToken())
// 	assert.Equal(t, botID, *tm.GetSlackBotID())
// 	assert.Equal(t, true, *tm.GetActive())
// 	assert.Equal(t, 11, *tm.GetCostPerUser())
// }

// func TestUpdateMany(t *testing.T) {
// 	db, close := DB(t)
// 	defer close()

// 	id := uuid.NewV4()
// 	teamname := uuid.NewV4().String()
// 	email := uuid.NewV4().String()
// 	teamID := uuid.NewV4().String()
// 	teamAccessToken := uuid.NewV4().String()
// 	botAccessToken := uuid.NewV4().String()
// 	botID := uuid.NewV4().String()

// 	team := teams.New().
// 		ID(id).
// 		TeamName(teamname).
// 		Email(email).
// 		SlackTeamID(teamID).
// 		SlackTeamAccessToken(teamAccessToken).
// 		SlackBotAccessToken(botAccessToken).
// 		SlackBotID(botID).
// 		Active(true).
// 		CostPerUser(13)

// 	_, err := teams.Insert(db, team)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	id2 := uuid.NewV4()
// 	teamname2 := uuid.NewV4().String()
// 	email2 := uuid.NewV4().String()
// 	teamID2 := uuid.NewV4().String()
// 	teamAccessToken2 := uuid.NewV4().String()
// 	botAccessToken2 := uuid.NewV4().String()
// 	botID2 := uuid.NewV4().String()

// 	team2 := teams.New().
// 		ID(id2).
// 		TeamName(teamname2).
// 		Email(email2).
// 		SlackTeamID(teamID2).
// 		SlackTeamAccessToken(teamAccessToken2).
// 		SlackBotAccessToken(botAccessToken2).
// 		SlackBotID(botID2).
// 		Active(true).
// 		CostPerUser(13)

// 	_, err = teams.Insert(db, team2)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	team3 := teams.New().Email("matt@gmail.com")

// 	tms, err := teams.UpdateMany(db, teams.Where("cost_per_user = $1", 13), team3)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.Equal(t, 2, len(tms))
// 	assert.Equal(t, id, *tms[0].GetID())
// 	assert.Equal(t, "matt@gmail.com", *tms[0].GetEmail())
// 	assert.Equal(t, id2, *tms[1].GetID())
// 	assert.Equal(t, "matt@gmail.com", *tms[1].GetEmail())
// }

// func TestDeleteMany(t *testing.T) {
// 	db, close := DB(t)
// 	defer close()

// 	id := uuid.NewV4()
// 	teamname := uuid.NewV4().String()
// 	email := uuid.NewV4().String()
// 	teamID := uuid.NewV4().String()
// 	teamAccessToken := uuid.NewV4().String()
// 	botAccessToken := uuid.NewV4().String()
// 	botID := uuid.NewV4().String()

// 	team := teams.New().
// 		ID(id).
// 		TeamName(teamname).
// 		Email(email).
// 		SlackTeamID(teamID).
// 		SlackTeamAccessToken(teamAccessToken).
// 		SlackBotAccessToken(botAccessToken).
// 		SlackBotID(botID).
// 		Active(true).
// 		CostPerUser(15)

// 	_, err := teams.Insert(db, team)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	id2 := uuid.NewV4()
// 	teamname2 := uuid.NewV4().String()
// 	email2 := uuid.NewV4().String()
// 	teamID2 := uuid.NewV4().String()
// 	teamAccessToken2 := uuid.NewV4().String()
// 	botAccessToken2 := uuid.NewV4().String()
// 	botID2 := uuid.NewV4().String()

// 	team2 := teams.New().
// 		ID(id2).
// 		TeamName(teamname2).
// 		Email(email2).
// 		SlackTeamID(teamID2).
// 		SlackTeamAccessToken(teamAccessToken2).
// 		SlackBotAccessToken(botAccessToken2).
// 		SlackBotID(botID2).
// 		Active(true).
// 		CostPerUser(15)

// 	_, err = teams.Insert(db, team2)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if e := teams.DeleteMany(db, teams.Where("cost_per_user = $1", 15)); e != nil {
// 		t.Fatal(e)
// 	}

// 	tms, err := teams.FindMany(db, teams.Where("cost_per_user = $1", 15))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.Equal(t, 0, len(tms))
// }

// func TestUpsert(t *testing.T) {
// 	db, close := DB(t)
// 	defer close()

// 	id := uuid.NewV4()
// 	teamname := uuid.NewV4().String()
// 	email := uuid.NewV4().String()
// 	teamID := uuid.NewV4().String()
// 	teamAccessToken := uuid.NewV4().String()
// 	botAccessToken := uuid.NewV4().String()
// 	botID := uuid.NewV4().String()

// 	team := teams.New().
// 		ID(id).
// 		TeamName(teamname).
// 		Email(email).
// 		SlackTeamID(teamID).
// 		SlackTeamAccessToken(teamAccessToken).
// 		SlackBotAccessToken(botAccessToken).
// 		SlackBotID(botID).
// 		Active(true).
// 		CostPerUser(16)

// 	tm, err := teams.Upsert(db, team)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.Equal(t, teamname, *tm.GetTeamName())
// 	assert.Equal(t, email, *tm.GetEmail())
// 	assert.Equal(t, teamID, *tm.GetSlackTeamID())
// 	assert.Equal(t, teamAccessToken, *tm.GetSlackTeamAccessToken())
// 	assert.Equal(t, botAccessToken, *tm.GetSlackBotAccessToken())
// 	assert.Equal(t, botID, *tm.GetSlackBotID())

// 	email2 := uuid.NewV4().String()

// 	team2 := teams.New().
// 		ID(id).
// 		Email(email2).
// 		TeamName(teamname).
// 		SlackTeamID(teamID).
// 		SlackTeamAccessToken(teamAccessToken).
// 		SlackBotAccessToken(botAccessToken).
// 		SlackBotID(botID).
// 		Active(false).
// 		CostPerUser(21)

// 	tm2, err := teams.Upsert(db, team2)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.Equal(t, teamname, *tm2.GetTeamName())
// 	assert.Equal(t, email2, *tm2.GetEmail())
// 	assert.Equal(t, teamID, *tm2.GetSlackTeamID())
// 	assert.Equal(t, teamAccessToken, *tm2.GetSlackTeamAccessToken())
// 	assert.Equal(t, botAccessToken, *tm2.GetSlackBotAccessToken())
// 	assert.Equal(t, botID, *tm2.GetSlackBotID())
// 	assert.Equal(t, false, *tm2.GetActive())
// 	assert.Equal(t, 21, *tm2.GetCostPerUser())
// }

// func TestUpsertBy(t *testing.T) {
// 	db, close := DB(t)
// 	defer close()

// 	id := uuid.NewV4()
// 	teamname := uuid.NewV4().String()
// 	email := uuid.NewV4().String()
// 	teamID := uuid.NewV4().String()
// 	teamAccessToken := uuid.NewV4().String()
// 	botAccessToken := uuid.NewV4().String()
// 	botID := uuid.NewV4().String()

// 	team := teams.New().
// 		ID(id).
// 		TeamName(teamname).
// 		Email(email).
// 		SlackTeamID(teamID).
// 		SlackTeamAccessToken(teamAccessToken).
// 		SlackBotAccessToken(botAccessToken).
// 		SlackBotID(botID).
// 		Active(true).
// 		CostPerUser(16)

// 	tm, err := teams.UpsertBySlackTeamID(db, team)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.Equal(t, teamname, *tm.GetTeamName())
// 	assert.Equal(t, email, *tm.GetEmail())
// 	assert.Equal(t, teamID, *tm.GetSlackTeamID())
// 	assert.Equal(t, teamAccessToken, *tm.GetSlackTeamAccessToken())
// 	assert.Equal(t, botAccessToken, *tm.GetSlackBotAccessToken())
// 	assert.Equal(t, botID, *tm.GetSlackBotID())

// 	email2 := uuid.NewV4().String()

// 	team2 := teams.New().
// 		Email(email2).
// 		TeamName(teamname).
// 		SlackTeamID(teamID).
// 		SlackTeamAccessToken(teamAccessToken).
// 		SlackBotAccessToken(botAccessToken).
// 		SlackBotID(botID)

// 	tm2, err := teams.UpsertBySlackTeamID(db, team2)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.Equal(t, teamname, *tm2.GetTeamName())
// 	assert.Equal(t, email2, *tm2.GetEmail())
// 	assert.Equal(t, teamID, *tm2.GetSlackTeamID())
// 	assert.Equal(t, teamAccessToken, *tm2.GetSlackTeamAccessToken())
// 	assert.Equal(t, botAccessToken, *tm2.GetSlackBotAccessToken())
// 	assert.Equal(t, botID, *tm2.GetSlackBotID())
// 	assert.Equal(t, true, *tm2.GetActive())
// 	assert.Equal(t, 16, *tm2.GetCostPerUser())
// }
