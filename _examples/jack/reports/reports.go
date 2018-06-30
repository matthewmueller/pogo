package reports

// GENERATED BY POGO. DO NOT EDIT.

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx"
	"github.com/matthewmueller/jack/pogo/enum"
	jack "github.com/matthewmueller/pogo/_examples/jack"
)

// ErrReportNotFound returned if the reports is not found
var ErrReportNotFound = errors.New("reports not found")

// ReportInput model for "jack"."reports"
type ReportInput struct {
	id        *string
	userID    *string
	timestamp *time.Time
	questions *json.RawMessage
	standupID *string
	status    *enum.ReportStatus
	createdAt *time.Time
	updatedAt *time.Time
}

// Report model for "jack"."reports"
type Report struct {
	ID        string
	UserID    string
	Timestamp time.Time
	Questions json.RawMessage
	StandupID string
	Status    enum.ReportStatus
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// New "jack"."reports" API
func New() *ReportInput {
	return &ReportInput{}
}

// ID sets the "id"
func (report *ReportInput) ID(id string) *ReportInput {
	report.id = &id
	return report
}

// UserID sets the "userID"
func (report *ReportInput) UserID(userID string) *ReportInput {
	report.userID = &userID
	return report
}

// Timestamp sets the "timestamp"
func (report *ReportInput) Timestamp(timestamp time.Time) *ReportInput {
	report.timestamp = &timestamp
	return report
}

// Questions sets the "questions"
func (report *ReportInput) Questions(questions json.RawMessage) *ReportInput {
	report.questions = &questions
	return report
}

// StandupID sets the "standupID"
func (report *ReportInput) StandupID(standupID string) *ReportInput {
	report.standupID = &standupID
	return report
}

// Status sets the "status"
func (report *ReportInput) Status(status enum.ReportStatus) *ReportInput {
	report.status = &status
	return report
}

// CreatedAt sets the "createdAt"
func (report *ReportInput) CreatedAt(createdAt time.Time) *ReportInput {
	report.createdAt = &createdAt
	return report
}

// UpdatedAt sets the "updatedAt"
func (report *ReportInput) UpdatedAt(updatedAt time.Time) *ReportInput {
	report.updatedAt = &updatedAt
	return report
}

// MarshalJSON marshals the "report" into JSON
func (report *ReportInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(report)
}

// UnmarshalJSON unmarshals json to a "report"
func (report *ReportInput) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, report)
}

func (report *ReportInput) String() string {
	return "report"
}

func getColumns(report *ReportInput) map[string]interface{} {
	columns := make(map[string]interface{})

	if report.id != nil {
		columns["id"] = *report.id
	}

	if report.userID != nil {
		columns["user_id"] = *report.userID
	}

	if report.timestamp != nil {
		columns["timestamp"] = *report.timestamp
	}

	if report.questions != nil {
		columns["questions"] = *report.questions
	}

	if report.standupID != nil {
		columns["standup_id"] = *report.standupID
	}

	if report.status != nil {
		columns["status"] = *report.status
	}

	if report.createdAt != nil {
		columns["created_at"] = *report.createdAt
	}

	if report.updatedAt != nil {
		columns["updated_at"] = *report.updatedAt
	}

	return columns
}

// WhereClause is a struct to handle where clauses
type WhereClause struct {
	condition string
	params    []interface{}
}

// Where specifies the conditions
func Where(condition string, params ...interface{}) *WhereClause {
	return &WhereClause{
		condition: condition,
		params:    params,
	}
}

// Insert a "report" into the "jack"."reports"
func Insert(ctx context.Context, db jack.DB, reportInput *ReportInput) (*Report, error) {
	// get all the non-nil columns and prepare them for the query
	_c, _i, _v := jack.Slice(getColumns(reportInput), 0)

	// sql insert query, primary key provided by sequence
	sqlstr := `
    INSERT INTO "jack"."reports" (` + strings.Join(_c, ", ") + `)
    VALUES (` + strings.Join(_i, ", ") + `)
    RETURNING "id", "user_id", "timestamp", "questions", "standup_id", "status", "created_at", "updated_at"
  `
	jack.Log(sqlstr, _v...)

	var report Report
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&report.ID, &report.UserID, &report.Timestamp, &report.Questions, &report.StandupID, &report.Status, &report.CreatedAt, &report.UpdatedAt); e != nil {
		return nil, e
	}

	return &report, nil
}

