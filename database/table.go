package database

import (
	"strconv"
	"strings"

	"github.com/matthewmueller/go-gen"
	"github.com/matthewmueller/pogo/util"
)

// Table struct
type Table struct {
	Type        byte   // type
	Name        string // table_name
	ManualPk    bool   // manual_pk
	Columns     []*Column
	ForeignKeys []*ForeignKey
	Indexes     []*Index
}

// Slug generates the slug case
func (t *Table) Slug() string {
	return gen.Lower(gen.Slug(util.Singular(t.Name)))
}

// Pascal generates the pascal case
func (t *Table) Pascal() string {
	return gen.Pascal(util.Singular(t.Name))
}

// Camel generates the camel case
func (t *Table) Camel() string {
	return gen.Camel(util.Singular(t.Name))
}

// PluralCamel generates the camel case
func (t *Table) PluralCamel() string {
	return gen.Camel(util.Plural(t.Name))
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

// Select builds the SQL SELECT string
func (t *Table) Select() string {
	var cols []string
	for _, col := range t.Columns {
		cols = append(cols, `"`+col.Name+`"`)
	}
	return strings.Join(cols, ", ")
}

// Returning builds the SQL RETURNING string
func (t *Table) Returning() string {
	var cols []string
	for _, col := range t.Columns {
		cols = append(cols, strconv.Quote(col.Name))
	}
	return strings.Join(cols, ", ")
}

// Scan builds the DB.Scan(...) params
func (t *Table) Scan() string {
	user := gen.Camel(util.Singular(t.Name))

	var cols []string
	for _, col := range t.Columns {
		cols = append(cols, `&`+user+`.`+gen.Pascal(col.Name))
	}
	return strings.Join(cols, ", ")
}

// IsManyToMany checks if the relationship is many-to-many
func (t *Table) IsManyToMany() bool {
	var pks []string

	for _, c := range t.Columns {
		if c.IsPrimaryKey {
			pks = append(pks, c.Name)
		}
	}
	if len(pks) > 1 {
		return true
	}

	return false
}
