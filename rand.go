package main

import (
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())
var randomizer = rand.New(source)
var maxNumber = 10

func randomPositive() int {
	return randomizer.Intn(maxNumber-1) + 1
}

func random() int {
	result := randomPositive() - randomPositive()
	for result == 0 {
		result = randomPositive() - randomPositive()
	}
	return result
}
