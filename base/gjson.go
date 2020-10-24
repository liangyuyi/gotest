package base

import "github.com/tidwall/gjson"

func GetA(str string) int {
	return int(gjson.Get(str, "a").Int())
}

func GetB(str string) int {
	return int(gjson.Get(str, "b").Int())
}
