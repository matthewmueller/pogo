package eventsusers

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx"
)

// GENERATED BY POGO. DO NOT EDIT.

// ErrEventUserNotFound returned if the events_users is not found
var ErrEventUserNotFound = errors.New("events_users not found")

// EventUserInput model for "public"."events_users"
type EventUserInput struct {
	userID    *string
	eventID   *string
	verb      *string
	createdAt *time.Time
}

// EventUser model for "public"."events_users"
type EventUser struct {
	UserID    string
	EventID   string
	Verb      string
	CreatedAt *time.Time
}

// New "public"."events_users" API
func New() *EventUserInput {
	return &EventUserInput{}
}

// UserID sets the "userID"
func (eventUser *EventUserInput) UserID(userID string) *EventUserInput {
	eventUser.userID = &userID
	return eventUser
}

// EventID sets the "eventID"
func (eventUser *EventUserInput) EventID(eventID string) *EventUserInput {
	eventUser.eventID = &eventID
	return eventUser
}

// Verb sets the "verb"
func (eventUser *EventUserInput) Verb(verb string) *EventUserInput {
	eventUser.verb = &verb
	return eventUser
}

// CreatedAt sets the "createdAt"
func (eventUser *EventUserInput) CreatedAt(createdAt time.Time) *EventUserInput {
	eventUser.createdAt = &createdAt
	return eventUser
}

// MarshalJSON marshals the "eventUser" into JSON
func (eventUser *EventUserInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(eventUser)
}

// UnmarshalJSON unmarshals json to a "eventUser"
func (eventUser *EventUserInput) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, eventUser)
}

func (eventUser *EventUserInput) String() string {
	return "eventUser"
}

func getColumns(eventUser *EventUserInput) map[string]interface{} {
	columns := make(map[string]interface{})

	if eventUser.userID != nil {
		columns["user_id"] = *eventUser.userID
	}

	if eventUser.eventID != nil {
		columns["event_id"] = *eventUser.eventID
	}

	if eventUser.verb != nil {
		columns["verb"] = *eventUser.verb
	}

	if eventUser.createdAt != nil {
		columns["created_at"] = *eventUser.createdAt
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

// Insert a "eventUser" into "public"."events_users"
func Insert(db digby.DB, eventUserInput *EventUserInput) (*EventUser, error) {
	// get all the non-nil fields and prepare them for the query
	_c, _i, _v := digby.Slice(getColumns(eventUserInput), 0)

	// sql insert query, primary key provided by sequence
	sqlstr := `
  INSERT INTO "public"."events_users" (` + strings.Join(_c, ", ") + `)
  VALUES (` + strings.Join(_i, ", ") + `)
  RETURNING "user_id", "event_id", "verb", "created_at"
  `
	digby.Log(sqlstr, _v...)

	// run the query
	var eventUser EventUser
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&eventUser.UserID, &eventUser.EventID, &eventUser.Verb, &eventUser.CreatedAt); e != nil {
		return nil, e
	}

	return &eventUser, nil
}

// Find a "EventUser" by its event_id and user_id
func Find(db digby.DB, eventID string, userID string) (*EventUser, error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT "user_id", "event_id", "verb", "created_at"
    FROM "public.events_users"
    WHERE "event_id" = $1 AND "user_id" = $2
  `
	digby.Log(sqlstr, eventID, userID)

	// run the query
	var eventUser EventUser
	row := db.QueryRow(sqlstr, eventID, userID)
	if e := row.Scan(&eventUser.UserID, &eventUser.EventID, &eventUser.Verb, &eventUser.CreatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrEventUserNotFound
		}
		return nil, e
	}

	return &eventUser, nil
}

// Update a "EventUser" by its event_id and user_id
func Update(db digby.DB, eventID string, userID string, eventUserInput *EventUserInput) (*EventUser, error) {
	fields := getColumns(eventUserInput)

	// don't update the foreign keys
	delete(fields, "event_id")
	delete(fields, "user_id")

	// prepare the slices
	_c, _i, _v := digby.Slice(fields, 2)

	// sql query
	sqlstr := `UPDATE "public"."events_users" SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `)
    WHERE "event_id" = $1 AND "user_id" = $2
    RETURNING "user_id", "event_id", "verb", "created_at"`

	// setup the query
	values := []interface{}{}
	values = append(values, eventID)
	values = append(values, userID)
	values = append(values, _v...)
	digby.Log(sqlstr, values...)

	// run the query
	var eventUser EventUser
	row := db.QueryRow(sqlstr, values...)
	if e := row.Scan(&eventUser.UserID, &eventUser.EventID, &eventUser.Verb, &eventUser.CreatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrEventUserNotFound
		}
		return nil, e
	}

	return &eventUser, nil
}

// Delete a "EventUser" by its event_id and user_id.
func Delete(db digby.DB, eventID string, userID string) error {
	// sql query
	const sqlstr = `
    DELETE FROM "public"."events_users"
    WHERE "event_id" = $1 AND "user_id" = $2
  `
	digby.Log(sqlstr, eventID, userID)

	// run query
	if _, e := db.Exec(sqlstr, eventID, userID); e != nil {
		if e == pgx.ErrNoRows {
			return ErrEventUserNotFound
		}
		return e
	}

	return nil
}