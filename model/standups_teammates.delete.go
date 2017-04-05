package model

// GENERATED BY POGO. DO NOT EDIT.

// Delete the StandupTeammate from the database.
func (st *StandupTeammates) Delete(StandupID *string, TeammateID *string) (err error) {
 // sql query
 const sqlstr = `
    DELETE FROM jack.standups_teammates
    WHERE standup_id = $1 AND teammate_id = $2
  `

 // run query
 XOLog(sqlstr, StandupID, TeammateID)
 _, err = st.DB.Exec(sqlstr, StandupID, TeammateID)
 if err != nil {
  return err
 }

 return nil
}
