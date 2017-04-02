package pogo

import (
	"fmt"

	"github.com/matthewmueller/pogo/db"
	"github.com/matthewmueller/pogo/postgres"
	"github.com/pkg/errors"
)

// Generate the database
func Generate(db db.DB) (err error) {
	tables, err := postgres.Tables(db, "jack")
	if err != nil {
		return errors.Wrap(err, "unable to lookup tables")
	}

	for _, table := range tables {
		fmt.Println(table.TableName)

		columns, err := postgres.Columns(db, "jack", table.TableName)
		if err != nil {
			return errors.Wrap(err, "unable to lookup columns")
		}

		for _, column := range columns {
			fmt.Printf("  %s:    %s\n", column.ColumnName, column.DataType)
		}
	}

	return nil
}

// Inspect the database
func Inspect(db db.DB) {

}
