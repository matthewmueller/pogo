package question

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/test/jack/pogo"
)

// GENERATED BY POGO. DO NOT EDIT.

// ErrQuestionNotFound returned if the questions is not found
var ErrQuestionNotFound = errors.New("question not found")

// Question result data for "jack"."questions"
type Question struct {
	ID        int    `json:"id,omitempty"`
	Order     int    `json:"order,omitempty"`
	StandupID int    `json:"standup_id,omitempty"`
	Question  string `json:"question,omitempty"`
}

// New "jack"."questions" input
func New() *Input {
	return &Input{}
}

// Input data for "jack"."questions"
type Input struct {
	id        *int
	order     *int
	standupID *int
	question  *string
}

// ID sets the "id"
func (q *Input) ID(id int) *Input {
	q.id = &id
	return q
}

// Order sets the "order"
func (q *Input) Order(order int) *Input {
	q.order = &order
	return q
}

// StandupID sets the "standup_id"
func (q *Input) StandupID(standupID int) *Input {
	q.standupID = &standupID
	return q
}

// Question sets the "question"
func (q *Input) Question(question string) *Input {
	q.question = &question
	return q
}

// MarshalJSON marshals the "question" into JSON
func (q *Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(q)
}

// UnmarshalJSON unmarshals json to a "question"
func (q *Input) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, q)
}

func (q *Input) String() string {
	return "questions"
}

func (q *Input) columns() map[string]interface{} {
	columns := make(map[string]interface{})

	if q.id != nil {
		columns["id"] = *q.id
	}

	if q.order != nil {
		columns["order"] = *q.order
	}

	if q.standupID != nil {
		columns["standup_id"] = *q.standupID
	}

	if q.question != nil {
		columns["question"] = *q.question
	}

	return columns
}

// NewFilter creates a new filter
func NewFilter() *Filter {
	return &Filter{}
}

// Filter filters for "jack"."questions"
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
func (f *Filter) IDIn(v ...int) *Filter {
	var rs []string
	for range v {
		rs = append(rs, "%s")
	}
	f.formats = append(f.formats, fmt.Sprintf(`id IN (%s)`, strings.Join(rs, `, `)))
	for _, i := range v {
		f.values = append(f.values, i)
	}
	return f
}

// IDNotIn id is not in
func (f *Filter) IDNotIn(v ...int) *Filter {
	var rs []string
	for range v {
		rs = append(rs, "%s")
	}
	f.formats = append(f.formats, fmt.Sprintf(`id NOT IN (%s)`, strings.Join(rs, `, `)))
	for _, i := range v {
		f.values = append(f.values, i)
	}
	return f
}

// Order order equals
func (f *Filter) Order(v int) *Filter {
	f.formats = append(f.formats, `order = %s`)
	f.values = append(f.values, v)
	return f
}

// OrderNot order doesn't equal
func (f *Filter) OrderNot(v int) *Filter {
	f.formats = append(f.formats, `order != %s`)
	f.values = append(f.values, v)
	return f
}

// OrderLt order is less than
func (f *Filter) OrderLt(v int) *Filter {
	f.formats = append(f.formats, `order < %s`)
	f.values = append(f.values, v)
	return f
}

// OrderLte order is less than or equal
func (f *Filter) OrderLte(v int) *Filter {
	f.formats = append(f.formats, `order <= %s`)
	f.values = append(f.values, v)
	return f
}

// OrderGt order is greater than
func (f *Filter) OrderGt(v int) *Filter {
	f.formats = append(f.formats, `order > %s`)
	f.values = append(f.values, v)
	return f
}

// OrderGte order is greater than or equal
func (f *Filter) OrderGte(v int) *Filter {
	f.formats = append(f.formats, `order >= %s`)
	f.values = append(f.values, v)
	return f
}

// OrderIn order is in
func (f *Filter) OrderIn(v ...int) *Filter {
	var rs []string
	for range v {
		rs = append(rs, "%s")
	}
	f.formats = append(f.formats, fmt.Sprintf(`order IN (%s)`, strings.Join(rs, `, `)))
	for _, i := range v {
		f.values = append(f.values, i)
	}
	return f
}

