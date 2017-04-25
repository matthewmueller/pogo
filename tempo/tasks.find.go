package tempo

import (
	"github.com/matthewmueller/pgx"
	"github.com/satori/go.uuid"
)

// GENERATED BY POGO. DO NOT EDIT.

// Find a Task by "id"
func (t *Tasks) Find(id *uuid.UUID) (task Task, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT "id", "key", "target", "rate", "offset", "timezone", "rate_type", "rate_options", "status", "target_type", "target_options", "user", "meta", "refreshed_at", "created_at", "updated_at"
    FROM public.tasks
    WHERE "id" = $1`

	DBLog(sqlstr, id)
	row := t.DB.QueryRow(sqlstr, id)
	err = row.Scan(&task.ID, &task.Key, &task.Target, &task.Rate, &task.Offset, &task.Timezone, &task.RateType, &task.RateOptions, &task.Status, &task.TargetType, &task.TargetOptions, &task.User, &task.Meta, &task.RefreshedAt, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return task, ErrTaskNotFound
		}
		return task, err
	}

	return task, nil
}

// FindByKeyAndUser find a Task
func (t *Tasks) FindByKeyAndUser(key *string, user *uuid.UUID) (task Task, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
		SELECT "id", "key", "target", "rate", "offset", "timezone", "rate_type", "rate_options", "status", "target_type", "target_options", "user", "meta", "refreshed_at", "created_at", "updated_at"
		FROM public.tasks
		WHERE "key" = $1 AND "user" = $2`

	DBLog(sqlstr, key, user)
	row := t.DB.QueryRow(sqlstr, key, user)
	err = row.Scan(&task.ID, &task.Key, &task.Target, &task.Rate, &task.Offset, &task.Timezone, &task.RateType, &task.RateOptions, &task.Status, &task.TargetType, &task.TargetOptions, &task.User, &task.Meta, &task.RefreshedAt, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return task, ErrTaskNotFound
		}
		return task, err
	}

	return task, nil
}