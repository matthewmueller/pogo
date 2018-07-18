package db

// Procedure represents a stored procedure.
type Procedure struct {
	Name       string // proc name
	Params     []*ProcedureParam
	ReturnType string // return type
}

// ProcedureParam represents a stored procedure.
type ProcedureParam struct {
	Name string // param name
	Type string // param type
}
