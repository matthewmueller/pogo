package model

import (
	"strconv"

	"github.com/matthewmueller/pgx"
)

// DB is the common interface for database operations that can be used with
// types from schema 'jack'.
//
// This should work with database/sql.DB and database/sql.Tx.
type DB interface {
	Exec(string, ...interface{}) (pgx.CommandTag, error)
	Query(string, ...interface{}) (*pgx.Rows, error)
	QueryRow(string, ...interface{}) *pgx.Row
}

// DBLog provides the log func used by generated queries.
var DBLog = func(string, ...interface{}) {}

func querySlices(fields map[string]interface{}, offset int) (c []string, i []string, v []interface{}) {
	n := offset + 1
	for col, val := range fields {
		c = append(c, col)
		i = append(i, "$"+strconv.Itoa(n))
		v = append(v, val)
		n++
	}
	return c, i, v
}

// // StringSlice is a slice of strings.
// type StringSlice []string
//
// // quoteEscapeRegex is the regex to match escaped characters in a string.
// var quoteEscapeRegex = regexp.MustCompile(`([^\\]([\\]{2})*)\\"`)
//
// // Scan satisfies the sql.Scanner interface for StringSlice.
// func (ss *StringSlice) Scan(src interface{}) error {
// 	buf, ok := src.([]byte)
// 	if !ok {
// 		return errors.New("invalid StringSlice")
// 	}
//
// 	// change quote escapes for csv parser
// 	str := quoteEscapeRegex.ReplaceAllString(string(buf), `$1""`)
// 	str = strings.Replace(str, `\\`, `\`, -1)
//
// 	// remove braces
// 	str = str[1 : len(str)-1]
//
// 	// bail if only one
// 	if len(str) == 0 {
// 		*ss = StringSlice([]string{})
// 		return nil
// 	}
//
// 	// parse with csv reader
// 	cr := csv.NewReader(strings.NewReader(str))
// 	slice, err := cr.Read()
// 	if err != nil {
// 		fmt.Printf("exiting!: %v\n", err)
// 		return err
// 	}
//
// 	*ss = StringSlice(slice)
//
// 	return nil
// }
//
// // Value satisfies the driver.Valuer interface for StringSlice.
// func (ss StringSlice) Value() (driver.Value, error) {
// 	v := make([]string, len(ss))
// 	for i, s := range ss {
// 		v[i] = `"` + strings.Replace(strings.Replace(s, `\`, `\\\`, -1), `"`, `\"`, -1) + `"`
// 	}
// 	return "{" + strings.Join(v, ",") + "}", nil
// }
//
// type ScannerValuer interface {
// 	sql.Scanner
// 	driver.Valuer
// }
//
// // Slice is a slice of ScannerValuers.
// type Slice []ScannerValuer
