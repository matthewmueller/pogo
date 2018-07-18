package report

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/test/jack/pogo"
	"github.com/matthewmueller/pogo/test/jack/pogo/enum"
)

// GENERATED BY POGO. DO NOT EDIT.

// ErrReportNotFound returned if the reports is not found
var ErrReportNotFound = errors.New("report not found")

// Report result data for "jack"."reports"
type Report struct {
	ID         int               `json:"id,omitempty"`
	TeammateID int               `json:"teammate_id,omitempty"`
	StandupID  int               `json:"standup_id,omitempty"`
	Status     enum.ReportStatus `json:"status,omitempty"`
	Timestamp  int               `json:"timestamp,omitempty"`
}

// New "jack"."reports" input
func New() *Input {
	return &Input{}
}

// Input data for "jack"."reports"
type Input struct {
	id         *int
	teammateID *int
	standupID  *int
	status     *enum.ReportStatus
	timestamp  *int
}

// ID sets the "id"
func (r *Input) ID(id int) *Input {
	r.id = &id
	return r
}

// TeammateID sets the "teammate_id"
func (r *Input) TeammateID(teammateID int) *Input {
	r.teammateID = &teammateID
	return r
}

// StandupID sets the "standup_id"
func (r *Input) StandupID(standupID int) *Input {
	r.standupID = &standupID
	return r
}

// Status sets the "status"
func (r *Input) Status(status enum.ReportStatus) *Input {
	r.status = &status
	return r
}

// Timestamp sets the "timestamp"
func (r *Input) Timestamp(timestamp int) *Input {
	r.timestamp = &timestamp
	return r
}

// MarshalJSON marshals the "report" into JSON
func (r *Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(r)
}

// UnmarshalJSON unmarshals json to a "report"
func (r *Input) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, r)
}

func (r *Input) String() string {
	return "reports"
}

func (r *Input) columns() map[string]interface{} {
	columns := make(map[string]interface{})

	if r.id != nil {
		columns["id"] = *r.id
	}

	if r.teammateID != nil {
		columns["teammate_id"] = *r.teammateID
	}

	if r.standupID != nil {
		columns["standup_id"] = *r.standupID
	}

	if r.status != nil {
		columns["status"] = *r.status
	}

	if r.timestamp != nil {
		columns["timestamp"] = *r.timestamp
	}

	return columns
}

// NewFilter creates a new filter
func NewFilter() *Filter {
	return &Filter{}
}

// Filter filters for "jack"."reports"
type Filter struct {
	formats []string
	values  []interface{}
}

var _ pogo.Condition = (*Filter)(nil)

// Clause fn
func (f *Filter) Clause() *pogo.Clause {
	return &pogo.Clause{
		Type:   "WHERE",
		Format: strings.Join(f.formats, " AND "),
		Params: f.values,
	}
}

// And filter
// func (f *Filter) And(filters ...*Filter) *Filter {
//   var clauses []string
//   for _, filter := range filters {
//     _ = filter
//     // clauses = append(clauses, string(filter.Condition()))
//   }
//   f.clauses = append(f.clauses, strings.Join(clauses, " AND "))
//   return f
// }

// Or filter
// func (f *Filter) Or(filters ...*Filter) *Filter {
//   var clauses []string
//   for _, filter := range filters {
//     _ = filter
//     // clauses = append(clauses, string(filter.Condition()))
//   }
//   f.clauses = append(f.clauses, strings.Join(clauses, " OR "))
//   return f
// }

// ID id equals
func (f *Filter) ID(v int) *Filter {
	f.formats = append(f.formats, `id = %s`)
	f.values = append(f.values, v)
	return f
}

// IDNot id doesn't equal
func (f *Filter) IDNot(v int) *Filter {
	f.formats = append(f.formats, `id != %s`)
	f.values = append(f.values, v)
	return f
}

// IDLt id is less than
func (f *Filter) IDLt(v int) *Filter {
	f.formats = append(f.formats, `id < %s`)
	f.values = append(f.values, v)
	return f
}

