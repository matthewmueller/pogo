package jack2

import (
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx"
	uuid "github.com/satori/go.uuid"
)

// GENERATED BY POGO. DO NOT EDIT.

// ErrTeamNotFound returned if the team is not found
var ErrTeamNotFound = errors.New("team not found")

// Teams class
type Teams struct {
	db DB
}

// Team model
type Team struct {
	ID                   *uuid.UUID `json:"id,omitempty"`
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

// team constructor
func team(db DB) *Teams {
	return &Teams{db}
}

// get all the non-nil fields
func (teams *Teams) fields(team *Team) map[string]interface{} {
	fields := make(map[string]interface{})

	if team.ID != nil {
		fields["id"] = team.ID
	}
	if team.SlackTeamID != nil {
		fields["slack_team_id"] = team.SlackTeamID
	}
	if team.SlackTeamAccessToken != nil {
		fields["slack_team_access_token"] = team.SlackTeamAccessToken
	}
	if team.SlackBotAccessToken != nil {
		fields["slack_bot_access_token"] = team.SlackBotAccessToken
	}
	if team.SlackBotID != nil {
		fields["slack_bot_id"] = team.SlackBotID
	}
	if team.TeamName != nil {
		fields["team_name"] = team.TeamName
	}
	if team.Scope != nil {
		fields["scope"] = team.Scope
	}
	if team.Email != nil {
		fields["email"] = team.Email
	}
	if team.StripeID != nil {
		fields["stripe_id"] = team.StripeID
	}
	if team.Active != nil {
		fields["active"] = team.Active
	}
	if team.FreeTeammates != nil {
		fields["free_teammates"] = team.FreeTeammates
	}
	if team.CostPerUser != nil {
		fields["cost_per_user"] = team.CostPerUser
	}
	if team.TrialEnds != nil {
		fields["trial_ends"] = team.TrialEnds
	}
	if team.CreatedAt != nil {
		fields["created_at"] = team.CreatedAt
	}
	if team.UpdatedAt != nil {
		fields["updated_at"] = team.UpdatedAt
	}

	return fields
}

// Find a team by "id"
func (teams *Teams) Find(id *uuid.UUID) (team *Team, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"
	FROM jack.teams
	WHERE "id" = $1
	`

	Log(sqlstr, id)
	row := teams.db.QueryRow(sqlstr, id)
	if e := row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrTeamNotFound
		}
		return nil, e
	}

	return team, nil
}

// FindByID find a team by `id`
func (teams *Teams) FindByID(iD *uuid.UUID) (team *Team, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"
	FROM jack.teams
	WHERE "id" = $1
	`

	Log(sqlstr, iD)
	row := teams.db.QueryRow(sqlstr, iD)
	err = row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return team, ErrTeamNotFound
		}
		return team, err
	}

	return team, nil
}

// FindBySlackBotAccessToken find a team by `slack_bot_access_token`
func (teams *Teams) FindBySlackBotAccessToken(slackBotAccessToken *string) (team *Team, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"
	FROM jack.teams
	WHERE "slack_bot_access_token" = $1
	`

	Log(sqlstr, slackBotAccessToken)
	row := teams.db.QueryRow(sqlstr, slackBotAccessToken)
	err = row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return team, ErrTeamNotFound
		}
		return team, err
	}

	return team, nil
}

// FindBySlackTeamAccessToken find a team by `slack_team_access_token`
func (teams *Teams) FindBySlackTeamAccessToken(slackTeamAccessToken *string) (team *Team, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"
	FROM jack.teams
	WHERE "slack_team_access_token" = $1
	`

	Log(sqlstr, slackTeamAccessToken)
	row := teams.db.QueryRow(sqlstr, slackTeamAccessToken)
	err = row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return team, ErrTeamNotFound
		}
		return team, err
	}

	return team, nil
}

