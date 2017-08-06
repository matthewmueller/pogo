package pogo

import (
	"github.com/apex/log"
	"github.com/matthewmueller/pogo/db"
	uuid "github.com/satori/go.uuid"
)

// Client is the struct containing all our models
type Client struct {
	Teams *Teams
	Team  *Team
}

type WhereCondition struct {
	Condition string
	Params    []interface{}
}

func Where(condition string, params ...interface{}) *WhereCondition {
	return &WhereCondition{
		Condition: condition,
		Params:    params,
	}
}

// New database client
func New(db db.DB) *Client {
	return &Client{
		Teams: team(db),
		Team:  &Team{},
	}
}

// model func
type Teams struct {
	db db.DB
}

func team(db db.DB) *Teams {
	return &Teams{db: db}
}

// functional options
type Team struct{}

func (*teamSettings) fields() map[string]interface{} {
	return nil
}

// input (private)
type teamSettings struct {
	ID *uuid.UUID
}

// output (public)
type TeamResult struct {
	ID uuid.UUID
}

func Test(db db.DB) {
	model := New(db)
	model.Teams.Insert(
		model.Team.ID(uuid.NewV4()),
	)
}

func (*Team) ID(id uuid.UUID) func(*teamSettings) {
	return func(team *teamSettings) {
		team.ID = &id
	}
}

func (teams *Teams) UpdateMany(where *WhereCondition) ([]Team, error) {
	return nil, nil
}

// func (standupsteammates *StandupsTeammates) Find(standupID *uuid.UUID, teammateID *uuid.UUID) (standupteammate *StandupTeammate, err error) {
// 	return nil, nil
// }
// func (standupsteammates *StandupsTeammates) Insert(standupteammate StandupTeammate) (*StandupTeammate, error) {
// 	return nil, nil
// }
// func (standupsteammates *StandupsTeammates) Update(standupID *uuid.UUID, teammateID *uuid.UUID, standupteammate StandupTeammate) (*StandupTeammate, error) {
// 	return nil, nil
// }
// func (standupsteammates *StandupsTeammates) Delete(standupID *uuid.UUID, teammateID *uuid.UUID) error {
// 	return nil
// }

func (teams *Teams) Find(id uuid.UUID) (team *TeamResult, err error) {
	return team, nil
}

func (teams *Teams) FindBySlackBotAccessToken(slackBotAccessToken *string) (team *Team, err error) {
	return nil, nil
}

func (t *Teams) Insert(fns ...func(*teamSettings)) (team Team, err error) {
	ts := &teamSettings{}
	for _, fn := range fns {
		fn(ts)
	}
	fields := ts.fields()
	log.Infof("%+v", fields)
	return team, nil
}

func (teams *Teams) Update(id uuid.UUID, fns ...func(*teamSettings)) (team Team, err error) {
	return team, nil
}

func (teams *Teams) FindBySlackTeamAccessToken(slackTeamAccessToken *string) (team *Team, err error) {
	return nil, nil
}
func (teams *Teams) FindBySlackTeamID(slackTeamID *string) (team *Team, err error) {
	return nil, nil
}
func (teams *Teams) FindMany(condition string, params ...interface{}) ([]*Team, error) {
	return nil, nil
}
func (teams *Teams) FindOne(condition string, params ...interface{}) (team *Team, err error) {
	return nil, nil
}
func (teams *Teams) Update(id *uuid.UUID, team Team) (*Team, error) {
	return nil, nil
}
func (teams *Teams) UpdateBySlackBotAccessToken(slackBotAccessToken *string, team Team) (*Team, error) {
	return nil, nil
}
func (teams *Teams) UpdateBySlackTeamAccessToken(slackTeamAccessToken *string, team Team) (*Team, error) {
	return nil, nil
}
func (teams *Teams) UpdateBySlackTeamID(team Team, slackTeamID *string) (*Team, error) {
	return nil, nil
}

func (teams *Teams) Delete(id *uuid.UUID) error {
	return nil
}
func (teams *Teams) DeleteBySlackBotAccessToken(slackBotAccessToken *string) error {
	return nil
}
func (teams *Teams) DeleteBySlackTeamAccessToken(slackTeamAccessToken *string) error {
	return nil
}
func (teams *Teams) DeleteBySlackTeamID(slackTeamID *string) error {
	return nil
}
func (teams *Teams) DeleteMany(condition string, params ...interface{}) error {
	return nil
}
func (teams *Teams) Upsert(team Team, action string) (*Team, error) {
	return nil, nil
}
func (teams *Teams) UpsertBySlackBotAccessToken(team Team, action string) (*Team, error) {
	return nil, nil
}
func (teams *Teams) UpsertBySlackTeamAccessToken(team Team, action string) (*Team, error) {
	return nil, nil
}
func (teams *Teams) UpsertBySlackTeamID(team Team, action string) (*Team, error) {
	return nil, nil
}
