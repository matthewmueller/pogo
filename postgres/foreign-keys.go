package postgres

import "github.com/matthewmueller/pogo/db"

// ForeignKey represents a foreign key.
type ForeignKey struct {
	ForeignKeyName string // foreign_key_name
	ColumnName     string // column_name
	RefIndexName   string // ref_index_name
	RefTableName   string // ref_table_name
	RefColumnName  string // ref_column_name
	KeyID          int    // key_id
	SeqNo          int    // seq_no
	OnUpdate       string // on_update
	OnDelete       string // on_delete
	Match          string // match
}

// ForeignKeys runs a custom query, returning results as ForeignKey.
func ForeignKeys(db db.DB, schema string, table string) (fks []*ForeignKey, err error) {
	// sql query

	const sqlstr = `
    SELECT r.conname, b.attname, i.relname, c.relname, d.attname, 0, 0, '', '', ''
    FROM pg_constraint r
    JOIN ONLY pg_class a ON a.oid = r.conrelid
    JOIN ONLY pg_attribute b ON b.attisdropped = false AND b.attnum = ANY(r.conkey) AND b.attrelid = r.conrelid
    JOIN ONLY pg_class i on i.oid = r.conindid
    JOIN ONLY pg_class c on c.oid = r.confrelid
    JOIN ONLY pg_attribute d ON d.attisdropped = false AND d.attnum = ANY(r.confkey) AND d.attrelid = r.confrelid
    JOIN ONLY pg_namespace n ON n.oid = r.connamespace
    WHERE r.contype = 'f' AND n.nspname = $1 AND a.relname = $2
    ORDER BY r.conname, b.attname
  `

	// run query
	// XOLog(sqlstr, schema, table)
	q, err := db.Query(sqlstr, schema, table)
	if err != nil {
		return fks, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		fk := ForeignKey{}

		// scan
		err = q.Scan(&fk.ForeignKeyName, &fk.ColumnName, &fk.RefIndexName, &fk.RefTableName, &fk.RefColumnName, &fk.KeyID, &fk.SeqNo, &fk.OnUpdate, &fk.OnDelete, &fk.Match)
		if err != nil {
			return fks, err
		}

		fks = append(fks, &fk)
	}

	return fks, nil
}