// FindBySlackTeamID find a team by `slack_team_id`
func (teams *Teams) FindBySlackTeamID(slackTeamID *string) (team *Team, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"
	FROM jack.teams
	WHERE "slack_team_id" = $1
	`

	Log(sqlstr, slackTeamID)
	row := teams.db.QueryRow(sqlstr, slackTeamID)
	err = row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return team, ErrTeamNotFound
		}
		return team, err
	}

	return team, nil
}

// FindMany find many `team`s by a given condition
func (teams *Teams) FindMany(condition string, params ...interface{}) ([]*Team, error) {
	var _o []*Team

	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"
	FROM jack.teams
	WHERE ` + condition

	Log(sqlstr, params...)
	rows, err := teams.db.Query(sqlstr, params...)
	if err != nil {
		return _o, err
	}
	defer rows.Close()

	for rows.Next() {
		var team *Team
		if e := rows.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt); e != nil {
			if e == pgx.ErrNoRows {
				return _o, ErrTeamNotFound
			}
			return _o, err
		}
		_o = append(_o, team)
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
func (teams *Teams) FindOne(condition string, params ...interface{}) (team *Team, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"
	FROM jack.teams
	WHERE ` + condition

	Log(sqlstr, params...)
	row := teams.db.QueryRow(sqlstr, params...)
	if e := row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrTeamNotFound
		}
		return nil, e
	}

	return team, nil
}

// Insert a `team` into the `jack.teams` table.
func (teams *Teams) Insert(team Team) (*Team, error) {
	// get all the non-nil fields and prepare them for the query
	_c, _i, _v := slice(teams.fields(&team), 0)

	// sql insert query, primary key provided by sequence
	sqlstr := `
	INSERT INTO jack.teams (` + strings.Join(_c, ", ") + `)
	VALUES (` + strings.Join(_i, ", ") + `)
	RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"
	`

	Log(sqlstr, _v...)
	row := teams.db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt); e != nil {
		return nil, e
	}

	return &team, nil
}

// Update a team by its `id`
func (teams *Teams) Update(team Team, id *uuid.UUID) (*Team, error) {
	fieldset := teams.fields(&team)

	// first check if we have the primary key
	if id == nil {
		return nil, errors.New(`primary key "id" must be non-nil`)
	}

	// don't update the primary key
	delete(fieldset, "id")

	// prepare the slices
	_c, _i, _v := slice(fieldset, 1)

	// sql query
	sqlstr := `UPDATE jack.teams SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `)
		WHERE "id" = $1
		RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"`

	// run query
	values := append([]interface{}{id}, _v...)
	Log(sqlstr, values...)

	row := teams.db.QueryRow(sqlstr, values...)
	if e := row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrTeamNotFound
		}
		return nil, e
	}

	return &team, nil
}

// UpdateByID find a Team
func (teams *Teams) UpdateByID(team Team, iD *uuid.UUID) (*Team, error) {
	fieldset := teams.fields(&team)

	// first check if we have all the keys we need
	if iD == nil {
		return nil, errors.New(`iD must be non-nil`)
	}

	// don't update the keys
	delete(fieldset, "iD")

	// prepare the slices
	_c, _i, _v := slice(fieldset, 1)

	// sql query
	sqlstr := `UPDATE jack.teams SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `) ` +
		`WHERE "id" = $1 ` +
		`RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"`

	// run query
	values := []interface{}{}
	values = append(values, iD)

	values = append(values, _v...)
	Log(sqlstr, values...)

	row := teams.db.QueryRow(sqlstr, values...)
	if e := row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrTeamNotFound
		}
		return nil, e
	}

	return &team, nil
}

// UpdateBySlackBotAccessToken find a Team
func (teams *Teams) UpdateBySlackBotAccessToken(team Team, slackBotAccessToken *string) (*Team, error) {
	fieldset := teams.fields(&team)

	// first check if we have all the keys we need
	if slackBotAccessToken == nil {
		return nil, errors.New(`slackBotAccessToken must be non-nil`)
	}

	// don't update the keys
	delete(fieldset, "slackBotAccessToken")

	// prepare the slices
	_c, _i, _v := slice(fieldset, 1)

	// sql query
	sqlstr := `UPDATE jack.teams SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `) ` +
		`WHERE "slack_bot_access_token" = $1 ` +
		`RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"`

	// run query
	values := []interface{}{}
	values = append(values, slackBotAccessToken)

	values = append(values, _v...)
	Log(sqlstr, values...)

	row := teams.db.QueryRow(sqlstr, values...)
	if e := row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrTeamNotFound
		}
		return nil, e
	}

	return &team, nil
}