// Find a `Report` by it's primary key `id`
func Find(ctx context.Context, db jack.DB, id string) (*Report, error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT "id", "user_id", "timestamp", "questions", "standup_id", "status", "created_at", "updated_at"
    FROM "jack"."reports"
    WHERE "id" = $1
  `
	jack.Log(sqlstr, &id)

	var report Report
	row := db.QueryRow(sqlstr, &id)
	if e := row.Scan(&report.ID, &report.UserID, &report.Timestamp, &report.Questions, &report.StandupID, &report.Status, &report.CreatedAt, &report.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrReportNotFound
		}
		return nil, e
	}

	return &report, nil
}

// FindOne find one report by a condition
func FindOne(ctx context.Context, db jack.DB, where *WhereClause) (*Report, error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
  SELECT "id", "user_id", "timestamp", "questions", "standup_id", "status", "created_at", "updated_at"
  FROM "jack"."reports"
  WHERE ` + where.condition
	jack.Log(sqlstr, where.params...)

	var report Report
	row := db.QueryRow(sqlstr, where.params...)
	if e := row.Scan(&report.ID, &report.UserID, &report.Timestamp, &report.Questions, &report.StandupID, &report.Status, &report.CreatedAt, &report.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrReportNotFound
		}
		return nil, e
	}

	return &report, nil
}

// FindWhere find many "report"'s by a given condition
func FindWhere(ctx context.Context, db jack.DB, where *WhereClause) ([]*Report, error) {
	reports := []*Report{}

	// sql select query, primary key provided by sequence
	sqlstr := `
  SELECT "id", "user_id", "timestamp", "questions", "standup_id", "status", "created_at", "updated_at"
  FROM "jack"."reports"
  WHERE ` + where.condition
	jack.Log(sqlstr, where.params...)

	rows, err := db.Query(sqlstr, where.params...)
	if err != nil {
		return reports, err
	}
	defer rows.Close()

	for rows.Next() {
		var report Report
		if e := rows.Scan(&report.ID, &report.UserID, &report.Timestamp, &report.Questions, &report.StandupID, &report.Status, &report.CreatedAt, &report.UpdatedAt); e != nil {
			if e == pgx.ErrNoRows {
				return reports, ErrReportNotFound
			}
			return reports, err
		}
		reports = append(reports, &report)
	}
	if rows.Err() != nil {
		return reports, rows.Err()
	}

	return reports, nil
}

// FindAll find all "report"'s
func FindAll(ctx context.Context, db jack.DB) ([]*Report, error) {
	reports := []*Report{}

	// sql select query, primary key provided by sequence
	sqlstr := `
  SELECT "id", "user_id", "timestamp", "questions", "standup_id", "status", "created_at", "updated_at"
  FROM "jack"."reports"`
	jack.Log(sqlstr)

	rows, err := db.Query(sqlstr)
	if err != nil {
		return reports, err
	}
	defer rows.Close()

	for rows.Next() {
		var report Report
		if e := rows.Scan(&report.ID, &report.UserID, &report.Timestamp, &report.Questions, &report.StandupID, &report.Status, &report.CreatedAt, &report.UpdatedAt); e != nil {
			if e == pgx.ErrNoRows {
				return reports, ErrReportNotFound
			}
			return reports, err
		}
		reports = append(reports, &report)
	}
	if rows.Err() != nil {
		return reports, rows.Err()
	}

	return reports, nil
}

