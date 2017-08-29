package teams

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/jack"
	uuid "github.com/satori/go.uuid"
)

// GENERATED BY POGO. DO NOT EDIT.

// ErrTeamNotFound returned if the team is not found
var ErrTeamNotFound = errors.New("team not found")

// columns in `jack.teams`
type columns struct {
	ID                   *string    `json:"id,omitempty"`
	SlackTeamID          *string    `json:"slack_team_id,omitempty"`
	SlackTeamAccessToken *string    `json:"slack_team_access_token,omitempty"`
	SlackBotAccessToken  *string    `json:"slack_bot_access_token,omitempty"`
	SlackBotID           *string    `json:"slack_bot_id,omitempty"`
	TeamName             *string    `json:"team_name,omitempty"`
	Scope                *[]string  `json:"scope,omitempty"`
	Email                *string    `json:"email,omitempty"`     // user email
	StripeID             *string    `json:"stripe_id,omitempty"` // user stripe id
	Active               *bool      `json:"active,omitempty"`
	FreeTeammates        *int       `json:"free_teammates,omitempty"`
	CostPerUser          *int       `json:"cost_per_user,omitempty"`
	TrialEnds            *time.Time `json:"trial_ends,omitempty"`
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
}

// Team fluent API
type Team struct {
	columns *columns
}

// New `jack.teams` API
func New() *Team {
	return &Team{&columns{}}
}

// ID sets the `id`
func (team *Team) ID(id uuid.UUID) *Team {
	*team.columns.ID = id.String()
	return team
}

// GetID returns the `id` if set
func (team *Team) GetID() (id *uuid.UUID) {
	if team.columns.ID == nil {
		return nil
	}

	_u, err := uuid.FromString(*team.columns.ID)
	if err != nil {
		return nil
	}

	return &_u
}

// SlackTeamID sets the `slack_team_id`
func (team *Team) SlackTeamID(slackTeamID string) *Team {
	team.columns.SlackTeamID = &slackTeamID
	return team
}

// GetSlackTeamID returns the `slack_team_id` if set
func (team *Team) GetSlackTeamID() (slackTeamID *string) {
	return team.columns.SlackTeamID
}

// SlackTeamAccessToken sets the `slack_team_access_token`
func (team *Team) SlackTeamAccessToken(slackTeamAccessToken string) *Team {
	team.columns.SlackTeamAccessToken = &slackTeamAccessToken
	return team
}

// GetSlackTeamAccessToken returns the `slack_team_access_token` if set
func (team *Team) GetSlackTeamAccessToken() (slackTeamAccessToken *string) {
	return team.columns.SlackTeamAccessToken
}

// SlackBotAccessToken sets the `slack_bot_access_token`
func (team *Team) SlackBotAccessToken(slackBotAccessToken string) *Team {
	team.columns.SlackBotAccessToken = &slackBotAccessToken
	return team
}

// GetSlackBotAccessToken returns the `slack_bot_access_token` if set
func (team *Team) GetSlackBotAccessToken() (slackBotAccessToken *string) {
	return team.columns.SlackBotAccessToken
}

// SlackBotID sets the `slack_bot_id`
func (team *Team) SlackBotID(slackBotID string) *Team {
	team.columns.SlackBotID = &slackBotID
	return team
}

// GetSlackBotID returns the `slack_bot_id` if set
func (team *Team) GetSlackBotID() (slackBotID *string) {
	return team.columns.SlackBotID
}

// TeamName sets the `team_name`
func (team *Team) TeamName(teamName string) *Team {
	team.columns.TeamName = &teamName
	return team
}

// GetTeamName returns the `team_name` if set
func (team *Team) GetTeamName() (teamName *string) {
	return team.columns.TeamName
}

// Scope sets the `scope`
func (team *Team) Scope(scope []string) *Team {
	team.columns.Scope = &scope
	return team
}