// OrderNotIn order is not in
func (f *Filter) OrderNotIn(v ...int) *Filter {
	var rs []string
	for range v {
		rs = append(rs, "%s")
	}
	f.formats = append(f.formats, fmt.Sprintf(`order NOT IN (%s)`, strings.Join(rs, `, `)))
	for _, i := range v {
		f.values = append(f.values, i)
	}
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
func (f *Filter) StandupIDIn(v ...int) *Filter {
	var rs []string
	for range v {
		rs = append(rs, "%s")
	}
	f.formats = append(f.formats, fmt.Sprintf(`standup_id IN (%s)`, strings.Join(rs, `, `)))
	for _, i := range v {
		f.values = append(f.values, i)
	}
	return f
}

// StandupIDNotIn standup_id is not in
func (f *Filter) StandupIDNotIn(v ...int) *Filter {
	var rs []string
	for range v {
		rs = append(rs, "%s")
	}
	f.formats = append(f.formats, fmt.Sprintf(`standup_id NOT IN (%s)`, strings.Join(rs, `, `)))
	for _, i := range v {
		f.values = append(f.values, i)
	}
	return f
}

// Question question equals
func (f *Filter) Question(v string) *Filter {
	f.formats = append(f.formats, `question = %s`)
	f.values = append(f.values, v)
	return f
}

// QuestionNot question doesn't equal
func (f *Filter) QuestionNot(v string) *Filter {
	f.formats = append(f.formats, `question != %s`)
	f.values = append(f.values, v)
	return f
}

// QuestionContains question contains
func (f *Filter) QuestionContains(v string) *Filter {
	f.formats = append(f.formats, `question LIKE '%%' || %s || '%%'`)
	f.values = append(f.values, v)
	return f
}

// QuestionNotContains question doesn't contain
func (f *Filter) QuestionNotContains(v string) *Filter {
	f.formats = append(f.formats, `question NOT LIKE '%%' || %s || '%%'`)
	f.values = append(f.values, v)
	return f
}

// QuestionStartsWith question starts with
func (f *Filter) QuestionStartsWith(v string) *Filter {
	f.formats = append(f.formats, `question LIKE %s || '%%'`)
	f.values = append(f.values, v)
	return f
}

// QuestionNotStartsWith question doesn't start with
func (f *Filter) QuestionNotStartsWith(v string) *Filter {
	f.formats = append(f.formats, `question NOT LIKE %s || '%%'`)
	f.values = append(f.values, v)
	return f
}

// QuestionEndsWith question ends with
func (f *Filter) QuestionEndsWith(v string) *Filter {
	f.formats = append(f.formats, `question LIKE '%%' || %s`)
	f.values = append(f.values, v)
	return f
}

// QuestionNotEndsWith question doesn't end with
func (f *Filter) QuestionNotEndsWith(v string) *Filter {
	f.formats = append(f.formats, `question NOT LIKE '%%' || %s`)
	f.values = append(f.values, v)
	return f
}

// QuestionLt question is less than
func (f *Filter) QuestionLt(v string) *Filter {
	f.formats = append(f.formats, `question < %s`)
	f.values = append(f.values, v)
	return f
}

// QuestionLte question is less than or equal
func (f *Filter) QuestionLte(v string) *Filter {
	f.formats = append(f.formats, `question <= %s`)
	f.values = append(f.values, v)
	return f
}

// QuestionGt question is greater than
func (f *Filter) QuestionGt(v string) *Filter {
	f.formats = append(f.formats, `question > %s`)
	f.values = append(f.values, v)
	return f
}

// QuestionGte question is greater than or equal
func (f *Filter) QuestionGte(v string) *Filter {
	f.formats = append(f.formats, `question >= %s`)
	f.values = append(f.values, v)
	return f
}

// QuestionIn question is in
func (f *Filter) QuestionIn(v ...string) *Filter {
	var rs []string
	for range v {
		rs = append(rs, "%s")
	}
	f.formats = append(f.formats, fmt.Sprintf(`question IN (%s)`, strings.Join(rs, `, `)))
	for _, i := range v {
		f.values = append(f.values, i)
	}
	return f
}

// QuestionNotIn question is not in
func (f *Filter) QuestionNotIn(v ...string) *Filter {
	var rs []string
	for range v {
		rs = append(rs, "%s")
	}
	f.formats = append(f.formats, fmt.Sprintf(`question NOT IN (%s)`, strings.Join(rs, `, `)))
	for _, i := range v {
		f.values = append(f.values, i)
	}
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

var _ pogo.Condition = (*Order)(nil)

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

// Order sorts `order` by an expression
func (o *Order) Order(order OrderBy) *Order {
	o.formats = append(o.formats, fmt.Sprintf(`"order" %s`, order))
	return o
}

// StandupID sorts `standup_id` by an expression
func (o *Order) StandupID(order OrderBy) *Order {
	o.formats = append(o.formats, fmt.Sprintf(`"standup_id" %s`, order))
	return o
}

// Question sorts `question` by an expression
func (o *Order) Question(order OrderBy) *Order {
	o.formats = append(o.formats, fmt.Sprintf(`"question" %s`, order))
	return o
}

// Insert a "questions" into the "jack"."questions"
func Insert(db pogo.DB, question *Input) (*Question, error) {
	// get all the non-nil columns and prepare them for the query
	_c, _i, _v := slice(question.columns(), 0)

	// sql insert query, primary key provided by sequence
	sqlstr := `
    INSERT INTO "jack"."questions" (` + strings.Join(_c, ", ") + `)
    VALUES (` + strings.Join(_i, ", ") + `)
    RETURNING "id", "order", "standup_id", "question"
  `

	// log query if we've provided a logger
	if pogo.Log != nil {
		pogo.Log(sqlstr, _v...)
	}

	var _question Question
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&_question.ID, &_question.Order, &_question.StandupID, &_question.Question); e != nil {
		return nil, e
	}

	return &_question, nil
}

// Find a `Question` by some conditions.
func Find(db pogo.DB, conds ...pogo.Condition) (*Question, error) {
	_s, _v, err := pogo.Conditions(conds...)
	if err != nil {
		return nil, err
	}

	// sql select query, primary key provided by sequence
	sqlstr := `SELECT "id", "order", "standup_id", "question" ` +
		`FROM "jack"."questions" ` +
		_s

	// log query if we've provided a logger
	if pogo.Log != nil {
		pogo.Log(sqlstr, _v...)
	}

	var _question Question
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&_question.ID, &_question.Order, &_question.StandupID, &_question.Question); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrQuestionNotFound
		}
		return nil, e
	}

	return &_question, nil
}

