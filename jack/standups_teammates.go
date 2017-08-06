package jack

import (
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx"
	uuid "github.com/satori/go.uuid"
)

// GENERATED BY POGO. DO NOT EDIT.

// ErrStandupTeammateNotFound returned if the standupteammate is not found
var ErrStandupTeammateNotFound = errors.New("standupteammate not found")

// StandupsTeammates class
type StandupsTeammates struct {
	db DB
}

// StandupTeammate model
type StandupTeammate struct {
	StandupID  *uuid.UUID `json:"standup_id,omitempty"`
	TeammateID *uuid.UUID `json:"teammate_id,omitempty"`
	TeamOwner  *bool      `json:"team_owner,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}

// standupteammate constructor
func standupteammate(db DB) *StandupsTeammates {
	return &StandupsTeammates{db}
}

// get all the non-nil fields
func fields(standupteammate *StandupTeammate) map[string]interface{} {
	fields := make(map[string]interface{})

	if standupteammate.StandupID != nil {
		fields["standup_id"] = standupteammate.StandupID
	}
	if standupteammate.TeammateID != nil {
		fields["teammate_id"] = standupteammate.TeammateID
	}
	if standupteammate.TeamOwner != nil {
		fields["team_owner"] = standupteammate.TeamOwner
	}
	if standupteammate.CreatedAt != nil {
		fields["created_at"] = standupteammate.CreatedAt
	}
	if standupteammate.UpdatedAt != nil {
		fields["updated_at"] = standupteammate.UpdatedAt
	}

	return fields
}

// Find a `standupteammate` by its `standup_id`, `teammate_id`
func (standupsteammates *StandupsTeammates) Find(standupID *uuid.UUID, teammateID *uuid.UUID) (standupteammate *StandupTeammate, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "standup_id", "teammate_id", "team_owner", "created_at", "updated_at"
	FROM jack.standups_teammates
	WHERE "standup_id" = $1 AND "teammate_id" = $2
	`

	Log(sqlstr, standupID, teammateID)
	row := standupsteammates.db.QueryRow(sqlstr, standupID, teammateID)
	if e := row.Scan(standupteammate.StandupID, standupteammate.TeammateID, standupteammate.TeamOwner, standupteammate.CreatedAt, standupteammate.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrStandupTeammateNotFound
		}
		return nil, e
	}

	return standupteammate, nil
}

// Insert a `standupteammate` into `jack.standups_teammates`
func (standupsteammates *StandupsTeammates) Insert(standupteammate StandupTeammate) (*StandupTeammate, error) {
	// get all the non-nil fields and prepare them for the query
	_c, _i, _v := slice(fields(&standupteammate), 0)

	// sql insert query, primary key provided by sequence
	sqlstr := `
	INSERT INTO jack.standups_teammates (` + strings.Join(_c, ", ") + `)
	VALUES (` + strings.Join(_i, ", ") + `)
	RETURNING "standup_id", "teammate_id", "team_owner", "created_at", "updated_at"
  `

	Log(sqlstr, _v...)
	row := standupsteammates.db.QueryRow(sqlstr, _v...)
	if e := row.Scan(standupteammate.StandupID, standupteammate.TeammateID, standupteammate.TeamOwner, standupteammate.CreatedAt, standupteammate.UpdatedAt); e != nil {
		return nil, e
	}

	return &standupteammate, nil
}

// Update a `StandupTeammate` by its `standup_id`, `teammate_id`
func (standupsteammates *StandupsTeammates) Update(standupID *uuid.UUID, teammateID *uuid.UUID, standupteammate StandupTeammate) (*StandupTeammate, error) {
	fieldset := fields(&standupteammate)

	// first check if we have the foreign keys
	if standupID == nil {
		return nil, errors.New(`"standupID" must be non-nil`)
	}
	if teammateID == nil {
		return nil, errors.New(`"teammateID" must be non-nil`)
	}

	// don't update the foreign keys
	delete(fieldset, "standup_id")
	delete(fieldset, "teammate_id")

	// prepare the slices
	_c, _i, _v := slice(fieldset, 2)

	// sql query
	sqlstr := `UPDATE jack.standups_teammates SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `)
		WHERE "standup_id" = $1 AND "teammate_id" = $2
		RETURNING "standup_id", "teammate_id", "team_owner", "created_at", "updated_at"`

	// run query
	values := []interface{}{}
	values = append(values, standupID)
	values = append(values, teammateID)

	values = append(values, _v...)
	Log(sqlstr, values...)

	row := standupsteammates.db.QueryRow(sqlstr, values...)
	if e := row.Scan(standupteammate.StandupID, standupteammate.TeammateID, standupteammate.TeamOwner, standupteammate.CreatedAt, standupteammate.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrStandupTeammateNotFound
		}
		return nil, e
	}

	return &standupteammate, nil
}

// Delete a `StandupTeammate` by its `standup_id`, `teammate_id`
func (standupsteammates *StandupsTeammates) Delete(standupID *uuid.UUID, teammateID *uuid.UUID) error {
	// sql query
	const sqlstr = `
	DELETE FROM jack.standups_teammates
	WHERE "standup_id" = $1 AND "teammate_id" = $2
	`

	// run query
	Log(sqlstr, standupID, teammateID)
	if _, e := standupsteammates.db.Exec(sqlstr, standupID, teammateID); e != nil {
		if e == pgx.ErrNoRows {
			return ErrStandupTeammateNotFound
		}
		return e
	}

	return nil
}
