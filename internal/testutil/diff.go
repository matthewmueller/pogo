package testutil

import (
	"bytes"
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
)

// Diff fn
func Diff(t, a interface{}, b interface{}) {

}

func diff(expected, actual string) string {

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(expected, actual, false)

	var buf bytes.Buffer
	for _, diff := range diffs {
		switch diff.Type {
		case diffmatchpatch.DiffInsert:
			buf.WriteString("\x1b[102m\x1b[30m")
			buf.WriteString(diff.Text)
			buf.WriteString("\x1b[0m")
		case diffmatchpatch.DiffDelete:
			buf.WriteString("\x1b[101m\x1b[30m")
			buf.WriteString(diff.Text)
			buf.WriteString("\x1b[0m")
		case diffmatchpatch.DiffEqual:
			buf.WriteString(diff.Text)
		}
	}

	result := buf.String()
	result = strings.Replace(result, "\\n", "\n", -1)
	result = strings.Replace(result, "\\t", "\t", -1)
	return result
}
