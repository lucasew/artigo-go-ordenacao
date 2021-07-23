package quick

import (
	"gosort/test_util"
	"testing"
)

func TestQuickSort(t *testing.T) {
	util.CheckOrdenacao(Quick)(t)
}