// GetScope returns the `scope` if set
func (team *Team) GetScope() (scope *[]string) {
	return team.columns.Scope
}

// Email sets the `email`
func (team *Team) Email(email string) *Team {
	team.columns.Email = &email
	return team
}

// GetEmail returns the `email` if set
func (team *Team) GetEmail() (email *string) {
	return team.columns.Email
}

// StripeID sets the `stripe_id`
func (team *Team) StripeID(stripeID string) *Team {
	team.columns.StripeID = &stripeID
	return team
}

// GetStripeID returns the `stripe_id` if set
func (team *Team) GetStripeID() (stripeID *string) {
	return team.columns.StripeID
}

// Active sets the `active`
func (team *Team) Active(active bool) *Team {
	team.columns.Active = &active
	return team
}

// GetActive returns the `active` if set
func (team *Team) GetActive() (active *bool) {
	return team.columns.Active
}

// FreeTeammates sets the `free_teammates`
func (team *Team) FreeTeammates(freeTeammates int) *Team {
	team.columns.FreeTeammates = &freeTeammates
	return team
}

// GetFreeTeammates returns the `free_teammates` if set
func (team *Team) GetFreeTeammates() (freeTeammates *int) {
	return team.columns.FreeTeammates
}

// CostPerUser sets the `cost_per_user`
func (team *Team) CostPerUser(costPerUser int) *Team {
	team.columns.CostPerUser = &costPerUser
	return team
}

// GetCostPerUser returns the `cost_per_user` if set
func (team *Team) GetCostPerUser() (costPerUser *int) {
	return team.columns.CostPerUser
}

// TrialEnds sets the `trial_ends`
func (team *Team) TrialEnds(trialEnds time.Time) *Team {
	team.columns.TrialEnds = &trialEnds
	return team
}

// GetTrialEnds returns the `trial_ends` if set
func (team *Team) GetTrialEnds() (trialEnds *time.Time) {
	return team.columns.TrialEnds
}

// CreatedAt sets the `created_at`
func (team *Team) CreatedAt(createdAt time.Time) *Team {
	team.columns.CreatedAt = &createdAt
	return team
}

// GetCreatedAt returns the `created_at` if set
func (team *Team) GetCreatedAt() (createdAt *time.Time) {
	return team.columns.CreatedAt
}

// UpdatedAt sets the `updated_at`
func (team *Team) UpdatedAt(updatedAt time.Time) *Team {
	team.columns.UpdatedAt = &updatedAt
	return team
}

// GetUpdatedAt returns the `updated_at` if set
func (team *Team) GetUpdatedAt() (updatedAt *time.Time) {
	return team.columns.UpdatedAt
}

// MarshalJSON marshals the `team` into JSON
func (team *Team) MarshalJSON() ([]byte, error) {
	return json.Marshal(team.columns)
}

// UnmarshalJSON unmarshals json to a `team`
func (team *Team) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, team.columns)
}

func (team *Team) String() string {
	return "team TODO"
}

// get all the non-nil columns
func getColumns(team *Team) map[string]interface{} {
	columns := make(map[string]interface{})

	if team.columns.ID != nil {
		columns["id"] = *team.columns.ID
	}
	if team.columns.SlackTeamID != nil {
		columns["slack_team_id"] = *team.columns.SlackTeamID
	}
	if team.columns.SlackTeamAccessToken != nil {
		columns["slack_team_access_token"] = *team.columns.SlackTeamAccessToken
	}
	if team.columns.SlackBotAccessToken != nil {
		columns["slack_bot_access_token"] = *team.columns.SlackBotAccessToken
	}
	if team.columns.SlackBotID != nil {
		columns["slack_bot_id"] = *team.columns.SlackBotID
	}
	if team.columns.TeamName != nil {
		columns["team_name"] = *team.columns.TeamName
	}
	if team.columns.Scope != nil {
		columns["scope"] = *team.columns.Scope
	}
	if team.columns.Email != nil {
		columns["email"] = *team.columns.Email
	}
	if team.columns.StripeID != nil {
		columns["stripe_id"] = *team.columns.StripeID
	}
	if team.columns.Active != nil {
		columns["active"] = *team.columns.Active
	}
	if team.columns.FreeTeammates != nil {
		columns["free_teammates"] = *team.columns.FreeTeammates
	}
	if team.columns.CostPerUser != nil {
		columns["cost_per_user"] = *team.columns.CostPerUser
	}
	if team.columns.TrialEnds != nil {
		columns["trial_ends"] = *team.columns.TrialEnds
	}
	if team.columns.CreatedAt != nil {
		columns["created_at"] = *team.columns.CreatedAt
	}
	if team.columns.UpdatedAt != nil {
		columns["updated_at"] = *team.columns.UpdatedAt
	}

	return columns
}

