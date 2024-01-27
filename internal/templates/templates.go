package templates

import (
	"embed"
	"io/fs"
)

//go:embed *.gotext
var fsys embed.FS

func AssetString(name string) (string, error) {
	code, err := fs.ReadFile(fsys, name)
	if err != nil {
		return "", err
	}
	return string(code), nil
}

func MustAssetString(name string) string {
	code, err := AssetString(name)
	if err != nil {
		panic(err)
	}
	return string(code)
}
