package db

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	gen "github.com/matthewmueller/go-gen"
)

// Index data
type Index struct {
	Name      string // index_name
	IsUnique  bool   // is_unique
	IsPrimary bool   // is_primary
	SeqNo     int    // seq_no
	Origin    string // origin
	IsPartial bool   // is_partial
	Columns   []*IndexColumn
}

// Method for the index
func (i *Index) Method() string {
	var cols []string
	for _, col := range i.Columns {
		cols = append(cols, gen.Pascal(col.Name))
	}
	sort.Strings(cols)
	return strings.Join(cols, "And")
}

// Description fn
func (i *Index) Description() string {
	var cols []string
	for _, col := range i.Columns {
		cols = append(cols, col.Name)
	}
	sort.Strings(cols)
	return strings.Join(cols, " and ")
}

// Params fn
func (i *Index) Params(schema *Schema) (string, error) {
	var cols []string
	for _, col := range i.Columns {
		typ, err := col.Type(schema)
		if err != nil {
			return "", err
		}
		cols = append(cols, gen.Camel(col.Name)+" "+typ)
	}
	sort.Strings(cols)
	return strings.Join(cols, ", "), nil
}

// Where fn
func (i *Index) Where() string {
	var cols []string

	// sort the column names
	for _, col := range i.Columns {
		cols = append(cols, col.Name)
	}
	sort.Strings(cols)
	for i, col := range cols {
		cols[i] = fmt.Sprintf("%s = $%d", strconv.Quote(col), i+1)
	}

	return strings.Join(cols, " AND ")
}

// Variables fn
func (i *Index) Variables() string {
	var cols []string
	for _, col := range i.Columns {
		cols = append(cols, gen.Camel(col.Name))
	}
	sort.Strings(cols)
	return strings.Join(cols, ", ")
}

// ColumnList is string-friendly the list of the columns
func (i *Index) ColumnList() string {
	var cols []string
	for _, col := range i.Columns {
		cols = append(cols, strconv.Quote(col.Name))
	}
	sort.Strings(cols)
	return strings.Join(cols, ", ")
}

// IndexColumn represents index column info.
type IndexColumn struct {
	SeqNo    int    // seq_no
	Cid      int    // cid
	Name     string // column_name
	NotNull  bool
	DataType string
}

// Pascal case
func (c *IndexColumn) Pascal() string {
	return gen.Pascal(c.Name)
}

// Camel case
func (c *IndexColumn) Camel() string {
	return gen.Camel(c.Name)
}

// Type of the column
func (c *IndexColumn) Type(schema *Schema) (string, error) {
	return coerce(schema, c.DataType)
}
