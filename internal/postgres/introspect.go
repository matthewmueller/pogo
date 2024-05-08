package postgres

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/matthewmueller/pogo/internal/schema"
	"github.com/pkg/errors"
)

// Table struct
type Table struct {
	Type     byte   // type
	ManualPk bool   // manual_pk
	Name     string // table_name
}

// Column struct
type Column struct {
	FieldOrdinal int
	Name         string
	DataType     string
	NotNull      bool
	Comment      *string
	DefaultValue *string
	IsPrimaryKey bool
}

// Index struct
type Index struct {
	Name      string
	IsUnique  bool
	IsPrimary bool
	SeqNo     int
	Origin    string
	IsPartial bool
}

// IndexColumn struct
type IndexColumn struct {
	SeqNo    int
	Cid      int
	Name     string
	DataType string
	NotNull  bool
}

// ForeignKey struct
type ForeignKey struct {
	Name          string
	FullName      string
	DataType      string
	RefIndexName  string
	RefTableName  string
	RefColumnName string
	KeyID         int
	SeqNo         int
	OnUpdate      string
	OnDelete      string
	Match         string
}

// Introspect a postgres database
func (d *DB) Introspect(schemaName string) (*schema.Schema, error) {
	var tables []*schema.Table

	paramPrefix := "$"

	tt, err := getTables(d.Conn, schemaName)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get tables from schema")
	}

	procedures, err := getProcedures(d.Conn, schemaName)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get the stored procedures")
	}

	// range over the procs and get the parameters
	for i, proc := range procedures {
		// get the params
		params, err := getProcedureParams(d.Conn, schemaName, proc.Name)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get the stored procedure params for '%s'", proc.Name)
		}
		procedures[i].Params = append(procedures[i].Params, params...)
	}

	// get enums
	enums, err := getEnums(d.Conn, schemaName)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get the enums")
	}

	for _, table := range tt {
		// get the columns
		cols, err := getColumns(d.Conn, schemaName, table.Name)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get columns for '%s' from schema", table.Name)
		}

		var columns []*schema.Column
		var pks []*schema.Column
		for _, col := range cols {
			dt, err := getType(enums, schemaName, col.DataType)
			if err != nil {
				return nil, err
			}
			column := schema.NewColumn(col.Name, "", dt, col.NotNull, nil, col.DefaultValue, col.IsPrimaryKey)
			columns = append(columns, column)
			if col.IsPrimaryKey {
				pks = append(pks, column)
			}
		}

		// group the primary key columns into a composite primary key
		pk := schema.NewPrimaryKey(pks, paramPrefix)

		// get the foreign keys
		fks, err := getForeignKeys(d.Conn, schemaName, table.Name)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get the foreign keys for '%s' from schema", table.Name)
		}

		// get the indexes
		idxs, err := getIndexes(d.Conn, schemaName, table.Name)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get the indexes for '%s' from schema", table.Name)
		}

		// get each of the index columns
		var indexes []*schema.Index
		for _, index := range idxs {
			// get the index columns
			columns, err := getIndexColumns(d.Conn, enums, schemaName, table.Name, index.Name)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to get index columns for %s", index.Name)
			}
			indexes = append(indexes, schema.NewIndex(index.Name, index.IsUnique, index.IsPrimary, paramPrefix, columns))
		}

		tables = append(tables, schema.NewTable(schemaName, table.Name, columns, pk, fks, indexes))
	}

	return schema.New(
		"postgres",
		schemaName,
		tables,
		enums,
		procedures,
	), nil
}

