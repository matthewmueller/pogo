package teammates

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx"
	uuid "github.com/satori/go.uuid"
)

// GENERATED BY POGO. DO NOT EDIT.

// ErrTeammateNotFound returned if the teammate is not found
var ErrTeammateNotFound = errors.New("teammate not found")

// columns in `jack.teammates`
type columns struct {
	ID        *uuid.UUID `json:"id,omitempty"`
	SlackID   *string    `json:"slack_id,omitempty"`
	Username  *string    `json:"username,omitempty"`
	FirstName *string    `json:"first_name,omitempty"`
	LastName  *string    `json:"last_name,omitempty"`
	Email     *string    `json:"email,omitempty"`
	Avatar    *string    `json:"avatar,omitempty"`
	Timezone  *string    `json:"timezone,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// Teammate fluent API
type Teammate struct {
	columns *columns
}

// New `jack.teammates` API
func New() *Teammate {
	return &Teammate{&columns{}}
}

// ID sets the `id`
func (teammate *Teammate) ID(id uuid.UUID) *Teammate {
	teammate.columns.ID = &id
	return teammate
}

// GetID returns the `id` if set
func (teammate *Teammate) GetID() (id *uuid.UUID) {
	return teammate.columns.ID
}

// SlackID sets the `slack_id`
func (teammate *Teammate) SlackID(slackID string) *Teammate {
	teammate.columns.SlackID = &slackID
	return teammate
}

// GetSlackID returns the `slack_id` if set
func (teammate *Teammate) GetSlackID() (slackID *string) {
	return teammate.columns.SlackID
}

// Username sets the `username`
func (teammate *Teammate) Username(username string) *Teammate {
	teammate.columns.Username = &username
	return teammate
}

// GetUsername returns the `username` if set
func (teammate *Teammate) GetUsername() (username *string) {
	return teammate.columns.Username
}

// FirstName sets the `first_name`
func (teammate *Teammate) FirstName(firstName string) *Teammate {
	teammate.columns.FirstName = &firstName
	return teammate
}

// GetFirstName returns the `first_name` if set
func (teammate *Teammate) GetFirstName() (firstName *string) {
	return teammate.columns.FirstName
}

// LastName sets the `last_name`
func (teammate *Teammate) LastName(lastName string) *Teammate {
	teammate.columns.LastName = &lastName
	return teammate
}

// GetLastName returns the `last_name` if set
func (teammate *Teammate) GetLastName() (lastName *string) {
	return teammate.columns.LastName
}

// Email sets the `email`
func (teammate *Teammate) Email(email string) *Teammate {
	teammate.columns.Email = &email
	return teammate
}

// GetEmail returns the `email` if set
func (teammate *Teammate) GetEmail() (email *string) {
	return teammate.columns.Email
}

// Avatar sets the `avatar`
func (teammate *Teammate) Avatar(avatar string) *Teammate {
	teammate.columns.Avatar = &avatar
	return teammate
}

// GetAvatar returns the `avatar` if set
func (teammate *Teammate) GetAvatar() (avatar *string) {
	return teammate.columns.Avatar
}

// Timezone sets the `timezone`
func (teammate *Teammate) Timezone(timezone string) *Teammate {
	teammate.columns.Timezone = &timezone
	return teammate
}

// GetTimezone returns the `timezone` if set
func (teammate *Teammate) GetTimezone() (timezone *string) {
	return teammate.columns.Timezone
}

// CreatedAt sets the `created_at`
func (teammate *Teammate) CreatedAt(createdAt time.Time) *Teammate {
	teammate.columns.CreatedAt = &createdAt
	return teammate
}

// GetCreatedAt returns the `created_at` if set
func (teammate *Teammate) GetCreatedAt() (createdAt *time.Time) {
	return teammate.columns.CreatedAt
}

// UpdatedAt sets the `updated_at`
func (teammate *Teammate) UpdatedAt(updatedAt time.Time) *Teammate {
	teammate.columns.UpdatedAt = &updatedAt
	return teammate
}

// GetUpdatedAt returns the `updated_at` if set
func (teammate *Teammate) GetUpdatedAt() (updatedAt *time.Time) {
	return teammate.columns.UpdatedAt
}

// MarshalJSON marshals the `teammate` into JSON
func (teammate *Teammate) MarshalJSON() ([]byte, error) {
	return json.Marshal(teammate.columns)
}

// UnmarshalJSON unmarshals json to a `teammate`
func (teammate *Teammate) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, teammate.columns)
}

func (teammate *Teammate) String() string {
	return "teammate TODO"
}

// get all the non-nil columns
func getColumns(teammate *Teammate) map[string]interface{} {
	columns := make(map[string]interface{})

	if teammate.columns.ID != nil {
		columns["id"] = teammate.ID
	}
	if teammate.columns.SlackID != nil {
		columns["slack_id"] = teammate.SlackID
	}
	if teammate.columns.Username != nil {
		columns["username"] = teammate.Username
	}
	if teammate.columns.FirstName != nil {
		columns["first_name"] = teammate.FirstName
	}
	if teammate.columns.LastName != nil {
		columns["last_name"] = teammate.LastName
	}
	if teammate.columns.Email != nil {
		columns["email"] = teammate.Email
	}
	if teammate.columns.Avatar != nil {
		columns["avatar"] = teammate.Avatar
	}
	if teammate.columns.Timezone != nil {
		columns["timezone"] = teammate.Timezone
	}
	if teammate.columns.CreatedAt != nil {
		columns["created_at"] = teammate.CreatedAt
	}
	if teammate.columns.UpdatedAt != nil {
		columns["updated_at"] = teammate.UpdatedAt
	}

	return columns
}

// Find a teammate by "id"
func Find(db jack.DB, id *uuid.UUID) (*Teammate, error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "slack_id", "username", "first_name", "last_name", "email", "avatar", "timezone", "created_at", "updated_at"
	FROM jack.teammates
	WHERE "id" = $1
	`
	jack.Log(sqlstr, id)

	var cols *columns
	row := db.QueryRow(sqlstr, id)
	if e := row.Scan(cols.ID, cols.SlackID, cols.Username, cols.FirstName, cols.LastName, cols.Email, cols.Avatar, cols.Timezone, cols.CreatedAt, cols.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrTeammateNotFound
		}
		return nil, e
	}

	return &Teammate{cols}, nil
}

