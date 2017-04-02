package postgres

import (
	"github.com/matthewmueller/pogo/db"
)

// Column represents column info.
type Column struct {
	FieldOrdinal int     // field_ordinal
	ColumnName   string  // column_name
	DataType     string  // data_type
	NotNull      bool    // not_null
	DefaultValue *string // default_value
	IsPrimaryKey bool    // is_primary_key
}

// Columns - get the columns of a table
func Columns(db db.DB, schema string, table string) (columns []*Column, err error) {

	// sql query
	// TODO: support onDelete and onUpdate
	const sqlstr = `SELECT ` +
		`a.attnum, ` + // ::integer AS field_ordinal
		`a.attname, ` + // ::varchar AS column_name
		`format_type(a.atttypid, a.atttypmod), ` + // ::varchar AS data_type
		`a.attnotnull, ` + // ::boolean AS not_null
		`pg_get_expr(ad.adbin, ad.adrelid), ` + // ::varchar AS default_value
		`COALESCE(ct.contype = 'p', false) ` + // ::boolean AS is_primary_key
		`FROM pg_attribute a ` +
		`JOIN ONLY pg_class c ON c.oid = a.attrelid ` +
		`JOIN ONLY pg_namespace n ON n.oid = c.relnamespace ` +
		`LEFT JOIN pg_constraint ct ON ct.conrelid = c.oid AND a.attnum = ANY(ct.conkey) AND ct.contype IN('p', 'u') ` +
		`LEFT JOIN pg_attrdef ad ON ad.adrelid = c.oid AND ad.adnum = a.attnum ` +
		`WHERE a.attisdropped = false AND n.nspname = $1 AND c.relname = $2 AND a.attnum > 0` +
		`ORDER BY a.attnum`

	// run query
	// XOLog(sqlstr, schema, table, sys)
	q, err := db.Query(sqlstr, schema, table)
	if err != nil {
		return columns, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		c := Column{}

		// scan
		err = q.Scan(&c.FieldOrdinal, &c.ColumnName, &c.DataType, &c.NotNull, &c.DefaultValue, &c.IsPrimaryKey)
		if err != nil {
			return columns, err
		}

		columns = append(columns, &c)
	}

	return columns, nil
}
