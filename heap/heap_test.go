package heap

import (
	"gosort/test_util"
	"testing"
)

func TestHeapSort(t *testing.T) {
	util.CheckOrdenacao(Heap)(t)
}
