package interpreter

type BinaryOperator int

const (
	OperatorPlus   BinaryOperator = iota
	OperatorTimes
	OperatorMinus
	OperatorDivide
)
