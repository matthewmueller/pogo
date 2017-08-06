package jack

import (
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx"
	uuid "github.com/satori/go.uuid"
)

// GENERATED BY POGO. DO NOT EDIT.

// ErrReportNotFound returned if the report is not found
var ErrReportNotFound = errors.New("report not found")

// Reports class
type Reports struct {
	db DB
}

// Report model
type Report struct {
	ID        *uuid.UUID              `json:"id,omitempty"`
	UserID    *uuid.UUID              `json:"user_id,omitempty"`
	Timestamp *time.Time              `json:"timestamp,omitempty"`
	Questions *map[string]interface{} `json:"questions,omitempty"`
	StandupID *uuid.UUID              `json:"standup_id,omitempty"`
	Status    *ReportStatus           `json:"status,omitempty"`
	CreatedAt *time.Time              `json:"created_at,omitempty"`
	UpdatedAt *time.Time              `json:"updated_at,omitempty"`
}

// report constructor
func report(db DB) *Reports {
	return &Reports{db}
}

// get all the non-nil fields
func (reports *Reports) fields(report *Report) map[string]interface{} {
	fields := make(map[string]interface{})

	if report.ID != nil {
		fields["id"] = report.ID
	}
	if report.UserID != nil {
		fields["user_id"] = report.UserID
	}
	if report.Timestamp != nil {
		fields["timestamp"] = report.Timestamp
	}
	if report.Questions != nil {
		fields["questions"] = report.Questions
	}
	if report.StandupID != nil {
		fields["standup_id"] = report.StandupID
	}
	if report.Status != nil {
		fields["status"] = report.Status
	}
	if report.CreatedAt != nil {
		fields["created_at"] = report.CreatedAt
	}
	if report.UpdatedAt != nil {
		fields["updated_at"] = report.UpdatedAt
	}

	return fields
}

// Find a report by "id"
func (reports *Reports) Find(id *uuid.UUID) (report *Report, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "user_id", "timestamp", "questions", "standup_id", "status", "created_at", "updated_at"
	FROM jack.reports
	WHERE "id" = $1
	`

	Log(sqlstr, id)
	row := reports.db.QueryRow(sqlstr, id)
	if e := row.Scan(report.ID, report.UserID, report.Timestamp, report.Questions, report.StandupID, report.Status, report.CreatedAt, report.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrReportNotFound
		}
		return nil, e
	}

	return report, nil
}

// FindMany find many `report`s by a given condition
func (reports *Reports) FindMany(condition string, params ...interface{}) ([]*Report, error) {
	var _o []*Report

	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "user_id", "timestamp", "questions", "standup_id", "status", "created_at", "updated_at"
	FROM jack.reports
	WHERE ` + condition

	Log(sqlstr, params...)
	rows, err := reports.db.Query(sqlstr, params...)
	if err != nil {
		return _o, err
	}
	defer rows.Close()

	for rows.Next() {
		var report *Report
		if e := rows.Scan(report.ID, report.UserID, report.Timestamp, report.Questions, report.StandupID, report.Status, report.CreatedAt, report.UpdatedAt); e != nil {
			if e == pgx.ErrNoRows {
				return _o, ErrReportNotFound
			}
			return _o, err
		}
		_o = append(_o, report)
	}
	if rows.Err() != nil {
		return _o, rows.Err()
	}

	// ensure we return an empty array
	// rather than nil when we marshal
	if len(_o) == 0 {
		return make([]*Report, 0), nil
	}

	return _o, nil
}

// FindOne find one report by a condition
func (reports *Reports) FindOne(condition string, params ...interface{}) (report *Report, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "user_id", "timestamp", "questions", "standup_id", "status", "created_at", "updated_at"
	FROM jack.reports
	WHERE ` + condition

	Log(sqlstr, params...)
	row := reports.db.QueryRow(sqlstr, params...)
	if e := row.Scan(report.ID, report.UserID, report.Timestamp, report.Questions, report.StandupID, report.Status, report.CreatedAt, report.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrReportNotFound
		}
		return nil, e
	}

	return report, nil
}

// Insert a `report` into the `jack.reports` table.
func (reports *Reports) Insert(report Report) (*Report, error) {
	// get all the non-nil fields and prepare them for the query
	_c, _i, _v := slice(reports.fields(&report), 0)

	// sql insert query, primary key provided by sequence
	sqlstr := `
	INSERT INTO jack.reports (` + strings.Join(_c, ", ") + `)
	VALUES (` + strings.Join(_i, ", ") + `)
	RETURNING "id", "user_id", "timestamp", "questions", "standup_id", "status", "created_at", "updated_at"
	`

	Log(sqlstr, _v...)
	row := reports.db.QueryRow(sqlstr, _v...)
	if e := row.Scan(report.ID, report.UserID, report.Timestamp, report.Questions, report.StandupID, report.Status, report.CreatedAt, report.UpdatedAt); e != nil {
		return nil, e
	}

	return &report, nil
}

// Update a report by its `id`
func (reports *Reports) Update(report Report, id *uuid.UUID) (*Report, error) {
	fieldset := reports.fields(&report)

	// first check if we have the primary key
	if id == nil {
		return nil, errors.New(`primary key "id" must be non-nil`)
	}

	// don't update the primary key
	delete(fieldset, "id")

	// prepare the slices
	_c, _i, _v := slice(fieldset, 1)

	// sql query
	sqlstr := `UPDATE jack.reports SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `)
		WHERE "id" = $1
		RETURNING "id", "user_id", "timestamp", "questions", "standup_id", "status", "created_at", "updated_at"`

	// run query
	values := append([]interface{}{id}, _v...)
	Log(sqlstr, values...)

	row := reports.db.QueryRow(sqlstr, values...)
	if e := row.Scan(report.ID, report.UserID, report.Timestamp, report.Questions, report.StandupID, report.Status, report.CreatedAt, report.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrReportNotFound
		}
		return nil, e
	}

	return &report, nil
}