// UpdateBySlackTeamAccessToken find a Team
func (teams *Teams) UpdateBySlackTeamAccessToken(team Team, slackTeamAccessToken *string) (*Team, error) {
	fieldset := teams.fields(&team)

	// first check if we have all the keys we need
	if slackTeamAccessToken == nil {
		return nil, errors.New(`slackTeamAccessToken must be non-nil`)
	}

	// don't update the keys
	delete(fieldset, "slackTeamAccessToken")

	// prepare the slices
	_c, _i, _v := slice(fieldset, 1)

	// sql query
	sqlstr := `UPDATE jack.teams SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `) ` +
		`WHERE "slack_team_access_token" = $1 ` +
		`RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"`

	// run query
	values := []interface{}{}
	values = append(values, slackTeamAccessToken)

	values = append(values, _v...)
	Log(sqlstr, values...)

	row := teams.db.QueryRow(sqlstr, values...)
	if e := row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrTeamNotFound
		}
		return nil, e
	}

	return &team, nil
}

// UpdateBySlackTeamID find a Team
func (teams *Teams) UpdateBySlackTeamID(team Team, slackTeamID *string) (*Team, error) {
	fieldset := teams.fields(&team)

	// first check if we have all the keys we need
	if slackTeamID == nil {
		return nil, errors.New(`slackTeamID must be non-nil`)
	}

	// don't update the keys
	delete(fieldset, "slackTeamID")

	// prepare the slices
	_c, _i, _v := slice(fieldset, 1)

	// sql query
	sqlstr := `UPDATE jack.teams SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `) ` +
		`WHERE "slack_team_id" = $1 ` +
		`RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"`

	// run query
	values := []interface{}{}
	values = append(values, slackTeamID)

	values = append(values, _v...)
	Log(sqlstr, values...)

	row := teams.db.QueryRow(sqlstr, values...)
	if e := row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrTeamNotFound
		}
		return nil, e
	}

	return &team, nil
}

// UpdateMany rows in `jack.teams` by a given condition
func (teams *Teams) UpdateMany(team *Team, condition string, params ...interface{}) ([]*Team, error) {
	var _o []*Team

	// get the non-nil fields
	fieldset := teams.fields(team)

	// prepare the slices
	_c, _i, _v := slice(fieldset, len(params))

	// sql query
	sqlstr := `UPDATE jack.teams SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `) ` +
		`WHERE ` + condition + ` ` +
		`RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"`

	values := []interface{}{}
	values = append(values, params...)
	values = append(values, _v...)

	// run query
	Log(sqlstr, values...)
	rows, err := teams.db.Query(sqlstr, values...)
	if err != nil {
		return _o, err
	}
	defer rows.Close()

	for rows.Next() {
		var team *Team
		if e := rows.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt); e != nil {
			if e == pgx.ErrNoRows {
				return _o, ErrTeamNotFound
			}
			return _o, err
		}
		_o = append(_o, team)
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
func (teams *Teams) Delete(id *uuid.UUID) error {
	// sql query
	sqlstr := `DELETE FROM jack.teams WHERE "id" = $1`

	// run query
	Log(sqlstr, id)
	if _, e := teams.db.Exec(sqlstr, id); e != nil {
		if e == pgx.ErrNoRows {
			return ErrTeamNotFound
		}
		return e
	}

	return nil
}