// FindBySlackID find a teammate by `slack_id`
func FindBySlackID(db jack.DB, slackID *string) (*Teammate, error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "slack_id", "username", "first_name", "last_name", "email", "avatar", "timezone", "created_at", "updated_at"
	FROM jack.teammates
	WHERE "slack_id" = $1
	`
	jack.Log(sqlstr, slackID)

	var cols *columns
	row := db.QueryRow(sqlstr, slackID)
	err := row.Scan(cols.ID, cols.SlackID, cols.Username, cols.FirstName, cols.LastName, cols.Email, cols.Avatar, cols.Timezone, cols.CreatedAt, cols.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrTeammateNotFound
		}
		return nil, err
	}

	return &Teammate{cols}, nil
}

// FindMany find many `teammate`s by a given condition
func FindMany(db jack.DB, condition string, params ...interface{}) ([]*Teammate, error) {
	var _o []*Teammate

	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "slack_id", "username", "first_name", "last_name", "email", "avatar", "timezone", "created_at", "updated_at"
	FROM jack.teammates
	WHERE ` + condition
	jack.Log(sqlstr, params...)

	rows, err := db.Query(sqlstr, params...)
	if err != nil {
		return _o, err
	}
	defer rows.Close()

	for rows.Next() {
		var cols *columns
		if e := rows.Scan(cols.ID, cols.SlackID, cols.Username, cols.FirstName, cols.LastName, cols.Email, cols.Avatar, cols.Timezone, cols.CreatedAt, cols.UpdatedAt); e != nil {
			if e == pgx.ErrNoRows {
				return _o, ErrTeammateNotFound
			}
			return _o, err
		}
		_o = append(_o, &Teammate{cols})
	}
	if rows.Err() != nil {
		return _o, rows.Err()
	}

	// ensure we return an empty array
	// rather than nil when we marshal
	if len(_o) == 0 {
		return make([]*Teammate, 0), nil
	}

	return _o, nil
}

