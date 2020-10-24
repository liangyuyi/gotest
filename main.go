package main

import (
	"fmt"
	"github.com/liangyuyi/gotest/run"
)

func main() {
	str := `{"a":1, "b":3}`
	fmt.Println(run.Result(str))
}
