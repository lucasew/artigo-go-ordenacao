package main

import (
	"html/template"
	"math/rand"
	"os"
)

func Generate(n int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = rand.Intn(1000000)
	}
	return ret
}

var tpl = `package vector

var Vector = map[int][]int{
    {{ range $k, $v := . }}
        {{ $k }}: []int{ {{ range $num := $v }} {{ $num }}, {{ end }} },
    {{end}}
}
`

func main() {
	t, err := template.New("vetores").Parse(tpl)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	cenarios := []int{}
	for i := 5000; i <= 200000; i += 5000 {
		cenarios = append(cenarios, i)
	}
	vetores := gerarVetores(cenarios)
	err = t.Execute(os.Stdout, vetores)
	if err != nil {
		panic(err)
	}
}

func gerarVetores(cenarios []int) map[int][]int {
	ret := map[int][]int{}
	for _, cenario := range cenarios {
		ret[cenario] = Generate(cenario)
	}
	return ret
}
