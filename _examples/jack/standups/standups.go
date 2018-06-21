package standups

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx"
	"github.com/matthewmueller/jack/pogo"
)

// GENERATED BY POGO. DO NOT EDIT.

// ErrStandupNotFound returned if the standups is not found
var ErrStandupNotFound = errors.New("standups not found")

// StandupInput model for "jack"."standups"
type StandupInput struct {
	id             *string
	name           *string
	slackChannelID *string
	time           *string
	timezone       *string
	questions      *json.RawMessage
	teamID         *string
	createdAt      *time.Time
	updatedAt      *time.Time
}

// Standup model for "jack"."standups"
type Standup struct {
	ID             string
	Name           string
	SlackChannelID string
	Time           string
	Timezone       string
	Questions      json.RawMessage
	TeamID         string
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

// New "jack"."standups" API
func New() *StandupInput {
	return &StandupInput{}
}

// ID sets the "id"
func (standup *StandupInput) ID(id string) *StandupInput {
	standup.id = &id
	return standup
}

// Name sets the "name"
func (standup *StandupInput) Name(name string) *StandupInput {
	standup.name = &name
	return standup
}

// SlackChannelID sets the "slackChannelID"
func (standup *StandupInput) SlackChannelID(slackChannelID string) *StandupInput {
	standup.slackChannelID = &slackChannelID
	return standup
}

// Time sets the "time"
func (standup *StandupInput) Time(time string) *StandupInput {
	standup.time = &time
	return standup
}

// Timezone sets the "timezone"
func (standup *StandupInput) Timezone(timezone string) *StandupInput {
	standup.timezone = &timezone
	return standup
}

// Questions sets the "questions"
func (standup *StandupInput) Questions(questions json.RawMessage) *StandupInput {
	standup.questions = &questions
	return standup
}

// TeamID sets the "teamID"
func (standup *StandupInput) TeamID(teamID string) *StandupInput {
	standup.teamID = &teamID
	return standup
}

// CreatedAt sets the "createdAt"
func (standup *StandupInput) CreatedAt(createdAt time.Time) *StandupInput {
	standup.createdAt = &createdAt
	return standup
}

// UpdatedAt sets the "updatedAt"
func (standup *StandupInput) UpdatedAt(updatedAt time.Time) *StandupInput {
	standup.updatedAt = &updatedAt
	return standup
}

// MarshalJSON marshals the "standup" into JSON
func (standup *StandupInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(standup)
}

// UnmarshalJSON unmarshals json to a "standup"
func (standup *StandupInput) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, standup)
}

func (standup *StandupInput) String() string {
	return "standup"
}

func getColumns(standup *StandupInput) map[string]interface{} {
	columns := make(map[string]interface{})

	if standup.id != nil {
		columns["id"] = *standup.id
	}

	if standup.name != nil {
		columns["name"] = *standup.name
	}

	if standup.slackChannelID != nil {
		columns["slack_channel_id"] = *standup.slackChannelID
	}

	if standup.time != nil {
		columns["time"] = *standup.time
	}

	if standup.timezone != nil {
		columns["timezone"] = *standup.timezone
	}

	if standup.questions != nil {
		columns["questions"] = *standup.questions
	}

	if standup.teamID != nil {
		columns["team_id"] = *standup.teamID
	}

	if standup.createdAt != nil {
		columns["created_at"] = *standup.createdAt
	}

	if standup.updatedAt != nil {
		columns["updated_at"] = *standup.updatedAt
	}

	return columns
}

// WhereClause is a struct to handle where clauses
type WhereClause struct {
	condition string
	params    []interface{}
}

// Where specifies the conditions
func Where(condition string, params ...interface{}) *WhereClause {
	return &WhereClause{
		condition: condition,
		params:    params,
	}
}

