package run

import "github.com/liangyuyi/gotest/base"

func Result(str string) int {
	a := base.GetA(str)
	b := base.GetB(str)
	return base.AddInt(a, b)
}
