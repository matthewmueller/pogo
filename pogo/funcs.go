package pogo

import (
	"sort"
	"strconv"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"

	"github.com/gedex/inflector"
	"github.com/knq/snaker"
	"github.com/matthewmueller/pogo/postgres"
)

var shortnameMap map[string]string

func init() {
	shortnameMap = map[string]string{}
}

// TemplateFuncs struct
type TemplateFuncs struct {
	Coerce *Coerce
}

// TemplateFunctions functions to pass in
func TemplateFunctions(coerce *Coerce) template.FuncMap {
	funcs := TemplateFuncs{
		Coerce: coerce,
	}

	return template.FuncMap{
		"shortname":      funcs.shortname,
		"classname":      funcs.classname,
		"classnameMM":    funcs.classnameMM,
		"modelname":      funcs.modelname,
		"modelnameMM":    funcs.modelnameMM,
		"modelreturn":    funcs.modelreturn,
		"modelreturnMM":  funcs.modelreturnMM,
		"field":          funcs.field,
		"param":          funcs.param,
		"fieldtype":      funcs.fieldtype,
		"schema":         funcs.schema,
		"fields":         funcs.fields,
		"gofields":       funcs.gofields,
		"primaryname":    funcs.primaryname,
		"primarytype":    funcs.primarytype,
		"primaryid":      funcs.primaryid,
		"fkparams":       funcs.fkparams,
		"fklist":         funcs.fklist,
		"fkwhere":        funcs.fkwhere,
		"fklength":       funcs.fklength,
		"indexmethod":    funcs.indexmethod,
		"indexparams":    funcs.indexparams,
		"indexvars":      funcs.indexvars,
		"indexparam":     funcs.indexparam,
		"indexparamlist": funcs.indexparamlist,
		"indexwhere":     funcs.indexwhere,
		"indexlength":    funcs.indexlength,
		// "indexreturning": funcs.indexreturning,
		// "indexscan":      funcs.indexscan,
	}
}

// reverseIndexRune finds the last rune r in s, returning -1 if not present.
func reverseIndexRune(s string, r rune) int {
	if s == "" {
		return -1
	}

	rs := []rune(s)
	for i := len(rs) - 1; i >= 0; i-- {
		if rs[i] == r {
			return i
		}
	}

	return -1
}

func (f *TemplateFuncs) classname(s string) string {
	return snaker.SnakeToCamelIdentifier(s)
}

func (f *TemplateFuncs) classnameMM(s string) string {
	if i := reverseIndexRune(s, '_'); i != -1 {
		s = inflector.Singularize(s[:i]) + "_" + s[i+1:]
	} else {
		s = inflector.Singularize(s)
	}

	return snaker.SnakeToCamelIdentifier(s)
}

func (f *TemplateFuncs) modelname(s string) string {
	if i := reverseIndexRune(s, '_'); i != -1 {
		s = s[:i] + "_" + inflector.Singularize(s[i+1:])
	} else {
		s = inflector.Singularize(s)
	}

	return snaker.SnakeToCamelIdentifier(s)
}

func (f *TemplateFuncs) modelnameMM(s string) string {
	if i := reverseIndexRune(s, '_'); i != -1 {
		s = inflector.Singularize(s[:i]) + "_" + inflector.Singularize(s[i+1:])
	} else {
		s = inflector.Singularize(s)
	}

	return snaker.SnakeToCamelIdentifier(s)
}

func (f *TemplateFuncs) modelreturn(s string) string {
	return strings.ToLower(f.modelname(s))
}

func (f *TemplateFuncs) modelreturnMM(s string) string {
	return strings.ToLower(f.modelnameMM(s))
}

func (f *TemplateFuncs) field(name string) string {
	return snaker.SnakeToCamelIdentifier(name)
}