func getTables(conn *pgx.Conn, schemaName string) (tables []*Table, err error) {
	// sql query
	const sqlstr = `
		SELECT
			c.relkind as type,
			c.relname as name
		FROM pg_class c
		JOIN ONLY pg_namespace n ON n.oid = c.relnamespace
		WHERE n.nspname = $1 AND c.relkind = $2
		ORDER BY c.relname
	`

	// run query
	// DBLog(sqlstr, schema, relkind)
	// "r" constant is for tables
	q, err := conn.Query(context.TODO(), sqlstr, schemaName, "r")
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		var t Table

		// scan
		err = q.Scan(&t.Type, &t.Name)
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

func getColumns(conn *pgx.Conn, schemaName string, table string) (columns []*Column, err error) {
	// sql query
	// TODO: support onDelete and onUpdate
	const sqlstr = `
		SELECT
			a.attnum as field_ordinal,
			a.attname as name,
			format_type(a.atttypid, a.atttypmod) as data_type,
			a.attnotnull as not_null,
			d.description as comment,
			pg_get_expr(ad.adbin, ad.adrelid) as default_value,
			COALESCE(ct.contype = 'p', false) as is_primary_key
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
	q, err := conn.Query(context.TODO(), sqlstr, schemaName, table)
	if err != nil {
		return columns, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		var c Column
		// scan
		err = q.Scan(&c.FieldOrdinal, &c.Name, &c.DataType, &c.NotNull, &c.Comment, &c.DefaultValue, &c.IsPrimaryKey)
		if err != nil {
			return columns, err
		}
		columns = append(columns, &c)
	}
	if e := q.Err(); e != nil {
		return nil, e
	}
	return columns, nil
}

func getForeignKeys(conn *pgx.Conn, schemaName string, table string) (fks []*schema.ForeignKey, err error) {
	// sql query
	const sqlstr = `
		SELECT
			b.attname as name,
			format_type(d.atttypid, d.atttypmod) as data_type,
			r.conname as full_name,
			i.relname as ref_index_name,
			c.relname as ref_table_name,
			d.attname as ref_column_name,
			0 as key_id,
			0 as seq_no,
			'' as on_update,
			'' as on_delete,
			'' as match
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
	q, err := conn.Query(context.TODO(), sqlstr, schemaName, table)
	if err != nil {
		return fks, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		var fk ForeignKey

		// scan
		err = q.Scan(&fk.Name, &fk.DataType, &fk.FullName, &fk.RefIndexName, &fk.RefTableName, &fk.RefColumnName, &fk.KeyID, &fk.SeqNo, &fk.OnUpdate, &fk.OnDelete, &fk.Match)
		if err != nil {
			return fks, err
		}

		// fk.DataType, err = getType(enums, schemaName, dt)
		// if err != nil {
		// 	return fks, err
		// }

		// fks = append(fks, &fk)
	}
	if e := q.Err(); e != nil {
		return nil, e
	}

	return fks, nil
}

func getIndexes(conn *pgx.Conn, schemaName string, table string) (indexes []*Index, err error) {
	// sql query
	const sqlstr = `
		SELECT DISTINCT
			ic.relname as name,
			i.indisunique as is_unique,
			i.indisprimary as is_primary,
			0 as seq_no,
			'' as origin,
			false as is_partial
		FROM pg_index i
		JOIN ONLY pg_class c ON c.oid = i.indrelid
		JOIN ONLY pg_namespace n ON n.oid = c.relnamespace
		JOIN ONLY pg_class ic ON ic.oid = i.indexrelid
		WHERE i.indkey <> '0' AND n.nspname = $1 AND c.relname = $2
	`

	// run query
	// DBLog(sqlstr, schema, table)
	q, err := conn.Query(context.TODO(), sqlstr, schemaName, table)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		var i Index

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
func getIndexColumns(conn *pgx.Conn, enums []*schema.Enum, schemaName string, table string, index string) (columns []*schema.IndexColumn, err error) {
	// load columns
	cols, err := indexColumns(conn, schemaName, index)
	if err != nil {
		return nil, err
	}

	// load col order
	colOrd, err := colOrder(conn, schemaName, index)
	if err != nil {
		return nil, err
	}

	// build schema name used in errors
	s := schemaName
	if s != "" {
		s = s + "."
	}

	// put cols in order using colOrder
	for _, v := range strings.Split(colOrd.Ord, " ") {
		cid, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("could not convert %s%s index %s column %s to int", s, table, index, v)
		}

		// find column
		found := false
		var c *IndexColumn
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

		dt, err := getType(enums, schemaName, c.DataType)
		if err != nil {
			return nil, err
		}

		columns = append(columns, schema.NewIndexColumn(c.Name, dt))
	}

	return columns, nil
}

// indexColumns runs a custom query, returning results as IndexColumn.
func indexColumns(conn *pgx.Conn, schemaName string, index string) (columns []*IndexColumn, err error) {
	// query the index columns
	const sqlstr = `
		SELECT
			(row_number() over()) as seq_no,
			a.attnum as cid,
			a.attname as name,
			format_type(a.atttypid, a.atttypmod) as data_type,
			a.attnotnull as not_null
		FROM pg_index i
		JOIN ONLY pg_class c ON c.oid = i.indrelid
		JOIN ONLY pg_namespace n ON n.oid = c.relnamespace
		JOIN ONLY pg_class ic ON ic.oid = i.indexrelid
		LEFT JOIN pg_attribute a ON i.indrelid = a.attrelid AND a.attnum = ANY(i.indkey) AND a.attisdropped = false
		WHERE i.indkey <> '0' AND n.nspname = $1 AND ic.relname = $2
	`

	// run query
	// DBLog(sqlstr, schema, index)
	q, err := conn.Query(context.TODO(), sqlstr, schemaName, index)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		var ic IndexColumn

		// scan
		err = q.Scan(&ic.SeqNo, &ic.Cid, &ic.Name, &ic.DataType, &ic.NotNull)
		if err != nil {
			return nil, err
		}

		columns = append(columns, &ic)
	}
	if e := q.Err(); e != nil {
		return nil, e
	}

	return columns, nil
}

