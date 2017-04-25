package jack

// GENERATED BY POGO. DO NOT EDIT.

import (
	"errors"
	"strings"

	"github.com/matthewmueller/pgx"
)

// Upsert the Cron by the Primary Key
func (c *Crons) Upsert(cc *Cron, action string) (cron Cron, err error) {
	fields := c.getFields(cc)

	// prepare the slices
	c, i, v := querySlices(fields, 0)

	// determine on conflict action
	var upsertAction string
	if action == UpsertDoUpdate {
		upsertAction = `DO UPDATE SET (` + strings.Join(c, ", ") + `) = ( EXCLUDED.` + strings.Join(c, ", EXCLUDED.") + `)`
	} else if action == UpsertDoNothing {
		upsertAction = UpsertDoNothing
	} else {
		return cron, errors.New("invalid upsert action")
	}

	// sql query
	sqlstr := `INSERT INTO public.crons (` + strings.Join(c, ", ") + `) ` +
		`VALUES (` + strings.Join(i, ", ") + `) ` +
		`ON CONFLICT ("id") ` +
		upsertAction + ` ` +
		`RETURNING "id", "frequency", "job", "tz", "start_time", "end_time", "retry", "created", "updated"`

		// run query
	DBLog(sqlstr, v...)
	row := c.DB.QueryRow(sqlstr, v...)
	err = row.Scan(&cron.ID, &cron.Frequency, &cron.Job, &cron.Tz, &cron.StartTime, &cron.EndTime, &cron.Retry, &cron.Created, &cron.Updated)
	if err != nil && err != pgx.ErrNoRows {
		return cron, err
	}

	return cron, nil
}

// UpsertByJob find a Cron
func (c *Crons) UpsertByJob(cc *Cron, action string) (cron Cron, err error) {
	fields := c.getFields(cc)

	// prepare the slices
	c, i, v := querySlices(fields, 0)

	// determine on conflict action
	var upsertAction string
	if action == UpsertDoUpdate {
		upsertAction = `DO UPDATE SET (` + strings.Join(c, ", ") + `) = ( EXCLUDED.` + strings.Join(c, ", EXCLUDED.") + `)`
	} else if action == UpsertDoNothing {
		upsertAction = UpsertDoNothing
	} else {
		return cron, errors.New("invalid upsert action")
	}

	// sql query
	sqlstr := `INSERT INTO public.crons (` + strings.Join(c, ", ") + `) ` +
		`VALUES (` + strings.Join(i, ", ") + `) ` +
		`ON CONFLICT ("job") ` +
		upsertAction + ` ` +
		`RETURNING "id", "frequency", "job", "tz", "start_time", "end_time", "retry", "created", "updated"`

		// run query
	DBLog(sqlstr, v...)
	row := c.DB.QueryRow(sqlstr, v...)
	err = row.Scan(&cron.ID, &cron.Frequency, &cron.Job, &cron.Tz, &cron.StartTime, &cron.EndTime, &cron.Retry, &cron.Created, &cron.Updated)
	if err != nil && err != pgx.ErrNoRows {
		return cron, err
	}

	return cron, nil
}