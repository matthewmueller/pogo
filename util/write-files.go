package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// WriteFiles fn
func WriteFiles(files map[string]string) error {
	// range over the files and create them
	for path, data := range files {
		dir := filepath.Dir(path)

		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}

		if err := ioutil.WriteFile(path, []byte(data), 0644); err != nil {
			return err
		}
	}
	return nil
}
