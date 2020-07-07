package main

import "fmt"

type ConstExpression struct {
	Value int
}

func (c ConstExpression) String() string {
	return fmt.Sprintf("%v", c.Value)
}

