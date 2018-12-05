package testutil

import (
	"bytes"
	"go/build"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	gen "github.com/matthewmueller/go-gen"
	"github.com/tj/assert"
)

// GoRun write a main.go file runs it, returning the result
func GoRun(t testing.TB, path, main string) (string, string, func()) {
	dir := filepath.Dir(path)

	code, err := gen.Format(main)
	assert.NoError(t, err)

	// ensure all the imports are coming this package
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, code, parser.ImportsOnly)
	assert.NoError(t, err)
	for _, imp := range f.Imports {
		path, err := strconv.Unquote(imp.Path.Value)
		assert.NoError(t, err)
		if !strings.Contains(path, "pogo") {
			continue
		}
		if !strings.Contains(path, filepath.Base(dir)) {
			t.Fatalf("path imported is not in the test directory: %s", path)
		}
	}

	err = os.MkdirAll(dir, 0755)
	assert.NoError(t, err)

	err = ioutil.WriteFile(path, []byte(code), 0644)
	assert.NoError(t, err)

	gobin, err := exec.LookPath("go")
	assert.NoError(t, err)

	// go run
	cmd := exec.Command(gobin, "run", path)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	// run the command
	if err := cmd.Run(); err != nil {
		t.Fatal(stderr.String())
	}

	return stdout.String(), stderr.String(), func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Fatal(err)
		}
	}
}

// GoImport fn
func GoImport(t testing.TB, abspath string) func(s string) string {
	gopath := build.Default.GOPATH
	importBase, err := filepath.Rel(filepath.Join(gopath, "src"), abspath)
	if err != nil {
		t.Fatal(err)
	}
	return func(path string) string {
		_, err := os.Stat(filepath.Join(abspath, path))
		if os.IsNotExist(err) {
			return ""
		} else if err != nil {
			t.Fatal(err)
		}

		return strconv.Quote(filepath.Join(importBase, path))
	}
}
