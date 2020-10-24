package run

import "gotest/base"

func Result(str string) int {
	a := base.GetA(str)
	b := base.GetB(str)
	return base.AddInt(a, b)
}
