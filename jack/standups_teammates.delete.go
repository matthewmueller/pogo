package jack

import "github.com/satori/go.uuid"

// GENERATED BY POGO. DO NOT EDIT.

// Delete the StandupTeammate from the database.
func (st *StandupTeammates) Delete(StandupID *uuid.UUID, TeammateID *uuid.UUID) (err error) {
	// sql query
	const sqlstr = `
    DELETE FROM jack.standups_teammates
    WHERE "standup_id" = $1 AND "teammate_id" = $2
  `

	// run query
	DBLog(sqlstr, StandupID, TeammateID)
	_, err = st.DB.Exec(sqlstr, StandupID, TeammateID)
	if err != nil {
		return err
	}

	return nil
}