// Find a team by "id"
func Find(db jack.DB, id *string) (*Team, error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"
	FROM jack.teams
	WHERE "id" = $1
	`
	jack.Log(sqlstr, id)

	var cols *columns
	row := db.QueryRow(sqlstr, id)
	if e := row.Scan(&cols.ID, &cols.SlackTeamID, &cols.SlackTeamAccessToken, &cols.SlackBotAccessToken, &cols.SlackBotID, &cols.TeamName, &cols.Scope, &cols.Email, &cols.StripeID, &cols.Active, &cols.FreeTeammates, &cols.CostPerUser, &cols.TrialEnds, &cols.CreatedAt, &cols.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrTeamNotFound
		}
		return nil, e
	}

	return &Team{cols}, nil
}

// FindBySlackBotAccessToken find a team by `slack_bot_access_token`
func FindBySlackBotAccessToken(db jack.DB, slackBotAccessToken *string) (*Team, error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"
	FROM jack.teams
	WHERE "slack_bot_access_token" = $1
	`
	jack.Log(sqlstr, slackBotAccessToken)

	var cols *columns
	row := db.QueryRow(sqlstr, slackBotAccessToken)
	err := row.Scan(&cols.ID, &cols.SlackTeamID, &cols.SlackTeamAccessToken, &cols.SlackBotAccessToken, &cols.SlackBotID, &cols.TeamName, &cols.Scope, &cols.Email, &cols.StripeID, &cols.Active, &cols.FreeTeammates, &cols.CostPerUser, &cols.TrialEnds, &cols.CreatedAt, &cols.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrTeamNotFound
		}
		return nil, err
	}

	return &Team{cols}, nil
}

// FindBySlackTeamAccessToken find a team by `slack_team_access_token`
func FindBySlackTeamAccessToken(db jack.DB, slackTeamAccessToken *string) (*Team, error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"
	FROM jack.teams
	WHERE "slack_team_access_token" = $1
	`
	jack.Log(sqlstr, slackTeamAccessToken)

	var cols *columns
	row := db.QueryRow(sqlstr, slackTeamAccessToken)
	err := row.Scan(&cols.ID, &cols.SlackTeamID, &cols.SlackTeamAccessToken, &cols.SlackBotAccessToken, &cols.SlackBotID, &cols.TeamName, &cols.Scope, &cols.Email, &cols.StripeID, &cols.Active, &cols.FreeTeammates, &cols.CostPerUser, &cols.TrialEnds, &cols.CreatedAt, &cols.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrTeamNotFound
		}
		return nil, err
	}

	return &Team{cols}, nil
}

// FindBySlackTeamID find a team by `slack_team_id`
func FindBySlackTeamID(db jack.DB, slackTeamID *string) (*Team, error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"
	FROM jack.teams
	WHERE "slack_team_id" = $1
	`
	jack.Log(sqlstr, slackTeamID)

	var cols *columns
	row := db.QueryRow(sqlstr, slackTeamID)
	err := row.Scan(&cols.ID, &cols.SlackTeamID, &cols.SlackTeamAccessToken, &cols.SlackBotAccessToken, &cols.SlackBotID, &cols.TeamName, &cols.Scope, &cols.Email, &cols.StripeID, &cols.Active, &cols.FreeTeammates, &cols.CostPerUser, &cols.TrialEnds, &cols.CreatedAt, &cols.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrTeamNotFound
		}
		return nil, err
	}

	return &Team{cols}, nil
}

