package bubble

import (
	"gosort/test_util"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	util.CheckOrdenacao(Bubble)(t)
}
