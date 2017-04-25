package jack

import "github.com/matthewmueller/pgx"

// GENERATED BY POGO. DO NOT EDIT.

// FindOne Report by a condition
func (r *Reports) FindOne(condition string, params ...interface{}) (report Report, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT "id", "user_id", "timestamp", "questions", "standup_id", "status"
    FROM public.reports
    WHERE ` + condition

	DBLog(sqlstr, params...)
	row := r.DB.QueryRow(sqlstr, params...)
	err = row.Scan(&report.ID, &report.UserID, &report.Timestamp, &report.Questions, &report.StandupID, &report.Status)
	if err != nil {
		if err == pgx.ErrNoRows {
			return report, ErrReportNotFound
		}
		return report, err
	}

	return report, nil
}