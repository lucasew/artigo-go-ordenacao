package radix

// Radix Sort
func Radix(valores []int) []int {
	dummyl := make([]int, len(valores))
	maior := valores[0]
	exp := 1
	for i := 0; i < len(valores); i++ {
		if valores[i] > maior {
			maior = valores[i]
		}
	}
	for maior/exp > 0 {
		buck := make([]int, 10)
		for i := 0; i < len(valores); i++ {
			buck[(valores[i]/exp)%10]++
		}
		for i := 1; i < 10; i++ {
			buck[i] += buck[i-1]
		}
		for i := len(valores) - 1; i >= 0; i-- {
			buck[(valores[i]/exp)%10]--
			dummyl[buck[(valores[i]/exp)%10]] = valores[i]
		}
		for i := 0; i < len(valores); i++ {
			valores[i] = dummyl[i]
		}
		exp *= 10
	}
	return valores
}
