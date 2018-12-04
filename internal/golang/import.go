package golang

import (
	"go/build"
	"os"
	"path/filepath"
)

var importBase string

func init() {
	gopath := build.Default.GOPATH
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	importBase, err = filepath.Rel(filepath.Join(gopath, "src"), cwd)
	if err != nil {
		panic(err)
	}
}

// Import path
func Import(path string) string {
	return filepath.Join(importBase, path)
}
