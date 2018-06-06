package standupsteammates

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/testjack"
)

// GENERATED BY POGO. DO NOT EDIT.

// ErrStandupTeammateNotFound returned if the standups_teammates is not found
var ErrStandupTeammateNotFound = errors.New("standups_teammates not found")

// StandupTeammateInput model for "jack"."standups_teammates"
type StandupTeammateInput struct {
	standupID  *string
	teammateID *string
	teamOwner  *bool
	createdAt  *time.Time
	updatedAt  *time.Time
}

// StandupTeammate model for "jack"."standups_teammates"
type StandupTeammate struct {
	StandupID  string
	TeammateID string
	TeamOwner  bool
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

// New "jack"."standups_teammates" API
func New() *StandupTeammateInput {
	return &StandupTeammateInput{}
}

// StandupID sets the "standupID"
func (standupTeammate *StandupTeammateInput) StandupID(standupID string) *StandupTeammateInput {
	standupTeammate.standupID = &standupID
	return standupTeammate
}

// TeammateID sets the "teammateID"
func (standupTeammate *StandupTeammateInput) TeammateID(teammateID string) *StandupTeammateInput {
	standupTeammate.teammateID = &teammateID
	return standupTeammate
}

// TeamOwner sets the "teamOwner"
func (standupTeammate *StandupTeammateInput) TeamOwner(teamOwner bool) *StandupTeammateInput {
	standupTeammate.teamOwner = &teamOwner
	return standupTeammate
}

// CreatedAt sets the "createdAt"
func (standupTeammate *StandupTeammateInput) CreatedAt(createdAt time.Time) *StandupTeammateInput {
	standupTeammate.createdAt = &createdAt
	return standupTeammate
}

// UpdatedAt sets the "updatedAt"
func (standupTeammate *StandupTeammateInput) UpdatedAt(updatedAt time.Time) *StandupTeammateInput {
	standupTeammate.updatedAt = &updatedAt
	return standupTeammate
}

// MarshalJSON marshals the "standupTeammate" into JSON
func (standupTeammate *StandupTeammateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(standupTeammate)
}

// UnmarshalJSON unmarshals json to a "standupTeammate"
func (standupTeammate *StandupTeammateInput) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, standupTeammate)
}

func (standupTeammate *StandupTeammateInput) String() string {
	return "standupTeammate"
}

func getColumns(standupTeammate *StandupTeammateInput) map[string]interface{} {
	columns := make(map[string]interface{})

	if standupTeammate.standupID != nil {
		columns["standup_id"] = *standupTeammate.standupID
	}

	if standupTeammate.teammateID != nil {
		columns["teammate_id"] = *standupTeammate.teammateID
	}

	if standupTeammate.teamOwner != nil {
		columns["team_owner"] = *standupTeammate.teamOwner
	}

	if standupTeammate.createdAt != nil {
		columns["created_at"] = *standupTeammate.createdAt
	}

	if standupTeammate.updatedAt != nil {
		columns["updated_at"] = *standupTeammate.updatedAt
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

// Insert a "standupTeammate" into "jack"."standups_teammates"
func Insert(db testjack.DB, standupTeammateInput *StandupTeammateInput) (*StandupTeammate, error) {
	// get all the non-nil fields and prepare them for the query
	_c, _i, _v := testjack.Slice(getColumns(standupTeammateInput), 0)

	// sql insert query, primary key provided by sequence
	sqlstr := `
  INSERT INTO jack"."standups_teammates (` + strings.Join(_c, ", ") + `)
  VALUES (` + strings.Join(_i, ", ") + `)
  RETURNING "standup_id", "teammate_id", "team_owner", "created_at", "updated_at"
  `
	testjack.Log(sqlstr, _v...)

	// run the query
	var standupTeammate StandupTeammate
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&standupTeammate.StandupID, &standupTeammate.TeammateID, &standupTeammate.TeamOwner, &standupTeammate.CreatedAt, &standupTeammate.UpdatedAt); e != nil {
		return nil, e
	}

	return &standupTeammate, nil
}

// Find a "StandupTeammate" by its standup_id and teammate_id
func Find(db testjack.DB, standupID string, teammateID string) (*StandupTeammate, error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT "standup_id", "teammate_id", "team_owner", "created_at", "updated_at"
    FROM "jack.standups_teammates"
    WHERE "standup_id" = $1 AND "teammate_id" = $2
  `
	testjack.Log(sqlstr, standupID, teammateID)

	// run the query
	var standupTeammate StandupTeammate
	row := db.QueryRow(sqlstr, standupID, teammateID)
	if e := row.Scan(&standupTeammate.StandupID, &standupTeammate.TeammateID, &standupTeammate.TeamOwner, &standupTeammate.CreatedAt, &standupTeammate.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrStandupTeammateNotFound
		}
		return nil, e
	}

	return &standupTeammate, nil
}

// Update a "StandupTeammate" by its standup_id and teammate_id
func Update(db testjack.DB, standupID string, teammateID string, standupTeammateInput *StandupTeammateInput) (*StandupTeammate, error) {
	fields := getColumns(standupTeammateInput)

	// don't update the foreign keys
	delete(fields, "standup_id")
	delete(fields, "teammate_id")

	// prepare the slices
	_c, _i, _v := testjack.Slice(fields, 2)

	// sql query
	sqlstr := `UPDATE "jack"."standups_teammates" SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `)
    WHERE "standup_id" = $1 AND "teammate_id" = $2
    RETURNING "standup_id", "teammate_id", "team_owner", "created_at", "updated_at"`

	// setup the query
	values := []interface{}{}
	values = append(values, standupID)
	values = append(values, teammateID)
	values = append(values, _v...)
	testjack.Log(sqlstr, values...)

	// run the query
	var standupTeammate StandupTeammate
	row := db.QueryRow(sqlstr, values...)
	if e := row.Scan(&standupTeammate.StandupID, &standupTeammate.TeammateID, &standupTeammate.TeamOwner, &standupTeammate.CreatedAt, &standupTeammate.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrStandupTeammateNotFound
		}
		return nil, e
	}

	return &standupTeammate, nil
}

// Delete a "StandupTeammate" by its standup_id and teammate_id.
func Delete(db testjack.DB, standupID string, teammateID string) error {
	// sql query
	const sqlstr = `
    DELETE FROM "jack"."standups_teammates"
    WHERE "standup_id" = $1 AND "teammate_id" = $2
  `
	testjack.Log(sqlstr, standupID, teammateID)

	// run query
	if _, e := db.Exec(sqlstr, standupID, teammateID); e != nil {
		if e == pgx.ErrNoRows {
			return ErrStandupTeammateNotFound
		}
		return e
	}

	return nil
}