// DeleteByID find a Team
func (teams *Teams) DeleteByID(iD *uuid.UUID) error {
	// sql delete query
	sqlstr := `DELETE FROM jack.teams WHERE "id" = $1`

	Log(sqlstr, iD)
	if _, e := teams.db.Exec(sqlstr, iD); e != nil {
		if e == pgx.ErrNoRows {
			return ErrTeamNotFound
		}
		return e
	}

	return nil
}

// DeleteBySlackBotAccessToken find a Team
func (teams *Teams) DeleteBySlackBotAccessToken(slackBotAccessToken *string) error {
	// sql delete query
	sqlstr := `DELETE FROM jack.teams WHERE "slack_bot_access_token" = $1`

	Log(sqlstr, slackBotAccessToken)
	if _, e := teams.db.Exec(sqlstr, slackBotAccessToken); e != nil {
		if e == pgx.ErrNoRows {
			return ErrTeamNotFound
		}
		return e
	}

	return nil
}

// DeleteBySlackTeamAccessToken find a Team
func (teams *Teams) DeleteBySlackTeamAccessToken(slackTeamAccessToken *string) error {
	// sql delete query
	sqlstr := `DELETE FROM jack.teams WHERE "slack_team_access_token" = $1`

	Log(sqlstr, slackTeamAccessToken)
	if _, e := teams.db.Exec(sqlstr, slackTeamAccessToken); e != nil {
		if e == pgx.ErrNoRows {
			return ErrTeamNotFound
		}
		return e
	}

	return nil
}

// DeleteBySlackTeamID find a Team
func (teams *Teams) DeleteBySlackTeamID(slackTeamID *string) error {
	// sql delete query
	sqlstr := `DELETE FROM jack.teams WHERE "slack_team_id" = $1`

	Log(sqlstr, slackTeamID)
	if _, e := teams.db.Exec(sqlstr, slackTeamID); e != nil {
		if e == pgx.ErrNoRows {
			return ErrTeamNotFound
		}
		return e
	}

	return nil
}

// DeleteMany delete many `team`'s by the given condition
func (teams *Teams) DeleteMany(condition string, params ...interface{}) error {
	// sql select query, primary key provided by sequence
	sqlstr := `DELETE FROM jack.teams WHERE ` + condition

	Log(sqlstr, params...)
	if _, e := teams.db.Exec(sqlstr, params...); e != nil {
		return e
	}

	return nil
}

// Upsert the `team` by its `id`.
func (teams *Teams) Upsert(team Team, action string) (*Team, error) {
	fieldset := teams.fields(&team)

	// prepare the slices
	_c, _i, _v := slice(fieldset, 0)

	// determine on conflict action
	var upsertAction string
	if action == UpsertDoUpdate {
		upsertAction = `DO UPDATE SET (` + strings.Join(_c, ", ") + `) = ( EXCLUDED.` + strings.Join(_c, ", EXCLUDED.") + `)`
	} else if action == UpsertDoNothing {
		upsertAction = UpsertDoNothing
	} else {
		return nil, errors.New("invalid upsert action")
	}

	// sql query
	sqlstr := `INSERT INTO jack.teams (` + strings.Join(_c, ", ") + `) ` +
		`VALUES (` + strings.Join(_i, ", ") + `) ` +
		`ON CONFLICT ("id") ` +
		upsertAction + ` ` +
		`RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"`

		// run query
	Log(sqlstr, _v...)
	row := teams.db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt); e != nil && e != pgx.ErrNoRows {
		return nil, e
	}

	return &team, nil
}

// UpsertByID find a Team
func (teams *Teams) UpsertByID(team Team, action string) (*Team, error) {
	// get all the non-nil fields and prepare them for the query
	_c, _i, _v := slice(teams.fields(&team), 0)

	// determine on conflict action
	var upsertAction string
	if action == UpsertDoUpdate {
		upsertAction = `DO UPDATE SET (` + strings.Join(_c, ", ") + `) = ( EXCLUDED.` + strings.Join(_c, ", EXCLUDED.") + `)`
	} else if action == UpsertDoNothing {
		upsertAction = UpsertDoNothing
	} else {
		return nil, errors.New("invalid upsert action")
	}

	// sql query
	sqlstr := `INSERT INTO jack.teams (` + strings.Join(_c, ", ") + `) ` +
		`VALUES (` + strings.Join(_i, ", ") + `) ` +
		`ON CONFLICT ("id") ` +
		upsertAction + ` ` +
		`RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"`

		// run query
	Log(sqlstr, _v...)
	row := teams.db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt); e != nil && e != pgx.ErrNoRows {
		return nil, e
	}

	return &team, nil
}

