package schema

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/matthewmueller/gotext"
)

// NewIndex fn
func NewIndex(
	name string,
	isUnique bool,
	isPrimary bool,
	paramPrefix string,
	columns []*IndexColumn,
) *Index {
	sort.Slice(columns, func(i, j int) bool { return columns[i].name < columns[j].name })
	return &Index{
		name,
		isUnique,
		isPrimary,
		paramPrefix,
		columns,
	}
}

// Index data
type Index struct {
	name      string // index_name
	isUnique  bool   // is_unique
	isPrimary bool   // is_primary

	paramPrefix string
	columns     []*IndexColumn
}

// IsUnique fn
func (i *Index) IsUnique() bool {
	return i.isUnique
}

// IsPrimary fn
func (i *Index) IsPrimary() bool {
	return i.isPrimary
}

// Columns fn
func (i *Index) Columns() []*IndexColumn {
	columns := make([]*IndexColumn, len(i.columns))
	copy(columns[:], i.columns[:])
	// sort the columns before using them
	sort.Slice(columns, func(i, j int) bool {
		return columns[i].name < columns[j].name
	})
	return columns
}

// Method for the index
func (i *Index) Method() string {
	var cols []string
	for _, col := range i.columns {
		cols = append(cols, gotext.Pascal(col.name))
	}
	sort.Strings(cols)
	return strings.Join(cols, "And")
}

// Description fn
func (i *Index) Description() string {
	var cols []string
	for _, col := range i.columns {
		cols = append(cols, col.name)
	}
	sort.Strings(cols)
	return strings.Join(cols, " and ")
}

// Params fn
func (i *Index) Params() (string, error) {
	var cols []string
	for _, col := range i.columns {
		cols = append(cols, gotext.Camel(col.name)+" "+col.dataType.String())
	}
	sort.Strings(cols)
	return strings.Join(cols, ", "), nil
}

// Where fn
func (i *Index) Where() string {
	var cols []string

	// sort the column names
	for _, col := range i.columns {
		cols = append(cols, col.name)
	}
	sort.Strings(cols)
	for j, col := range cols {
		cols[j] = fmt.Sprintf("%s = %s%d", strconv.Quote(col), i.paramPrefix, j+1)
	}
	return strings.Join(cols, " AND ")
}

// Variables fn
func (i *Index) Variables() string {
	var cols []string
	for _, col := range i.columns {
		cols = append(cols, gotext.Camel(col.name))
	}
	sort.Strings(cols)
	return strings.Join(cols, ", ")
}

// ColumnList is string-friendly the list of the columns
func (i *Index) ColumnList() string {
	var cols []string
	for _, col := range i.columns {
		cols = append(cols, strconv.Quote(col.name))
	}
	sort.Strings(cols)
	return strings.Join(cols, ", ")
}
