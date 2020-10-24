package main

import (
	"fmt"
	"gotest/run"
)

func main() {
	str := `{"a":1, "b":3}`
	fmt.Println(run.Result(str))
}