// IDLte id is less than or equal
func (f *Filter) IDLte(v int) *Filter {
	f.formats = append(f.formats, `id <= %s`)
	f.values = append(f.values, v)
	return f
}

// IDGt id is greater than
func (f *Filter) IDGt(v int) *Filter {
	f.formats = append(f.formats, `id > %s`)
	f.values = append(f.values, v)
	return f
}

// IDGte id is greater than or equal
func (f *Filter) IDGte(v int) *Filter {
	f.formats = append(f.formats, `id >= %s`)
	f.values = append(f.values, v)
	return f
}

// IDIn id is in
func (f *Filter) IDIn(v int) *Filter {
	f.formats = append(f.formats, `id IN (%s)`)
	f.values = append(f.values, v)
	return f
}

// IDNotIn id is not in
func (f *Filter) IDNotIn(v int) *Filter {
	f.formats = append(f.formats, `id NOT IN (%s)`)
	f.values = append(f.values, v)
	return f
}

// TeammateID teammate_id equals
func (f *Filter) TeammateID(v int) *Filter {
	f.formats = append(f.formats, `teammate_id = %s`)
	f.values = append(f.values, v)
	return f
}

// TeammateIDNot teammate_id doesn't equal
func (f *Filter) TeammateIDNot(v int) *Filter {
	f.formats = append(f.formats, `teammate_id != %s`)
	f.values = append(f.values, v)
	return f
}

// TeammateIDLt teammate_id is less than
func (f *Filter) TeammateIDLt(v int) *Filter {
	f.formats = append(f.formats, `teammate_id < %s`)
	f.values = append(f.values, v)
	return f
}

// TeammateIDLte teammate_id is less than or equal
func (f *Filter) TeammateIDLte(v int) *Filter {
	f.formats = append(f.formats, `teammate_id <= %s`)
	f.values = append(f.values, v)
	return f
}

// TeammateIDGt teammate_id is greater than
func (f *Filter) TeammateIDGt(v int) *Filter {
	f.formats = append(f.formats, `teammate_id > %s`)
	f.values = append(f.values, v)
	return f
}

// TeammateIDGte teammate_id is greater than or equal
func (f *Filter) TeammateIDGte(v int) *Filter {
	f.formats = append(f.formats, `teammate_id >= %s`)
	f.values = append(f.values, v)
	return f
}

// TeammateIDIn teammate_id is in
func (f *Filter) TeammateIDIn(v int) *Filter {
	f.formats = append(f.formats, `teammate_id IN (%s)`)
	f.values = append(f.values, v)
	return f
}

// TeammateIDNotIn teammate_id is not in
func (f *Filter) TeammateIDNotIn(v int) *Filter {
	f.formats = append(f.formats, `teammate_id NOT IN (%s)`)
	f.values = append(f.values, v)
	return f
}

// StandupID standup_id equals
func (f *Filter) StandupID(v int) *Filter {
	f.formats = append(f.formats, `standup_id = %s`)
	f.values = append(f.values, v)
	return f
}

// StandupIDNot standup_id doesn't equal
func (f *Filter) StandupIDNot(v int) *Filter {
	f.formats = append(f.formats, `standup_id != %s`)
	f.values = append(f.values, v)
	return f
}

// StandupIDLt standup_id is less than
func (f *Filter) StandupIDLt(v int) *Filter {
	f.formats = append(f.formats, `standup_id < %s`)
	f.values = append(f.values, v)
	return f
}

// StandupIDLte standup_id is less than or equal
func (f *Filter) StandupIDLte(v int) *Filter {
	f.formats = append(f.formats, `standup_id <= %s`)
	f.values = append(f.values, v)
	return f
}

// StandupIDGt standup_id is greater than
func (f *Filter) StandupIDGt(v int) *Filter {
	f.formats = append(f.formats, `standup_id > %s`)
	f.values = append(f.values, v)
	return f
}

// StandupIDGte standup_id is greater than or equal
func (f *Filter) StandupIDGte(v int) *Filter {
	f.formats = append(f.formats, `standup_id >= %s`)
	f.values = append(f.values, v)
	return f
}

