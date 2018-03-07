package interpreter

import (
	"errors"
	"fmt"
)

func update(table *Table, key string, value int) *Table {
	return &Table{key, value, table}
}

func lookup(table *Table, key string) (int, error) {
	head := table
	for head != nil {
		if key == head.id {
			return head.value, nil
		}
	}
	return 0, errors.New("no such identifier")
}

func interpretExpression(expression Expression, table *Table) (int, *Table, error) {
	switch expression.(type) {
	case IdentifierExpression:
		value, err := lookup(table, expression.(IdentifierExpression).Id)
		return value, table, err
	case NumberExpression:
		return expression.(NumberExpression).Number, table, nil
	case OperationExpression:
		operationExpression := expression.(OperationExpression)
		leftValue, leftTable, leftErr := interpretExpression(operationExpression.Left, table)
		if leftErr != nil {
			return 0, nil, leftErr
		}
		rightValue, rightTable, rightErr := interpretExpression(operationExpression.Right, leftTable)
		if rightErr != nil {
			return 0, nil, rightErr
		}
		result := 0
		switch operationExpression.Operator {
		case OperatorPlus:
			result = leftValue + rightValue
		case OperatorTimes:
			result = leftValue * rightValue
		case OperatorMinus:
			result = leftValue - rightValue
		case OperatorDivide:
			result = leftValue / rightValue
		}

		return result, rightTable, nil
	case SequenceExpression:
		sequenceExpression := expression.(SequenceExpression)
		table, err := interpretStatement(sequenceExpression.Statement, table)
		if err != nil {
			return 0, nil, err
		}

		return interpretExpression(sequenceExpression.NextExpression, table)
	}
	return 0, nil, errors.New("not implemented")
}

func interpretStatement(statement Statement, table *Table) (*Table, error) {
	switch statement.(type) {
	case AssignStatement:
		assignStatement := statement.(AssignStatement)
		value, table, err := interpretExpression(assignStatement.Expression, table)
		if err != nil {
			return table, err
		}
		return update(table, assignStatement.Id, value), nil
	case CompoundStatement:
		compoundStatement := statement.(CompoundStatement)
		leftTable, leftErr := interpretStatement(compoundStatement.StatementLeft, table)
		if leftErr != nil {
			return nil, leftErr
		}
		return interpretStatement(compoundStatement.StatementRight, leftTable)
	case PrintStatement:
		printStatement := statement.(PrintStatement)
		expressionList := printStatement.ExpressionList
		newTable := table

		for _, ok := expressionList.(LastExpressionList); !ok; _, ok = expressionList.(LastExpressionList) {

			pairExpressionList, ok := expressionList.(PairExpressionList)

			if !ok {
				return nil, errors.New("should be pairExpressionList")
			}

			value, newTmpTable, err := interpretExpression(pairExpressionList.Head, newTable)
			if err != nil {
				return nil, err
			}
			fmt.Printf("%d ", value)
			newTable = newTmpTable
			expressionList = pairExpressionList.Tail
		}

		lastExpressionList, _ := expressionList.(LastExpressionList)
		value, newTmpTable, err := interpretExpression(lastExpressionList.Head, newTable)
		if err == nil {
			fmt.Printf("%d \n", value)
		}
		return newTmpTable, err
	}
	return nil, errors.New("not implemented")
}

func Interpret(statement Statement) {
	_, err := interpretStatement(statement, nil)
	if err != nil {
		fmt.Println(err)
	}
}