func (f *TemplateFuncs) param(name string) string {
	s := snaker.SnakeToCamelIdentifier(name)
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

func (f *TemplateFuncs) fieldtype(kind string) string {
	return f.Coerce.Coerce(kind)
	// fmt.Println(name, t)
	// return snaker.SnakeToCamelIdentifier(name) + " " + t + "`json:\"" + name + ",omitempty\"`"
}

// schema takes a series of names and joins them with the schema name.
func (f *TemplateFuncs) schema(s string, names ...string) string {
	n := strings.Join(names, ".")

	if s == "" && n == "" {
		return ""
	}

	if s != "" && n != "" {
		s = s + "."
	}

	return s + n
}

// shortname generates a safe Go identifier for typ. typ is first checked
// against ArgType.ShortNameTypeMap, and if not found, then the value is
// calculated and stored in the ShortNameTypeMap for future use.
//
// A shortname is the concatentation of the lowercase of the first character in
// the words comprising the name. For example, "MyCustomName" will have have
// the shortname of "mcn".
//
// If a generated shortname conflicts with a Go reserved name, then the
// corresponding value in goReservedNames map will be used.
//
// Generated shortnames that have conflicts with any scopeConflicts member will
// have ArgType.NameConflictSuffix appended.
//
// Note: recognized types for scopeConflicts are string, []*Field,
// []*QueryParam.

func (f *TemplateFuncs) shortname(typ string, scopeConflicts ...interface{}) string {
	var v string
	var ok bool

	// check short name map
	if v, ok = shortnameMap[typ]; !ok {
		// calc the short name
		u := []string{}
		for _, s := range strings.Split(strings.ToLower(snaker.CamelToSnake(typ)), "_") {
			if len(s) > 0 && s != "id" {
				u = append(u, s[:1])
			}
		}
		v = strings.Join(u, "")

		// check go reserved names
		if n, ok := goReservedNames[v]; ok {
			v = n
		}

		// store back to short name map
		shortnameMap[typ] = v
	}

	// initial conflicts are the default imported packages from
	// xo_package.go.tpl
	conflicts := map[string]bool{
		"sql":     true,
		"driver":  true,
		"csv":     true,
		"errors":  true,
		"fmt":     true,
		"regexp":  true,
		"strings": true,
		"time":    true,
	}

	// add scopeConflicts to conflicts
	for _, c := range scopeConflicts {
		switch k := c.(type) {
		case string:
			conflicts[k] = true

		// case []*Field:
		// 	for _, f := range k {
		// 		conflicts[f.Name] = true
		// 	}
		// case []*QueryParam:
		// 	for _, f := range k {
		// 		conflicts[f.Name] = true
		// 	}

		default:
			panic("not implemented")
		}
	}

	// append suffix if conflict exists
	if _, ok := conflicts[v]; ok {
		v = v + v
	}

	return v
}

func (f *TemplateFuncs) fields(columns []*postgres.Column) string {
	cols := []string{}
	for _, col := range columns {
		cols = append(cols, col.ColumnName)
	}
	return "\"" + strings.Join(cols, "\", \"") + "\""
}

func (f *TemplateFuncs) gofields(columns []*postgres.Column, key string) string {
	cols := []string{}
	for _, col := range columns {
		cols = append(cols, "&"+key+"."+f.field(col.ColumnName))
	}
	return strings.Join(cols, ", ")
}

func (f *TemplateFuncs) indexmethod(index *Index) string {
	cols := []string{}

	for _, col := range index.Columns {
		if index.IsUnique && !index.IsPrimary {
			cols = append(cols, f.field(col.ColumnName))
		}
	}

	sort.Strings(cols)
	return strings.Join(cols, "And")
}

func (f *TemplateFuncs) indexparams(index *Index) string {
	cols := []string{}

	for _, col := range index.Columns {
		if index.IsUnique && !index.IsPrimary {
			cols = append(cols, strings.ToLower(f.field(col.ColumnName))+" "+f.Coerce.Coerce(col.DataType))
		}
	}

	sort.Strings(cols)
	return strings.Join(cols, ", ")
}

func (f *TemplateFuncs) indexvars(index *Index) string {
	cols := []string{}

	for _, col := range index.Columns {
		if index.IsUnique && !index.IsPrimary {
			cols = append(cols, strings.ToLower(f.field(col.ColumnName)))
		}
	}

	sort.Strings(cols)
	return strings.Join(cols, ", ")
}

func (f *TemplateFuncs) indexparam(colname string) string {
	return strings.ToLower(f.field(colname))
}

func (f *TemplateFuncs) indexparamlist(index *Index) string {
	cols := []string{}

	for _, col := range index.Columns {
		if index.IsUnique && !index.IsPrimary {
			cols = append(cols, strings.ToLower(f.field(col.ColumnName)))
		}
	}

	sort.Strings(cols)
	return `"` + strings.Join(cols, "\", \"") + `"`
}

func (f *TemplateFuncs) indexwhere(index *Index) string {
	cols := []string{}
	i := 1

	for _, col := range index.Columns {
		if index.IsUnique && !index.IsPrimary {
			cols = append(cols, "\""+col.ColumnName+"\" = $"+strconv.Itoa(i))
			i++
		}
	}

	return strings.Join(cols, " AND ")
}

// func (f *TemplateFuncs) indexreturning(index *Index) string {
// 	cols := []string{}
// 	for _, col := range index.Columns {
// 		cols = append(cols, col.ColumnName)
// 	}
// 	return "\"" + strings.Join(cols, "\", \"") + "\""
// }
//
// func (f *TemplateFuncs) indexscan(columns []*postgres.IndexColumn, key string) string {
// 	cols := []string{}
// 	for _, col := range columns {
// 		cols = append(cols, "&"+key+"."+f.field(col.ColumnName))
// 	}
// 	return strings.Join(cols, ", ")
// }

func (f *TemplateFuncs) indexlength(index *Index) string {
	i := 0

	for range index.Columns {
		if index.IsUnique && !index.IsPrimary {
			i++
		}
	}

	return strconv.Itoa(i)
}

func (f *TemplateFuncs) primaryname(columns []*postgres.Column) string {
	for _, col := range columns {
		if col.IsPrimaryKey {
			return col.ColumnName
		}
	}
	return ""
}

func (f *TemplateFuncs) primarytype(columns []*postgres.Column) string {
	for _, col := range columns {
		if col.IsPrimaryKey {
			return f.Coerce.Coerce(col.DataType)
		}
	}
	return ""
}

func (f *TemplateFuncs) primaryid(columns []*postgres.Column) string {
	for _, col := range columns {
		if col.IsPrimaryKey {
			return f.field(col.ColumnName)
		}
	}
	return ""
}

func (f *TemplateFuncs) fkparams(fks []*postgres.ForeignKey, cols []*postgres.Column) string {
	out := []string{}
	for _, fk := range fks {
		colname := fk.ColumnName
		for _, col := range cols {
			if colname == col.ColumnName {
				out = append(out, f.field(fk.ColumnName)+" "+f.Coerce.Coerce(col.DataType))
			}
		}
	}
	return strings.Join(out, ", ")
}

func (f *TemplateFuncs) fklist(fks []*postgres.ForeignKey) string {
	out := []string{}
	for _, fk := range fks {
		out = append(out, f.field(fk.ColumnName))
	}
	return strings.Join(out, ", ")
}

func (f *TemplateFuncs) fklength(fks []*postgres.ForeignKey) string {
	i := 0

	for range fks {
		i++
	}

	return strconv.Itoa(i)
}

func (f *TemplateFuncs) fkwhere(fks []*postgres.ForeignKey) string {
	out := []string{}
	for i, fk := range fks {
		out = append(out, "\""+fk.ColumnName+"\" = $"+strconv.Itoa(i+1))
	}
	return strings.Join(out, " AND ")
}

// func fields ()

// SinguralizeIdentifier will singularize a identifier, returning it in
// CamelCase.
// func govar(s string) string {
// 	if i := reverseIndexRune(s, '_'); i != -1 {
// 		s = s[:i] + "_" + inflector.Singularize(s[i+1:])
// 	} else {
// 		s = inflector.Singularize(s)
// 	}
//
// 	return snaker.SnakeToCamelIdentifier(s)
// }

//
// // retype checks typ against known types, and prefixing
// // ArgType.CustomTypePackage (if applicable).
// func (a *ArgType) retype(typ string) string {
// 	if strings.Contains(typ, ".") {
// 		return typ
// 	}
//
// 	prefix := ""
// 	for strings.HasPrefix(typ, "[]") {
// 		typ = typ[2:]
// 		prefix = prefix + "[]"
// 	}
//
// 	if _, ok := a.KnownTypeMap[typ]; !ok {
// 		pkg := a.CustomTypePackage
// 		if pkg != "" {
// 			pkg = pkg + "."
// 		}
//
// 		return prefix + pkg + typ
// 	}
//
// 	return prefix + typ
// }
//
// // reniltype checks typ against known nil types (similar to retype), prefixing
// // ArgType.CustomTypePackage (if applicable).
// func (a *ArgType) reniltype(typ string) string {
// 	if strings.Contains(typ, ".") {
// 		return typ
// 	}
//
// 	if strings.HasSuffix(typ, "{}") {
// 		if _, ok := a.KnownTypeMap[typ[:len(typ)-2]]; ok {
// 			return typ
// 		}
//
// 		pkg := a.CustomTypePackage
// 		if pkg != "" {
// 			pkg = pkg + "."
// 		}
//
// 		return pkg + typ
// 	}
//
// 	return typ
// }
//
// // shortname generates a safe Go identifier for typ. typ is first checked
// // against ArgType.ShortNameTypeMap, and if not found, then the value is
// // calculated and stored in the ShortNameTypeMap for future use.
// //
// // A shortname is the concatentation of the lowercase of the first character in
// // the words comprising the name. For example, "MyCustomName" will have have
// // the shortname of "mcn".
// //
// // If a generated shortname conflicts with a Go reserved name, then the
// // corresponding value in goReservedNames map will be used.
// //
// // Generated shortnames that have conflicts with any scopeConflicts member will
// // have ArgType.NameConflictSuffix appended.
// //
// // Note: recognized types for scopeConflicts are string, []*Field,
// // []*QueryParam.
// func (a *ArgType) shortname(typ string, scopeConflicts ...interface{}) string {
// 	var v string
// 	var ok bool
//
// 	// check short name map
// 	if v, ok = a.ShortNameTypeMap[typ]; !ok {
// 		// calc the short name
// 		u := []string{}
// 		for _, s := range strings.Split(strings.ToLower(snaker.CamelToSnake(typ)), "_") {
// 			if len(s) > 0 && s != "id" {
// 				u = append(u, s[:1])
// 			}
// 		}
// 		v = strings.Join(u, "")
//
// 		// check go reserved names
// 		if n, ok := goReservedNames[v]; ok {
// 			v = n
// 		}
//
// 		// store back to short name map
// 		a.ShortNameTypeMap[typ] = v
// 	}
//
// 	// initial conflicts are the default imported packages from
// 	// xo_package.go.tpl
// 	conflicts := map[string]bool{
// 		"sql":     true,
// 		"driver":  true,
// 		"csv":     true,
// 		"errors":  true,
// 		"fmt":     true,
// 		"regexp":  true,
// 		"strings": true,
// 		"time":    true,
// 	}
//
// 	// add scopeConflicts to conflicts
// 	for _, c := range scopeConflicts {
// 		switch k := c.(type) {
// 		case string:
// 			conflicts[k] = true
//
// 		case []*Field:
// 			for _, f := range k {
// 				conflicts[f.Name] = true
// 			}
// 		case []*QueryParam:
// 			for _, f := range k {
// 				conflicts[f.Name] = true
// 			}
//
// 		default:
// 			panic("not implemented")
// 		}
// 	}
//
// 	// append suffix if conflict exists
// 	if _, ok := conflicts[v]; ok {
// 		v = v + a.NameConflictSuffix
// 	}
//
// 	return v
// }
//
// // colnames creates a list of the column names found in fields, excluding any
// // Field with Name contained in ignoreNames.
// //
// // Used to present a comma separated list of column names, that can be used in
// // a SELECT, or UPDATE, or other SQL clause requiring an list of identifiers
// // (ie, "field_1, field_2, field_3, ...").
// func (a *ArgType) colnames(fields []*Field, ignoreNames ...string) string {
// 	ignore := map[string]bool{}
// 	for _, n := range ignoreNames {
// 		ignore[n] = true
// 	}
//
// 	str := ""
// 	i := 0
// 	for _, f := range fields {
// 		if ignore[f.Name] {
// 			continue
// 		}
//
// 		if i != 0 {
// 			str = str + ", "
// 		}
// 		str = str + a.colname(f.Col)
// 		i++
// 	}
//
// 	return str
// }
//
// // colnamesquery creates a list of the column names in fields as a query and
// // joined by sep, excluding any Field with Name contained in ignoreNames.
// //
// // Used to create a list of column names in a WHERE clause (ie, "field_1 = $1
// // AND field_2 = $2 AND ...") or in an UPDATE clause (ie, "field = $1, field =
// // $2, ...").
// func (a *ArgType) colnamesquery(fields []*Field, sep string, ignoreNames ...string) string {
// 	ignore := map[string]bool{}
// 	for _, n := range ignoreNames {
// 		ignore[n] = true
// 	}
//
// 	str := ""
// 	i := 0
// 	for _, f := range fields {
// 		if ignore[f.Name] {
// 			continue
// 		}
//
// 		if i != 0 {
// 			str = str + sep
// 		}
// 		str = str + a.colname(f.Col) + " = " + a.Loader.NthParam(i)
// 		i++
// 	}
//
// 	return str
// }
//
// // colprefixnames creates a list of the column names found in fields with the
// // supplied prefix, excluding any Field with Name contained in ignoreNames.
// //
// // Used to present a comma separated list of column names with a prefix. Used in
// // a SELECT, or UPDATE (ie, "t.field_1, t.field_2, t.field_3, ...").
// func (a *ArgType) colprefixnames(fields []*Field, prefix string, ignoreNames ...string) string {
// 	ignore := map[string]bool{}
// 	for _, n := range ignoreNames {
// 		ignore[n] = true
// 	}
//
// 	str := ""
// 	i := 0
// 	for _, f := range fields {
// 		if ignore[f.Name] {
// 			continue
// 		}
//
// 		if i != 0 {
// 			str = str + ", "
// 		}
// 		str = str + prefix + "." + a.colname(f.Col)
// 		i++
// 	}
//
// 	return str
// }
//
// // colvals creates a list of value place holders for fields excluding any Field
// // with Name contained in ignoreNames.
// //
// // Used to present a comma separated list of column place holders, used in a
// // SELECT or UPDATE statement (ie, "$1, $2, $3 ...").
// func (a *ArgType) colvals(fields []*Field, ignoreNames ...string) string {
// 	ignore := map[string]bool{}
// 	for _, n := range ignoreNames {
// 		ignore[n] = true
// 	}
//
// 	str := ""
// 	i := 0
// 	for _, f := range fields {
// 		if ignore[f.Name] {
// 			continue
// 		}
//
// 		if i != 0 {
// 			str = str + ", "
// 		}
// 		str = str + a.Loader.NthParam(i)
// 		i++
// 	}
//
// 	return str
// }
//
// // fieldnames creates a list of field names from fields of the adding the
// // provided prefix, and excluding any Field with Name contained in ignoreNames.
// //
// // Used to present a comma separated list of field names, ie in a Go statement
// // (ie, "t.Field1, t.Field2, t.Field3 ...")
// func (a *ArgType) fieldnames(fields []*Field, prefix string, ignoreNames ...string) string {
// 	ignore := map[string]bool{}
// 	for _, n := range ignoreNames {
// 		ignore[n] = true
// 	}
//
// 	str := ""
// 	i := 0
// 	for _, f := range fields {
// 		if ignore[f.Name] {
// 			continue
// 		}
//
// 		if i != 0 {
// 			str = str + ", "
// 		}
// 		str = str + prefix + "." + f.Name
// 		i++
// 	}
//
// 	return str
// }
//
// // colcount returns the 1-based count of fields, excluding any Field with Name
// // contained in ignoreNames.
// //
// // Used to get the count of fields, and useful for specifying the last SQL
// // parameter.
// func (a *ArgType) colcount(fields []*Field, ignoreNames ...string) int {
// 	ignore := map[string]bool{}
// 	for _, n := range ignoreNames {
// 		ignore[n] = true
// 	}
//
// 	i := 1
// 	for _, f := range fields {
// 		if ignore[f.Name] {
// 			continue
// 		}
//
// 		i++
// 	}
// 	return i
// }
//
// // goReservedNames is a map of of go reserved names to "safe" names.
// var goReservedNames = map[string]string{
// 	"break":       "brk",
// 	"case":        "cs",
// 	"chan":        "chn",
// 	"const":       "cnst",
// 	"continue":    "cnt",
// 	"default":     "def",
// 	"defer":       "dfr",
// 	"else":        "els",
// 	"fallthrough": "flthrough",
// 	"for":         "fr",
// 	"func":        "fn",
// 	"go":          "goVal",
// 	"goto":        "gt",
// 	"if":          "ifVal",
// 	"import":      "imp",
// 	"interface":   "iface",
// 	"map":         "mp",
// 	"package":     "pkg",
// 	"range":       "rnge",
// 	"return":      "ret",
// 	"select":      "slct",
// 	"struct":      "strct",
// 	"switch":      "swtch",
// 	"type":        "typ",
// 	"var":         "vr",
//
// 	// go types
// 	"error":      "e",
// 	"bool":       "b",
// 	"string":     "str",
// 	"byte":       "byt",
// 	"rune":       "r",
// 	"uintptr":    "uptr",
// 	"int":        "i",
// 	"int8":       "i8",
// 	"int16":      "i16",
// 	"int32":      "i32",
// 	"int64":      "i64",
// 	"uint":       "u",
// 	"uint8":      "u8",
// 	"uint16":     "u16",
// 	"uint32":     "u32",
// 	"uint64":     "u64",
// 	"float32":    "z",
// 	"float64":    "f",
// 	"complex64":  "c",
// 	"complex128": "c128",
// }
//
// // goparamlist converts a list of fields into their named Go parameters,
// // skipping any Field with Name contained in ignoreNames. addType will cause
// // the go Type to be added after each variable name. addPrefix will cause the
// // returned string to be prefixed with ", " if the generated string is not
// // empty.
// //
// // Any field name encountered will be checked against goReservedNames, and will
// // have its name substituted by its corresponding looked up value.
// //
// // Used to present a comma separated list of Go variable names for use with as
// // either a Go func parameter list, or in a call to another Go func.
// // (ie, ", a, b, c, ..." or ", a T1, b T2, c T3, ...").
// func (a *ArgType) goparamlist(fields []*Field, addPrefix bool, addType bool, ignoreNames ...string) string {
// 	ignore := map[string]bool{}
// 	for _, n := range ignoreNames {
// 		ignore[n] = true
// 	}
//
// 	i := 0
// 	vals := []string{}
// 	for _, f := range fields {
// 		if ignore[f.Name] {
// 			continue
// 		}
//
// 		s := "v" + strconv.Itoa(i)
// 		if len(f.Name) > 0 {
// 			n := strings.Split(snaker.CamelToSnake(f.Name), "_")
// 			s = strings.ToLower(n[0]) + f.Name[len(n[0]):]
// 		}
//
// 		// check go reserved names
// 		if r, ok := goReservedNames[strings.ToLower(s)]; ok {
// 			s = r
// 		}
//
// 		// add the go type
// 		if addType {
// 			s += " " + a.retype(f.Type)
// 		}
//
// 		// add to vals
// 		vals = append(vals, s)
//
// 		i++
// 	}
//
// 	// concat generated values
// 	str := strings.Join(vals, ", ")
// 	if addPrefix && str != "" {
// 		return ", " + str
// 	}
//
// 	return str
// }
//
// // convext generates the Go conversion for f in order for it to be assignable
// // to t.
// //
// // FIXME: this should be a better name, like "goconversion" or some such.
// func (a *ArgType) convext(prefix string, f *Field, t *Field) string {
// 	expr := prefix + "." + f.Name
// 	if f.Type == t.Type {
// 		return expr
// 	}
//
// 	ft := f.Type
// 	if strings.HasPrefix(ft, "sql.Null") {
// 		expr = expr + "." + f.Type[8:]
// 		ft = strings.ToLower(f.Type[8:])
// 	}
//
// 	if t.Type != ft {
// 		expr = t.Type + "(" + expr + ")"
// 	}
//
// 	return expr
// }
//
// // schemafn takes a series of names and joins them with the schema name.
// func (a *ArgType) schemafn(s string, names ...string) string {
// 	// escape table names
// 	if a.EscapeTableNames {
// 		for i, t := range names {
// 			names[i] = a.Loader.Escape(TableEsc, t)
// 		}
// 	}
//
// 	n := strings.Join(names, ".")
//
// 	if s == "" && n == "" {
// 		return ""
// 	}
//
// 	if s != "" && n != "" {
// 		if a.EscapeSchemaName {
// 			s = a.Loader.Escape(SchemaEsc, s)
// 		}
// 		s = s + "."
// 	}
//
// 	return s + n
// }
//
// // colname returns the ColumnName of col, optionally escaping it if
// // ArgType.EscapeColumnNames is toggled.
// func (a *ArgType) colname(col *postgres.Column) string {
// 	if a.EscapeColumnNames {
// 		return a.Loader.Escape(ColumnEsc, col.ColumnName)
// 	}
//
// 	return col.ColumnName
// }
//
// // hascolumn takes a list of fields and determines if field with the specified
// // column name is in the list.
// func (a *ArgType) hascolumn(fields []*Field, name string) bool {
// 	for _, f := range fields {
// 		if f.Col.ColumnName == name {
// 			return true
// 		}
// 	}
//
// 	return false
// }
//
// // hasfield takes a list of fields and determines if field with the specified
// // field name is in the list.
// func (a *ArgType) hasfield(fields []*Field, name string) bool {
// 	for _, f := range fields {
// 		if f.Name == name {
// 			return true
// 		}
// 	}
//
// 	return false
// }

// goReservedNames is a map of of go reserved names to "safe" names.
var goReservedNames = map[string]string{
	"break":       "brk",
	"case":        "cs",
	"chan":        "chn",
	"const":       "cnst",
	"continue":    "cnt",
	"default":     "def",
	"defer":       "dfr",
	"else":        "els",
	"fallthrough": "flthrough",
	"for":         "fr",
	"func":        "fn",
	"go":          "goVal",
	"goto":        "gt",
	"if":          "ifVal",
	"import":      "imp",
	"interface":   "iface",
	"map":         "mp",
	"package":     "pkg",
	"range":       "rnge",
	"return":      "ret",
	"select":      "slct",
	"struct":      "strct",
	"switch":      "swtch",
	"type":        "typ",
	"var":         "vr",

	// go types
	"error":      "e",
	"bool":       "b",
	"string":     "str",
	"byte":       "byt",
	"rune":       "r",
	"uintptr":    "uptr",
	"int":        "i",
	"int8":       "i8",
	"int16":      "i16",
	"int32":      "i32",
	"int64":      "i64",
	"uint":       "u",
	"uint8":      "u8",
	"uint16":     "u16",
	"uint32":     "u32",
	"uint64":     "u64",
	"float32":    "z",
	"float64":    "f",
	"complex64":  "c",
	"complex128": "c128",
}