// FindMany find many `team`s by a given condition
func FindMany(db jack.DB, condition string, params ...interface{}) ([]*Team, error) {
	var _o []*Team

	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"
	FROM jack.teams
	WHERE ` + condition
	jack.Log(sqlstr, params...)

	rows, err := db.Query(sqlstr, params...)
	if err != nil {
		return _o, err
	}
	defer rows.Close()

	for rows.Next() {
		var cols *columns
		if e := rows.Scan(&cols.ID, &cols.SlackTeamID, &cols.SlackTeamAccessToken, &cols.SlackBotAccessToken, &cols.SlackBotID, &cols.TeamName, &cols.Scope, &cols.Email, &cols.StripeID, &cols.Active, &cols.FreeTeammates, &cols.CostPerUser, &cols.TrialEnds, &cols.CreatedAt, &cols.UpdatedAt); e != nil {
			if e == pgx.ErrNoRows {
				return _o, ErrTeamNotFound
			}
			return _o, err
		}
		_o = append(_o, &Team{cols})
	}
	if rows.Err() != nil {
		return _o, rows.Err()
	}

	// ensure we return an empty array
	// rather than nil when we marshal
	if len(_o) == 0 {
		return make([]*Team, 0), nil
	}

	return _o, nil
}

// FindOne find one team by a condition
func FindOne(db jack.DB, condition string, params ...interface{}) (*Team, error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"
	FROM jack.teams
	WHERE ` + condition
	jack.Log(sqlstr, params...)

	var cols *columns
	row := db.QueryRow(sqlstr, params...)
	if e := row.Scan(&cols.ID, &cols.SlackTeamID, &cols.SlackTeamAccessToken, &cols.SlackBotAccessToken, &cols.SlackBotID, &cols.TeamName, &cols.Scope, &cols.Email, &cols.StripeID, &cols.Active, &cols.FreeTeammates, &cols.CostPerUser, &cols.TrialEnds, &cols.CreatedAt, &cols.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrTeamNotFound
		}
		return nil, e
	}

	return &Team{cols}, nil
}

// Insert a `team` into the `jack.teams` table.
func Insert(db jack.DB, team *Team) (*Team, error) {
	// get all the non-nil columns and prepare them for the query
	_c, _i, _v := jack.Slice(getColumns(team), 0)

	// sql insert query, primary key provided by sequence
	sqlstr := `
	INSERT INTO jack.teams (` + strings.Join(_c, ", ") + `)
	VALUES (` + strings.Join(_i, ", ") + `)
	RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"
	`
	jack.Log(sqlstr, _v...)

	cols := &columns{}
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&cols.ID, &cols.SlackTeamID, &cols.SlackTeamAccessToken, &cols.SlackBotAccessToken, &cols.SlackBotID, &cols.TeamName, &cols.Scope, &cols.Email, &cols.StripeID, &cols.Active, &cols.FreeTeammates, &cols.CostPerUser, &cols.TrialEnds, &cols.CreatedAt, &cols.UpdatedAt); e != nil {
		return nil, e
	}

	return &Team{cols}, nil
}

