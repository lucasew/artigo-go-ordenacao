package bubble

// Bubble sort
func Bubble(valores []int) []int {
	tamanho := len(valores)
	isAlterado := true
	for isAlterado {
		isAlterado = false
		for i := 0; i < tamanho-1; i++ {
			if valores[i] > valores[i+1] {
				dummy := valores[i]
				valores[i] = valores[i+1]
				valores[i+1] = dummy
				isAlterado = true
			}
		}
	}
	return valores
}
