package model

// Delete deletes the Team from the database.
func (m *Teams) Delete(ID *string) (err error) {
	// sql query
	const sqlstr = `
    DELETE FROM jack.teams
    WHERE id = $1
  `

	// run query
	XOLog(sqlstr, ID)
	_, err = m.DB.Exec(sqlstr, ID)
	if err != nil {
		return err
	}

	return nil
}
