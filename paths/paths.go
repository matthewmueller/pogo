package paths

import (
	"errors"
	"go/build"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	text "github.com/matthewmueller/go-text"
)

// Load the path
func Load(base string) (*Path, error) {
	var path Path
	var err error

	// get the cwd
	if path.cwd, err = os.Getwd(); err != nil {
		return nil, err
	}

	// get the project root
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("couldn't get the current file")
	}
	path.root = filepath.Join(file, "..", "..")

	// get the gosrc
	gopath := build.Default.GOPATH
	path.gosrc = filepath.Join(gopath, "src")

	// get the absolute path
	if path.abs, err = filepath.Abs(base); err != nil {
		return nil, err
	}

	// relpath
	if path.rel, err = filepath.Rel(path.cwd, path.abs); err != nil {
		return nil, err
	}

	// import path
	if path.imp, err = filepath.Rel(path.gosrc, path.abs); err != nil {
		return nil, err
	}

	return &path, nil
}

// Path struct
type Path struct {
	cwd   string
	root  string
	gosrc string

	abs string
	rel string
	imp string
}

// Package gets the path's package name
func (p *Path) Package() string {
	return text.Lower(text.Camel(strings.TrimPrefix(filepath.Base(p.abs), "go-")))
}

// New relative subpath
func (p *Path) New(relpath string) (*Path, error) {
	var path Path
	var err error

	// get the absolute path
	path.abs = filepath.Join(p.abs, relpath)

	// update the relative path
	if path.rel, err = filepath.Rel(p.cwd, path.abs); err != nil {
		return nil, err
	}

	// update the import path
	if path.imp, err = filepath.Rel(p.gosrc, path.abs); err != nil {
		return nil, err
	}

	return &path, nil
}

// Root gets the gumbo's root
func (p *Path) Root() string {
	return p.root
}

// Cwd gets the current working directory
func (p *Path) Cwd() string {
	return p.cwd
}

// Import gets the import path
func (p *Path) Import(paths ...string) string {
	paths = append([]string{p.imp}, paths...)
	return filepath.Join(paths...)
}

// Rel gets the relative path
func (p *Path) Rel(paths ...string) string {
	paths = append([]string{p.rel}, paths...)
	return filepath.Join(paths...)
}

// Abs gets the absolute path
func (p *Path) Abs(paths ...string) string {
	paths = append([]string{p.abs}, paths...)
	return filepath.Join(paths...)
}
