package orders

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx"
	"github.com/matthewmueller/gambit/pogo/enum"
	"github.com/matthewmueller/pogo/testgambit"
	uuid "github.com/satori/go.uuid"
)

// GENERATED BY POGO. DO NOT EDIT.

// ErrOrderNotFound returned if the order is not found
var ErrOrderNotFound = errors.New("order not found")

// columns in `"1"."orders"`
type columns struct {
	ID        *string           `json:"id,omitempty"`
	Timestamp *time.Time        `json:"timestamp,omitempty"`
	Exchange  *enum.Exchange    `json:"exchange,omitempty"`
	Currency  *enum.Currency    `json:"currency,omitempty"`
	Side      *enum.OrderSide   `json:"side,omitempty"`
	Status    *enum.OrderStatus `json:"status,omitempty"`
	Type      *enum.OrderType   `json:"type,omitempty"`
	Price     *int              `json:"price,omitempty"`
	Size      *int              `json:"size,omitempty"`
	CreatedAt *time.Time        `json:"created_at,omitempty"`
	UpdatedAt *time.Time        `json:"updated_at,omitempty"`
}

// Order fluent API
type Order struct {
	columns *columns
}

// New `"1"."orders"` API
func New() *Order {
	return &Order{&columns{}}
}

// ID sets the `id`
func (order *Order) ID(id uuid.UUID) *Order {
	order.columns.ID = testgambit.DecodeUUID(id)
	return order
}

// GetID returns the `id` if set
func (order *Order) GetID() (id *uuid.UUID) {
	return testgambit.EncodeUUID(order.columns.ID)
}

// Timestamp sets the `timestamp`
func (order *Order) Timestamp(timestamp time.Time) *Order {
	order.columns.Timestamp = &timestamp
	return order
}

// GetTimestamp returns the `timestamp` if set
func (order *Order) GetTimestamp() (timestamp *time.Time) {
	return order.columns.Timestamp
}

// Exchange sets the `exchange`
func (order *Order) Exchange(exchange enum.Exchange) *Order {
	order.columns.Exchange = &exchange
	return order
}

// GetExchange returns the `exchange` if set
func (order *Order) GetExchange() (exchange *enum.Exchange) {
	return order.columns.Exchange
}

// Currency sets the `currency`
func (order *Order) Currency(currency enum.Currency) *Order {
	order.columns.Currency = &currency
	return order
}

// GetCurrency returns the `currency` if set
func (order *Order) GetCurrency() (currency *enum.Currency) {
	return order.columns.Currency
}

// Side sets the `side`
func (order *Order) Side(side enum.OrderSide) *Order {
	order.columns.Side = &side
	return order
}

// GetSide returns the `side` if set
func (order *Order) GetSide() (side *enum.OrderSide) {
	return order.columns.Side
}

// Status sets the `status`
func (order *Order) Status(status enum.OrderStatus) *Order {
	order.columns.Status = &status
	return order
}

// GetStatus returns the `status` if set
func (order *Order) GetStatus() (status *enum.OrderStatus) {
	return order.columns.Status
}

// Type sets the `type`
func (order *Order) Type(kind enum.OrderType) *Order {
	order.columns.Type = &kind
	return order
}

// GetType returns the `type` if set
func (order *Order) GetType() (kind *enum.OrderType) {
	return order.columns.Type
}

// Price sets the `price`
func (order *Order) Price(price int) *Order {
	order.columns.Price = &price
	return order
}

// GetPrice returns the `price` if set
func (order *Order) GetPrice() (price *int) {
	return order.columns.Price
}

// Size sets the `size`
func (order *Order) Size(size int) *Order {
	order.columns.Size = &size
	return order
}

// GetSize returns the `size` if set
func (order *Order) GetSize() (size *int) {
	return order.columns.Size
}

// CreatedAt sets the `created_at`
func (order *Order) CreatedAt(createdAt time.Time) *Order {
	order.columns.CreatedAt = &createdAt
	return order
}

// GetCreatedAt returns the `created_at` if set
func (order *Order) GetCreatedAt() (createdAt *time.Time) {
	return order.columns.CreatedAt
}

// UpdatedAt sets the `updated_at`
func (order *Order) UpdatedAt(updatedAt time.Time) *Order {
	order.columns.UpdatedAt = &updatedAt
	return order
}

// GetUpdatedAt returns the `updated_at` if set
func (order *Order) GetUpdatedAt() (updatedAt *time.Time) {
	return order.columns.UpdatedAt
}

// MarshalJSON marshals the `order` into JSON
func (order *Order) MarshalJSON() ([]byte, error) {
	return json.Marshal(order.columns)
}

// UnmarshalJSON unmarshals json to a `order`
func (order *Order) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, order.columns)
}

