package generator

import (
	"math/rand"
)

func Generate(n int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = rand.Int()
	}
	return ret
}
