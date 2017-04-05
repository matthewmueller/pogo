package postgres

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/matthewmueller/pogo/db"
)

// IndexColumn represents index column info.
type IndexColumn struct {
	SeqNo      int    // seq_no
	Cid        int    // cid
	ColumnName string // column_name
}

// IndexColumns returns the column list for an index.
func IndexColumns(db db.DB, schema string, table string, index string) ([]*IndexColumn, error) {
	var err error

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
	ret := []*IndexColumn{}
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

		ret = append(ret, c)
	}

	return ret, nil
}

// PgColOrder represents index column order.
type PgColOrder struct {
	Ord string // ord
}

// indexColumns runs a custom query, returning results as IndexColumn.
func indexColumns(db db.DB, schema string, index string) ([]*IndexColumn, error) {
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
	res := []*IndexColumn{}
	for q.Next() {
		ic := IndexColumn{}

		// scan
		err = q.Scan(&ic.SeqNo, &ic.Cid, &ic.ColumnName)
		if err != nil {
			return nil, err
		}

		res = append(res, &ic)
	}

	return res, nil
}

// colOrder runs a custom query, returning results as PgColOrder.
func colOrder(db db.DB, schema string, index string) (*PgColOrder, error) {
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
	var pco PgColOrder
	err = db.QueryRow(sqlstr, schema, index).Scan(&pco.Ord)
	if err != nil {
		return nil, err
	}

	return &pco, nil
}
