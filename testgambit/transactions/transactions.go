package transactions

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/testgambit"
	"github.com/matthewmueller/pogo/testgambit/enum"
	uuid "github.com/satori/go.uuid"
)

// GENERATED BY POGO. DO NOT EDIT.

// ErrTransactionNotFound returned if the transaction is not found
var ErrTransactionNotFound = errors.New("transaction not found")

// columns in `"1"."transactions"`
type columns struct {
	ID        *string        `json:"id,omitempty"`
	Timestamp *time.Time     `json:"timestamp,omitempty"`
	Exchange  *enum.Exchange `json:"exchange,omitempty"`
	Currency  *enum.Currency `json:"currency,omitempty"`
	High      *int           `json:"high,omitempty"`
	Low       *int           `json:"low,omitempty"`
	Open      *int           `json:"open,omitempty"`
	Close     *int           `json:"close,omitempty"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
}

// Transaction fluent API
type Transaction struct {
	columns *columns
}

// New `"1"."transactions"` API
func New() *Transaction {
	return &Transaction{&columns{}}
}

// ID sets the `id`
func (transaction *Transaction) ID(id uuid.UUID) *Transaction {
	transaction.columns.ID = testgambit.DecodeUUID(id)
	return transaction
}

// GetID returns the `id` if set
func (transaction *Transaction) GetID() (id *uuid.UUID) {
	return testgambit.EncodeUUID(transaction.columns.ID)
}

// Timestamp sets the `timestamp`
func (transaction *Transaction) Timestamp(timestamp time.Time) *Transaction {
	transaction.columns.Timestamp = &timestamp
	return transaction
}

// GetTimestamp returns the `timestamp` if set
func (transaction *Transaction) GetTimestamp() (timestamp *time.Time) {
	return transaction.columns.Timestamp
}

// Exchange sets the `exchange`
func (transaction *Transaction) Exchange(exchange enum.Exchange) *Transaction {
	transaction.columns.Exchange = &exchange
	return transaction
}

// GetExchange returns the `exchange` if set
func (transaction *Transaction) GetExchange() (exchange *enum.Exchange) {
	return transaction.columns.Exchange
}

// Currency sets the `currency`
func (transaction *Transaction) Currency(currency enum.Currency) *Transaction {
	transaction.columns.Currency = &currency
	return transaction
}

// GetCurrency returns the `currency` if set
func (transaction *Transaction) GetCurrency() (currency *enum.Currency) {
	return transaction.columns.Currency
}

// High sets the `high`
func (transaction *Transaction) High(high int) *Transaction {
	transaction.columns.High = &high
	return transaction
}

// GetHigh returns the `high` if set
func (transaction *Transaction) GetHigh() (high *int) {
	return transaction.columns.High
}

// Low sets the `low`
func (transaction *Transaction) Low(low int) *Transaction {
	transaction.columns.Low = &low
	return transaction
}

// GetLow returns the `low` if set
func (transaction *Transaction) GetLow() (low *int) {
	return transaction.columns.Low
}

// Open sets the `open`
func (transaction *Transaction) Open(open int) *Transaction {
	transaction.columns.Open = &open
	return transaction
}

// GetOpen returns the `open` if set
func (transaction *Transaction) GetOpen() (open *int) {
	return transaction.columns.Open
}

// Close sets the `close`
func (transaction *Transaction) Close(cls int) *Transaction {
	transaction.columns.Close = &cls
	return transaction
}

// GetClose returns the `close` if set
func (transaction *Transaction) GetClose() (cls *int) {
	return transaction.columns.Close
}

// CreatedAt sets the `created_at`
func (transaction *Transaction) CreatedAt(createdAt time.Time) *Transaction {
	transaction.columns.CreatedAt = &createdAt
	return transaction
}

// GetCreatedAt returns the `created_at` if set
func (transaction *Transaction) GetCreatedAt() (createdAt *time.Time) {
	return transaction.columns.CreatedAt
}

// MarshalJSON marshals the `transaction` into JSON
func (transaction *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(transaction.columns)
}

// UnmarshalJSON unmarshals json to a `transaction`
func (transaction *Transaction) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, transaction.columns)
}

func (transaction *Transaction) String() string {
	return "transaction TODO"
}

// get all the non-nil columns
func getColumns(transaction *Transaction) map[string]interface{} {
	columns := make(map[string]interface{})

	if transaction.columns.ID != nil {
		columns["id"] = *transaction.columns.ID
	}
	if transaction.columns.Timestamp != nil {
		columns["timestamp"] = *transaction.columns.Timestamp
	}
	if transaction.columns.Exchange != nil {
		columns["exchange"] = *transaction.columns.Exchange
	}
	if transaction.columns.Currency != nil {
		columns["currency"] = *transaction.columns.Currency
	}
	if transaction.columns.High != nil {
		columns["high"] = *transaction.columns.High
	}
	if transaction.columns.Low != nil {
		columns["low"] = *transaction.columns.Low
	}
	if transaction.columns.Open != nil {
		columns["open"] = *transaction.columns.Open
	}
	if transaction.columns.Close != nil {
		columns["close"] = *transaction.columns.Close
	}
	if transaction.columns.CreatedAt != nil {
		columns["created_at"] = *transaction.columns.CreatedAt
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

// Find a transaction by "id"
func Find(db testgambit.DB, id uuid.UUID) (*Transaction, error) {
	_id := testgambit.DecodeUUID(id)

	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "timestamp", "exchange", "currency", "high", "low", "open", "close", "created_at"
	FROM "1"."transactions"
	WHERE "id" = $1
	`
	testgambit.Log(sqlstr, _id)

	cols := &columns{}
	row := db.QueryRow(sqlstr, _id)
	if e := row.Scan(&cols.ID, &cols.Timestamp, &cols.Exchange, &cols.Currency, &cols.High, &cols.Low, &cols.Open, &cols.Close, &cols.CreatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrTransactionNotFound
		}
		return nil, e
	}

	return &Transaction{cols}, nil
}

