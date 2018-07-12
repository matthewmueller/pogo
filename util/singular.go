package util

import (
	"strings"

	gen "github.com/matthewmueller/go-gen"
)

// Singular makes each segment singular
func Singular(s string) string {
	segments := strings.Split(gen.Snake(s), "_")
	var singles []string
	for _, segment := range segments {
		singles = append(singles, gen.Pascal(gen.Singular(segment)))
	}
	return strings.Join(singles, "")
}
