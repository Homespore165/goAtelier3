package main

import (
	"math/rand/v2"
)

func ArrayGenerate() []int {
	var arr [1000]int
	for i := range arr {
		arr[i] = rand.IntN(10000)
	}
	return arr[:]
}
