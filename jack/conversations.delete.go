package jack

import (
	"github.com/matthewmueller/pgx"
)

// GENERATED BY POGO. DO NOT EDIT.

// Delete the Conversation from the database.
func (c *Conversations) Delete(id *int) (err error) {
	// sql query
	sqlstr := `DELETE FROM public.conversations WHERE "id" = $1`

	// run query
	DBLog(sqlstr, id)
	_, err = c.DB.Exec(sqlstr, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrConversationNotFound
		}
		return err
	}

	return nil
}