// Update a team by its `id`
func Update(db jack.DB, team *Team, id *string) (*Team, error) {
	fields := getColumns(team)

	// first check if we have the primary key
	if id == nil {
		return nil, errors.New(`primary key "id" must be non-nil`)
	}

	// don't update the primary key
	delete(fields, "id")

	// prepare the slices
	_c, _i, _v := jack.Slice(fields, 1)

	// sql query
	sqlstr := `UPDATE jack.teams SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `)
		WHERE "id" = $1
		RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"`

	// setup query
	values := append([]interface{}{id}, _v...)
	jack.Log(sqlstr, values...)

	// run the query
	var cols *columns
	row := db.QueryRow(sqlstr, values...)
	if e := row.Scan(&cols.ID, &cols.SlackTeamID, &cols.SlackTeamAccessToken, &cols.SlackBotAccessToken, &cols.SlackBotID, &cols.TeamName, &cols.Scope, &cols.Email, &cols.StripeID, &cols.Active, &cols.FreeTeammates, &cols.CostPerUser, &cols.TrialEnds, &cols.CreatedAt, &cols.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrTeamNotFound
		}
		return nil, e
	}

	return &Team{cols}, nil
}

// UpdateBySlackBotAccessToken find a Team
func UpdateBySlackBotAccessToken(db jack.DB, team *Team, slackBotAccessToken *string) (*Team, error) {
	fields := getColumns(team)

	// first check if we have all the keys we need
	if slackBotAccessToken == nil {
		return nil, errors.New(`slackBotAccessToken must be non-nil`)
	}

	// don't update the keys
	delete(fields, "slackBotAccessToken")

	// prepare the slices
	_c, _i, _v := jack.Slice(fields, 1)

	// sql query
	sqlstr := `UPDATE jack.teams SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `) ` +
		`WHERE "slack_bot_access_token" = $1 ` +
		`RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"`

	// setup the query
	values := []interface{}{}
	values = append(values, slackBotAccessToken)

	values = append(values, _v...)
	jack.Log(sqlstr, values...)

	// run the query
	var cols *columns
	row := db.QueryRow(sqlstr, values...)
	if e := row.Scan(&cols.ID, &cols.SlackTeamID, &cols.SlackTeamAccessToken, &cols.SlackBotAccessToken, &cols.SlackBotID, &cols.TeamName, &cols.Scope, &cols.Email, &cols.StripeID, &cols.Active, &cols.FreeTeammates, &cols.CostPerUser, &cols.TrialEnds, &cols.CreatedAt, &cols.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrTeamNotFound
		}
		return nil, e
	}

	return &Team{cols}, nil
}

// UpdateBySlackTeamAccessToken find a Team
func UpdateBySlackTeamAccessToken(db jack.DB, team *Team, slackTeamAccessToken *string) (*Team, error) {
	fields := getColumns(team)

	// first check if we have all the keys we need
	if slackTeamAccessToken == nil {
		return nil, errors.New(`slackTeamAccessToken must be non-nil`)
	}

	// don't update the keys
	delete(fields, "slackTeamAccessToken")

	// prepare the slices
	_c, _i, _v := jack.Slice(fields, 1)

	// sql query
	sqlstr := `UPDATE jack.teams SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `) ` +
		`WHERE "slack_team_access_token" = $1 ` +
		`RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"`

	// setup the query
	values := []interface{}{}
	values = append(values, slackTeamAccessToken)

	values = append(values, _v...)
	jack.Log(sqlstr, values...)

	// run the query
	var cols *columns
	row := db.QueryRow(sqlstr, values...)
	if e := row.Scan(&cols.ID, &cols.SlackTeamID, &cols.SlackTeamAccessToken, &cols.SlackBotAccessToken, &cols.SlackBotID, &cols.TeamName, &cols.Scope, &cols.Email, &cols.StripeID, &cols.Active, &cols.FreeTeammates, &cols.CostPerUser, &cols.TrialEnds, &cols.CreatedAt, &cols.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrTeamNotFound
		}
		return nil, e
	}

	return &Team{cols}, nil
}

