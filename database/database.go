package database

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/matthewmueller/go-gen"
)

// Database interface
type Database interface {
	Introspect(schema string) (*Schema, error)
}

// Schema introspection
type Schema struct {
	Name   string
	Tables []*Table
	Enums  []*Enum
}

// Table represents table info.
type Table struct {
	Type        byte   // type
	Name        string // table_name
	ManualPk    bool   // manual_pk
	Columns     []*Column
	ForeignKeys []*ForeignKey
	Indexes     []*Index
}

// Model name
func (t *Table) Model() string {
	parts := strings.Split(gen.Base(t.Name), " ")
	for i, w := range parts {
		parts[i] = gen.Singular(w)
	}
	return gen.Pascal(strings.Join(parts, " "))
}

// PrimaryKey fn
func (t *Table) PrimaryKey() *Column {
	for _, col := range t.Columns {
		if col.IsPrimaryKey {
			return col
		}
	}
	return nil
}

// Uniques fn
func (t *Table) Uniques() (uniques []*Index) {
	for _, idx := range t.Indexes {
		if idx.IsPrimary || !idx.IsUnique {
			continue
		}
		uniques = append(uniques, idx)
	}
	return uniques
}

// Composite fn
func (t *Table) Composite() (foreign *Composite) {
	var cols []*Column
	for _, fk := range t.ForeignKeys {
		for _, col := range t.Columns {
			if fk.Name != col.Name {
				continue
			}
			cols = append(cols, col)
		}
	}

	if len(cols) == 0 {
		return nil
	}
	return &Composite{cols}
}

// SQLSelect generator
func (t *Table) SQLSelect() string {
	var cols []string
	for _, col := range t.Columns {
		cols = append(cols, `"`+col.Name+`"`)
	}
	return strings.Join(cols, ", ")
}

// GoScan generator
func (t *Table) GoScan() string {
	parts := strings.Split(gen.Base(t.Name), " ")
	for i, w := range parts {
		parts[i] = gen.Singular(w)
	}
	user := gen.Camel(strings.Join(parts, " "))

	var cols []string
	for _, col := range t.Columns {
		cols = append(cols, `&`+user+`.`+gen.Pascal(col.Name))
	}
	return strings.Join(cols, ", ")
}

// SQLReturn generator
func (t *Table) SQLReturn() string {
	var cols []string
	for _, col := range t.Columns {
		cols = append(cols, strconv.Quote(col.Name))
	}
	return strings.Join(cols, ", ")
}

// Column represents column info.
type Column struct {
	FieldOrdinal int     // field_ordinal
	Name         string  // column_name
	DataType     string  // data_type
	NotNull      bool    // not_null
	Comment      *string // description
	DefaultValue *string // default_value
	IsPrimaryKey bool    // is_primary_key

	// we add this on during introspection
	GoType string
}

// ForeignKey represents a foreign key.
type ForeignKey struct {
	ForeignKeyName string // foreign_key_name
	Name           string // column_name
	RefIndexName   string // ref_index_name
	RefTableName   string // ref_table_name
	RefColumnName  string // ref_column_name
	KeyID          int    // key_id
	SeqNo          int    // seq_no
	OnUpdate       string // on_update
	OnDelete       string // on_delete
	Match          string // match
}

// Composite is just for template
type Composite struct {
	Columns []*Column
}

// NumColumns fn
func (c *Composite) NumColumns() int {
	return len(c.Columns)
}

// Description fn
func (c *Composite) Description() string {
	var cols []string
	for _, col := range c.Columns {
		cols = append(cols, col.Name)
	}
	sort.Strings(cols)
	return strings.Join(cols, " and ")
}

// SQLWhere fn
func (c *Composite) SQLWhere() string {
	var cols []string
	for i, col := range c.Columns {
		cols = append(cols, fmt.Sprintf("\"%s\" = $%d", col.Name, i+1))
	}
	sort.Strings(cols)
	return strings.Join(cols, " AND ")
}

// GoParams fn
func (c *Composite) GoParams() string {
	var cols []string
	for _, col := range c.Columns {
		cols = append(cols, gen.Camel(col.Name)+" "+col.GoType)
	}
	sort.Strings(cols)
	return strings.Join(cols, ", ")
}

// GoVars fn
func (c *Composite) GoVars() string {
	var cols []string
	for _, col := range c.Columns {
		cols = append(cols, gen.Camel(col.Name))
	}
	sort.Strings(cols)
	return strings.Join(cols, ", ")
}

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

func (i *Index) NumColumns() int {
	return len(i.Columns)
}

func (i *Index) GoMethod() string {
	var cols []string
	for _, col := range i.Columns {
		cols = append(cols, gen.Pascal(col.Name))
	}
	sort.Strings(cols)
	return strings.Join(cols, "And")
}

func (i *Index) Description() string {
	var cols []string
	for _, col := range i.Columns {
		cols = append(cols, col.Name)
	}
	sort.Strings(cols)
	return strings.Join(cols, " and ")
}

func (i *Index) GoParams() string {
	var cols []string
	for _, col := range i.Columns {
		cols = append(cols, gen.Camel(col.Name)+" "+col.GoType)
	}
	sort.Strings(cols)
	return strings.Join(cols, ", ")
}

func (i *Index) SQLColumns() string {
	var cols []string
	for _, col := range i.Columns {
		cols = append(cols, strconv.Quote(col.Name))
	}
	sort.Strings(cols)
	return strings.Join(cols, ", ")
}

func (i *Index) SQLWhere() string {
	var cols []string
	for i, col := range i.Columns {
		cols = append(cols, fmt.Sprintf("%s = $%d", strconv.Quote(col.Name), i+1))
	}
	sort.Strings(cols)
	return strings.Join(cols, " AND ")
}

func (i *Index) GoVars() string {
	var cols []string
	for _, col := range i.Columns {
		cols = append(cols, gen.Camel(col.Name))
	}
	sort.Strings(cols)
	return strings.Join(cols, ", ")
}

// IndexColumn represents index column info.
type IndexColumn struct {
	SeqNo    int    // seq_no
	Cid      int    // cid
	Name     string // column_name
	DataType string

	// we add this on during introspection
	GoType string
}

// Enum struct
type Enum struct {
	Name   string
	Values []*Value
}

// Value represents a enum value.
type Value struct {
	Label string
	Order int
}

// PgColOrder represents index column order.
type PgColOrder struct {
	Ord string // ord
}
