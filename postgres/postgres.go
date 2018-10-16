package postgres

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/db"
	"github.com/pkg/errors"
)

// Introspect a postgres database
func Introspect(conn *pgx.Conn, schema string) (*db.Schema, error) {
	tables, err := getTables(conn, schema)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get tables from schema")
	}

	for _, table := range tables {
		// get the columns
		columns, err := getColumns(conn, schema, table.Name)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get columns for '%s' from schema", table.Name)
		}
		table.Columns = columns

		// get the foreign keys
		fks, err := getForeignKeys(conn, schema, table.Name)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get the foreign keys for '%s' from schema", table.Name)
		}
		table.ForeignKeys = fks

		// get the indexes
		indexes, err := getIndexes(conn, schema, table.Name)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get the indexes for '%s' from schema", table.Name)
		}

		// get each of the index columns
		for _, index := range indexes {
			// get the index columns
			icols, err := getIndexColumns(conn, schema, table.Name, index.Name)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to get index columns for %s", index.Name)
			}
			index.Columns = icols
		}

		table.Indexes = indexes
	}

	procedures, err := getProcedures(conn, schema)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get the stored procedures")
	}

	enums, err := getEnums(conn, schema)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get the enums")
	}

	return &db.Schema{
		Name:       schema,
		Tables:     tables,
		Enums:      enums,
		Procedures: procedures,
	}, nil
}

func getTables(conn *pgx.Conn, schema string) (tables []*db.Table, err error) {
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
	q, err := conn.Query(sqlstr, schema, "r")
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		t := db.Table{}

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

func getColumns(conn *pgx.Conn, schema string, table string) (columns []*db.Column, err error) {
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
	q, err := conn.Query(sqlstr, schema, table)
	if err != nil {
		return columns, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		c := db.Column{}

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

func getForeignKeys(conn *pgx.Conn, schema string, table string) (fks []*db.ForeignKey, err error) {
	// sql query

	const sqlstr = `
    SELECT b.attname, format_type(d.atttypid, d.atttypmod), r.conname, i.relname, c.relname, d.attname, 0, 0, '', '', ''
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
	q, err := conn.Query(sqlstr, schema, table)
	if err != nil {
		return fks, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		fk := db.ForeignKey{}

		// scan
		err = q.Scan(&fk.Name, &fk.DataType, &fk.ForeignKeyName, &fk.RefIndexName, &fk.RefTableName, &fk.RefColumnName, &fk.KeyID, &fk.SeqNo, &fk.OnUpdate, &fk.OnDelete, &fk.Match)
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

func getIndexes(conn *pgx.Conn, schema string, table string) (indexes []*db.Index, err error) {
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
	q, err := conn.Query(sqlstr, schema, table)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		i := db.Index{}

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
func getIndexColumns(conn *pgx.Conn, schema string, table string, index string) ([]*db.IndexColumn, error) {
	// load columns
	cols, err := indexColumns(conn, schema, index)
	if err != nil {
		return nil, err
	}

	// load col order
	colOrd, err := colOrder(conn, schema, index)
	if err != nil {
		return nil, err
	}

	// build schema name used in errors
	s := schema
	if s != "" {
		s = s + "."
	}

	// put cols in order using colOrder
	ret := []*db.IndexColumn{}
	for _, v := range strings.Split(colOrd.Ord, " ") {
		cid, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("could not convert %s%s index %s column %s to int", s, table, index, v)
		}

		// find column
		found := false
		var c *db.IndexColumn
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

	return ret, nil
}

// indexColumns runs a custom query, returning results as IndexColumn.
func indexColumns(conn *pgx.Conn, schema string, index string) ([]*db.IndexColumn, error) {
	var err error

	// query the index columns
	const sqlstr = `
		SELECT (row_number() over()), a.attnum, a.attname, format_type(a.atttypid, a.atttypmod), a.attnotnull FROM pg_index i
		JOIN ONLY pg_class c ON c.oid = i.indrelid
		JOIN ONLY pg_namespace n ON n.oid = c.relnamespace
		JOIN ONLY pg_class ic ON ic.oid = i.indexrelid
		LEFT JOIN pg_attribute a ON i.indrelid = a.attrelid AND a.attnum = ANY(i.indkey) AND a.attisdropped = false
		WHERE i.indkey <> '0' AND n.nspname = $1 AND ic.relname = $2
	`

	// run query
	// DBLog(sqlstr, schema, index)
	q, err := conn.Query(sqlstr, schema, index)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*db.IndexColumn{}
	for q.Next() {
		ic := db.IndexColumn{}

		// scan
		err = q.Scan(&ic.SeqNo, &ic.Cid, &ic.Name, &ic.DataType, &ic.NotNull)
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
func colOrder(conn *pgx.Conn, schema string, index string) (*columnOrder, error) {
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
	err = conn.QueryRow(sqlstr, schema, index).Scan(&pco.Ord)
	if err != nil {
		return nil, err
	}

	return &pco, nil
}

func getProcedures(conn *pgx.Conn, schema string) (procs []*db.Procedure, err error) {
	// sql query
	const sqlstr = `SELECT ` +
		`p.proname, ` + // ::varchar AS proc_name
		`pg_get_function_result(p.oid) ` + // ::varchar AS return_type
		`FROM pg_proc p ` +
		`JOIN ONLY pg_namespace n ON p.pronamespace = n.oid ` +
		`WHERE n.nspname = $1`

	// run query
	q, err := conn.Query(sqlstr, schema)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		p := db.Procedure{}

		// scan
		err = q.Scan(&p.Name, &p.ReturnType)
		if err != nil {
			return nil, err
		}

		procs = append(procs, &p)
	}

	// range over the procs and get the parameters
	for i, proc := range procs {
		// get the params
		params, err := getProcedureParams(conn, schema, proc.Name)
		if err != nil {
			return procs, err
		}
		procs[i].Params = append(procs[i].Params, params...)
	}

	return procs, nil
}

func getProcedureParams(conn *pgx.Conn, schema, procedure string) (params []*db.ProcedureParam, err error) {
	// sql query
	const sqlstr = `SELECT ` +
		`UNNEST(p.proargnames), ` + // ::varchar as name
		`UNNEST(STRING_TO_ARRAY(oidvectortypes(p.proargtypes), ', ')) ` + // ::varchar AS param_type
		`FROM pg_proc p ` +
		`JOIN ONLY pg_namespace n ON p.pronamespace = n.oid ` +
		`WHERE n.nspname = $1 AND p.proname = $2`

	q, err := conn.Query(sqlstr, schema, procedure)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		pp := db.ProcedureParam{}

		// scan
		err = q.Scan(&pp.Name, &pp.Type)
		if err != nil {
			return nil, err
		}

		params = append(params, &pp)
	}

	return params, nil
}

// getTable function
func getEnums(conn *pgx.Conn, schema string) (enums []*db.Enum, err error) {
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
	q, err := conn.Query(sqlstr, schema)
	if err != nil {
		return nil, err
	}

	// load results
	for q.Next() {
		e := db.Enum{}

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
		values, err := getEnumValues(conn, schema, re.Name)
		if err != nil {
			return nil, err
		}
		re.Values = values
	}

	return enums, nil
}

// Values runs a custom query, returning results as Value.
func getEnumValues(conn *pgx.Conn, schema string, enum string) ([]*db.EnumValue, error) {
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
	q, err := conn.Query(sqlstr, schema, enum)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*db.EnumValue{}
	for q.Next() {
		ev := db.EnumValue{}

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
