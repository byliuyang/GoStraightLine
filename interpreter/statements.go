package interpreter

type Statement interface{}

type CompoundStatement struct {
	Statement
	StatementLeft  Statement
	StatementRight Statement
}

type AssignStatement struct {
	Statement
	Id         string
	Expression Expression
}

type PrintStatement struct {
	Statement
	ExpressionList ExpressionList
}
