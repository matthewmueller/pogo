package tempo

// GENERATED BY POGO. DO NOT EDIT.

import (
	"errors"
	"strings"

	"github.com/satori/go.uuid"
)

// Update the Active by the Primary Key
func (a *Actives) Update(aa *Active, id *uuid.UUID) (active Active, err error) {
	fields := a.getFields(aa)

	// first check if we have the primary key
	if id == nil {
		return active, errors.New(`primary key "id" must be non-nil`)
	}

	// don't update the primary key
	delete(fields, "id")

	// prepare the slices
	c, i, v := querySlices(fields, 1)

	// sql query
	sqlstr := `UPDATE public.actives SET (` +
		strings.Join(c, ", ") + `) = (` +
		strings.Join(i, ", ") + `)
		WHERE "id" = $1
		RETURNING "id", "token", "state", "used", "created_at", "updated_at"`

	// run query
	values := append([]interface{}{id}, v...)
	DBLog(sqlstr, values...)

	row := a.DB.QueryRow(sqlstr, values...)
	err = row.Scan(&active.ID, &active.Token, &active.State, &active.Used, &active.CreatedAt, &active.UpdatedAt)
	if err != nil {
		return active, err
	}

	return active, nil
}

// UpdateByState find a Active
func (a *Actives) UpdateByState(aa *Active, state *uuid.UUID) (active Active, err error) {
	fields := a.getFields(aa)

	// first check if we have all the keys we need
	if state == nil {
		return active, errors.New(`state must be non-nil`)
	}

	// don't update the keys
	delete(fields, "state")

	// prepare the slices
	c, i, v := querySlices(fields, 1)

	// sql query
	sqlstr := `UPDATE public.actives SET (` +
		strings.Join(c, ", ") + `) = (` +
		strings.Join(i, ", ") + `) ` +
		`WHERE "state" = $1 ` +
		`RETURNING "id", "token", "state", "used", "created_at", "updated_at"`

	// run query
	values := []interface{}{}
	values = append(values, state)

	values = append(values, v...)
	DBLog(sqlstr, values...)

	row := a.DB.QueryRow(sqlstr, values...)
	err = row.Scan(&active.ID, &active.Token, &active.State, &active.Used, &active.CreatedAt, &active.UpdatedAt)
	if err != nil {
		return active, err
	}

	return active, nil
}

// UpdateByToken find a Active
func (a *Actives) UpdateByToken(aa *Active, token *uuid.UUID) (active Active, err error) {
	fields := a.getFields(aa)

	// first check if we have all the keys we need
	if token == nil {
		return active, errors.New(`token must be non-nil`)
	}

	// don't update the keys
	delete(fields, "token")

	// prepare the slices
	c, i, v := querySlices(fields, 1)

	// sql query
	sqlstr := `UPDATE public.actives SET (` +
		strings.Join(c, ", ") + `) = (` +
		strings.Join(i, ", ") + `) ` +
		`WHERE "token" = $1 ` +
		`RETURNING "id", "token", "state", "used", "created_at", "updated_at"`

	// run query
	values := []interface{}{}
	values = append(values, token)

	values = append(values, v...)
	DBLog(sqlstr, values...)

	row := a.DB.QueryRow(sqlstr, values...)
	err = row.Scan(&active.ID, &active.Token, &active.State, &active.Used, &active.CreatedAt, &active.UpdatedAt)
	if err != nil {
		return active, err
	}

	return active, nil
}