package insert

// Insert sort
func Insert(vetor []int) []int {
	tamanho := len(vetor)
	for i := 1; i < tamanho; i++ {
		dummy := vetor[i] // Valor da reserva
		j := i - 1
		for ; j >= 0 && vetor[j] > dummy; j-- {
			vetor[j+1] = vetor[j]
		}
		vetor[j+1] = dummy
	}
	return vetor
}
