package postgres

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/database"
	"github.com/pkg/errors"
)

// Introspect a postgres database
func Introspect(db *pgx.Conn, schema string) (*database.Schema, error) {
	tables, err := getTables(db, schema)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get tables from schema")
	}

	for _, table := range tables {
		// get the columns
		columns, err := getColumns(db, schema, table.Name)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get columns for '%s' from schema", table.Name)
		}
		table.Columns = columns

		// get the foreign keys
		fks, err := getForeignKeys(db, schema, table.Name)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get the foreign keys for '%s' from schema", table.Name)
		}
		table.ForeignKeys = fks

		// get the indexes
		indexes, err := getIndexes(db, schema, table.Name)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get the indexes for '%s' from schema", table.Name)
		}

		// get each of the index columns
		for _, index := range indexes {
			// get the index columns
			icols, err := getIndexColumns(db, schema, table.Name, index.Name)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to get index columns for %s", index.Name)
			}
			index.Columns = icols
		}

		table.Indexes = indexes
	}

	enums, err := getEnums(db, schema)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get the enums")
	}

	return &database.Schema{
		Name:   schema,
		Tables: tables,
		Enums:  enums,
	}, nil
}

func getTables(db *pgx.Conn, schema string) (tables []*database.Table, err error) {
	// sql query
	const sqlstr = `
	SELECT c.relkind, c.relname, false
	FROM pg_class c
	JOIN ONLY pg_namespace n ON n.oid = c.relnamespace
	WHERE n.nspname = $1 AND c.relkind = $2
	ORDER BY c.relname
`

	// run query
	// DBLog(sqlstr, schema, relkind)
	// "r" constant is for tables
	q, err := db.Query(sqlstr, schema, "r")
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		t := database.Table{}

		// scan
		err = q.Scan(&t.Type, &t.Name, &t.ManualPk)
		if err != nil {
			return nil, err
		}

		tables = append(tables, &t)
	}
	if e := q.Err(); e != nil {
		return nil, e
	}

	return tables, nil
}

func getColumns(db *pgx.Conn, schema string, table string) (columns []*database.Column, err error) {
	// sql query
	// TODO: support onDelete and onUpdate
	const sqlstr = `
	SELECT
	a.attnum,
	a.attname,
	format_type(a.atttypid, a.atttypmod),
	a.attnotnull,
	d.description,
	pg_get_expr(ad.adbin, ad.adrelid),
	COALESCE(ct.contype = 'p', false)
	FROM pg_attribute a
	JOIN ONLY pg_class c ON c.oid = a.attrelid
	JOIN ONLY pg_namespace n ON n.oid = c.relnamespace
	LEFT JOIN pg_constraint ct ON ct.conrelid = c.oid AND a.attnum = ANY(ct.conkey) AND ct.contype IN('p', 'u')
	LEFT JOIN pg_attrdef ad ON ad.adrelid = c.oid AND ad.adnum = a.attnum
	LEFT JOIN pg_description d ON d.objoid = a.attrelid AND d.objsubid = a.attnum
	WHERE a.attisdropped = false AND n.nspname = $1 AND c.relname = $2 AND a.attnum > 0
	ORDER BY a.attnum
	`

	// run query
	// DBLog(sqlstr, schema, table, sys)
	q, err := db.Query(sqlstr, schema, table)
	if err != nil {
		return columns, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		c := database.Column{}

		// scan
		err = q.Scan(&c.FieldOrdinal, &c.Name, &c.DataType, &c.NotNull, &c.Comment, &c.DefaultValue, &c.IsPrimaryKey)
		if err != nil {
			return columns, err
		}

		// coerce column type into
		// something we understand

		columns = append(columns, &c)
	}
	if e := q.Err(); e != nil {
		return nil, e
	}

	return columns, nil
}

