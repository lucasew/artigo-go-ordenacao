package util

import (
	"testing"
)

func TestIsOrdenado(t *testing.T) {
	truthCases := [][]int{
		{2, 3, 4, 5, 6},
		{1, 66, 777, 832},
		{1, 2, 3, 4},
		{1},
	}
	for i := range truthCases {
		if !IsOrdenado(truthCases[i]) {
			t.Fail()
		}
	}
	falsyCases := [][]int{
		{5, 4, 3, 2, 1},
		{99, 0},
		{9, 1},
	}
	for i := range falsyCases {
		if IsOrdenado(falsyCases[i]) {
			t.Fail()
		}
	}
}
