package testutil

import (
	"testing"

	"github.com/matthewmueller/go-text"
	"github.com/matthewmueller/pogo/internal/tests"
	"github.com/tj/assert"
)

// Name of the test
func Name(t testing.TB, test *tests.Test) (name string) {
	switch {
	case test.Title != "":
		return text.Lower(text.Snake(test.Title))
	case test.QueryGo != "":
		return text.Lower(text.Snake(Truncate(test.QueryGo, 80)))
	case test.QueryTS != "":
		return text.Lower(text.Snake(Truncate(test.QueryTS, 80)))
	default:
		assert.Fail(t, "no test name found")
	}
	return name
}
