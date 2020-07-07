package main

import (
	"fmt"
)

func Add(exp *Equation, random func() int) {
	value := random()
	exp.Right = addInternal(exp.Right, value)
	exp.Left = addInternal(exp.Left, value)
}

func addInternal(exp Expression, value int) Expression {

	var result Expression
	switch v := exp.(type) {
	case *ConstExpression:
		v.Value = v.Value + value
		result = v
	case *AddConstantExpression:
		v.Value = v.Value + value
		result = v
	default:
		result = &AddConstantExpression{
			Parent: exp,
			Value:  value,
		}
	}
	return result
}

type AddConstantExpression struct {
	Parent Expression
	Value  int
}

func (a *AddConstantExpression) String() string {

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var parentString string
	if a.Parent == nil {
		parentString = "NaN"
	} else {
		parentString = a.Parent.String()
	}

	if a.Value >= 0 {
		return fmt.Sprintf("%v + %v", parentString, abs(a.Value))
	} else {
		return fmt.Sprintf("%v - %v", parentString, abs(a.Value))
	}

}
