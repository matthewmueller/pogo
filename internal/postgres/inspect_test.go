package postgres_test

import (
	"os"
	"testing"

	"github.com/matthewmueller/diff"
	"github.com/matthewmueller/pogo/internal/postgres"
	"github.com/matthewmueller/pogo/internal/tests"
	"github.com/matthewmueller/pogo/internal/testutil"
	"github.com/tj/assert"
)

func dial(t testing.TB, url string) (*postgres.Pool, func()) {
	pool, err := postgres.Dial(url)
	assert.NoError(t, err)
	return pool, func() {
		pool.Close()
	}
}

func TestPG(t *testing.T) {
	url := os.Getenv("POSTGRES_URL")
	assert.NotEmpty(t, url)
	pool, err := postgres.Dial(url)
	assert.NoError(t, err)
	defer pool.Close()
	for _, test := range tests.Postgres {
		t.Run(testutil.Name(t, test), func(t *testing.T) {
			conn, err := pool.Acquire()
			assert.NoError(t, err)
			defer conn.Release()
			// migrate down
			if test.Down != "" {
				conn, err := pool.Acquire()
				defer conn.Release()
				_, err = conn.Exec(test.Down)
				assert.NoError(t, err)
			}
			// migrate up
			if test.Up != "" {
				_, err := conn.Exec(test.Up)
				assert.NoError(t, err)
			}
			conn.Release()
			if test.Down != "" {
				defer func() {
					// migrate down after
					conn, err := pool.Acquire()
					defer conn.Release()
					_, err = conn.Exec(test.Down)
					assert.NoError(t, err)
				}()
			}
			conn, err = pool.Acquire()
			assert.NoError(t, err)
			defer conn.Release()
			schema, err := postgres.Inspect(conn)
			assert.NoError(t, err)
			diff.Test(t, test.Schema, schema.String())
		})
	}
}
