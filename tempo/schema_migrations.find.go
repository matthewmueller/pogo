package tempo

import (
	"github.com/matthewmueller/pgx"
)

// GENERATED BY POGO. DO NOT EDIT.

// Find a SchemaMigration by "version"
func (sm *SchemaMigrations) Find(version *int) (schemamigration SchemaMigration, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT "version"
    FROM public.schema_migrations
    WHERE "version" = $1`

	DBLog(sqlstr, version)
	row := sm.DB.QueryRow(sqlstr, version)
	err = row.Scan(&schemamigration.Version)
	if err != nil {
		if err == pgx.ErrNoRows {
			return schemamigration, ErrSchemaMigrationNotFound
		}
		return schemamigration, err
	}

	return schemamigration, nil
}