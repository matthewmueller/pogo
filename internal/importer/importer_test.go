package importer_test

import (
	"testing"

	"github.com/matthewmueller/pogo/internal/importer"
	"github.com/tj/assert"
)

func TestPath(t *testing.T) {
	imp := importer.New("./user/user.go")
	path, err := imp.Import("../../pogo/pogo.go")
	assert.NoError(t, err)
	assert.Contains(t, path, "github.com")
	assert.Contains(t, path, "pogo/internal/importer/pogo/pogo.go")
}

func TestDir(t *testing.T) {
	imp := importer.New("./")
	path, err := imp.Import("./pogo/pogo.go")
	assert.NoError(t, err)
	assert.Contains(t, path, "github.com")
	assert.Contains(t, path, "pogo/internal/importer/pogo/pogo.go")
}
