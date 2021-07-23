package shell

import (
	"gosort/test_util"
	"testing"
)

func TestShellSort(t *testing.T) {
	util.CheckOrdenacao(Shell)(t)
}
