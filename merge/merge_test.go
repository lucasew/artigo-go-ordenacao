package merge

import (
	"gosort/test_util"
	"testing"
)

func TestMergeSort(t *testing.T) {
	util.CheckOrdenacao(Merge)(t)
}
