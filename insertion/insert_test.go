package insert

import (
	"gosort/test_util"
	"testing"
)

func TestInsertSort(t *testing.T) {
	util.CheckOrdenacao(Insert)(t)
}