// Update a "report" in "jack"."reports" by its "id"
func Update(ctx context.Context, db jack.DB, id string, reportInput *ReportInput) (*Report, error) {
	fields := getColumns(reportInput)

	// don't update the primary key
	delete(fields, "id")

	// prepare the slices
	_c, _i, _v := jack.Slice(fields, 1)

	// sql query
	sqlstr := `UPDATE "jack"."reports" SET (` +
		strings.Join(_c, ", ") +
		`) = (` +
		strings.Join(_i, ", ") +
		`)
    WHERE "id" = $1
    RETURNING "id", "user_id", "timestamp", "questions", "standup_id", "status", "created_at", "updated_at"`

	// setup query
	values := append([]interface{}{&id}, _v...)
	jack.Log(sqlstr, values...)

	// run the query
	var report Report
	row := db.QueryRow(sqlstr, values...)
	if e := row.Scan(&report.ID, &report.UserID, &report.Timestamp, &report.Questions, &report.StandupID, &report.Status, &report.CreatedAt, &report.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrReportNotFound
		}
		return nil, e
	}

	return &report, nil
}

// UpdateWhere rows in "jack"."reports" by a given condition
func UpdateWhere(ctx context.Context, db jack.DB, where *WhereClause, reportInput *ReportInput) ([]*Report, error) {
	reports := []*Report{}

	// prepare the slices
	_c, _i, _v := jack.Slice(getColumns(reportInput), len(where.params))

	// sql query
	sqlstr := `UPDATE "jack"."reports" SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `) ` +
		`WHERE ` + where.condition + ` ` +
		`RETURNING "id", "user_id", "timestamp", "questions", "standup_id", "status", "created_at", "updated_at"`

	// setup the query
	values := []interface{}{}
	values = append(values, where.params...)
	values = append(values, _v...)
	jack.Log(sqlstr, values...)

	// run query
	rows, err := db.Query(sqlstr, values...)
	if err != nil {
		return reports, err
	}
	defer rows.Close()

	for rows.Next() {
		var report Report
		if e := rows.Scan(&report.ID, &report.UserID, &report.Timestamp, &report.Questions, &report.StandupID, &report.Status, &report.CreatedAt, &report.UpdatedAt); e != nil {
			if e == pgx.ErrNoRows {
				return reports, ErrReportNotFound
			}
			return reports, err
		}
		reports = append(reports, &report)
	}
	if rows.Err() != nil {
		return reports, rows.Err()
	}

	return reports, nil
}

// Delete a "report" from the "jack"."reports" table
func Delete(ctx context.Context, db jack.DB, id string) error {
	// sql query
	sqlstr := `DELETE FROM "jack"."reports" WHERE "id" = $1`
	jack.Log(sqlstr, id)

	// run query
	if _, e := db.Exec(sqlstr, id); e != nil {
		if e == pgx.ErrNoRows {
			return ErrReportNotFound
		}
		return e
	}

	return nil
}

// DeleteWhere delete many "report"'s by the given condition
func DeleteWhere(ctx context.Context, db jack.DB, where *WhereClause) error {
	// sql select query, primary key provided by sequence
	sqlstr := `DELETE FROM "jack"."reports" WHERE ` + where.condition
	jack.Log(sqlstr, where.params...)

	if _, e := db.Exec(sqlstr, where.params...); e != nil {
		return e
	}

	return nil
}

// Upsert the "report" by its "id".
func Upsert(ctx context.Context, db jack.DB, reportInput *ReportInput) (*Report, error) {
	// get all the non-nil columns and prepare them for the query
	_c, _i, _v := jack.Slice(getColumns(reportInput), 0)

	// sql query
	sqlstr := `INSERT INTO "jack"."reports" (` + strings.Join(_c, ", ") + `) ` +
		`VALUES (` + strings.Join(_i, ", ") + `) ` +
		`ON CONFLICT ("id") ` +
		`DO UPDATE SET (` + strings.Join(_c, ", ") + `) = ( EXCLUDED.` + strings.Join(_c, ", EXCLUDED.") + `) ` +
		`RETURNING "id", "user_id", "timestamp", "questions", "standup_id", "status", "created_at", "updated_at"`
	jack.Log(sqlstr, _v...)

	// run query
	var report Report
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&report.ID, &report.UserID, &report.Timestamp, &report.Questions, &report.StandupID, &report.Status, &report.CreatedAt, &report.UpdatedAt); e != nil {
		return nil, e
	}

	return &report, nil
}
