package jack

import (
	"github.com/matthewmueller/pgx"
)

// GENERATED BY POGO. DO NOT EDIT.

// Find a Cron by "id"
func (c *Crons) Find(id *int) (cron Cron, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT "id", "frequency", "job", "tz", "start_time", "end_time", "retry", "created", "updated"
    FROM public.crons
    WHERE "id" = $1`

	DBLog(sqlstr, id)
	row := c.DB.QueryRow(sqlstr, id)
	err = row.Scan(&cron.ID, &cron.Frequency, &cron.Job, &cron.Tz, &cron.StartTime, &cron.EndTime, &cron.Retry, &cron.Created, &cron.Updated)
	if err != nil {
		if err == pgx.ErrNoRows {
			return cron, ErrCronNotFound
		}
		return cron, err
	}

	return cron, nil
}

// FindByJob find a Cron
func (c *Crons) FindByJob(job *string) (cron Cron, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
		SELECT "id", "frequency", "job", "tz", "start_time", "end_time", "retry", "created", "updated"
		FROM public.crons
		WHERE "job" = $1`

	DBLog(sqlstr, job)
	row := c.DB.QueryRow(sqlstr, job)
	err = row.Scan(&cron.ID, &cron.Frequency, &cron.Job, &cron.Tz, &cron.StartTime, &cron.EndTime, &cron.Retry, &cron.Created, &cron.Updated)
	if err != nil {
		if err == pgx.ErrNoRows {
			return cron, ErrCronNotFound
		}
		return cron, err
	}

	return cron, nil
}
