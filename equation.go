package main

import "fmt"

type Equation struct {
	Left Expression
	Right Expression
}

func (e *Equation) String() string {
	return fmt.Sprintf("%v = %v", e.Left, e.Right)
}