// UpdateBySlackTeamID find a Team
func UpdateBySlackTeamID(db jack.DB, team *Team, slackTeamID *string) (*Team, error) {
	fields := getColumns(team)

	// first check if we have all the keys we need
	if slackTeamID == nil {
		return nil, errors.New(`slackTeamID must be non-nil`)
	}

	// don't update the keys
	delete(fields, "slackTeamID")

	// prepare the slices
	_c, _i, _v := jack.Slice(fields, 1)

	// sql query
	sqlstr := `UPDATE jack.teams SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `) ` +
		`WHERE "slack_team_id" = $1 ` +
		`RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"`

	// setup the query
	values := []interface{}{}
	values = append(values, slackTeamID)

	values = append(values, _v...)
	jack.Log(sqlstr, values...)

	// run the query
	var cols *columns
	row := db.QueryRow(sqlstr, values...)
	if e := row.Scan(&cols.ID, &cols.SlackTeamID, &cols.SlackTeamAccessToken, &cols.SlackBotAccessToken, &cols.SlackBotID, &cols.TeamName, &cols.Scope, &cols.Email, &cols.StripeID, &cols.Active, &cols.FreeTeammates, &cols.CostPerUser, &cols.TrialEnds, &cols.CreatedAt, &cols.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrTeamNotFound
		}
		return nil, e
	}

	return &Team{cols}, nil
}

// UpdateMany rows in `jack.teams` by a given condition
func UpdateMany(db jack.DB, team *Team, condition string, params ...interface{}) ([]*Team, error) {
	var _o []*Team

	// prepare the slices
	_c, _i, _v := jack.Slice(getColumns(team), len(params))

	// sql query
	sqlstr := `UPDATE jack.teams SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `) ` +
		`WHERE ` + condition + ` ` +
		`RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"`

		// setup the query
	values := []interface{}{}
	values = append(values, params...)
	values = append(values, _v...)
	jack.Log(sqlstr, values...)

	// run query
	rows, err := db.Query(sqlstr, values...)
	if err != nil {
		return _o, err
	}
	defer rows.Close()

	for rows.Next() {
		var cols *columns
		if e := rows.Scan(&cols.ID, &cols.SlackTeamID, &cols.SlackTeamAccessToken, &cols.SlackBotAccessToken, &cols.SlackBotID, &cols.TeamName, &cols.Scope, &cols.Email, &cols.StripeID, &cols.Active, &cols.FreeTeammates, &cols.CostPerUser, &cols.TrialEnds, &cols.CreatedAt, &cols.UpdatedAt); e != nil {
			if e == pgx.ErrNoRows {
				return _o, ErrTeamNotFound
			}
			return _o, err
		}
		_o = append(_o, &Team{cols})
	}
	if rows.Err() != nil {
		return _o, rows.Err()
	}

	// ensure we return an empty array
	// rather than nil when we marshal
	if len(_o) == 0 {
		return make([]*Team, 0), nil
	}

	return _o, nil
}

// Delete a `team` from the `jack.teams` table
func Delete(db jack.DB, id *string) error {
	// sql query
	sqlstr := `DELETE FROM jack.teams WHERE "id" = $1`
	jack.Log(sqlstr, id)

	// run query
	if _, e := db.Exec(sqlstr, id); e != nil {
		if e == pgx.ErrNoRows {
			return ErrTeamNotFound
		}
		return e
	}

	return nil
}

// DeleteBySlackBotAccessToken find a Team
func DeleteBySlackBotAccessToken(db jack.DB, slackBotAccessToken *string) error {
	// sql delete query
	sqlstr := `DELETE FROM jack.teams WHERE "slack_bot_access_token" = $1`
	jack.Log(sqlstr, slackBotAccessToken)

	if _, e := db.Exec(sqlstr, slackBotAccessToken); e != nil {
		if e == pgx.ErrNoRows {
			return ErrTeamNotFound
		}
		return e
	}

	return nil
}