// FindOne find one teammate by a condition
func FindOne(db jack.DB, condition string, params ...interface{}) (*Teammate, error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "slack_id", "username", "first_name", "last_name", "email", "avatar", "timezone", "created_at", "updated_at"
	FROM jack.teammates
	WHERE ` + condition
	jack.Log(sqlstr, params...)

	var cols *columns
	row := db.QueryRow(sqlstr, params...)
	if e := row.Scan(cols.ID, cols.SlackID, cols.Username, cols.FirstName, cols.LastName, cols.Email, cols.Avatar, cols.Timezone, cols.CreatedAt, cols.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrTeammateNotFound
		}
		return nil, e
	}

	return &Teammate{cols}, nil
}

// Insert a `teammate` into the `jack.teammates` table.
func Insert(db jack.DB, teammate Teammate) (*Teammate, error) {
	// get all the non-nil columns and prepare them for the query
	_c, _i, _v := jack.Slice(getColumns(&teammate), 0)

	// sql insert query, primary key provided by sequence
	sqlstr := `
	INSERT INTO jack.teammates (` + strings.Join(_c, ", ") + `)
	VALUES (` + strings.Join(_i, ", ") + `)
	RETURNING "id", "slack_id", "username", "first_name", "last_name", "email", "avatar", "timezone", "created_at", "updated_at"
	`
	jack.Log(sqlstr, _v...)

	var cols *columns
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(cols.ID, cols.SlackID, cols.Username, cols.FirstName, cols.LastName, cols.Email, cols.Avatar, cols.Timezone, cols.CreatedAt, cols.UpdatedAt); e != nil {
		return nil, e
	}

	return &Teammate{cols}, nil
}

// Update a teammate by its `id`
func Update(db jack.DB, teammate Teammate, id *uuid.UUID) (*Teammate, error) {
	fields := getColumns(&teammate)

	// first check if we have the primary key
	if id == nil {
		return nil, errors.New(`primary key "id" must be non-nil`)
	}

	// don't update the primary key
	delete(fields, "id")

	// prepare the slices
	_c, _i, _v := jack.Slice(fields, 1)

	// sql query
	sqlstr := `UPDATE jack.teammates SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `)
		WHERE "id" = $1
		RETURNING "id", "slack_id", "username", "first_name", "last_name", "email", "avatar", "timezone", "created_at", "updated_at"`

	// setup query
	values := append([]interface{}{id}, _v...)
	jack.Log(sqlstr, values...)

	// run the query
	var cols *columns
	row := db.QueryRow(sqlstr, values...)
	if e := row.Scan(cols.ID, cols.SlackID, cols.Username, cols.FirstName, cols.LastName, cols.Email, cols.Avatar, cols.Timezone, cols.CreatedAt, cols.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrTeammateNotFound
		}
		return nil, e
	}

	return &Teammate{cols}, nil
}

// UpdateBySlackID find a Teammate
func UpdateBySlackID(db jack.DB, teammate Teammate, slackID *string) (*Teammate, error) {
	fields := getColumns(&teammate)

	// first check if we have all the keys we need
	if slackID == nil {
		return nil, errors.New(`slackID must be non-nil`)
	}

	// don't update the keys
	delete(fields, "slackID")

	// prepare the slices
	_c, _i, _v := jack.Slice(fields, 1)

	// sql query
	sqlstr := `UPDATE jack.teammates SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `) ` +
		`WHERE "slack_id" = $1 ` +
		`RETURNING "id", "slack_id", "username", "first_name", "last_name", "email", "avatar", "timezone", "created_at", "updated_at"`

	// setup the query
	values := []interface{}{}
	values = append(values, slackID)

	values = append(values, _v...)
	jack.Log(sqlstr, values...)

	// run the query
	var cols *columns
	row := db.QueryRow(sqlstr, values...)
	if e := row.Scan(cols.ID, cols.SlackID, cols.Username, cols.FirstName, cols.LastName, cols.Email, cols.Avatar, cols.Timezone, cols.CreatedAt, cols.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrTeammateNotFound
		}
		return nil, e
	}

	return &Teammate{cols}, nil
}

// UpdateMany rows in `jack.teammates` by a given condition
func UpdateMany(db jack.DB, teammate *Teammate, condition string, params ...interface{}) ([]*Teammate, error) {
	var _o []*Teammate

	// prepare the slices
	_c, _i, _v := jack.Slice(getColumns(teammate), len(params))

	// sql query
	sqlstr := `UPDATE jack.teammates SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `) ` +
		`WHERE ` + condition + ` ` +
		`RETURNING "id", "slack_id", "username", "first_name", "last_name", "email", "avatar", "timezone", "created_at", "updated_at"`

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
		if e := rows.Scan(cols.ID, cols.SlackID, cols.Username, cols.FirstName, cols.LastName, cols.Email, cols.Avatar, cols.Timezone, cols.CreatedAt, cols.UpdatedAt); e != nil {
			if e == pgx.ErrNoRows {
				return _o, ErrTeammateNotFound
			}
			return _o, err
		}
		_o = append(_o, &Teammate{cols})
	}
	if rows.Err() != nil {
		return _o, rows.Err()
	}

	// ensure we return an empty array
	// rather than nil when we marshal
	if len(_o) == 0 {
		return make([]*Teammate, 0), nil
	}

	return _o, nil
}

