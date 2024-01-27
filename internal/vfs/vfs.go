package vfs

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sort"

	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/httpfs"
	"golang.org/x/tools/godoc/vfs/mapfs"
)

// FileSystem alias
type FileSystem = vfs.FileSystem

// Map the files to a filesystem
func Map(files map[string]string) vfs.FileSystem {
	return mapfs.New(files)
}

// Write the filesystem to disk
func Write(fs vfs.FileSystem, dir string) error {
	return Walk(fs, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		} else if info.IsDir() {
			return nil
		}
		abspath := filepath.Join(dir, path)
		absdir := filepath.Dir(abspath)
		if err := os.MkdirAll(absdir, 0755); err != nil {
			return err
		}
		f, err := fs.Open(path)
		if err != nil {
			return err
		}
		data, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		if err := ioutil.WriteFile(abspath, data, 0644); err != nil {
			return err
		}
		return nil
	})
}

// Walk fn
// TODO: not a big deal, but it'd be better to transition to vfs.FileSystem
func Walk(fs vfs.FileSystem, walkFn filepath.WalkFunc) error {
	return walk(httpfs.New(fs), string(filepath.Separator), walkFn)
}

// Walk walks the filesystem rooted at root, calling walkFn for each file or
// directory in the filesystem, including root. All errors that arise visiting files
// and directories are filtered by walkFn. The files are walked in lexical
// order.
func walk(fs http.FileSystem, root string, walkFn filepath.WalkFunc) error {
	info, err := stat(fs, root)
	if err != nil {
		return walkFn(root, nil, err)
	}
	return walk2(fs, root, info, walkFn)
}

// walk recursively descends path, calling walkFn.
func walk2(fs http.FileSystem, p string, info os.FileInfo, walkFn filepath.WalkFunc) error {
	err := walkFn(p, info, nil)
	if err != nil {
		if info.IsDir() && err == filepath.SkipDir {
			return nil
		}
		return err
	}

	if !info.IsDir() {
		return nil
	}

	names, err := readDirNames(fs, p)
	if err != nil {
		return walkFn(p, info, err)
	}

	for _, name := range names {
		filename := path.Join(p, name)
		fileInfo, err := stat(fs, filename)
		if err != nil {
			if err := walkFn(filename, fileInfo, err); err != nil && err != filepath.SkipDir {
				return err
			}
		} else {
			err = walk2(fs, filename, fileInfo, walkFn)
			if err != nil {
				if !fileInfo.IsDir() || err != filepath.SkipDir {
					return err
				}
			}
		}
	}
	return nil
}

// readDirNames reads the directory named by dirname and returns
// a sorted list of directory entries.
func readDirNames(fs http.FileSystem, dirname string) ([]string, error) {
	fis, err := readDir(fs, dirname)
	if err != nil {
		return nil, err
	}
	names := make([]string, len(fis))
	for i := range fis {
		names[i] = fis[i].Name()
	}
	sort.Strings(names)
	return names, nil
}

// Stat returns the FileInfo structure describing file.
func stat(fs http.FileSystem, name string) (os.FileInfo, error) {
	f, err := fs.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return f.Stat()
}

// ReadDir reads the contents of the directory associated with file and
// returns a slice of FileInfo values in directory order.
func readDir(fs http.FileSystem, name string) ([]os.FileInfo, error) {
	f, err := fs.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return f.Readdir(0)
}
