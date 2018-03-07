package interpreter

type Expression interface{}

type IdentifierExpression struct {
	Expression
	Id string
}

type NumberExpression struct {
	Expression
	Number int
}

type OperationExpression struct {
	Expression
	Left     Expression
	Operator BinaryOperator
	Right    Expression
}

type SequenceExpression struct {
	Expression
	Statement      Statement
	NextExpression Expression
}
