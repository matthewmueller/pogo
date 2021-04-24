package schema

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/matthewmueller/gotext"
	"github.com/matthewmueller/text"
)

// NewTable fn
func NewTable(
	schema string,
	name string,
	columns []*Column,
	pk *PrimaryKey,
	fks []*ForeignKey,
	idxs []*Index,
) *Table {
	sort.Slice(columns, func(i, j int) bool { return columns[i].name < columns[j].name })
	sort.Slice(idxs, func(i, j int) bool { return idxs[i].name < idxs[j].name })
	sort.Slice(fks, func(i, j int) bool { return fks[i].name < fks[j].name })
	return &Table{
		schema,
		name,
		columns,
		pk,
		fks,
		idxs,
	}
}

// Table struct
type Table struct {
	schema  string
	name    string
	columns []*Column
	pk      *PrimaryKey
	fks     []*ForeignKey
	idxs    []*Index
}

// SQLName is an SQL-friendly name of the table
// e.g "public"."blogs"
func (t *Table) SQLName() string {
	if t.schema == "" {
		return fmt.Sprintf("%q", t.name)
	}
	return fmt.Sprintf("%q.%q", t.schema, t.name)
}

// Columns of the table
func (t *Table) Columns() []*Column {
	return t.columns
}

// Indexes of the table
func (t *Table) Indexes() []*Index {
	return t.idxs
}

// Slug generates the slug case
func (t *Table) Slug() string {
	return gotext.Lower(text.Pascal(singular(t.name)))
}

// Pascal generates the pascal case
func (t *Table) Pascal() string {
	return gotext.Pascal(singular(t.name))
}

// Short generates a short variable
func (t *Table) Short() string {
	return gotext.Lower(text.Short(singular(t.name)))
}

// Camel generates the camel case
func (t *Table) Camel() string {
	return gotext.Camel(singular(t.name))
}

// PluralCamel generates the camel case
func (t *Table) PluralCamel() string {
	return gotext.Camel(plural(t.name))
}

// PrimaryKey fn
func (t *Table) PrimaryKey() *PrimaryKey {
	if len(t.pk.columns) == 0 {
		return nil
	}
	return t.pk
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
	for _, col := range t.columns {
		cols = append(cols, `"`+col.SQLName()+`"`)
	}
	return strings.Join(cols, ", ")
}

// Returning builds the SQL RETURNING string
func (t *Table) Returning() string {
	var cols []string
	for _, col := range t.columns {
		cols = append(cols, strconv.Quote(col.SQLName()))
	}
	return strings.Join(cols, ", ")
}

// Scan builds the DB.Scan(...) params
func (t *Table) Scan() string {
	camel := gotext.Camel(singular(t.name))
	var cols []string
	for _, col := range t.columns {
		cols = append(cols, `&_`+camel+`.`+col.Pascal())
	}
	return strings.Join(cols, ", ")
}

// Filters fn
func (t *Table) Filters() (filters []*Filter) {
	for _, col := range t.columns {
		filters = append(filters, newFilter(col.name, col.dataType, col.notNull))
	}
	return filters
}

// Orders fn
func (t *Table) Orders() (orders []*OrderField) {
	for _, col := range t.columns {
		orders = append(orders, newOrder(col.name, col.comment, col.dataType))
	}
	return orders
}
