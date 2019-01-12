package postgres

import (
	"database/sql"
	"os/exec"
	"strings"
	"testing"

	_ "github.com/lib/pq"
	"github.com/matthewmueller/pogo/internal/databases/postgres"

	"github.com/tj/assert"
)

var connStr = "postgres://localhost:5432/pogo_test?sslmode=disable"

func TestIntrospect(t *testing.T) {
	// url := os.Getenv("POSTGRES_URL")
	assert.NotEmpty(t, connStr)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			exec.Command("dropdb", "pogo_test").Run()
			assert.NoError(t, exec.Command("createdb", "pogo_test").Run())
			defer func() {
				if !t.Failed() {
					assert.NoError(t, exec.Command("dropdb", "pogo_test").Run())
				}
			}()

			db, err := sql.Open("postgres", connStr)
			assert.NoError(t, err)
			defer db.Close()

			_, err = db.Exec(test.up)
			assert.NoError(t, err)

			pogo, err := postgres.Introspect(db)
			assert.NoError(t, err)

			assert.Equal(t, strings.TrimSpace(test.expected), pogo.TestString())
		})
	}

}

var tests = []struct {
	name     string
	up       string
	expected string
}{
	{
		name: "2 schemas",
		up: `
			create schema a;
			create schema b;
		`,
		expected: `
			DATABASE(SCHEMA(A),SCHEMA(B))
		`,
	},
}
