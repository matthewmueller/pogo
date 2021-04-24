package schema

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/matthewmueller/gotext"
)

// NewPrimaryKey fn
func NewPrimaryKey(
	columns []*Column,
	paramPrefix string,
) *PrimaryKey {
	sort.Slice(columns, func(i, j int) bool { return columns[i].name < columns[j].name })
	return &PrimaryKey{columns, paramPrefix}
}

// PrimaryKey struct
type PrimaryKey struct {
	columns     []*Column
	paramPrefix string
}

// Columns fn
func (pk *PrimaryKey) Columns() []*Column {
	return pk.columns
}

// Method for the index
func (pk *PrimaryKey) Method() string {
	var cols []string
	for _, col := range pk.columns {
		pascal := gotext.Pascal(col.name)
		// fix sqlite's "rowid"
		if col.name == "rowid" {
			pascal = "RowID"
		}
		cols = append(cols, pascal)
	}
	sort.Strings(cols)
	return strings.Join(cols, "And")
}

// Params fn
func (pk *PrimaryKey) Params() (string, error) {
	var cols []string
	for _, col := range pk.columns {
		cols = append(cols, gotext.Camel(col.name)+" "+col.dataType.String())
	}
	sort.Strings(cols)
	return strings.Join(cols, ", "), nil
}

// Where fn
func (pk *PrimaryKey) Where() string {
	var cols []string

	// sort the column names
	for _, col := range pk.columns {
		cols = append(cols, col.name)
	}
	sort.Strings(cols)
	for j, col := range cols {
		cols[j] = fmt.Sprintf("%s = %s%d", strconv.Quote(col), pk.paramPrefix, j+1)
	}
	return strings.Join(cols, " AND ")
}

// Variables fn
func (pk *PrimaryKey) Variables() string {
	var cols []string
	for _, col := range pk.columns {
		cols = append(cols, gotext.Camel(col.name))
	}
	sort.Strings(cols)
	return strings.Join(cols, ", ")
}

// ColumnList is string-friendly the list of the columns
func (pk *PrimaryKey) ColumnList() string {
	var cols []string
	for _, col := range pk.columns {
		cols = append(cols, strconv.Quote(col.name))
	}
	sort.Strings(cols)
	return strings.Join(cols, ", ")
}