// StandupIDIn standup_id is in
func (f *Filter) StandupIDIn(v int) *Filter {
	f.formats = append(f.formats, `standup_id IN (%s)`)
	f.values = append(f.values, v)
	return f
}

// StandupIDNotIn standup_id is not in
func (f *Filter) StandupIDNotIn(v int) *Filter {
	f.formats = append(f.formats, `standup_id NOT IN (%s)`)
	f.values = append(f.values, v)
	return f
}

// Status
func (f *Filter) Status(v enum.ReportStatus) *Filter {
	f.formats = append(f.formats, ``)
	f.values = append(f.values, v)
	return f
}

// StatusNot
func (f *Filter) StatusNot(v enum.ReportStatus) *Filter {
	f.formats = append(f.formats, ``)
	f.values = append(f.values, v)
	return f
}

// StatusIn
func (f *Filter) StatusIn(v enum.ReportStatus) *Filter {
	f.formats = append(f.formats, ``)
	f.values = append(f.values, v)
	return f
}

// StatusNotIn
func (f *Filter) StatusNotIn(v enum.ReportStatus) *Filter {
	f.formats = append(f.formats, ``)
	f.values = append(f.values, v)
	return f
}

// Timestamp timestamp equals
func (f *Filter) Timestamp(v int) *Filter {
	f.formats = append(f.formats, `timestamp = %s`)
	f.values = append(f.values, v)
	return f
}

// TimestampNot timestamp doesn't equal
func (f *Filter) TimestampNot(v int) *Filter {
	f.formats = append(f.formats, `timestamp != %s`)
	f.values = append(f.values, v)
	return f
}

// TimestampLt timestamp is less than
func (f *Filter) TimestampLt(v int) *Filter {
	f.formats = append(f.formats, `timestamp < %s`)
	f.values = append(f.values, v)
	return f
}

// TimestampLte timestamp is less than or equal
func (f *Filter) TimestampLte(v int) *Filter {
	f.formats = append(f.formats, `timestamp <= %s`)
	f.values = append(f.values, v)
	return f
}

// TimestampGt timestamp is greater than
func (f *Filter) TimestampGt(v int) *Filter {
	f.formats = append(f.formats, `timestamp > %s`)
	f.values = append(f.values, v)
	return f
}

// TimestampGte timestamp is greater than or equal
func (f *Filter) TimestampGte(v int) *Filter {
	f.formats = append(f.formats, `timestamp >= %s`)
	f.values = append(f.values, v)
	return f
}

// TimestampIn timestamp is in
func (f *Filter) TimestampIn(v int) *Filter {
	f.formats = append(f.formats, `timestamp IN (%s)`)
	f.values = append(f.values, v)
	return f
}

// TimestampNotIn timestamp is not in
func (f *Filter) TimestampNotIn(v int) *Filter {
	f.formats = append(f.formats, `timestamp NOT IN (%s)`)
	f.values = append(f.values, v)
	return f
}

// OrderBy specificies the ORDERBy BY <order>
type OrderBy string

const (
	// Asc sorts by ascending order
	ASC OrderBy = "ASC"

	// Desc sorts by descending order
	DESC OrderBy = "DESC"
)

// NewOrder fn
func NewOrder() *Order {
	return &Order{}
}

// Order orders the given fields
type Order struct {
	formats []string
}

// Clause fn
func (o *Order) Clause() *pogo.Clause {
	return &pogo.Clause{
		Type:   "ORDER BY",
		Format: strings.Join(o.formats, ", "),
	}
}

// ID sorts `id` by an expression
func (o *Order) ID(order OrderBy) *Order {
	o.formats = append(o.formats, fmt.Sprintf(`"id" %s`, order))
	return o
}

// TeammateID sorts `teammate_id` by an expression
func (o *Order) TeammateID(order OrderBy) *Order {
	o.formats = append(o.formats, fmt.Sprintf(`"teammate_id" %s`, order))
	return o
}

// StandupID sorts `standup_id` by an expression
func (o *Order) StandupID(order OrderBy) *Order {
	o.formats = append(o.formats, fmt.Sprintf(`"standup_id" %s`, order))
	return o
}

