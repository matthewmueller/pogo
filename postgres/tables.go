package postgres

import "github.com/matthewmueller/pogo/db"

// Table represents table info.
type Table struct {
	Type      *string // type
	TableName *string // table_name
	ManualPk  *bool   // manual_pk
}

// Tables get all the postgres tables
func Tables(db db.DB, schema string) (tables []*Table, err error) {
	// sql query
	const sqlstr = `
    SELECT c.relkind, c.relname, false
    FROM pg_class c
    JOIN ONLY pg_namespace n ON n.oid = c.relnamespace
    WHERE n.nspname = $1 AND c.relkind = $2
    ORDER BY c.relname
  `

	// run query
	// DBLog(sqlstr, schema, relkind)
	// "r" constant is for tables
	q, err := db.Query(sqlstr, schema, "r")
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	for q.Next() {
		t := Table{}

		// scan
		err = q.Scan(&t.Type, &t.TableName, &t.ManualPk)
		if err != nil {
			return nil, err
		}

		tables = append(tables, &t)
	}

	return tables, nil
}
