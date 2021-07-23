package merge

// Merge sort
func Merge(valores []int) []int {
	tamanho := len(valores)
	if tamanho <= 1 {
		return valores
	}
	metade := tamanho / 2
	va := valores[0:metade]
	vb := valores[metade:tamanho]
	Merge(va)
	Merge(vb)
	ia := 0
	ib := 0
	tmp := make([]int, tamanho)
	for i := 0; i < tamanho; i++ {
		if ib == len(vb) {
			tmp[i] = va[ia]
			ia++
			continue
		}
		if ia < len(va) && va[ia] <= vb[ib] {
			tmp[i] = va[ia]
			ia++
		} else if ib < len(vb) {
			tmp[i] = vb[ib]
			ib++
		}
	}
	for i := 0; i < tamanho; i++ {
		valores[i] = tmp[i]
	}
	tmp = nil
	return valores
}