func getForeignKeys(db *pgx.Conn, schema string, table string) (fks []*database.ForeignKey, err error) {
	// sql query

	const sqlstr = `
    SELECT r.conname, b.attname, i.relname, c.relname, d.attname, 0, 0, '', '', ''
    FROM pg_constraint r
    JOIN ONLY pg_class a ON a.oid = r.conrelid
    JOIN ONLY pg_attribute b ON b.attisdropped = false AND b.attnum = ANY(r.conkey) AND b.attrelid = r.conrelid
    JOIN ONLY pg_class i on i.oid = r.conindid
    JOIN ONLY pg_class c on c.oid = r.confrelid
    JOIN ONLY pg_attribute d ON d.attisdropped = false AND d.attnum = ANY(r.confkey) AND d.attrelid = r.confrelid
    JOIN ONLY pg_namespace n ON n.oid = r.connamespace
    WHERE r.contype = 'f' AND n.nspname = $1 AND a.relname = $2
    ORDER BY r.conname, b.attname
  `

	// run query
	// DBLog(sqlstr, schema, table)
	q, err := db.Query(sqlstr, schema, table)
	if err != nil {
		return fks, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		fk := database.ForeignKey{}

		// scan
		err = q.Scan(&fk.ForeignKeyName, &fk.Name, &fk.RefIndexName, &fk.RefTableName, &fk.RefColumnName, &fk.KeyID, &fk.SeqNo, &fk.OnUpdate, &fk.OnDelete, &fk.Match)
		if err != nil {
			return fks, err
		}

		fks = append(fks, &fk)
	}
	if e := q.Err(); e != nil {
		return nil, e
	}

	return fks, nil
}

func getIndexes(db *pgx.Conn, schema string, table string) (indexes []*database.Index, err error) {
	// sql query
	const sqlstr = `SELECT ` +
		`DISTINCT ic.relname, ` + // ::varchar AS index_name
		`i.indisunique, ` + // ::boolean AS is_unique
		`i.indisprimary, ` + // ::boolean AS is_primary
		`0, ` + // ::integer AS seq_no
		`'', ` + // ::varchar AS origin
		`false ` + // ::boolean AS is_partial
		`FROM pg_index i ` +
		`JOIN ONLY pg_class c ON c.oid = i.indrelid ` +
		`JOIN ONLY pg_namespace n ON n.oid = c.relnamespace ` +
		`JOIN ONLY pg_class ic ON ic.oid = i.indexrelid ` +
		`WHERE i.indkey <> '0' AND n.nspname = $1 AND c.relname = $2`

	// run query
	// DBLog(sqlstr, schema, table)
	q, err := db.Query(sqlstr, schema, table)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		i := database.Index{}

		// scan
		err = q.Scan(&i.Name, &i.IsUnique, &i.IsPrimary, &i.SeqNo, &i.Origin, &i.IsPartial)
		if err != nil {
			return nil, err
		}

		indexes = append(indexes, &i)
	}
	if e := q.Err(); e != nil {
		return nil, e
	}

	return indexes, nil
}

// get the column indexes
func getIndexColumns(db *pgx.Conn, schema string, table string, index string) ([]*database.IndexColumn, error) {
	allcols, err := getColumns(db, schema, table)
	if err != nil {
		return nil, err
	}

	// load columns
	cols, err := indexColumns(db, schema, index)
	if err != nil {
		return nil, err
	}

	// load col order
	colOrd, err := colOrder(db, schema, index)
	if err != nil {
		return nil, err
	}

	// build schema name used in errors
	s := schema
	if s != "" {
		s = s + "."
	}

	// put cols in order using colOrder
	ret := []*database.IndexColumn{}
	for _, v := range strings.Split(colOrd.Ord, " ") {
		cid, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("could not convert %s%s index %s column %s to int", s, table, index, v)
		}

		// find column
		found := false
		var c *database.IndexColumn
		for _, ic := range cols {
			if cid == ic.Cid {
				found = true
				c = ic
				break
			}
		}

		// sanity check
		if !found {
			return nil, fmt.Errorf("could not find %s%s index %s column id %d", s, table, index, cid)
		}

		ret = append(ret, c)
	}

	for _, col := range ret {
		for _, allcol := range allcols {
			if col.Name == allcol.Name {
				col.DataType = allcol.DataType
			}
		}
	}

	return ret, nil
}

