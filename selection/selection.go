package selection

// Selection sort
func Selection(valores []int) []int {
	for i := 0; i < len(valores)-1; i++ {
		minimo := i
		for j := i; j < len(valores); j++ {
			if valores[minimo] > valores[j] {
				minimo = j
			}
		}
		if i != minimo {
			swp := valores[i]
			valores[i] = valores[minimo]
			valores[minimo] = swp
		}
	}
	return valores
}
