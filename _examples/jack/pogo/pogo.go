package pogo

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx"
)

// GENERATED BY POGO. DO NOT EDIT.

// DB is the common interface for database operations that can be used with
// types from schema `jack`. Note that this is
// also copied into each of the table packages.
//
// This should work with database/sql.DB and database/sql.Tx.
type DB interface {
	Exec(string, ...interface{}) (pgx.CommandTag, error)
	Query(string, ...interface{}) (*pgx.Rows, error)
	QueryRow(string, ...interface{}) *pgx.Row
}

// Log function. Override this with the logger of your choice
var Log = func(string, ...interface{}) {}

// Conditions turns a list of conditions into
// sql clauses and params
func Conditions(conds ...Condition) (sql string, params []interface{}, err error) {
	var wheres []string
	var groupbys []string
	var orderbys []string
	var limits []string
	var ith int

	for _, cond := range conds {
		clause := cond.Clause()
		switch clause.Type {
		case "WHERE":
			var refs []interface{}
			for _, param := range clause.Params {
				ith++
				refs = append(refs, "$"+strconv.Itoa(ith))
				params = append(params, param)
			}
			wheres = append(wheres, fmt.Sprintf(clause.Format, refs...))
		case "GROUP BY":
			var refs []interface{}
			for _, param := range clause.Params {
				ith++
				refs = append(refs, "$"+strconv.Itoa(ith))
				params = append(params, param)
			}
			groupbys = append(groupbys, fmt.Sprintf(clause.Format, refs...))
		case "ORDER BY":
			var refs []interface{}
			for _, param := range clause.Params {
				ith++
				refs = append(refs, "$"+strconv.Itoa(ith))
				params = append(params, param)
			}
			orderbys = append(orderbys, fmt.Sprintf(clause.Format, refs...))
		case "LIMIT":
			var refs []interface{}
			for _, param := range clause.Params {
				ith++
				refs = append(refs, "$"+strconv.Itoa(ith))
				params = append(params, param)
			}
			limits = append(limits, fmt.Sprintf(clause.Format, refs...))
		default:
			return sql, params, fmt.Errorf("unknown condition type")
		}
	}

	// put everything together
	var out []string
	if len(wheres) > 0 {
		out = append(out, "WHERE "+strings.Join(wheres, " AND "))
	}
	if len(groupbys) > 0 {
		out = append(out, "GROUP BY "+strings.Join(groupbys, ", "))
	}
	if len(orderbys) > 0 {
		out = append(out, "ORDER BY "+strings.Join(orderbys, ", "))
	}
	if len(limits) > 0 {
		out = append(out, "LIMIT "+strings.Join(limits, ", "))
	}

	return strings.Join(out, " "), params, nil
}

// Condition interface
type Condition interface {
	Clause() *Clause
}

// Clause struct
type Clause struct {
	Type   string
	Format string
	Params []interface{}
}