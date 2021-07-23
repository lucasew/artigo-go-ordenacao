package benchmark

import (
	"gosort/vector"
	"testing"
)

func BenchmarkMethod(f func([]int) []int) map[int]testing.BenchmarkResult {
	results := map[int]testing.BenchmarkResult{}
	for k, v := range vector.Vector {
		benchfunc := BenchmarkifyCase(f, v)
		results[k] = ResultifyBenchmark(benchfunc)
	}
	return results
}

func ResultifyBenchmark(f func(b *testing.B)) testing.BenchmarkResult {
	return testing.Benchmark(f)
}

func BenchmarkifyCase(f func([]int) []int, vet []int) func(b *testing.B) {
	ord := make([]int, len(vet))
	return func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			copy(ord, vet) // Copia o vetor salvo lá para outro lugar, já que passamos ele como referencia
			f(ord)
		}
	}
}
