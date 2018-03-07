package interpreter

type ExpressionList interface{}

type PairExpressionList struct {
	ExpressionList
	Head Expression
	Tail ExpressionList
}

type LastExpressionList struct {
	ExpressionList
	Head Expression
}
