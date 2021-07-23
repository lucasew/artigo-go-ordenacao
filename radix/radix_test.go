package radix

import (
	"gosort/test_util"
	"testing"
)

func TestRadixSort(t *testing.T) {
	util.CheckOrdenacao(Radix)(t)
}
