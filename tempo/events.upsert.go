package tempo

// GENERATED BY POGO. DO NOT EDIT.

import (
	"errors"
	"strings"

	"github.com/matthewmueller/pgx"
)

// Upsert the Event by the Primary Key
func (e *Events) Upsert(ee *Event, action string) (event Event, err error) {
	fields := e.getFields(ee)

	// prepare the slices
	c, i, v := querySlices(fields, 0)

	// determine on conflict action
	var upsertAction string
	if action == UpsertDoUpdate {
		upsertAction = `DO UPDATE SET (` + strings.Join(c, ", ") + `) = ( EXCLUDED.` + strings.Join(c, ", EXCLUDED.") + `)`
	} else if action == UpsertDoNothing {
		upsertAction = UpsertDoNothing
	} else {
		return event, errors.New("invalid upsert action")
	}

	// sql query
	sqlstr := `INSERT INTO public.events (` + strings.Join(c, ", ") + `) ` +
		`VALUES (` + strings.Join(i, ", ") + `) ` +
		`ON CONFLICT ("id") ` +
		upsertAction + ` ` +
		`RETURNING "id", "time", "task", "created_at", "triggered_at", "status", "response", "attempts"`

		// run query
	DBLog(sqlstr, v...)
	row := e.DB.QueryRow(sqlstr, v...)
	err = row.Scan(&event.ID, &event.Time, &event.Task, &event.CreatedAt, &event.TriggeredAt, &event.Status, &event.Response, &event.Attempts)
	if err != nil && err != pgx.ErrNoRows {
		return event, err
	}

	return event, nil
}

// UpsertByTaskAndTime find a Event
func (e *Events) UpsertByTaskAndTime(ee *Event, action string) (event Event, err error) {
	fields := e.getFields(ee)

	// prepare the slices
	c, i, v := querySlices(fields, 0)

	// determine on conflict action
	var upsertAction string
	if action == UpsertDoUpdate {
		upsertAction = `DO UPDATE SET (` + strings.Join(c, ", ") + `) = ( EXCLUDED.` + strings.Join(c, ", EXCLUDED.") + `)`
	} else if action == UpsertDoNothing {
		upsertAction = UpsertDoNothing
	} else {
		return event, errors.New("invalid upsert action")
	}

	// sql query
	sqlstr := `INSERT INTO public.events (` + strings.Join(c, ", ") + `) ` +
		`VALUES (` + strings.Join(i, ", ") + `) ` +
		`ON CONFLICT ("task", "time") ` +
		upsertAction + ` ` +
		`RETURNING "id", "time", "task", "created_at", "triggered_at", "status", "response", "attempts"`

		// run query
	DBLog(sqlstr, v...)
	row := e.DB.QueryRow(sqlstr, v...)
	err = row.Scan(&event.ID, &event.Time, &event.Task, &event.CreatedAt, &event.TriggeredAt, &event.Status, &event.Response, &event.Attempts)
	if err != nil && err != pgx.ErrNoRows {
		return event, err
	}

	return event, nil
}
