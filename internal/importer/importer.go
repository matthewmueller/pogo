package importer

import (
	"go/build"
	"path/filepath"
)

// New importer
func New(templatePath string) *Importer {
	return &Importer{
		templatePath,
	}
}

// Importer struct
type Importer struct {
	templatePath string
}

// Import gets the import path based on the output path
func (i *Importer) Import(paths ...string) (string, error) {
	absDir, err := filepath.Abs(i.templatePath)
	if err != nil {
		return "", err
	}

	// get the gosrc
	gopath := build.Default.GOPATH
	gosrc := filepath.Join(gopath, "src")

	// get the path relative to $GOPATH/src
	// TODO: this probably doesn't work with go mod
	relDir, err := filepath.Rel(gosrc, absDir)
	if err != nil {
		return "", err
	}

	// join all teh paths together
	paths = append([]string{relDir}, paths...)
	return filepath.Join(paths...), nil
}
