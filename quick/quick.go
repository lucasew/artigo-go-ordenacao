package quick

// Quick sort Mesmo que Partition Exchange Sort
func Quick(valores []int) []int {
	qsort(valores, 0, len(valores)-1)
	return valores
}

func qsort(valores []int, low, high int) {
	if low < high {
		pi := partition(valores, low, high)
		qsort(valores, low, pi-1)
		qsort(valores, pi+1, high)
	}
}

func partition(valores []int, low, high int) int {
	pivot := valores[high]
	i := (low - 1)
	for j := low; j <= high-1; j++ {
		if valores[j] < pivot {
			i++
			swp := valores[i]
			valores[i] = valores[j]
			valores[j] = swp
		}
	}
	swp := valores[i+1]
	valores[i+1] = valores[high]
	valores[high] = swp
	return i + 1
}
