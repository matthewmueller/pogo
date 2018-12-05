package schema

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/matthewmueller/go-gen"
)

// NewTable fn
func NewTable(
	schema string,
	name string,
	cols []*Column,
	fks []*ForeignKey,
	idxs []*Index,
) *Table {
	return &Table{schema, name, cols, fks, idxs}
}

// Table struct
type Table struct {
	schema string
	name   string
	cols   []*Column
	fks    []*ForeignKey
	idxs   []*Index
}

// SQLName is an SQL-friendly name of the table
// e.g "public"."blogs"
func (t *Table) SQLName() string {
	if t.schema == "" {
		return fmt.Sprintf("%q", t.schema)
	}
	return fmt.Sprintf("%q.%q", t.schema, t.name)
}

// func (t *Table) Name() string {

// }

// Columns of the table
func (t *Table) Columns() []*Column {
	return t.cols
}

// Indexes of the table
func (t *Table) Indexes() []*Index {
	return t.idxs
}

// Slug generates the slug case
func (t *Table) Slug() string {
	return gen.Lower(gen.Pascal(singular(t.name)))
}

// Pascal generates the pascal case
func (t *Table) Pascal() string {
	return gen.Pascal(singular(t.name))
}

// Short generates a short variable
func (t *Table) Short() string {
	return gen.Lower(gen.Short(singular(t.name)))
}

// Camel generates the camel case
func (t *Table) Camel() string {
	return gen.Camel(singular(t.name))
}

// PluralCamel generates the camel case
func (t *Table) PluralCamel() string {
	return gen.Camel(plural(t.name))
}

// PrimaryKey fn
func (t *Table) PrimaryKey() *Column {
	for _, col := range t.cols {
		if col.isPrimaryKey {
			return col
		}
	}
	return nil
}

// Uniques fn
func (t *Table) Uniques() (uniques []*Index) {
	for _, idx := range t.idxs {
		if idx.isPrimary || !idx.isUnique {
			continue
		}
		uniques = append(uniques, idx)
	}
	return uniques
}

// Select builds the SQL SELECT string
func (t *Table) Select() string {
	var cols []string
	for _, col := range t.cols {
		cols = append(cols, `"`+col.Snake()+`"`)
	}
	return strings.Join(cols, ", ")
}

// Returning builds the SQL RETURNING string
func (t *Table) Returning() string {
	var cols []string
	for _, col := range t.cols {
		cols = append(cols, strconv.Quote(col.Snake()))
	}
	return strings.Join(cols, ", ")
}

// Scan builds the DB.Scan(...) params
func (t *Table) Scan() string {
	camel := gen.Camel(singular(t.name))
	var cols []string
	for _, col := range t.cols {
		cols = append(cols, `&_`+camel+`.`+col.Pascal())
	}
	return strings.Join(cols, ", ")
}

// Filters fn
func (t *Table) Filters() (filters []*Filter) {
	for _, col := range t.cols {
		filters = append(filters, newFilter(col.name, col.dataType, col.notNull))
	}
	return filters
}

// Orders fn
func (t *Table) Orders() (orders []*OrderField) {
	for _, col := range t.cols {
		orders = append(orders, newOrder(col.name, col.comment, col.dataType))
	}
	return orders
}
