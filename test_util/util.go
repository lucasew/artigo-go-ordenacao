package util

import (
	"gosort/rand"
	"testing"
)

var cases = map[string][]int{
	"123":          {1, 2, 3},
	"321":          {3, 2, 1},
	"alternado":    {1, 3, 5, 7, 9, 2, 4, 6, 8},
	"aleatorio10":  generator.Generate(10),
	"aleatorio100": generator.Generate(100),
}

// IsOrdenado o vetor está ordenado?
func IsOrdenado(vet []int) bool {
	for i := 0; i < len(vet)-1; i++ {
		if vet[i] > vet[i+1] {
			return false
		}
	}
	return true
}

// CheckOrdenacao retorna test case checando ordenacao
func CheckOrdenacao(f func([]int) []int) func(*testing.T) {
	return func(t *testing.T) {
		for k, v := range cases {
			this := make([]int, len(v))
			copy(this, v)
			t.Run(k, func(t *testing.T) {
				isOrdenado := IsOrdenado(f(this))
				if !isOrdenado {
					t.Errorf("%#v nao foi ordenado corretamente\n", this)
				}
			})
		}
	}
}