// Delete a `teammate` from the `jack.teammates` table
func Delete(db jack.DB, id *uuid.UUID) error {
	// sql query
	sqlstr := `DELETE FROM jack.teammates WHERE "id" = $1`
	jack.Log(sqlstr, id)

	// run query
	if _, e := db.Exec(sqlstr, id); e != nil {
		if e == pgx.ErrNoRows {
			return ErrTeammateNotFound
		}
		return e
	}

	return nil
}

// DeleteBySlackID find a Teammate
func DeleteBySlackID(db jack.DB, slackID *string) error {
	// sql delete query
	sqlstr := `DELETE FROM jack.teammates WHERE "slack_id" = $1`
	jack.Log(sqlstr, slackID)

	if _, e := db.Exec(sqlstr, slackID); e != nil {
		if e == pgx.ErrNoRows {
			return ErrTeammateNotFound
		}
		return e
	}

	return nil
}

// DeleteMany delete many `teammate`'s by the given condition
func DeleteMany(db jack.DB, condition string, params ...interface{}) error {
	// sql select query, primary key provided by sequence
	sqlstr := `DELETE FROM jack.teammates WHERE ` + condition
	jack.Log(sqlstr, params...)

	if _, e := db.Exec(sqlstr, params...); e != nil {
		return e
	}

	return nil
}

// Upsert the `teammate` by its `id`.
func Upsert(db jack.DB, teammate Teammate, action string) (*Teammate, error) {
	// prepare the slices
	_c, _i, _v := jack.Slice(getColumns(&teammate), 0)

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
	sqlstr := `INSERT INTO jack.teammates (` + strings.Join(_c, ", ") + `) ` +
		`VALUES (` + strings.Join(_i, ", ") + `) ` +
		`ON CONFLICT ("id") ` +
		upsertAction + ` ` +
		`RETURNING "id", "slack_id", "username", "first_name", "last_name", "email", "avatar", "timezone", "created_at", "updated_at"`
	jack.Log(sqlstr, _v...)

	// run query
	var cols *columns
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(cols.ID, cols.SlackID, cols.Username, cols.FirstName, cols.LastName, cols.Email, cols.Avatar, cols.Timezone, cols.CreatedAt, cols.UpdatedAt); e != nil && e != pgx.ErrNoRows {
		return nil, e
	}

	return &Teammate{cols}, nil
}

// UpsertBySlackID find a Teammate
func UpsertBySlackID(db jack.DB, teammate Teammate, action string) (*Teammate, error) {
	// get all the non-nil columns and prepare them for the query
	_c, _i, _v := jack.Slice(getColumns(&teammate), 0)

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
	sqlstr := `INSERT INTO jack.teammates (` + strings.Join(_c, ", ") + `) ` +
		`VALUES (` + strings.Join(_i, ", ") + `) ` +
		`ON CONFLICT ("slack_id") ` +
		upsertAction + ` ` +
		`RETURNING "id", "slack_id", "username", "first_name", "last_name", "email", "avatar", "timezone", "created_at", "updated_at"`
	jack.Log(sqlstr, _v...)

	// run query
	var cols *columns
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(cols.ID, cols.SlackID, cols.Username, cols.FirstName, cols.LastName, cols.Email, cols.Avatar, cols.Timezone, cols.CreatedAt, cols.UpdatedAt); e != nil && e != pgx.ErrNoRows {
		return nil, e
	}

	return &Teammate{cols}, nil
}
