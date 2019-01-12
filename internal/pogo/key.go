package pogo

// Key interface
type Key interface {
}

// PrimaryKey struct
type PrimaryKey struct {
	Columns []Column
}

// ForeignKey struct
type ForeignKey struct {
	Column Column
}
