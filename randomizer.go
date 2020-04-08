package main

import (
	"math/rand"
	"time"
)

var generator = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomize(percentage float64) bool {
	num := generator.Float64()
	if num < percentage {
		return true
	} else {
		return false
	}
}
