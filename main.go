package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ops = map[int]func(exp *Equation, random func() int){
	0: Add,
	1: Multiply,
}

func main() {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for try := 0; try < 10; try++ {
		e := &Equation{
			Left:  &VariableExpression{},
			Right: &ConstExpression{Value: 3},
		}
		for i := 0; i < 5; i++ {
			randomOp := ops[r1.Intn(len(ops))]
			randomOp(e, random)
		}

		fmt.Println(e.String())
	}

}
