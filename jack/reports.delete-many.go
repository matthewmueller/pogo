package jack

// GENERATED BY POGO. DO NOT EDIT.

// DeleteMany delete many Reports by a condition
func (r *Reports) DeleteMany(condition string, params ...interface{}) (err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `DELETE FROM jack.reports WHERE ` + condition

	DBLog(sqlstr, params...)
	_, err = r.DB.Exec(sqlstr, params...)
	if err != nil {
		return err
	}

	return nil
}
