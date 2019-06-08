package sqlite

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/matthewmueller/pogo/internal/schema"
	"github.com/pkg/errors"
)

// Table struct
type Table struct {
	Name string
}

// Index struct
type Index struct {
	IndexName string // index_name
	IsUnique  bool   // is_unique
	SeqNo     int    // seq_no
	Origin    string // origin
	IsPartial bool   // is_partial
}

// IndexColumn struct
type IndexColumn struct {
	SeqNo      int    // seq_no
	Cid        int    // cid
	ColumnName string // column_name
}

// Column struct
type Column struct {
	FieldOrdinal int     // field_ordinal
	ColumnName   string  // column_name
	DataType     string  // data_type
	NotNull      bool    // not_null
	DefaultValue *string // default_value
	IsPrimaryKey bool    // pk_col_index

	Alias string
}

// ForeignKey struct
type ForeignKey struct {
	ColumnName    string  // column_name
	RefIndexName  string  // ref_index_name
	RefTableName  string  // ref_table_name
	RefColumnName *string // ref_column_name
	KeyID         int     // key_id
	SeqNo         int     // seq_no
	OnUpdate      string  // on_update
	OnDelete      string  // on_delete
	Match         string  // match
}

// Introspect a sqlite database
// TODO: support views
func (d *DB) Introspect(schemaName string) (*schema.Schema, error) {
	var tables []*schema.Table

	paramPrefix := "?"

	// get all the tables
	tt, err := d.getTables(schemaName, "table")
	if err != nil {
		return nil, errors.Wrap(err, "unable to get tables from schema")
	}

	// build a map of tables to a list of columns
	colmap := make(map[string][]*Column)
	for _, table := range tt {
		cols, err := d.getColumns(schemaName, table.Name)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get cols for '%s' from schema", table.Name)
		}
		for _, col := range cols {
			colmap[table.Name] = append(colmap[table.Name], col)
		}
	}

	// get the indexes
	for _, table := range tt {
		// turn columns into schema.Column
		var columns []*schema.Column
		var pks []*schema.Column
		for _, col := range colmap[table.Name] {
			dt, err := getType(schemaName, col.DataType)
			if err != nil {
				return nil, err
			}
			column := schema.NewColumn(col.ColumnName, col.Alias, dt, col.NotNull, nil, col.DefaultValue, col.IsPrimaryKey)
			columns = append(columns, column)
			if col.IsPrimaryKey {
				pks = append(pks, column)
			}
		}

		// group the primary key columns into a composite primary key
		pk := schema.NewPrimaryKey(pks, paramPrefix)

		// get the foreign keys
		fks, err := d.getForeignKeys(schemaName, table.Name, colmap)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get the foreign keys for '%s' from schema", table.Name)
		}

		idxs, err := d.getIndexes(schemaName, table.Name)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get the indexes for '%s' from schema", table.Name)
		}
		// get each of the index columns
		var indexes []*schema.Index
		for _, index := range idxs {
			icols, err := d.getIndexColumns(schemaName, table.Name, colmap[table.Name], index.IndexName)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to get index columns for %s", index.IndexName)
			}
			indexes = append(indexes, schema.NewIndex(index.IndexName, index.IsUnique, index.Origin == "pk", paramPrefix, icols))
		}
		tables = append(tables, schema.NewTable(schemaName, table.Name, columns, pk, fks, indexes))
	}

	return schema.New(
		"sqlite",
		schemaName,
		tables,
		[]*schema.Enum{},
		[]*schema.Procedure{},
	), nil
}

func (d *DB) getTables(schemaName, relkind string) (tables []*Table, err error) {
	conn := d.DB

	// sql query
	const sqlstr = `SELECT
		tbl_name AS table_name
		FROM sqlite_master
		WHERE tbl_name NOT LIKE 'sqlite_%' AND type = ?`

	// run query
	q, err := conn.Query(sqlstr, relkind)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		var t Table
		// scan
		err = q.Scan(&t.Name)
		if err != nil {
			return nil, err
		}
		tables = append(tables, &t)
	}

	return tables, nil
}

func (d *DB) getColumns(schemaName string, table string) (cols []*Column, err error) {
	conn := d.DB

	// sql query
	var sqlstr = `PRAGMA table_info(` + table + `)`

	// run query
	q, err := conn.Query(sqlstr)
	if err != nil {
		return cols, err
	}
	defer q.Close()

	pks := 0

	// load results
	for q.Next() {
		var c Column
		var defaultValue sql.NullString
		var primaryKey int

		// scan
		err = q.Scan(&c.FieldOrdinal, &c.ColumnName, &c.DataType, &c.NotNull, &defaultValue, &primaryKey)
		if err != nil {
			return cols, err
		}

		if defaultValue.Valid {
			// TODO: not sure why i need to copy it in first
			s := defaultValue.String
			c.DefaultValue = &s
		}

		if primaryKey > 0 {
			c.IsPrimaryKey = true
			pks++
		}

		cols = append(cols, &c)
	}
	if e := q.Err(); e != nil {
		return nil, e
	}

	// if we don't have an explicit primary key,
	// sqlite assigns a 64bit integer named "rowid"
	// prepend it to the existing columns
	if pks == 0 {
		cols = append([]*Column{
			{
				FieldOrdinal: 0,
				ColumnName:   "rowid",
				DataType:     "INTEGER",
				NotNull:      true,
				DefaultValue: nil,
				IsPrimaryKey: true,
			},
		}, cols...)
		return cols, nil
	}

	return cols, nil
}

