package model

import (
 "errors"
 "strings"
)

// GENERATED BY POGO. DO NOT EDIT.

// Update the StandupTeammate by the Primary Key
func (st *StandupTeammates) Update(StandupID *string, TeammateID *string, stst *StandupTeammate) (standupteammate StandupTeammate, err error) {
 fields := st.getFields(stst)

 // first check if we have the foreign keys

 if StandupID == nil {
  return standupteammate, errors.New(`"standup_id" must be non-nil`)
 }

 if TeammateID == nil {
  return standupteammate, errors.New(`"teammate_id" must be non-nil`)
 }

 // don't update the foreign keys

 delete(fields, "standup_id")

 delete(fields, "teammate_id")

 // prepare the slices
 c, i, v := querySlices(fields, 1)

 // sql query
 sqlstr := `UPDATE jack.standups_teammates SET (` +
  strings.Join(c, ", ") + `) = (` +
  strings.Join(i, ", ") + `)
		WHERE standup_id = $1 AND teammate_id = $2
		RETURNING standup_id, teammate_id, team_owner, created_at, updated_at`

 // run query
 values := []interface{}{}

 values = append(values, StandupID)

 values = append(values, TeammateID)

 values = append(values, v...)
 DBLog(sqlstr, values...)

 row := st.DB.QueryRow(sqlstr, values...)
 err = row.Scan(&standupteammate.StandupID, &standupteammate.TeammateID, &standupteammate.TeamOwner, &standupteammate.CreatedAt, &standupteammate.UpdatedAt)
 if err != nil {
  return standupteammate, err
 }

 return standupteammate, nil
}