// DeleteBySlackTeamAccessToken find a Team
func DeleteBySlackTeamAccessToken(db jack.DB, slackTeamAccessToken *string) error {
	// sql delete query
	sqlstr := `DELETE FROM jack.teams WHERE "slack_team_access_token" = $1`
	jack.Log(sqlstr, slackTeamAccessToken)

	if _, e := db.Exec(sqlstr, slackTeamAccessToken); e != nil {
		if e == pgx.ErrNoRows {
			return ErrTeamNotFound
		}
		return e
	}

	return nil
}

// DeleteBySlackTeamID find a Team
func DeleteBySlackTeamID(db jack.DB, slackTeamID *string) error {
	// sql delete query
	sqlstr := `DELETE FROM jack.teams WHERE "slack_team_id" = $1`
	jack.Log(sqlstr, slackTeamID)

	if _, e := db.Exec(sqlstr, slackTeamID); e != nil {
		if e == pgx.ErrNoRows {
			return ErrTeamNotFound
		}
		return e
	}

	return nil
}

// DeleteMany delete many `team`'s by the given condition
func DeleteMany(db jack.DB, condition string, params ...interface{}) error {
	// sql select query, primary key provided by sequence
	sqlstr := `DELETE FROM jack.teams WHERE ` + condition
	jack.Log(sqlstr, params...)

	if _, e := db.Exec(sqlstr, params...); e != nil {
		return e
	}

	return nil
}

// Upsert the `team` by its `id`.
func Upsert(db jack.DB, team *Team, action string) (*Team, error) {
	// prepare the slices
	_c, _i, _v := jack.Slice(getColumns(team), 0)

	// determine on conflict action
	var upsertAction string
	if action == jack.UpsertDoUpdate {
		upsertAction = `DO UPDATE SET (` + strings.Join(_c, ", ") + `) = ( EXCLUDED.` + strings.Join(_c, ", EXCLUDED.") + `)`
	} else if action == jack.UpsertDoNothing {
		upsertAction = jack.UpsertDoNothing
	} else {
		return nil, errors.New("invalid upsert action")
	}

	// sql query
	sqlstr := `INSERT INTO jack.teams (` + strings.Join(_c, ", ") + `) ` +
		`VALUES (` + strings.Join(_i, ", ") + `) ` +
		`ON CONFLICT ("id") ` +
		upsertAction + ` ` +
		`RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"`
	jack.Log(sqlstr, _v...)

	// run query
	var cols *columns
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&cols.ID, &cols.SlackTeamID, &cols.SlackTeamAccessToken, &cols.SlackBotAccessToken, &cols.SlackBotID, &cols.TeamName, &cols.Scope, &cols.Email, &cols.StripeID, &cols.Active, &cols.FreeTeammates, &cols.CostPerUser, &cols.TrialEnds, &cols.CreatedAt, &cols.UpdatedAt); e != nil && e != pgx.ErrNoRows {
		return nil, e
	}

	return &Team{cols}, nil
}

// UpsertBySlackBotAccessToken find a Team
func UpsertBySlackBotAccessToken(db jack.DB, team *Team, action string) (*Team, error) {
	// get all the non-nil columns and prepare them for the query
	_c, _i, _v := jack.Slice(getColumns(team), 0)

	// determine on conflict action
	var upsertAction string
	if action == jack.UpsertDoUpdate {
		upsertAction = `DO UPDATE SET (` + strings.Join(_c, ", ") + `) = ( EXCLUDED.` + strings.Join(_c, ", EXCLUDED.") + `)`
	} else if action == jack.UpsertDoNothing {
		upsertAction = jack.UpsertDoNothing
	} else {
		return nil, errors.New("invalid upsert action")
	}

	// sql query
	sqlstr := `INSERT INTO jack.teams (` + strings.Join(_c, ", ") + `) ` +
		`VALUES (` + strings.Join(_i, ", ") + `) ` +
		`ON CONFLICT ("slack_bot_access_token") ` +
		upsertAction + ` ` +
		`RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"`
	jack.Log(sqlstr, _v...)

	// run query
	var cols *columns
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&cols.ID, &cols.SlackTeamID, &cols.SlackTeamAccessToken, &cols.SlackBotAccessToken, &cols.SlackBotID, &cols.TeamName, &cols.Scope, &cols.Email, &cols.StripeID, &cols.Active, &cols.FreeTeammates, &cols.CostPerUser, &cols.TrialEnds, &cols.CreatedAt, &cols.UpdatedAt); e != nil && e != pgx.ErrNoRows {
		return nil, e
	}

	return &Team{cols}, nil
}