// UpsertBySlackBotAccessToken find a Team
func (teams *Teams) UpsertBySlackBotAccessToken(team Team, action string) (*Team, error) {
	// get all the non-nil fields and prepare them for the query
	_c, _i, _v := slice(teams.fields(&team), 0)

	// determine on conflict action
	var upsertAction string
	if action == UpsertDoUpdate {
		upsertAction = `DO UPDATE SET (` + strings.Join(_c, ", ") + `) = ( EXCLUDED.` + strings.Join(_c, ", EXCLUDED.") + `)`
	} else if action == UpsertDoNothing {
		upsertAction = UpsertDoNothing
	} else {
		return nil, errors.New("invalid upsert action")
	}

	// sql query
	sqlstr := `INSERT INTO jack.teams (` + strings.Join(_c, ", ") + `) ` +
		`VALUES (` + strings.Join(_i, ", ") + `) ` +
		`ON CONFLICT ("slack_bot_access_token") ` +
		upsertAction + ` ` +
		`RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"`

		// run query
	Log(sqlstr, _v...)
	row := teams.db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt); e != nil && e != pgx.ErrNoRows {
		return nil, e
	}

	return &team, nil
}

// UpsertBySlackTeamAccessToken find a Team
func (teams *Teams) UpsertBySlackTeamAccessToken(team Team, action string) (*Team, error) {
	// get all the non-nil fields and prepare them for the query
	_c, _i, _v := slice(teams.fields(&team), 0)

	// determine on conflict action
	var upsertAction string
	if action == UpsertDoUpdate {
		upsertAction = `DO UPDATE SET (` + strings.Join(_c, ", ") + `) = ( EXCLUDED.` + strings.Join(_c, ", EXCLUDED.") + `)`
	} else if action == UpsertDoNothing {
		upsertAction = UpsertDoNothing
	} else {
		return nil, errors.New("invalid upsert action")
	}

	// sql query
	sqlstr := `INSERT INTO jack.teams (` + strings.Join(_c, ", ") + `) ` +
		`VALUES (` + strings.Join(_i, ", ") + `) ` +
		`ON CONFLICT ("slack_team_access_token") ` +
		upsertAction + ` ` +
		`RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"`

		// run query
	Log(sqlstr, _v...)
	row := teams.db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt); e != nil && e != pgx.ErrNoRows {
		return nil, e
	}

	return &team, nil
}

// UpsertBySlackTeamID find a Team
func (teams *Teams) UpsertBySlackTeamID(team Team, action string) (*Team, error) {
	// get all the non-nil fields and prepare them for the query
	_c, _i, _v := slice(teams.fields(&team), 0)

	// determine on conflict action
	var upsertAction string
	if action == UpsertDoUpdate {
		upsertAction = `DO UPDATE SET (` + strings.Join(_c, ", ") + `) = ( EXCLUDED.` + strings.Join(_c, ", EXCLUDED.") + `)`
	} else if action == UpsertDoNothing {
		upsertAction = UpsertDoNothing
	} else {
		return nil, errors.New("invalid upsert action")
	}

	// sql query
	sqlstr := `INSERT INTO jack.teams (` + strings.Join(_c, ", ") + `) ` +
		`VALUES (` + strings.Join(_i, ", ") + `) ` +
		`ON CONFLICT ("slack_team_id") ` +
		upsertAction + ` ` +
		`RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"`

		// run query
	Log(sqlstr, _v...)
	row := teams.db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt); e != nil && e != pgx.ErrNoRows {
		return nil, e
	}

	return &team, nil
}