// Insert a "standup" into the "jack"."standups"
func Insert(db jack.DB, standupInput *StandupInput) (*Standup, error) {
	// get all the non-nil columns and prepare them for the query
	_c, _i, _v := jack.Slice(getColumns(standupInput), 0)

	// sql insert query, primary key provided by sequence
	sqlstr := `
    INSERT INTO "jack"."standups" (` + strings.Join(_c, ", ") + `)
    VALUES (` + strings.Join(_i, ", ") + `)
    RETURNING "id", "name", "slack_channel_id", "time", "timezone", "questions", "team_id", "created_at", "updated_at"
  `
	jack.Log(sqlstr, _v...)

	var standup Standup
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&standup.ID, &standup.Name, &standup.SlackChannelID, &standup.Time, &standup.Timezone, &standup.Questions, &standup.TeamID, &standup.CreatedAt, &standup.UpdatedAt); e != nil {
		return nil, e
	}

	return &standup, nil
}

// Find a "Standup" by "id"
func Find(db pogo.DB, id string) (*Standup, error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT "id", "name", "slack_channel_id", "time", "timezone", "questions", "team_id", "created_at", "updated_at"
    FROM "jack"."standups"
    WHERE "id" = $1
  `
	pogo.Log(sqlstr, &id)

	var standup Standup
	row := db.QueryRow(sqlstr, &id)
	if e := row.Scan(&standup.ID, &standup.Name, &standup.SlackChannelID, &standup.Time, &standup.Timezone, &standup.Questions, &standup.TeamID, &standup.CreatedAt, &standup.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrStandupNotFound
		}
		return nil, e
	}

	return &standup, nil
}

// FindByID find a standup by id
func FindByID(db jack.DB, id string) (*Standup, error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT "id", "name", "slack_channel_id", "time", "timezone", "questions", "team_id", "created_at", "updated_at"
    FROM "jack"."standups"
    WHERE "id" = $1
  `
	jack.Log(sqlstr, id)

	var standup Standup
	row := db.QueryRow(sqlstr, id)
	err := row.Scan(&standup.ID, &standup.Name, &standup.SlackChannelID, &standup.Time, &standup.Timezone, &standup.Questions, &standup.TeamID, &standup.CreatedAt, &standup.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrStandupNotFound
		}
		return nil, err
	}

	return &standup, nil
}

// FindBySlackChannelID find a standup by slack_channel_id
func FindBySlackChannelID(db jack.DB, slackChannelID string) (*Standup, error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT "id", "name", "slack_channel_id", "time", "timezone", "questions", "team_id", "created_at", "updated_at"
    FROM "jack"."standups"
    WHERE "slack_channel_id" = $1
  `
	jack.Log(sqlstr, slackChannelID)

	var standup Standup
	row := db.QueryRow(sqlstr, slackChannelID)
	err := row.Scan(&standup.ID, &standup.Name, &standup.SlackChannelID, &standup.Time, &standup.Timezone, &standup.Questions, &standup.TeamID, &standup.CreatedAt, &standup.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrStandupNotFound
		}
		return nil, err
	}

	return &standup, nil
}

// FindOne find one standup by a condition
func FindOne(db jack.DB, where *WhereClause) (*Standup, error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
  SELECT "id", "name", "slack_channel_id", "time", "timezone", "questions", "team_id", "created_at", "updated_at"
  FROM "jack"."standups"
  WHERE ` + where.condition
	jack.Log(sqlstr, where.params...)

	var standup Standup
	row := db.QueryRow(sqlstr, where.params...)
	if e := row.Scan(&standup.ID, &standup.Name, &standup.SlackChannelID, &standup.Time, &standup.Timezone, &standup.Questions, &standup.TeamID, &standup.CreatedAt, &standup.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrStandupNotFound
		}
		return nil, e
	}

	return &standup, nil
}