// PgColOrder represents index column order.
type columnOrder struct {
	Ord string // ord
}

// colOrder runs a custom query, returning results as columnOrder.
func colOrder(conn *pgx.Conn, schemaName string, index string) (*columnOrder, error) {
	var err error
	// sql query
	const sqlstr = `
		SELECT
			i.indkey as ord
		FROM pg_index i
		JOIN ONLY pg_class c ON c.oid = i.indrelid
		JOIN ONLY pg_namespace n ON n.oid = c.relnamespace
		JOIN ONLY pg_class ic ON ic.oid = i.indexrelid
		WHERE n.nspname = $1 AND ic.relname = $2
	`
	// run query
	var pco columnOrder
	err = conn.QueryRow(context.TODO(), sqlstr, schemaName, index).Scan(&pco.Ord)
	if err != nil {
		return nil, err
	}

	return &pco, nil
}

func getProcedures(conn *pgx.Conn, schemaName string) (procs []*schema.Procedure, err error) {
	// sql query
	const sqlstr = `
		SELECT
			p.proname as name,
			pg_get_function_result(p.oid) as return_type
		FROM pg_proc p
		JOIN ONLY pg_namespace n ON p.pronamespace = n.oid
		WHERE n.nspname = $1
	`
	// run query
	q, err := conn.Query(context.TODO(), sqlstr, schemaName)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		var p schema.Procedure

		// scan
		var name *string
		err = q.Scan(&name, &p.ReturnType)
		if err != nil {
			return nil, err
		}
		// skip procedures that don't have a name
		if name == nil {
			continue
		}
		p.Name = *name

		procs = append(procs, &p)
	}

	return procs, nil
}

