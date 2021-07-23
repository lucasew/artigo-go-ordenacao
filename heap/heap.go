package heap

// Heap Sort
func Heap(valores []int) []int {
	for max := len(valores) - 1; max > 0; max-- {
		imax := max
		i := 0
		for ; i < max; i++ {
			if valores[i] > valores[imax] {
				imax = i
			}
		}
		dummy := valores[max]
		valores[max] = valores[imax]
		valores[imax] = dummy

	}
	return valores
}
