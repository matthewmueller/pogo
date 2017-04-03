package model

import (
	"errors"
	"strings"
)

// Update the Team by the Primary Key
func (t *Team) Update(db DB) error {
	var err error

	fields := t.getFields()

	// first check if we have the primary key
	if t.ID == nil {
		return errors.New(`primary key "id" must be non-nil`)
	}

	// don't update the primary key
	delete(fields, "id")

	// prepare the slices
	c, i, v := querySlices(fields, 1)

	// sql query
	sqlstr := `UPDATE jack.teams SET (` +
		strings.Join(c, ", ") + `) = (` +
		strings.Join(i, ", ") + `)` +
		` WHERE id = $1`

	// run query
	values := append([]interface{}{t.ID}, v...)
	XOLog(sqlstr, values...)
	_, err = db.Exec(sqlstr, values...)
	return err
}
