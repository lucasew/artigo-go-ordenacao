package selection

import (
	"gosort/test_util"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	util.CheckOrdenacao(Selection)(t)
}
