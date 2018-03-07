package interpreter

type Table struct {
	id    string
	value int
	tail  *Table
}