// FindByID a `Question` by some conditions.
func FindByID(db pogo.DB, id int) (*Question, error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT "id", "order", "standup_id", "question"
    FROM "jack"."questions"
    WHERE "id" = $1
  `

	// log query if we've provided a logger
	if pogo.Log != nil {
		pogo.Log(sqlstr, &id)
	}

	var _question Question
	row := db.QueryRow(sqlstr, &id)
	if e := row.Scan(&_question.ID, &_question.Order, &_question.StandupID, &_question.Question); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrQuestionNotFound
		}
		return nil, e
	}

	return &_question, nil
}

// FindMany finds many "jack"."questions" by a condition
func FindMany(db pogo.DB, conds ...pogo.Condition) ([]*Question, error) {
	var questions []*Question

	_s, _v, err := pogo.Conditions(conds...)
	if err != nil {
		return questions, err
	}

	// sql select query, primary key provided by sequence
	sqlstr := `SELECT "id", "order", "standup_id", "question" ` +
		`FROM "jack"."questions" ` +
		_s

	// log query if we've provided a logger
	if pogo.Log != nil {
		pogo.Log(sqlstr, _v...)
	}

	rows, err := db.Query(sqlstr, _v...)
	if err != nil {
		return questions, err
	}
	defer rows.Close()

	for rows.Next() {
		var _question Question
		if e := rows.Scan(&_question.ID, &_question.Order, &_question.StandupID, &_question.Question); e != nil {
			if e == pgx.ErrNoRows {
				return questions, ErrQuestionNotFound
			}
			return questions, err
		}
		questions = append(questions, &_question)
	}
	if rows.Err() != nil {
		return questions, rows.Err()
	}

	return questions, nil
}

// Update "questions" rows in "jack"."questions" by a condition, returning 1 result
func Update(db pogo.DB, question *Input, conds ...pogo.Condition) (*Question, error) {
	_s, _v, err := pogo.Conditions(conds...)
	if err != nil {
		return nil, err
	}

	fields := question.columns()

	// prepare the slices
	_c, _i, _v2 := slice(fields, len(_v))
	_v = append(_v, _v2...)

	// setup the update fields
	var _u []string
	for i, c := range _c {
		_u = append(_u, c+" = "+_i[i])
	}

	// return an error if no update input is provided
	if len(_u) == 0 {
		return nil, errors.New("question.Update: no input provided")
	}

	// sql query
	sqlstr := `UPDATE "jack"."questions" SET ` +
		strings.Join(_u, ", ") + ` ` +
		_s + ` ` +
		`RETURNING "id", "order", "standup_id", "question"`

	// log query if we've provided a logger
	if pogo.Log != nil {
		pogo.Log(sqlstr, _v...)
	}

	// run the query
	var _question Question
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&_question.ID, &_question.Order, &_question.StandupID, &_question.Question); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrQuestionNotFound
		}
		return nil, e
	}

	return &_question, nil
}

// UpdateMany updates "questions" rows in "jack"."questions" by conditions, returning all results
func UpdateMany(db pogo.DB, question *Input, conds ...pogo.Condition) ([]*Question, error) {
	var questions []*Question

	_s, _v, err := pogo.Conditions(conds...)
	if err != nil {
		return nil, err
	}

	fields := question.columns()

	// prepare the slices
	_c, _i, _v2 := slice(fields, len(_v))
	_v = append(_v, _v2...)

	// setup the update fields
	var _u []string
	for i, c := range _c {
		_u = append(_u, c+" = "+_i[i])
	}

	// return an error if no update input is provided
	if len(_u) == 0 {
		return nil, errors.New("question.UpdateMany: no input provided")
	}

	// sql query
	sqlstr := `UPDATE "jack"."questions" SET ` +
		strings.Join(_u, ", ") + ` ` +
		_s + ` ` +
		`RETURNING "id", "order", "standup_id", "question"`

	// log query if we've provided a logger
	if pogo.Log != nil {
		pogo.Log(sqlstr, _v...)
	}

	rows, err := db.Query(sqlstr, _v...)
	if err != nil {
		return questions, err
	}
	defer rows.Close()

	for rows.Next() {
		var _question Question
		if e := rows.Scan(&_question.ID, &_question.Order, &_question.StandupID, &_question.Question); e != nil {
			if e == pgx.ErrNoRows {
				return questions, ErrQuestionNotFound
			}
			return questions, err
		}
		questions = append(questions, &_question)
	}
	if rows.Err() != nil {
		return questions, rows.Err()
	}

	return questions, nil
}

// UpdateByID a "question" in "jack"."questions" by its "id"
func UpdateByID(db pogo.DB, id int, question *Input) (*Question, error) {
	fields := question.columns()

	// don't update the primary key
	delete(fields, "id")

	// prepare the slices
	_c, _i, _v := slice(fields, 1)

	// setup the update fields
	var _u []string
	for i, c := range _c {
		_u = append(_u, c+" = "+_i[i])
	}

	// return an error if no update input is provided
	if len(_u) == 0 {
		return nil, errors.New("question.UpdateByID: no input provided")
	}

	// sql query
	sqlstr := `UPDATE "jack"."questions" ` +
		`SET ` + strings.Join(_u, ", ") + ` ` +
		`WHERE "id" = $1 ` +
		`RETURNING "id", "order", "standup_id", "question"`

	// setup query
	values := append([]interface{}{&id}, _v...)

	// log query if we've provided a logger
	if pogo.Log != nil {
		pogo.Log(sqlstr, values...)
	}

	// run the query
	var _question Question
	row := db.QueryRow(sqlstr, values...)
	if e := row.Scan(&_question.ID, &_question.Order, &_question.StandupID, &_question.Question); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrQuestionNotFound
		}
		return nil, e
	}

	return &_question, nil
}

// Delete `Question`s by some conditions. Returns the first result.
func Delete(db pogo.DB, conds ...pogo.Condition) (*Question, error) {
	_s, _v, err := pogo.Conditions(conds...)
	if err != nil {
		return nil, err
	}

	// sql delete query
	sqlstr := `DELETE FROM "jack"."questions" ` +
		_s + ` ` +
		`RETURNING "id", "order", "standup_id", "question"`

	// log query if we've provided a logger
	if pogo.Log != nil {
		pogo.Log(sqlstr, _v...)
	}

	var _question Question
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&_question.ID, &_question.Order, &_question.StandupID, &_question.Question); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrQuestionNotFound
		}
		return nil, e
	}

	return &_question, nil
}

// DeleteMany `Question`s by some conditions, returning all results.
func DeleteMany(db pogo.DB, conds ...pogo.Condition) ([]*Question, error) {
	var questions []*Question

	_s, _v, err := pogo.Conditions(conds...)
	if err != nil {
		return questions, err
	}

	// sql delete query
	sqlstr := `DELETE FROM "jack"."questions" ` +
		_s + ` ` +
		`RETURNING "id", "order", "standup_id", "question"`

	// log query if we've provided a logger
	if pogo.Log != nil {
		pogo.Log(sqlstr, _v...)
	}

	rows, err := db.Query(sqlstr, _v...)
	if err != nil {
		return questions, err
	}
	defer rows.Close()

	for rows.Next() {
		var _question Question
		if e := rows.Scan(&_question.ID, &_question.Order, &_question.StandupID, &_question.Question); e != nil {
			if e == pgx.ErrNoRows {
				return questions, ErrQuestionNotFound
			}
			return questions, err
		}
		questions = append(questions, &_question)
	}
	if rows.Err() != nil {
		return questions, rows.Err()
	}

	return questions, nil
}

// DeleteByID a "question" from the "jack"."questions" table
func DeleteByID(db pogo.DB, id int) (*Question, error) {
	// sql delete query
	sqlstr := `DELETE FROM "jack"."questions" ` +
		`WHERE "id" = $1 ` +
		`RETURNING "id", "order", "standup_id", "question"`

	// log query if we've provided a logger
	if pogo.Log != nil {
		pogo.Log(sqlstr, id)
	}

	// run the query
	var _question Question
	row := db.QueryRow(sqlstr, id)
	if e := row.Scan(&_question.ID, &_question.Order, &_question.StandupID, &_question.Question); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrQuestionNotFound
		}
		return nil, e
	}

	return &_question, nil
}

// UpsertByID inserts a `questions`, updating the row if `ID` already exists.
func UpsertByID(db pogo.DB, id int, question *Input) (*Question, error) {
	// add values to input, overriding existing keys if present in the input
	question = question.ID(id)

	fields := question.columns()

	// prepare the slices
	_c, _i, _v := slice(fields, 0)

	// setup the update fields
	var _u []string
	for _, c := range _c {
		_u = append(_u, c+" = EXCLUDED."+c)
	}

	// return an error if no update input is provided
	if len(_u) == 0 {
		return nil, errors.New("question.UpsertByID: no input provided")
	}

	// sql query
	sqlstr := `INSERT INTO "jack"."questions" (` + strings.Join(_c, ", ") + `) ` +
		`VALUES (` + strings.Join(_i, ", ") + `) ` +
		`ON CONFLICT (id) ` +
		`DO UPDATE SET ` + strings.Join(_u, ", ") + ` ` +
		`RETURNING "id", "order", "standup_id", "question"`

	// log query if we've provided a logger
	if pogo.Log != nil {
		pogo.Log(sqlstr, _v...)
	}

	// run query
	var _question Question
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&_question.ID, &_question.Order, &_question.StandupID, &_question.Question); e != nil && e != pgx.ErrNoRows {
		return nil, e
	}

	return &_question, nil
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