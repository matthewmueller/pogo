package tempo

import (
	"github.com/matthewmueller/pgx"
)

// GENERATED BY POGO. DO NOT EDIT.

// Delete the SchemaMigration from the database.
func (sm *SchemaMigrations) Delete(version *int) (err error) {
	// sql query
	sqlstr := `DELETE FROM public.schema_migrations WHERE "version" = $1`

	// run query
	DBLog(sqlstr, version)
	_, err = sm.DB.Exec(sqlstr, version)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrSchemaMigrationNotFound
		}
		return err
	}

	return nil
}