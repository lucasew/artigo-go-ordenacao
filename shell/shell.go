package shell

import (
	"gosort/insertion"
)

func Shell(valores []int) []int {
	hn := roundupDivision(len(valores))
	for ; hn > 1; hn = roundupDivision(hn) {
		for i := 0; i < len(valores)-hn; i++ {
			if valores[i] > valores[i+hn] {
				swp := valores[i]
				valores[i] = valores[i+hn]
				valores[i+hn] = swp
			}
		}
	}
	return insert.Insert(valores)
}

func roundupDivision(n int) int {
	ret := n / 2
	if (n % 2) != 0 {
		ret++
	}
	return ret
}