// FindMany find many "standup"'s by a given condition
func FindMany(db jack.DB, where *WhereClause) ([]*Standup, error) {
	standups := []*Standup{}

	// sql select query, primary key provided by sequence
	sqlstr := `
  SELECT "id", "name", "slack_channel_id", "time", "timezone", "questions", "team_id", "created_at", "updated_at"
  FROM "jack"."standups"
  WHERE ` + where.condition
	jack.Log(sqlstr, where.params...)

	rows, err := db.Query(sqlstr, where.params...)
	if err != nil {
		return standups, err
	}
	defer rows.Close()

	for rows.Next() {
		var standup Standup
		if e := rows.Scan(&standup.ID, &standup.Name, &standup.SlackChannelID, &standup.Time, &standup.Timezone, &standup.Questions, &standup.TeamID, &standup.CreatedAt, &standup.UpdatedAt); e != nil {
			if e == pgx.ErrNoRows {
				return standups, ErrStandupNotFound
			}
			return standups, err
		}
		standups = append(standups, &standup)
	}
	if rows.Err() != nil {
		return standups, rows.Err()
	}

	return standups, nil
}

// Update a "standup" in "jack"."standups" by its "id"
func Update(db jack.DB, id string, standupInput *StandupInput) (*Standup, error) {
	fields := getColumns(standupInput)

	// don't update the primary key
	delete(fields, "id")

	// prepare the slices
	_c, _i, _v := jack.Slice(fields, 1)

	// sql query
	sqlstr := `UPDATE "jack"."standups" SET (` +
		strings.Join(_c, ", ") +
		`) = (` +
		strings.Join(_i, ", ") +
		`)
    WHERE "id" = $1
    RETURNING "id", "name", "slack_channel_id", "time", "timezone", "questions", "team_id", "created_at", "updated_at"`

	// setup query
	values := append([]interface{}{&id}, _v...)
	jack.Log(sqlstr, values...)

	// run the query
	var standup Standup
	row := db.QueryRow(sqlstr, values...)
	if e := row.Scan(&standup.ID, &standup.Name, &standup.SlackChannelID, &standup.Time, &standup.Timezone, &standup.Questions, &standup.TeamID, &standup.CreatedAt, &standup.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrStandupNotFound
		}
		return nil, e
	}

	return &standup, nil
}

// UpdateBySlackChannelID find a Standup
func UpdateBySlackChannelID(db jack.DB, slackChannelID string, standupInput *StandupInput) (*Standup, error) {
	fields := getColumns(standupInput)

	// don't update the keys
	delete(fields, "slack_channel_id")

	// prepare the slices
	_c, _i, _v := jack.Slice(fields, 1)

	// sql query
	sqlstr := `UPDATE "jack"."standups" SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `) ` +
		`WHERE "slack_channel_id" = $1 ` +
		`RETURNING "id", "name", "slack_channel_id", "time", "timezone", "questions", "team_id", "created_at", "updated_at"`

	// setup the query
	values := []interface{}{}
	values = append(values, slackChannelID)
	values = append(values, _v...)
	jack.Log(sqlstr, values...)

	// run the query
	var standup Standup
	row := db.QueryRow(sqlstr, values...)
	if e := row.Scan(&standup.ID, &standup.Name, &standup.SlackChannelID, &standup.Time, &standup.Timezone, &standup.Questions, &standup.TeamID, &standup.CreatedAt, &standup.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrStandupNotFound
		}
		return nil, e
	}

	return &standup, nil
}

// UpdateMany rows in "jack"."standups" by a given condition
func UpdateMany(db jack.DB, where *WhereClause, standupInput *StandupInput) ([]*Standup, error) {
	standups := []*Standup{}

	// prepare the slices
	_c, _i, _v := jack.Slice(getColumns(standupInput), len(where.params))

	// sql query
	sqlstr := `UPDATE "jack"."standups" SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `) ` +
		`WHERE ` + where.condition + ` ` +
		`RETURNING "id", "name", "slack_channel_id", "time", "timezone", "questions", "team_id", "created_at", "updated_at"`

	// setup the query
	values := []interface{}{}
	values = append(values, where.params...)
	values = append(values, _v...)
	jack.Log(sqlstr, values...)

	// run query
	rows, err := db.Query(sqlstr, values...)
	if err != nil {
		return standups, err
	}
	defer rows.Close()

	for rows.Next() {
		var standup Standup
		if e := rows.Scan(&standup.ID, &standup.Name, &standup.SlackChannelID, &standup.Time, &standup.Timezone, &standup.Questions, &standup.TeamID, &standup.CreatedAt, &standup.UpdatedAt); e != nil {
			if e == pgx.ErrNoRows {
				return standups, ErrStandupNotFound
			}
			return standups, err
		}
		standups = append(standups, &standup)
	}
	if rows.Err() != nil {
		return standups, rows.Err()
	}

	return standups, nil
}

