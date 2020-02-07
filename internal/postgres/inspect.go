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
		join only pg_namespace n ON n.oid = t.typnamespace
		join only pg_enum e ON t.oid = e.enumtypid
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
		enums = append(enums, enum)
	}
	return enums, nil
}

func (i *Inspector) enumValues(enum *schema.Enum) (values []*schema.EnumValue, err error) {
	const sql = `
		SELECT
			e.enumsortorder::int as order,
			e.enumlabel as value
		from pg_type t
		join only pg_namespace n ON n.oid = t.typnamespace
		left join pg_enum e ON t.oid = e.enumtypid
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
		join only pg_namespace n ON n.oid = c.relnamespace
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
	for _, table := range tables {
		// get the columns
		table.Columns, err = i.columns(table, enums)
		if err != nil {
			return nil, err
		}
		// get the primary index
		table.Primary, err = i.primary(table, table.Columns)
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
		join only pg_class c ON c.oid = a.attrelid
		join only pg_namespace n ON n.oid = c.relnamespace
		left join pg_attrdef ad ON ad.adrelid = c.oid and ad.adnum = a.attnum
		left join pg_description d ON d.objoid = a.attrelid and d.objsubid = a.attnum
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
		join only pg_class a ON a.oid = r.conrelid
		join only pg_namespace n ON n.oid = r.connamespace
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
		primary.Schema = i.Schema
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