func getProcedureParams(conn *pgx.Conn, schemaName, procedure string) (params []*schema.ProcedureParam, err error) {
	// sql query
	const sqlstr = `
		SELECT
			UNNEST(p.proargnames) as name,
			UNNEST(STRING_TO_ARRAY(oidvectortypes(p.proargtypes), ', ')) as param_type
		FROM pg_proc p
		JOIN ONLY pg_namespace n ON p.pronamespace = n.oid
		WHERE n.nspname = $1 AND p.proname = $2
	`
	// run query
	q, err := conn.Query(context.TODO(), sqlstr, schemaName, procedure)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		pp := schema.ProcedureParam{}

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
func getEnums(conn *pgx.Conn, schemaName string) (enums []*schema.Enum, err error) {
	// sql query
	const sqlstr = `
    SELECT DISTINCT
			t.typname as name
		FROM pg_type t
		JOIN ONLY pg_namespace n ON n.oid = t.typnamespace
		JOIN ONLY pg_enum e ON t.oid = e.enumtypid
		WHERE n.nspname = $1
  `
	// run query
	// DBLog(sqlstr, schema)
	q, err := conn.Query(context.TODO(), sqlstr, schemaName)
	if err != nil {
		return nil, err
	}

	// load results
	for q.Next() {
		e := schema.Enum{}

		// scan
		err = q.Scan(&e.Name)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get enums")
		}

		enums = append(enums, &e)
	}
	if e := q.Err(); e != nil {
		return nil, e
	}
	q.Close()

	for _, re := range enums {
		values, err := getEnumValues(conn, schemaName, re.Name)
		if err != nil {
			return nil, err
		}
		re.Values = values
	}

	return enums, nil
}

// Values runs a custom query, returning results as Value.
func getEnumValues(conn *pgx.Conn, schemaName string, enum string) ([]*schema.EnumValue, error) {
	var err error

	// sql query
	const sqlstr = `
    SELECT
			e.enumlabel as label,
			e.enumsortorder::int as order
		FROM pg_type t
		JOIN ONLY pg_namespace n ON n.oid = t.typnamespace
		LEFT JOIN pg_enum e ON t.oid = e.enumtypid
		WHERE n.nspname = $1 AND t.typname = $2
  `

	// run query
	// DBLog(sqlstr, schema, enum)
	q, err := conn.Query(context.TODO(), sqlstr, schemaName, enum)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*schema.EnumValue{}
	for q.Next() {
		ev := schema.EnumValue{}

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

// getType takes an SQL type and returns a schema.Type
func getType(enums []*schema.Enum, schemaName, sqlType string) (schema.DataType, error) {
	// handle SETOF
	if strings.HasPrefix(sqlType, "SETOF ") {
		t, err := getType(enums, schemaName, sqlType[len("SETOF "):])
		if err != nil {
			return nil, err
		}
		return &schema.List{DataType: t}, nil
	}

	// determine if it's an array
	if strings.HasSuffix(sqlType, "[]") {
		sqlType = sqlType[:len(sqlType)-2]
		t, err := getType(enums, schemaName, sqlType)
		if err != nil {
			return nil, err
		}
		return &schema.List{DataType: t}, nil
	}

	// treat numerics differently
	// TODO: distinguish between different numerics
	// e.g numeric(5, 3), numeric(7, 1), etc.
	if strings.HasPrefix(sqlType, "numeric") {
		return &schema.Float64{}, nil
	}

	switch sqlType {
	case "text", "uuid", "citext":
		return &schema.String{}, nil
	case "boolean":
		return &schema.Boolean{}, nil
	case "integer", "smallint":
		return &schema.Int{}, nil
	case "bigint":
		return &schema.Int64{}, nil
	case "real", "double", "float":
		// TODO distinguish float32, float64, etc. with new types
		return &schema.Float64{}, nil
	case "time with time zone", "time without time zone":
		return &schema.String{}, nil
	case "date", "timestamp", "timestamp with time zone", "timestamp without time zone":
		return &schema.DateTime{}, nil
	case "json", "jsonb":
		return &schema.JSON{}, nil
	}

	// handle enums
	unquoted := strings.Replace(sqlType, "\"", "", -1)
	for _, enum := range enums {
		name := enum.Name
		if schemaName != "" && schemaName != "public" {
			name = fmt.Sprintf(`%s.%s`, schemaName, name)
		}
		if name == unquoted {
			return &schema.Enumerable{
				Schema: schemaName,
				Name:   enum.Name,
			}, nil
		}
	}

	return nil, fmt.Errorf(`postgres getType: unhandled data type: %q`, sqlType)
}
