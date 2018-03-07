// GoStraightLine is a straight line program interpreter
package main

import "./interpreter"

func main() {
	program := interpreter.CompoundStatement{
		StatementLeft: interpreter.AssignStatement{
			Id: "a",
			Expression: interpreter.OperationExpression{
				Left:     interpreter.NumberExpression{Number: 5},
				Operator: interpreter.OperatorPlus,
				Right:    interpreter.NumberExpression{Number: 3},
			},
		},
		StatementRight: interpreter.CompoundStatement{
			StatementLeft: interpreter.AssignStatement{
				Id: "b",
				Expression: interpreter.SequenceExpression{
					Statement: interpreter.PrintStatement{
						ExpressionList: interpreter.PairExpressionList{
							Head: interpreter.IdentifierExpression{Id: "a"},
							Tail: interpreter.LastExpressionList{
								Head: interpreter.OperationExpression{
									Left: interpreter.IdentifierExpression{
										Id: "a",
									},
									Operator: interpreter.OperatorMinus,
									Right: interpreter.NumberExpression{
										Number: 1,
									},
								},
							},
						},
					},
					NextExpression: interpreter.OperationExpression{
						Left:     interpreter.NumberExpression{Number: 10},
						Operator: interpreter.OperatorTimes,
						Right: interpreter.IdentifierExpression{
							Id: "a",
						},
					},
				},
			},
			StatementRight: interpreter.PrintStatement{
				ExpressionList: interpreter.LastExpressionList{
					Head: interpreter.IdentifierExpression{
						Id: "b",
					},
				},
			},
		},
	}

	interpreter.Interpret(program)
}