// indexColumns runs a custom query, returning results as IndexColumn.
func indexColumns(db *pgx.Conn, schema string, index string) ([]*database.IndexColumn, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`(row_number() over()), ` + // ::integer AS seq_no
		`a.attnum, ` + // ::integer AS cid
		`a.attname ` + // ::varchar AS column_name
		`FROM pg_index i ` +
		`JOIN ONLY pg_class c ON c.oid = i.indrelid ` +
		`JOIN ONLY pg_namespace n ON n.oid = c.relnamespace ` +
		`JOIN ONLY pg_class ic ON ic.oid = i.indexrelid ` +
		`LEFT JOIN pg_attribute a ON i.indrelid = a.attrelid AND a.attnum = ANY(i.indkey) AND a.attisdropped = false ` +
		`WHERE i.indkey <> '0' AND n.nspname = $1 AND ic.relname = $2`

	// run query
	// DBLog(sqlstr, schema, index)
	q, err := db.Query(sqlstr, schema, index)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*database.IndexColumn{}
	for q.Next() {
		ic := database.IndexColumn{}

		// scan
		err = q.Scan(&ic.SeqNo, &ic.Cid, &ic.Name)
		if err != nil {
			return nil, err
		}

		res = append(res, &ic)
	}
	if e := q.Err(); e != nil {
		return nil, e
	}

	return res, nil
}

// PgColOrder represents index column order.
type columnOrder struct {
	Ord string // ord
}

// colOrder runs a custom query, returning results as columnOrder.
func colOrder(db *pgx.Conn, schema string, index string) (*columnOrder, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`i.indkey ` + // ::varchar AS ord
		`FROM pg_index i ` +
		`JOIN ONLY pg_class c ON c.oid = i.indrelid ` +
		`JOIN ONLY pg_namespace n ON n.oid = c.relnamespace ` +
		`JOIN ONLY pg_class ic ON ic.oid = i.indexrelid ` +
		`WHERE n.nspname = $1 AND ic.relname = $2`

	// run query
	// DBLog(sqlstr, schema, index)
	var pco columnOrder
	err = db.QueryRow(sqlstr, schema, index).Scan(&pco.Ord)
	if err != nil {
		return nil, err
	}

	return &pco, nil
}

// getTable function
func getEnums(db *pgx.Conn, schema string) (enums []*database.Enum, err error) {
	// sql query
	const sqlstr = `
    SELECT DISTINCT
		t.typname
		FROM pg_type t
		JOIN ONLY pg_namespace n ON n.oid = t.typnamespace
		JOIN ONLY pg_enum e ON t.oid = e.enumtypid
		WHERE n.nspname = $1
  `

	// run query
	// DBLog(sqlstr, schema)
	q, err := db.Query(sqlstr, schema)
	if err != nil {
		return nil, err
	}

	// load results
	for q.Next() {
		e := database.Enum{}

		// scan
		err = q.Scan(&e.Name)
		if err != nil {
			return nil, err
		}

		enums = append(enums, &e)
	}
	if e := q.Err(); e != nil {
		return nil, e
	}
	q.Close()

	for _, re := range enums {
		values, err := getEnumValues(db, schema, re.Name)
		if err != nil {
			return nil, err
		}
		re.Values = values
	}

	return enums, nil
}

// Values runs a custom query, returning results as Value.
func getEnumValues(db *pgx.Conn, schema string, enum string) ([]*database.EnumValue, error) {
	var err error

	// sql query
	const sqlstr = `
    SELECT
		e.enumlabel,
		e.enumsortorder::int
		FROM pg_type t
		JOIN ONLY pg_namespace n ON n.oid = t.typnamespace
		LEFT JOIN pg_enum e ON t.oid = e.enumtypid
		WHERE n.nspname = $1 AND t.typname = $2
  `

	// run query
	// DBLog(sqlstr, schema, enum)
	q, err := db.Query(sqlstr, schema, enum)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*database.EnumValue{}
	for q.Next() {
		ev := database.EnumValue{}

		// scan
		err = q.Scan(&ev.Label, &ev.Order)
		if err != nil {
			return nil, err
		}

		res = append(res, &ev)
	}
	if e := q.Err(); e != nil {
		return nil, e
	}

	return res, nil
}