func (order *Order) String() string {
	return "order"
}

// get all the non-nil columns
func getColumns(order *Order) map[string]interface{} {
	columns := make(map[string]interface{})

	if order.columns.ID != nil {
		columns["id"] = *order.columns.ID
	}
	if order.columns.Timestamp != nil {
		columns["timestamp"] = *order.columns.Timestamp
	}
	if order.columns.Exchange != nil {
		columns["exchange"] = *order.columns.Exchange
	}
	if order.columns.Currency != nil {
		columns["currency"] = *order.columns.Currency
	}
	if order.columns.Side != nil {
		columns["side"] = *order.columns.Side
	}
	if order.columns.Status != nil {
		columns["status"] = *order.columns.Status
	}
	if order.columns.Type != nil {
		columns["type"] = *order.columns.Type
	}
	if order.columns.Price != nil {
		columns["price"] = *order.columns.Price
	}
	if order.columns.Size != nil {
		columns["size"] = *order.columns.Size
	}
	if order.columns.CreatedAt != nil {
		columns["created_at"] = *order.columns.CreatedAt
	}
	if order.columns.UpdatedAt != nil {
		columns["updated_at"] = *order.columns.UpdatedAt
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

// Find a order by "id"
func Find(db testgambit.DB, id uuid.UUID) (*Order, error) {
	_id := testgambit.DecodeUUID(id)

	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "timestamp", "exchange", "currency", "side", "status", "type", "price", "size", "created_at", "updated_at"
	FROM "1"."orders"
	WHERE "id" = $1
	`
	testgambit.Log(sqlstr, _id)

	cols := &columns{}
	row := db.QueryRow(sqlstr, _id)
	if e := row.Scan(&cols.ID, &cols.Timestamp, &cols.Exchange, &cols.Currency, &cols.Side, &cols.Status, &cols.Type, &cols.Price, &cols.Size, &cols.CreatedAt, &cols.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrOrderNotFound
		}
		return nil, e
	}

	return &Order{cols}, nil
}

// FindMany find many `order`s by a given condition
func FindMany(db testgambit.DB, where *WhereClause) ([]*Order, error) {
	var _o []*Order

	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "timestamp", "exchange", "currency", "side", "status", "type", "price", "size", "created_at", "updated_at"
	FROM "1"."orders"
	WHERE ` + where.condition
	testgambit.Log(sqlstr, where.params...)

	rows, err := db.Query(sqlstr, where.params...)
	if err != nil {
		return _o, err
	}
	defer rows.Close()

	for rows.Next() {
		cols := &columns{}
		if e := rows.Scan(&cols.ID, &cols.Timestamp, &cols.Exchange, &cols.Currency, &cols.Side, &cols.Status, &cols.Type, &cols.Price, &cols.Size, &cols.CreatedAt, &cols.UpdatedAt); e != nil {
			if e == pgx.ErrNoRows {
				return _o, ErrOrderNotFound
			}
			return _o, err
		}
		_o = append(_o, &Order{cols})
	}
	if rows.Err() != nil {
		return _o, rows.Err()
	}

	// ensure we return an empty array
	// rather than nil when we marshal
	if len(_o) == 0 {
		return make([]*Order, 0), nil
	}

	return _o, nil
}

// FindOne find one order by a condition
func FindOne(db testgambit.DB, where *WhereClause) (*Order, error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "timestamp", "exchange", "currency", "side", "status", "type", "price", "size", "created_at", "updated_at"
	FROM "1"."orders"
	WHERE ` + where.condition
	testgambit.Log(sqlstr, where.params...)

	cols := &columns{}
	row := db.QueryRow(sqlstr, where.params...)
	if e := row.Scan(&cols.ID, &cols.Timestamp, &cols.Exchange, &cols.Currency, &cols.Side, &cols.Status, &cols.Type, &cols.Price, &cols.Size, &cols.CreatedAt, &cols.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrOrderNotFound
		}
		return nil, e
	}

	return &Order{cols}, nil
}

// Insert a `order` into the `"1"."orders"` table.
func Insert(db testgambit.DB, order *Order) (*Order, error) {
	// get all the non-nil columns and prepare them for the query
	_c, _i, _v := testgambit.Slice(getColumns(order), 0)

	// sql insert query, primary key provided by sequence
	sqlstr := `
	INSERT INTO "1"."orders" (` + strings.Join(_c, ", ") + `)
	VALUES (` + strings.Join(_i, ", ") + `)
	RETURNING "id", "timestamp", "exchange", "currency", "side", "status", "type", "price", "size", "created_at", "updated_at"
	`
	testgambit.Log(sqlstr, _v...)

	cols := &columns{}
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&cols.ID, &cols.Timestamp, &cols.Exchange, &cols.Currency, &cols.Side, &cols.Status, &cols.Type, &cols.Price, &cols.Size, &cols.CreatedAt, &cols.UpdatedAt); e != nil {
		return nil, e
	}

	return &Order{cols}, nil
}

// Update a order by its `id`
func Update(db testgambit.DB, id uuid.UUID, order *Order) (*Order, error) {
	_id := testgambit.DecodeUUID(id)
	fields := getColumns(order)

	// don't update the primary key
	delete(fields, "id")

	// prepare the slices
	_c, _i, _v := testgambit.Slice(fields, 1)

	// sql query
	sqlstr := `UPDATE "1"."orders" SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `)
		WHERE "id" = $1
		RETURNING "id", "timestamp", "exchange", "currency", "side", "status", "type", "price", "size", "created_at", "updated_at"`

	// setup query
	values := append([]interface{}{_id}, _v...)
	testgambit.Log(sqlstr, values...)

	// run the query
	cols := &columns{}
	row := db.QueryRow(sqlstr, values...)
	if e := row.Scan(&cols.ID, &cols.Timestamp, &cols.Exchange, &cols.Currency, &cols.Side, &cols.Status, &cols.Type, &cols.Price, &cols.Size, &cols.CreatedAt, &cols.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrOrderNotFound
		}
		return nil, e
	}

	return &Order{cols}, nil
}

// UpdateMany rows in `"1"."orders"` by a given condition
func UpdateMany(db testgambit.DB, where *WhereClause, order *Order) ([]*Order, error) {
	var _o []*Order

	// prepare the slices
	_c, _i, _v := testgambit.Slice(getColumns(order), len(where.params))

	// sql query
	sqlstr := `UPDATE "1"."orders" SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `) ` +
		`WHERE ` + where.condition + ` ` +
		`RETURNING "id", "timestamp", "exchange", "currency", "side", "status", "type", "price", "size", "created_at", "updated_at"`

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
		if e := rows.Scan(&cols.ID, &cols.Timestamp, &cols.Exchange, &cols.Currency, &cols.Side, &cols.Status, &cols.Type, &cols.Price, &cols.Size, &cols.CreatedAt, &cols.UpdatedAt); e != nil {
			if e == pgx.ErrNoRows {
				return _o, ErrOrderNotFound
			}
			return _o, err
		}
		_o = append(_o, &Order{cols})
	}
	if rows.Err() != nil {
		return _o, rows.Err()
	}

	// ensure we return an empty array
	// rather than nil when we marshal
	if len(_o) == 0 {
		return make([]*Order, 0), nil
	}

	return _o, nil
}

// Delete a `order` from the `"1"."orders"` table
func Delete(db testgambit.DB, id uuid.UUID) error {
	_id := testgambit.DecodeUUID(id)

	// sql query
	sqlstr := `DELETE FROM "1"."orders" WHERE "id" = $1`
	testgambit.Log(sqlstr, _id)

	// run query
	if _, e := db.Exec(sqlstr, _id); e != nil {
		if e == pgx.ErrNoRows {
			return ErrOrderNotFound
		}
		return e
	}

	return nil
}

// DeleteMany delete many `order`'s by the given condition
func DeleteMany(db testgambit.DB, where *WhereClause) error {
	// sql select query, primary key provided by sequence
	sqlstr := `DELETE FROM "1"."orders" WHERE ` + where.condition
	testgambit.Log(sqlstr, where.params...)

	if _, e := db.Exec(sqlstr, where.params...); e != nil {
		return e
	}

	return nil
}

// Upsert the `order` by its `id`.
func Upsert(db testgambit.DB, order *Order) (*Order, error) {
	// get all the non-nil columns and prepare them for the query
	_c, _i, _v := testgambit.Slice(getColumns(order), 0)

	// sql query
	sqlstr := `INSERT INTO "1"."orders" (` + strings.Join(_c, ", ") + `) ` +
		`VALUES (` + strings.Join(_i, ", ") + `) ` +
		`ON CONFLICT ("id") ` +
		`DO UPDATE SET (` + strings.Join(_c, ", ") + `) = ( EXCLUDED.` + strings.Join(_c, ", EXCLUDED.") + `) ` +
		`RETURNING "id", "timestamp", "exchange", "currency", "side", "status", "type", "price", "size", "created_at", "updated_at"`
	testgambit.Log(sqlstr, _v...)

	// run query
	cols := &columns{}
	row := db.QueryRow(sqlstr, _v...)
	if e := row.Scan(&cols.ID, &cols.Timestamp, &cols.Exchange, &cols.Currency, &cols.Side, &cols.Status, &cols.Type, &cols.Price, &cols.Size, &cols.CreatedAt, &cols.UpdatedAt); e != nil {
		return nil, e
	}

	return &Order{cols}, nil
}
