package sqlite

import (
	"database/sql"
	"strings"

	"github.com/matthewmueller/errors"
	"github.com/matthewmueller/pogo/internal/schema"
)

// Introspect a sqlite database
// TODO: support views
func (d *DB) Introspect(schemaName string) (*schema.Schema, error) {
	tables, err := d.getTables(schemaName, "table")
	if err != nil {
		return nil, errors.Wrap(err, "unable to get tables from schema")
	}

	// get all columns in all tables first
	for _, table := range tables {
		// get the columns
		columns, err := d.getColumns(schemaName, table.Name)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get columns for '%s' from schema", table.Name)
		}
		table.Columns = columns
	}

	// get the foreign keys
	for _, table := range tables {
		fks, err := d.getForeignKeys(tables, schemaName, table.Name)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get the foreign keys for '%s' from schema", table.Name)
		}
		table.ForeignKeys = fks
	}

	// get the indexes
	for _, table := range tables {
		indexes, err := d.getIndexes(schemaName, table.Name)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get the indexes for '%s' from schema", table.Name)
		}
		// get each of the index columns
		for _, index := range indexes {
			// get the index columns
			icols, err := d.getIndexColumns(tables, schemaName, table.Name, index.Name)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to get index columns for %s", index.Name)
			}
			index.Columns = icols
		}
		table.Indexes = indexes
	}

	return &schema.Schema{
		Name:   schemaName,
		Tables: tables,
	}, nil
}

func (d *DB) getTables(schemaName, relkind string) (tables []*schema.Table, err error) {
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
	var res []*schema.Table
	for q.Next() {
		var t schema.Table
		// scan
		err = q.Scan(&t.Name)
		if err != nil {
			return nil, err
		}
		res = append(res, &t)
	}

	return res, nil
}

func (d *DB) getColumns(schemaName string, table string) (cols []*schema.Column, err error) {
	conn := d.DB

	// sql query
	var sqlstr = `PRAGMA table_info(` + table + `)`

	// run query
	q, err := conn.Query(sqlstr)
	if err != nil {
		return cols, err
	}
	defer q.Close()

	type column struct {
		FieldOrdinal int            // field_ordinal
		ColumnName   string         // column_name
		DataType     string         // data_type
		NotNull      bool           // not_null
		DefaultValue sql.NullString // default_value
		PkColIndex   int            // pk_col_index
	}
	var cc []column
	var hasPrimaryKey bool

	// load results
	for q.Next() {
		var c column

		// scan
		err = q.Scan(&c.FieldOrdinal, &c.ColumnName, &c.DataType, &c.NotNull, &c.DefaultValue, &c.PkColIndex)
		if err != nil {
			return cols, err
		}

		if c.PkColIndex == 1 {
			hasPrimaryKey = true
		}

		cc = append(cc, c)
	}
	if e := q.Err(); e != nil {
		return nil, e
	}

	// shift the fields
	shift := 0

	// if we don't have an explicit primary key,
	// sqlite assigns a 64bit integer named "rowid"
	if !hasPrimaryKey {
		shift++
		cols = append(cols, &schema.Column{
			// TODO: consider an alias to id
			FieldOrdinal: shift,
			Name:         "rowid",
			DataType:     "INTEGER",
			NotNull:      true,
			IsPrimaryKey: true,
		})
	}

	// map columns into []schema.Column
	for _, c := range cc {
		var col schema.Column
		col.FieldOrdinal = shift + c.FieldOrdinal
		col.Name = c.ColumnName
		col.DataType = strings.ToUpper(c.DataType)
		col.NotNull = c.NotNull
		col.IsPrimaryKey = c.PkColIndex == 1
		if c.DefaultValue.Valid {
			// TODO: not sure why i need to copy it in first
			s := c.DefaultValue.String
			col.DefaultValue = &s
		}
		cols = append(cols, &col)
	}

	return cols, nil
}

func (d *DB) getForeignKeys(allTables []*schema.Table, schemaName string, table string) (fks []*schema.ForeignKey, err error) {
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
		var f struct {
			ColumnName    string // column_name
			RefIndexName  string // ref_index_name
			RefTableName  string // ref_table_name
			RefColumnName string // ref_column_name
			KeyID         int    // key_id
			SeqNo         int    // seq_no
			OnUpdate      string // on_update
			OnDelete      string // on_delete
			Match         string // match
		}

		// scan
		err = q.Scan(&f.KeyID, &f.SeqNo, &f.RefTableName, &f.ColumnName, &f.RefColumnName, &f.OnUpdate, &f.OnDelete, &f.Match)
		if err != nil {
			return nil, err
		}

		// map to schema.ForeignKey
		var fk schema.ForeignKey
		fk.Name = f.ColumnName
		fk.RefIndexName = f.RefIndexName
		fk.RefTableName = f.RefTableName
		fk.RefColumnName = f.RefColumnName
		fk.KeyID = f.KeyID
		fk.SeqNo = f.SeqNo
		fk.OnUpdate = f.OnUpdate
		fk.OnDelete = f.OnDelete
		fk.Match = f.Match

		// find the datatype
		for _, t := range allTables {
			if t.Name != table {
				continue
			}
			for _, col := range t.Columns {
				if col.Name != f.RefColumnName {
					continue
				}
				fk.DataType = col.DataType
			}
		}

		fks = append(fks, &fk)
	}

	return fks, nil
}

func (d *DB) getIndexes(schemaName string, table string) (idxs []*schema.Index, err error) {
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
		var i struct {
			IndexName string // index_name
			IsUnique  bool   // is_unique
			SeqNo     int    // seq_no
			Origin    string // origin
			IsPartial bool   // is_partial
		}

		// scan
		err = q.Scan(&i.SeqNo, &i.IndexName, &i.IsUnique, &i.Origin, &i.IsPartial)
		if err != nil {
			return nil, err
		}

		// map to schema.Index
		var idx schema.Index
		idx.Name = i.IndexName
		idx.IsUnique = i.IsUnique
		idx.IsPrimary = false
		idx.SeqNo = i.SeqNo
		idx.Origin = i.Origin
		idx.IsPartial = i.IsPartial
		idxs = append(idxs, &idx)
	}
	if e := q.Err(); e != nil {
		return nil, e
	}

	return idxs, nil
}

// get the column indexes
func (d *DB) getIndexColumns(allTables []*schema.Table, schemaName string, table string, index string) (ics []*schema.IndexColumn, err error) {
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
		var ic struct {
			SeqNo      int    // seq_no
			Cid        int    // cid
			ColumnName string // column_name
		}

		// scan
		err = q.Scan(&ic.SeqNo, &ic.Cid, &ic.ColumnName)
		if err != nil {
			return nil, err
		}

		var idxc schema.IndexColumn
		idxc.Name = ic.ColumnName
		idxc.SeqNo = ic.SeqNo
		idxc.Cid = ic.Cid

		// find the datatype
		for _, t := range allTables {
			if t.Name != table {
				continue
			}
			for _, col := range t.Columns {
				if col.Name != ic.ColumnName {
					continue
				}
				idxc.DataType = col.DataType
				idxc.NotNull = col.NotNull
			}
		}

		ics = append(ics, &idxc)
	}
	if e := q.Err(); e != nil {
		return nil, e
	}

	return ics, nil
}