// UpsertBySlackTeamAccessToken find a Team
func UpsertBySlackTeamAccessToken(db jack.DB, team *Team, action string) (*Team, error) {
	// get all the non-nil columns and prepare them for the query
	_c, _i, _v := jack.Slice(getColumns(team), 0)

	// determine on conflict action
	var upsertAction string
	if action == jack.UpsertDoUpdate {
		upsertAction = `DO UPDATE SET (` + strings.Join(_c, ", ") + `) = ( EXCLUDED.` + strings.Join(_c, ", EXCLUDED.") + `)`
	} else if action == jack.UpsertDoNothing {
		upsertAction = jack.UpsertDoNothing
	} else {
		return nil, errors.New("invalid upsert action")
	}

	// sql query
	sqlstr := `INSERT INTO jack.teams (` + strings.Join(_c, ", ") + `) ` +
		`VALUES (` + strings.Join(_i, ", ") + `) ` +
		`ON CONFLICT ("slack_team_access_token") ` +
		upsertAction + ` ` +
		`RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"`
	jack.Log(sqlstr, _v...)

	// run query
	var cols *columns
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&cols.ID, &cols.SlackTeamID, &cols.SlackTeamAccessToken, &cols.SlackBotAccessToken, &cols.SlackBotID, &cols.TeamName, &cols.Scope, &cols.Email, &cols.StripeID, &cols.Active, &cols.FreeTeammates, &cols.CostPerUser, &cols.TrialEnds, &cols.CreatedAt, &cols.UpdatedAt); e != nil && e != pgx.ErrNoRows {
		return nil, e
	}

	return &Team{cols}, nil
}

// UpsertBySlackTeamID find a Team
func UpsertBySlackTeamID(db jack.DB, team *Team, action string) (*Team, error) {
	// get all the non-nil columns and prepare them for the query
	_c, _i, _v := jack.Slice(getColumns(team), 0)

	// determine on conflict action
	var upsertAction string
	if action == jack.UpsertDoUpdate {
		upsertAction = `DO UPDATE SET (` + strings.Join(_c, ", ") + `) = ( EXCLUDED.` + strings.Join(_c, ", EXCLUDED.") + `)`
	} else if action == jack.UpsertDoNothing {
		upsertAction = jack.UpsertDoNothing
	} else {
		return nil, errors.New("invalid upsert action")
	}

	// sql query
	sqlstr := `INSERT INTO jack.teams (` + strings.Join(_c, ", ") + `) ` +
		`VALUES (` + strings.Join(_i, ", ") + `) ` +
		`ON CONFLICT ("slack_team_id") ` +
		upsertAction + ` ` +
		`RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"`
	jack.Log(sqlstr, _v...)

	// run query
	var cols *columns
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&cols.ID, &cols.SlackTeamID, &cols.SlackTeamAccessToken, &cols.SlackBotAccessToken, &cols.SlackBotID, &cols.TeamName, &cols.Scope, &cols.Email, &cols.StripeID, &cols.Active, &cols.FreeTeammates, &cols.CostPerUser, &cols.TrialEnds, &cols.CreatedAt, &cols.UpdatedAt); e != nil && e != pgx.ErrNoRows {
		return nil, e
	}

	return &Team{cols}, nil
}