// Status sorts `status` by an expression
func (o *Order) Status(order OrderBy) *Order {
	o.formats = append(o.formats, fmt.Sprintf(`"status" %s`, order))
	return o
}

// Timestamp sorts `timestamp` by an expression
func (o *Order) Timestamp(order OrderBy) *Order {
	o.formats = append(o.formats, fmt.Sprintf(`"timestamp" %s`, order))
	return o
}

// Insert a "reports" into the "jack"."reports"
func Insert(db pogo.DB, report *Input) (*Report, error) {
	// get all the non-nil columns and prepare them for the query
	_c, _i, _v := slice(report.columns(), 0)

	// sql insert query, primary key provided by sequence
	sqlstr := `
    INSERT INTO "jack"."reports" (` + strings.Join(_c, ", ") + `)
    VALUES (` + strings.Join(_i, ", ") + `)
    RETURNING "id", "teammate_id", "standup_id", "status", "timestamp"
  `

	// log query if we've provided a logger
	if pogo.Log != nil {
		pogo.Log(sqlstr, _v...)
	}

	var _report Report
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&_report.ID, &_report.TeammateID, &_report.StandupID, &_report.Status, &_report.Timestamp); e != nil {
		return nil, e
	}

	return &_report, nil
}

// Find a `Report` by some conditions.
func Find(db pogo.DB, conds ...pogo.Condition) (*Report, error) {
	_s, _v, err := pogo.Conditions(conds...)
	if err != nil {
		return nil, err
	}

	// sql select query, primary key provided by sequence
	sqlstr := `SELECT "id", "teammate_id", "standup_id", "status", "timestamp" ` +
		`FROM "jack"."reports" ` +
		_s

	// log query if we've provided a logger
	if pogo.Log != nil {
		pogo.Log(sqlstr, _v...)
	}

	var _report Report
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&_report.ID, &_report.TeammateID, &_report.StandupID, &_report.Status, &_report.Timestamp); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrReportNotFound
		}
		return nil, e
	}

	return &_report, nil
}