// FindMany find many `transaction`s by a given condition
func FindMany(db testgambit.DB, where *WhereClause) ([]*Transaction, error) {
	var _o []*Transaction

	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "timestamp", "exchange", "currency", "high", "low", "open", "close", "created_at"
	FROM "1"."transactions"
	WHERE ` + where.condition
	testgambit.Log(sqlstr, where.params...)

	rows, err := db.Query(sqlstr, where.params...)
	if err != nil {
		return _o, err
	}
	defer rows.Close()

	for rows.Next() {
		cols := &columns{}
		if e := rows.Scan(&cols.ID, &cols.Timestamp, &cols.Exchange, &cols.Currency, &cols.High, &cols.Low, &cols.Open, &cols.Close, &cols.CreatedAt); e != nil {
			if e == pgx.ErrNoRows {
				return _o, ErrTransactionNotFound
			}
			return _o, err
		}
		_o = append(_o, &Transaction{cols})
	}
	if rows.Err() != nil {
		return _o, rows.Err()
	}

	// ensure we return an empty array
	// rather than nil when we marshal
	if len(_o) == 0 {
		return make([]*Transaction, 0), nil
	}

	return _o, nil
}

// FindOne find one transaction by a condition
func FindOne(db testgambit.DB, where *WhereClause) (*Transaction, error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "timestamp", "exchange", "currency", "high", "low", "open", "close", "created_at"
	FROM "1"."transactions"
	WHERE ` + where.condition
	testgambit.Log(sqlstr, where.params...)

	cols := &columns{}
	row := db.QueryRow(sqlstr, where.params...)
	if e := row.Scan(&cols.ID, &cols.Timestamp, &cols.Exchange, &cols.Currency, &cols.High, &cols.Low, &cols.Open, &cols.Close, &cols.CreatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrTransactionNotFound
		}
		return nil, e
	}

	return &Transaction{cols}, nil
}

// Insert a `transaction` into the `"1"."transactions"` table.
func Insert(db testgambit.DB, transaction *Transaction) (*Transaction, error) {
	// get all the non-nil columns and prepare them for the query
	_c, _i, _v := testgambit.Slice(getColumns(transaction), 0)

	// sql insert query, primary key provided by sequence
	sqlstr := `
	INSERT INTO "1"."transactions" (` + strings.Join(_c, ", ") + `)
	VALUES (` + strings.Join(_i, ", ") + `)
	RETURNING "id", "timestamp", "exchange", "currency", "high", "low", "open", "close", "created_at"
	`
	testgambit.Log(sqlstr, _v...)

	cols := &columns{}
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&cols.ID, &cols.Timestamp, &cols.Exchange, &cols.Currency, &cols.High, &cols.Low, &cols.Open, &cols.Close, &cols.CreatedAt); e != nil {
		return nil, e
	}

	return &Transaction{cols}, nil
}

