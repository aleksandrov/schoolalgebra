package main

import (
	"fmt"
)

func Multiply(exp *Equation) {
	value := random()
	exp.Right = multiplyInternal(exp.Right, value)
	exp.Left = multiplyInternal(exp.Left, value)
}

func multiplyInternal(exp Expression, value int) Expression {
	var result Expression
	switch v := exp.(type) {
	case *ConstExpression:
		v.Value = v.Value * value
		result = v
	case *MultiplyExpression:
		v.Value = v.Value * value
		result = v
	default:
		result = &MultiplyExpression{
			Parent: exp,
			Value:  value,
		}
	}
	return result
}

type MultiplyExpression struct {
	Parent Expression
	Value  int
}

func (m MultiplyExpression) String() string {
	var result string
	switch v := m.Parent.(type) {
	case *VariableExpression:
		result = fmt.Sprintf("%v%v", m.Value, v.String())
	default:
		result = fmt.Sprintf("%v * (%v)", m.Value, m.Parent.String())
	}
	return result
}