// FindByID a `Report` by some conditions.
func FindByID(db pogo.DB, id int) (*Report, error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT "id", "teammate_id", "standup_id", "status", "timestamp"
    FROM "jack"."reports"
    WHERE "id" = $1
  `

	// log query if we've provided a logger
	if pogo.Log != nil {
		pogo.Log(sqlstr, &id)
	}

	var _report Report
	row := db.QueryRow(sqlstr, &id)
	if e := row.Scan(&_report.ID, &_report.TeammateID, &_report.StandupID, &_report.Status, &_report.Timestamp); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrReportNotFound
		}
		return nil, e
	}

	return &_report, nil
}

// FindMany finds many "jack"."reports" by a condition
func FindMany(db pogo.DB, conds ...pogo.Condition) ([]*Report, error) {
	var reports []*Report

	_s, _v, err := pogo.Conditions(conds...)
	if err != nil {
		return reports, err
	}

	// sql select query, primary key provided by sequence
	sqlstr := `SELECT "id", "teammate_id", "standup_id", "status", "timestamp" ` +
		`FROM "jack"."reports" ` +
		_s

	// log query if we've provided a logger
	if pogo.Log != nil {
		pogo.Log(sqlstr, _v...)
	}

	rows, err := db.Query(sqlstr, _v...)
	if err != nil {
		return reports, err
	}
	defer rows.Close()

	for rows.Next() {
		var _report Report
		if e := rows.Scan(&_report.ID, &_report.TeammateID, &_report.StandupID, &_report.Status, &_report.Timestamp); e != nil {
			if e == pgx.ErrNoRows {
				return reports, ErrReportNotFound
			}
			return reports, err
		}
		reports = append(reports, &_report)
	}
	if rows.Err() != nil {
		return reports, rows.Err()
	}

	return reports, nil
}

// UpdateByID a "report" in "jack"."reports" by its "id"
func UpdateByID(db pogo.DB, id int, report *Input) (*Report, error) {
	fields := report.columns()

	// don't update the primary key
	delete(fields, "id")

	// prepare the slices
	_c, _i, _v := slice(fields, 1)

	// setup the update fields
	var _u []string
	for i, c := range _c {
		_u = append(_u, c+" = "+_i[i])
	}

	// sql query
	sqlstr := `UPDATE "jack"."reports" SET ` +
		strings.Join(_u, ", ") + ` ` +
		`WHERE "id" = $1 ` +
		`RETURNING "id", "teammate_id", "standup_id", "status", "timestamp"`

	// setup query
	values := append([]interface{}{&id}, _v...)

	// log query if we've provided a logger
	if pogo.Log != nil {
		pogo.Log(sqlstr, values...)
	}

	// run the query
	var _report Report
	row := db.QueryRow(sqlstr, values...)
	if e := row.Scan(&_report.ID, &_report.TeammateID, &_report.StandupID, &_report.Status, &_report.Timestamp); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrReportNotFound
		}
		return nil, e
	}

	return &_report, nil
}

// Delete `Report`s by some conditions. Returns the first result.
func Delete(db pogo.DB, conds ...pogo.Condition) (*Report, error) {
	_s, _v, err := pogo.Conditions(conds...)
	if err != nil {
		return nil, err
	}

	// sql delete query
	sqlstr := `DELETE FROM "jack"."reports" ` +
		_s + ` ` +
		`RETURNING "id", "teammate_id", "standup_id", "status", "timestamp"`

	// log query if we've provided a logger
	if pogo.Log != nil {
		pogo.Log(sqlstr, _v...)
	}

	var _report Report
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&_report.ID, &_report.TeammateID, &_report.StandupID, &_report.Status, &_report.Timestamp); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrReportNotFound
		}
		return nil, e
	}

	return &_report, nil
}

// DeleteMany `Report`s by some conditions, returning all results.
func DeleteMany(db pogo.DB, conds ...pogo.Condition) ([]*Report, error) {
	var reports []*Report

	_s, _v, err := pogo.Conditions(conds...)
	if err != nil {
		return reports, err
	}

	// sql delete query
	sqlstr := `DELETE FROM "jack"."reports" ` +
		_s + ` ` +
		`RETURNING "id", "teammate_id", "standup_id", "status", "timestamp"`

	// log query if we've provided a logger
	if pogo.Log != nil {
		pogo.Log(sqlstr, _v...)
	}

	rows, err := db.Query(sqlstr, _v...)
	if err != nil {
		return reports, err
	}
	defer rows.Close()

	for rows.Next() {
		var _report Report
		if e := rows.Scan(&_report.ID, &_report.TeammateID, &_report.StandupID, &_report.Status, &_report.Timestamp); e != nil {
			if e == pgx.ErrNoRows {
				return reports, ErrReportNotFound
			}
			return reports, err
		}
		reports = append(reports, &_report)
	}
	if rows.Err() != nil {
		return reports, rows.Err()
	}

	return reports, nil
}

// DeleteByID a "report" from the "jack"."reports" table
func DeleteByID(db pogo.DB, id int) (*Report, error) {
	// sql delete query
	sqlstr := `DELETE FROM "jack"."reports" ` +
		`WHERE "id" = $1 ` +
		`RETURNING "id", "teammate_id", "standup_id", "status", "timestamp"`

	// log query if we've provided a logger
	if pogo.Log != nil {
		pogo.Log(sqlstr, id)
	}

	// run the query
	var _report Report
	row := db.QueryRow(sqlstr, id)
	if e := row.Scan(&_report.ID, &_report.TeammateID, &_report.StandupID, &_report.Status, &_report.Timestamp); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrReportNotFound
		}
		return nil, e
	}

	return &_report, nil
}

// slice converts our columns into something the sql driver can understand
func slice(columns map[string]interface{}, offset int) (c []string, i []string, v []interface{}) {
	n := offset + 1
	for col, val := range columns {
		c = append(c, `"`+col+`"`)
		i = append(i, "$"+strconv.Itoa(n))
		v = append(v, val)
		n++
	}
	return c, i, v
}