func (d *DB) getForeignKeys(schemaName string, table string, colmap map[string][]*Column) (fks []*schema.ForeignKey, err error) {
	conn := d.DB

	// sql query
	var sqlstr = `PRAGMA foreign_key_list(` + table + `)`

	// run query
	q, err := conn.Query(sqlstr, table)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		var f ForeignKey

		// scan
		err = q.Scan(&f.KeyID, &f.SeqNo, &f.RefTableName, &f.ColumnName, &f.RefColumnName, &f.OnUpdate, &f.OnDelete, &f.Match)
		if err != nil {
			return nil, err
		}

		// make introspection a bit stricter for better compatibility with other dbs
		cols, ok := colmap[f.RefTableName]
		if !ok {
			return nil, fmt.Errorf("sqlite introspect: couldn't find foreign table: %q", f.RefTableName)
		}

		var c *Column
		for _, col := range cols {
			if col.ColumnName != *f.RefColumnName {
				continue
			}
			c = col
			break
		}
		if c == nil {
			return nil, fmt.Errorf("sqlite introspect: couldn't find foreign column: %q.(%q)", table, f.ColumnName)
		}

		dt, err := getType(schemaName, c.DataType)
		if err != nil {
			return nil, err
		}

		fks = append(fks, schema.NewForeignKey(f.ColumnName, dt))
	}

	return fks, nil
}

func (d *DB) getIndexes(schemaName string, table string) (idxs []*Index, err error) {
	conn := d.DB

	// sql query
	sqlstr := `PRAGMA index_list(` + table + `)`

	// run query
	// DBLog(sqlstr, schema, table)
	q, err := conn.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		var i Index

		// scan
		err = q.Scan(&i.SeqNo, &i.IndexName, &i.IsUnique, &i.Origin, &i.IsPartial)
		if err != nil {
			return nil, err
		}
		idxs = append(idxs, &i)
	}
	if e := q.Err(); e != nil {
		return nil, e
	}

	return idxs, nil
}

// get the column indexes
func (d *DB) getIndexColumns(schemaName string, table string, cols []*Column, index string) (ics []*schema.IndexColumn, err error) {
	conn := d.DB

	// query the index columns
	sqlstr := `PRAGMA index_info(` + index + `)`

	// run query
	q, err := conn.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		var ic IndexColumn

		// scan
		err = q.Scan(&ic.SeqNo, &ic.Cid, &ic.ColumnName)
		if err != nil {
			return nil, err
		}

		// find the datatype
		var c *Column
		for _, col := range cols {
			if col.ColumnName != ic.ColumnName {
				continue
			}
			c = col
			break
		}
		if c == nil {
			return nil, fmt.Errorf("sqlite introspect: couldn't find referenced column: %q.(%q) while getting the index columns", table, ic.ColumnName)
		}

		dt, err := getType(schemaName, c.DataType)
		if err != nil {
			return nil, err
		}
		ics = append(ics, schema.NewIndexColumn(ic.ColumnName, dt))
	}
	if e := q.Err(); e != nil {
		return nil, e
	}

	return ics, nil
}

// getType takes an SQL type and returns a schema.Type
func getType(schemaName, sqlType string) (schema.DataType, error) {
	// handle SETOF
	if strings.HasPrefix(sqlType, "SETOF ") {
		t, err := getType(schemaName, sqlType[len("SETOF "):])
		if err != nil {
			return nil, err
		}
		return &schema.List{DataType: t}, nil
	}

	// determine if it's an array
	if strings.HasSuffix(sqlType, "[]") {
		sqlType = sqlType[:len(sqlType)-2]
		t, err := getType(schemaName, sqlType)
		if err != nil {
			return nil, err
		}
		return &schema.List{DataType: t}, nil
	}

	switch strings.ToLower(sqlType) {
	case "text", "uuid", "citext":
		return &schema.String{}, nil
	case "boolean":
		return &schema.Boolean{}, nil
	case "integer", "smallint", "bigint":
		// TODO distinguish int32, int64, etc. with new types
		return &schema.Integer{}, nil
	case "real", "double", "float":
		// TODO distinguish float32, float64, etc. with new types
		return &schema.Float32{}, nil
	case "time with time zone", "time without time zone":
		return &schema.String{}, nil
	case "date", "timestamp", "timestamp with time zone", "timestamp without time zone":
		return &schema.DateTime{}, nil
	case "json", "jsonb":
		return &schema.JSON{}, nil
	}

	return nil, fmt.Errorf(`sqlite getType: unhandled data type: %q`, sqlType)
}
