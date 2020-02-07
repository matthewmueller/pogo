package postgres_test

import (
	"os"
	"testing"

	"github.com/matthewmueller/pogo/internal/postgres"
	"github.com/tj/assert"
)

func getURL(t testing.TB) string {
	url := os.Getenv("POSTGRES_URL")
	assert.NotEmpty(t, url)
	return url
}

func TestEnv(t *testing.T) {
	getURL(t)
}

func TestPool(t *testing.T) {
	pool, err := postgres.Dial(getURL(t))
	assert.NoError(t, err)
	conn, err := pool.Acquire()
	assert.NoError(t, err)
	conn.Release()
	pool.Close()
}
