package template

import (
	"bytes"
	"text/template"
)

// Map is just an alias
type Map map[string]interface{}

// MustCompile with panic if there's a compilation error
func MustCompile(name, text string) func(data interface{}) (string, error) {
	t, err := Compile(name, text)
	if err != nil {
		panic(err)
	}
	return t
}

// Compile the template into a function
func Compile(name, text string) (func(data interface{}) (string, error), error) {
	template, err := template.New(name).Parse(text)
	if err != nil {
		panic(err)
	}
	return func(data interface{}) (string, error) {
		var b bytes.Buffer
		if err := template.Execute(&b, data); err != nil {
			return "", err
		}
		return b.String(), nil
	}, nil
}
