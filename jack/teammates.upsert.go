package jack

// GENERATED BY POGO. DO NOT EDIT.

import (
	"errors"
	"strings"

	"github.com/matthewmueller/pgx"
)

// Upsert the Teammate by the Primary Key
func (t *Teammates) Upsert(tt *Teammate, action string) (teammate Teammate, err error) {
	fields := t.getFields(tt)

	// prepare the slices
	_c, _i, _v := querySlices(fields, 0)

	// determine on conflict action
	var upsertAction string
	if action == UpsertDoUpdate {
		upsertAction = `DO UPDATE SET (` + strings.Join(_c, ", ") + `) = ( EXCLUDED.` + strings.Join(_c, ", EXCLUDED.") + `)`
	} else if action == UpsertDoNothing {
		upsertAction = UpsertDoNothing
	} else {
		return teammate, errors.New("invalid upsert action")
	}

	// sql query
	sqlstr := `INSERT INTO jack.teammates (` + strings.Join(_c, ", ") + `) ` +
		`VALUES (` + strings.Join(_i, ", ") + `) ` +
		`ON CONFLICT ("id") ` +
		upsertAction + ` ` +
		`RETURNING "id", "slack_id", "username", "first_name", "last_name", "email", "avatar", "timezone", "created_at", "updated_at"`

		// run query
	DBLog(sqlstr, _v...)
	row := t.DB.QueryRow(sqlstr, _v...)
	err = row.Scan(&teammate.ID, &teammate.SlackID, &teammate.Username, &teammate.FirstName, &teammate.LastName, &teammate.Email, &teammate.Avatar, &teammate.Timezone, &teammate.CreatedAt, &teammate.UpdatedAt)
	if err != nil && err != pgx.ErrNoRows {
		return teammate, err
	}

	return teammate, nil
}

// UpsertBySlackID find a Teammate
func (t *Teammates) UpsertBySlackID(tt *Teammate, action string) (teammate Teammate, err error) {
	fields := t.getFields(tt)

	// prepare the slices
	_c, _i, _v := querySlices(fields, 0)

	// determine on conflict action
	var upsertAction string
	if action == UpsertDoUpdate {
		upsertAction = `DO UPDATE SET (` + strings.Join(_c, ", ") + `) = ( EXCLUDED.` + strings.Join(_c, ", EXCLUDED.") + `)`
	} else if action == UpsertDoNothing {
		upsertAction = UpsertDoNothing
	} else {
		return teammate, errors.New("invalid upsert action")
	}

	// sql query
	sqlstr := `INSERT INTO jack.teammates (` + strings.Join(_c, ", ") + `) ` +
		`VALUES (` + strings.Join(_i, ", ") + `) ` +
		`ON CONFLICT ("slackid") ` +
		upsertAction + ` ` +
		`RETURNING "id", "slack_id", "username", "first_name", "last_name", "email", "avatar", "timezone", "created_at", "updated_at"`

		// run query
	DBLog(sqlstr, _v...)
	row := t.DB.QueryRow(sqlstr, _v...)
	err = row.Scan(&teammate.ID, &teammate.SlackID, &teammate.Username, &teammate.FirstName, &teammate.LastName, &teammate.Email, &teammate.Avatar, &teammate.Timezone, &teammate.CreatedAt, &teammate.UpdatedAt)
	if err != nil && err != pgx.ErrNoRows {
		return teammate, err
	}

	return teammate, nil
}
