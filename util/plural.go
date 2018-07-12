package util

import (
	"strings"

	gen "github.com/matthewmueller/go-gen"
)

// Plural makes each segment singular
func Plural(s string) string {
	segments := strings.Split(gen.Snake(s), "_")
	var singles []string
	for _, segment := range segments {
		singles = append(singles, gen.Pascal(gen.Plural(segment)))
	}
	return strings.Join(singles, "")
}