// Delete a "standup" from the "jack"."standups" table
func Delete(db jack.DB, id string) error {
	// sql query
	sqlstr := `DELETE FROM "jack"."standups" WHERE "id" = $1`
	jack.Log(sqlstr, id)

	// run query
	if _, e := db.Exec(sqlstr, id); e != nil {
		if e == pgx.ErrNoRows {
			return ErrStandupNotFound
		}
		return e
	}

	return nil
}

// DeleteBySlackChannelID deletes a "standup"
func DeleteBySlackChannelID(db jack.DB, slackChannelID string) error {
	// sql delete query
	sqlstr := `DELETE FROM "jack"."standups" WHERE "slack_channel_id" = $1`
	jack.Log(sqlstr, slackChannelID)

	if _, e := db.Exec(sqlstr, slackChannelID); e != nil {
		if e == pgx.ErrNoRows {
			return ErrStandupNotFound
		}
		return e
	}

	return nil
}

// DeleteMany delete many "standup"'s by the given condition
func DeleteMany(db jack.DB, where *WhereClause) error {
	// sql select query, primary key provided by sequence
	sqlstr := `DELETE FROM "jack"."standups" WHERE ` + where.condition
	jack.Log(sqlstr, where.params...)

	if _, e := db.Exec(sqlstr, where.params...); e != nil {
		return e
	}

	return nil
}

// Upsert the "standup" by its "id".
func Upsert(db jack.DB, standupInput *StandupInput) (*Standup, error) {
	// get all the non-nil columns and prepare them for the query
	_c, _i, _v := jack.Slice(getColumns(standupInput), 0)

	// sql query
	sqlstr := `INSERT INTO "jack"."standups" (` + strings.Join(_c, ", ") + `) ` +
		`VALUES (` + strings.Join(_i, ", ") + `) ` +
		`ON CONFLICT ("id") ` +
		`DO UPDATE SET (` + strings.Join(_c, ", ") + `) = ( EXCLUDED.` + strings.Join(_c, ", EXCLUDED.") + `) ` +
		`RETURNING "id", "name", "slack_channel_id", "time", "timezone", "questions", "team_id", "created_at", "updated_at"`
	jack.Log(sqlstr, _v...)

	// run query
	var standup Standup
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&standup.ID, &standup.Name, &standup.SlackChannelID, &standup.Time, &standup.Timezone, &standup.Questions, &standup.TeamID, &standup.CreatedAt, &standup.UpdatedAt); e != nil {
		return nil, e
	}

	return &standup, nil
}

// UpsertBySlackChannelID find a "Standup"
func UpsertBySlackChannelID(db jack.DB, standupInput *StandupInput) (*Standup, error) {
	// get all the non-nil columns and prepare them for the query
	_c, _i, _v := jack.Slice(getColumns(standupInput), 0)

	// sql query
	sqlstr := `INSERT INTO "jack"."standups" (` + strings.Join(_c, ", ") + `) ` +
		`VALUES (` + strings.Join(_i, ", ") + `) ` +
		`ON CONFLICT (slack_channel_id) ` +
		`DO UPDATE SET (` + strings.Join(_c, ", ") + `) = ( EXCLUDED.` + strings.Join(_c, ", EXCLUDED.") + `) ` +
		`RETURNING "id", "name", "slack_channel_id", "time", "timezone", "questions", "team_id", "created_at", "updated_at"`
	jack.Log(sqlstr, _v...)

	// run query
	var standup Standup
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&standup.ID, &standup.Name, &standup.SlackChannelID, &standup.Time, &standup.Timezone, &standup.Questions, &standup.TeamID, &standup.CreatedAt, &standup.UpdatedAt); e != nil && e != pgx.ErrNoRows {
		return nil, e
	}

	return &standup, nil
}
