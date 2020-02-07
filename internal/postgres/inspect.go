package postgres

import (
	"fmt"

	"github.com/matthewmueller/pogo/internal/schema"
)

// Inspect fn
func Inspect(db *Conn) (s *schema.Schema, err error) {
	i := &Inspector{db, "public"}
	return i.Inspect()
}

// Inspector struct
type Inspector struct {
	DB     DB
	Schema string
}

// Inspect function
func (i *Inspector) Inspect() (scheme *schema.Schema, err error) {
	scheme = new(schema.Schema)
	// get the enums
	scheme.Enums, err = i.enums()
	if err != nil {
		return nil, err
	}
	// get the tables
	scheme.Tables, err = i.tables(scheme.Enums)
	if err != nil {
		return nil, err
	}
	return scheme, nil
}

func (i *Inspector) enums() (enums []*schema.Enum, err error) {
	// sql query
	const sql = `
		SELECT DISTINCT
			t.typname as name
		from pg_type t
		join only pg_namespace n on n.oid = t.typnamespace
		join only pg_enum e on t.oid = e.enumtypid
		where n.nspname = $1
	`
	rows, err := i.DB.Query(sql, i.Schema)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		// scan
		var enum schema.Enum
		enum.Schema = i.Schema
		err = rows.Scan(&enum.Name)
		if err != nil {
			return nil, err
		}
		enums = append(enums, &enum)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	// explicitly close to avoid "conn busy"
	rows.Close()
	for _, enum := range enums {
		// add the enum values
		enum.Values, err = i.enumValues(enum)
		if err != nil {
			return nil, err
		}
	}
	return enums, nil
}