// Update a transaction by its `id`
func Update(db testgambit.DB, id uuid.UUID, transaction *Transaction) (*Transaction, error) {
	_id := testgambit.DecodeUUID(id)
	fields := getColumns(transaction)

	// don't update the primary key
	delete(fields, "id")

	// prepare the slices
	_c, _i, _v := testgambit.Slice(fields, 1)

	// sql query
	sqlstr := `UPDATE "1"."transactions" SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `)
		WHERE "id" = $1
		RETURNING "id", "timestamp", "exchange", "currency", "high", "low", "open", "close", "created_at"`

	// setup query
	values := append([]interface{}{_id}, _v...)
	testgambit.Log(sqlstr, values...)

	// run the query
	cols := &columns{}
	row := db.QueryRow(sqlstr, values...)
	if e := row.Scan(&cols.ID, &cols.Timestamp, &cols.Exchange, &cols.Currency, &cols.High, &cols.Low, &cols.Open, &cols.Close, &cols.CreatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrTransactionNotFound
		}
		return nil, e
	}

	return &Transaction{cols}, nil
}

// UpdateMany rows in `"1"."transactions"` by a given condition
func UpdateMany(db testgambit.DB, where *WhereClause, transaction *Transaction) ([]*Transaction, error) {
	var _o []*Transaction

	// prepare the slices
	_c, _i, _v := testgambit.Slice(getColumns(transaction), len(where.params))

	// sql query
	sqlstr := `UPDATE "1"."transactions" SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `) ` +
		`WHERE ` + where.condition + ` ` +
		`RETURNING "id", "timestamp", "exchange", "currency", "high", "low", "open", "close", "created_at"`

		// setup the query
	values := []interface{}{}
	values = append(values, where.params...)
	values = append(values, _v...)
	testgambit.Log(sqlstr, values...)

	// run query
	rows, err := db.Query(sqlstr, values...)
	if err != nil {
		return _o, err
	}
	defer rows.Close()

	for rows.Next() {
		cols := &columns{}
		if e := rows.Scan(&cols.ID, &cols.Timestamp, &cols.Exchange, &cols.Currency, &cols.High, &cols.Low, &cols.Open, &cols.Close, &cols.CreatedAt); e != nil {
			if e == pgx.ErrNoRows {
				return _o, ErrTransactionNotFound
			}
			return _o, err
		}
		_o = append(_o, &Transaction{cols})
	}
	if rows.Err() != nil {
		return _o, rows.Err()
	}

	// ensure we return an empty array
	// rather than nil when we marshal
	if len(_o) == 0 {
		return make([]*Transaction, 0), nil
	}

	return _o, nil
}

// Delete a `transaction` from the `"1"."transactions"` table
func Delete(db testgambit.DB, id uuid.UUID) error {
	_id := testgambit.DecodeUUID(id)

	// sql query
	sqlstr := `DELETE FROM "1"."transactions" WHERE "id" = $1`
	testgambit.Log(sqlstr, _id)

	// run query
	if _, e := db.Exec(sqlstr, _id); e != nil {
		if e == pgx.ErrNoRows {
			return ErrTransactionNotFound
		}
		return e
	}

	return nil
}

// DeleteMany delete many `transaction`'s by the given condition
func DeleteMany(db testgambit.DB, where *WhereClause) error {
	// sql select query, primary key provided by sequence
	sqlstr := `DELETE FROM "1"."transactions" WHERE ` + where.condition
	testgambit.Log(sqlstr, where.params...)

	if _, e := db.Exec(sqlstr, where.params...); e != nil {
		return e
	}

	return nil
}

// Upsert the `transaction` by its `id`.
func Upsert(db testgambit.DB, transaction *Transaction) (*Transaction, error) {
	// get all the non-nil columns and prepare them for the query
	_c, _i, _v := testgambit.Slice(getColumns(transaction), 0)

	// sql query
	sqlstr := `INSERT INTO "1"."transactions" (` + strings.Join(_c, ", ") + `) ` +
		`VALUES (` + strings.Join(_i, ", ") + `) ` +
		`ON CONFLICT ("id") ` +
		`DO UPDATE SET (` + strings.Join(_c, ", ") + `) = ( EXCLUDED.` + strings.Join(_c, ", EXCLUDED.") + `) ` +
		`RETURNING "id", "timestamp", "exchange", "currency", "high", "low", "open", "close", "created_at"`
	testgambit.Log(sqlstr, _v...)

	// run query
	cols := &columns{}
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&cols.ID, &cols.Timestamp, &cols.Exchange, &cols.Currency, &cols.High, &cols.Low, &cols.Open, &cols.Close, &cols.CreatedAt); e != nil {
		return nil, e
	}

	return &Transaction{cols}, nil
}
