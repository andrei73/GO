package random

import (
	"math/rand"
)

func Generate(n int) []int{
	var a []int
	for i:=0; i<n; i++ {
		a = append(a, rand.Intn(1000))
	}
	return a
}