func (i *Inspector) enumValues(enum *schema.Enum) (values []*schema.EnumValue, err error) {
	const sql = `
		SELECT
			e.enumsortorder::int as order,
			e.enumlabel as value
		from pg_type t
		join only pg_namespace n on n.oid = t.typnamespace
		left join pg_enum e on t.oid = e.enumtypid
		where n.nspname = $1 and t.typname = $2
	`
	rows, err := i.DB.Query(sql, i.Schema, enum.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		// scan
		var value schema.EnumValue
		err = rows.Scan(&value.Order, &value.Label)
		if err != nil {
			return nil, err
		}
		values = append(values, &value)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	rows.Close()
	return values, nil
}

func (i *Inspector) tables(enums []*schema.Enum) (tables []*schema.Table, err error) {
	// sql query, "r" constant is for tables
	const sql = `
		SELECT
			c.relname as name
		from pg_class c
		join only pg_namespace n on n.oid = c.relnamespace
		where n.nspname = $1 and c.relkind = 'r'
		order by c.relname
	`
	rows, err := i.DB.Query(sql, i.Schema)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		// scan
		table := new(schema.Table)
		table.Schema = i.Schema
		err = rows.Scan(&table.Name)
		if err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	rows.Close()
	// fill up the columns first
	for _, table := range tables {
		table.Columns, err = i.columns(table, enums)
		if err != nil {
			return nil, err
		}
	}
	// work on the indexes that depend on columns
	for _, table := range tables {
		// get the primary index
		table.Primary, err = i.primary(table, table.Columns)
		if err != nil {
			return nil, err
		}
		// get the foreign keys
		table.Foreigns, err = i.foreigns(table, tables)
		if err != nil {
			return nil, err
		}
		// get the unique indexes
		table.Uniques, err = i.uniques(table, table.Columns)
		if err != nil {
			return nil, err
		}
	}
	return tables, nil
}

func (i *Inspector) columns(table *schema.Table, enums []*schema.Enum) (columns []*schema.Column, err error) {
	// sql query
	const sql = `
		select
			a.attnum as order,
			a.attname as name,
			format_type(a.atttypid, a.atttypmod) as data_type,
			a.attnotnull as not_null,
			d.description as comment,
			pg_get_expr(ad.adbin, ad.adrelid) as default_value
		from pg_attribute a
		join only pg_class c on c.oid = a.attrelid
		join only pg_namespace n on n.oid = c.relnamespace
		left join pg_attrdef ad on ad.adrelid = c.oid and ad.adnum = a.attnum
		left join pg_description d on d.objoid = a.attrelid and d.objsubid = a.attnum
		where a.attisdropped = false and n.nspname = $1 and c.relname = $2 and a.attnum > 0
		order by a.attnum
	`
	rows, err := i.DB.Query(sql, i.Schema, table.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		// scan
		column := new(schema.Column)
		err = rows.Scan(&column.Order, &column.Name, &column.Type, &column.NotNull, &column.Comment, &column.Default)
		if err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	rows.Close()
	return columns, nil
}

// find the primary key
func (i *Inspector) primary(table *schema.Table, columns []*schema.Column) (primary *schema.Primary, err error) {
	// sql query
	const sql = `
		select
			conname as name,
			conkey as columns
		from pg_constraint r
		join only pg_class a on a.oid = r.conrelid
		join only pg_namespace n on n.oid = r.connamespace
		where r.contype = 'p' and n.nspname = $1 and a.relname = $2
		order by conname;
	`
	rows, err := i.DB.Query(sql, i.Schema, table.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		primary = new(schema.Primary)
		var ids []int
		// scan
		err = rows.Scan(&primary.Name, &ids)
		if err != nil {
			return nil, err
		}
		// link primary key columns to table columns
		primary.Columns = make([]*schema.Column, len(ids))
		for i, id := range ids {
			for j, column := range columns {
				if j+1 != id {
					continue
				}
				primary.Columns[i] = column
				break
			}
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	// we only expect a single primary key
	if rows.Next() {
		return nil, fmt.Errorf("postgres.Inspector.primary(%q) has more than 1 primary key", table.Name)
	}
	rows.Close()
	return primary, nil
}

// find the primary key
func (i *Inspector) foreigns(table *schema.Table, tables []*schema.Table) (foreigns []*schema.Foreign, err error) {
	// sql query
	const sql = `
		select
			r.conname as name,
			r.conkey as columns,
			r.confkey as ref_columns,
			c.relname as ref_table,
			n.nspname as ref_schema
		from pg_constraint r
		join only pg_namespace n on n.oid = r.connamespace
		join only pg_class a on a.oid = r.conrelid
		join only pg_class c on c.oid = r.confrelid
		join only pg_namespace rn on rn.oid = c.relnamespace
		where r.contype = 'f' and n.nspname = $1 and a.relname = $2;
	`
	rows, err := i.DB.Query(sql, i.Schema, table.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var foreign schema.Foreign
		var columns []int
		var refColumns []int
		// scan
		err = rows.Scan(&foreign.Name, &columns, &refColumns, &foreign.RefTable, &foreign.RefSchema)
		if err != nil {
			return nil, err
		}
		// link foreign columns to table columns
		foreign.Columns = make([]*schema.Column, len(columns))
		for i, id := range columns {
			for j, column := range table.Columns {
				if j+1 != id {
					continue
				}
				foreign.Columns[i] = column
			}
		}
		// link foreign reference columns to the referenced table columns
		foreign.RefColumns = make([]*schema.Column, len(refColumns))
		for _, refTable := range tables {
			if refTable.Name != foreign.RefTable || refTable.Schema != foreign.RefSchema {
				continue
			}
			for i, id := range refColumns {
				for j, column := range refTable.Columns {
					if j+1 != id {
						continue
					}
					foreign.RefColumns[i] = column
				}
			}
		}
		foreigns = append(foreigns, &foreign)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	rows.Close()
	return foreigns, nil
}

// find the unique indexes
func (i *Inspector) uniques(table *schema.Table, columns []*schema.Column) (uniques []*schema.Unique, err error) {
	// sql query
	const sql = `
		select
			conname as name,
			conkey as columns
		from pg_constraint r
		join only pg_class a on a.oid = r.conrelid
		join only pg_namespace n on n.oid = r.connamespace
		where r.contype = 'u' and n.nspname = $1 and a.relname = $2
		order by conname;
	`
	rows, err := i.DB.Query(sql, i.Schema, table.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var unique schema.Unique
		var columns []int
		// scan
		err = rows.Scan(&unique.Name, &columns)
		if err != nil {
			return nil, err
		}
		// link foreign columns to table columns
		unique.Columns = make([]*schema.Column, len(columns))
		for i, id := range columns {
			for j, column := range table.Columns {
				if j+1 != id {
					continue
				}
				unique.Columns[i] = column
			}
		}
		uniques = append(uniques, &unique)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	rows.Close()
	return uniques, nil
}
