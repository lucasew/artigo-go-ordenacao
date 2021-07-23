package main

import (
	"fmt"
	"gosort/benchmark"
	"gosort/bubble"
	"gosort/heap"
	"gosort/insertion"
	"gosort/merge"
	"gosort/quick"
	"gosort/radix"
	"gosort/selection"
	"gosort/shell"
	"gosort/vector"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

var comandos = map[string]func([]int) []int{
	"bubble":    bubble.Bubble,
	"heap":      heap.Heap,
	"insertion": insert.Insert,
	"radix":     radix.Radix,
	"selection": selection.Selection,
	"shell":     shell.Shell,
	"quick":     quick.Quick,
	"merge":     merge.Merge,
}

var operacoes = map[string]func(){
	"sort":      cmdSort,
	"bench":     cmdBench,
	"stepBench": cmdStepBench,
	"autoBench": cmdAutoBench,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Comandos disponiveis: ")
		for k := range operacoes {
			fmt.Printf(" %s", k)
		}
		fmt.Printf("\n")
		panic("Argumentos insuficientes")
	}
	operacoes[os.Args[1]]()
}

func cmdSort() {
	if len(os.Args) < 3 {
		fmt.Printf("Comandos disponiveis: ")
		for k := range comandos {
			fmt.Printf(" %s", k)
		}
		fmt.Printf("\n")
		panic("Argumentos insuficientes")
	}
	cmd := comandos[os.Args[2]]
	var err error
	numeros := make([]int, len(os.Args)-3)
	for i := 3; i < len(os.Args); i++ {
		numeros[i-3], err = strconv.Atoi(os.Args[i])
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("%v\n", cmd(numeros))
}

var isStepBench = false

func cmdStepBench() {
	isStepBench = true
	cmdBench()
}

func getTemperatura() int {
	strtemp, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		panic(err)
	}
	strtemp = strtemp[0 : len(strtemp)-1] // Tira o \n no final da string, evitando o erro ao obter em número
	temp, err := strconv.Atoi(string(strtemp))
	if err != nil {
		panic(err)
	}
	return temp / 1000
}

func cmdAutoBench() {
	isStepBench = true
	printCSVHeader()
	for i := 0; i < 3; i++ {
		for k, v := range comandos {
			handleBenchMethod(k, v)
		}
	}
}

const esperarTemperatura = 43 // Esperar até a máquina chegar a x graus para testar o próximo vetor
func cmdBench() {
	if len(os.Args) < 3 {
		fmt.Printf("Comandos disponiveis: ")
		for k := range comandos {
			fmt.Printf(" %s", k)
		}
		fmt.Printf("\n")
		panic("Argumentos insuficientes")
	}
	f := comandos[os.Args[2]]
	printCSVHeader()
	handleBenchMethod(os.Args[2], f)
}

func printCSVHeader() {
	fmt.Println("metodo,tamanhovetor,tempo")
}

func handleBenchMethod(algorithm string, f func([]int) []int) {
	ticker := time.Tick(time.Second)
	for k, v := range vector.Vector {
		res := benchmark.ResultifyBenchmark(benchmark.BenchmarkifyCase(f, v))
		fmt.Printf("%s,%d,%.2f\n", algorithm, k, float64(res.T.Microseconds())/float64(res.N))
		for isStepBench && getTemperatura() > esperarTemperatura { // Espera até a temperatura ficar menor que esperarTemperatura, automações...
			<-ticker
		}
	}
}