// UpdateMany rows in `jack.reports` by a given condition
func (reports *Reports) UpdateMany(report *Report, condition string, params ...interface{}) ([]*Report, error) {
	var _o []*Report

	// get the non-nil fields
	fieldset := reports.fields(report)

	// prepare the slices
	_c, _i, _v := slice(fieldset, len(params))

	// sql query
	sqlstr := `UPDATE jack.reports SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `) ` +
		`WHERE ` + condition + ` ` +
		`RETURNING "id", "user_id", "timestamp", "questions", "standup_id", "status", "created_at", "updated_at"`

	values := []interface{}{}
	values = append(values, params...)
	values = append(values, _v...)

	// run query
	Log(sqlstr, values...)
	rows, err := reports.db.Query(sqlstr, values...)
	if err != nil {
		return _o, err
	}
	defer rows.Close()

	for rows.Next() {
		var report *Report
		if e := rows.Scan(report.ID, report.UserID, report.Timestamp, report.Questions, report.StandupID, report.Status, report.CreatedAt, report.UpdatedAt); e != nil {
			if e == pgx.ErrNoRows {
				return _o, ErrReportNotFound
			}
			return _o, err
		}
		_o = append(_o, report)
	}
	if rows.Err() != nil {
		return _o, rows.Err()
	}

	// ensure we return an empty array
	// rather than nil when we marshal
	if len(_o) == 0 {
		return make([]*Report, 0), nil
	}

	return _o, nil
}

// Delete a `report` from the `jack.reports` table
func (reports *Reports) Delete(id *uuid.UUID) error {
	// sql query
	sqlstr := `DELETE FROM jack.reports WHERE "id" = $1`

	// run query
	Log(sqlstr, id)
	if _, e := reports.db.Exec(sqlstr, id); e != nil {
		if e == pgx.ErrNoRows {
			return ErrReportNotFound
		}
		return e
	}

	return nil
}

// DeleteMany delete many `report`'s by the given condition
func (reports *Reports) DeleteMany(condition string, params ...interface{}) error {
	// sql select query, primary key provided by sequence
	sqlstr := `DELETE FROM jack.reports WHERE ` + condition

	Log(sqlstr, params...)
	if _, e := reports.db.Exec(sqlstr, params...); e != nil {
		return e
	}

	return nil
}

// Upsert the `report` by its `id`.
func (reports *Reports) Upsert(report Report, action string) (*Report, error) {
	fieldset := reports.fields(&report)

	// prepare the slices
	_c, _i, _v := slice(fieldset, 0)

	// determine on conflict action
	var upsertAction string
	if action == UpsertDoUpdate {
		upsertAction = `DO UPDATE SET (` + strings.Join(_c, ", ") + `) = ( EXCLUDED.` + strings.Join(_c, ", EXCLUDED.") + `)`
	} else if action == UpsertDoNothing {
		upsertAction = UpsertDoNothing
	} else {
		return nil, errors.New("invalid upsert action")
	}

	// sql query
	sqlstr := `INSERT INTO jack.reports (` + strings.Join(_c, ", ") + `) ` +
		`VALUES (` + strings.Join(_i, ", ") + `) ` +
		`ON CONFLICT ("id") ` +
		upsertAction + ` ` +
		`RETURNING "id", "user_id", "timestamp", "questions", "standup_id", "status", "created_at", "updated_at"`

		// run query
	Log(sqlstr, _v...)
	row := reports.db.QueryRow(sqlstr, _v...)
	if e := row.Scan(report.ID, report.UserID, report.Timestamp, report.Questions, report.StandupID, report.Status, report.CreatedAt, report.UpdatedAt); e != nil && e != pgx.ErrNoRows {
		return nil, e
	}

	return &report, nil
}
