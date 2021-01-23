package main

import "C"
import (
	"github.com/el-ideal-ideas/ellib/str"
	"github.com/el-ideal-ideas/ellib/utils"
)

//export similarity
func similarity(s, t *C.char) float64 {
	rate, _ := str.Similarity(C.GoString(s), C.GoString(t), 0)
	return float64(rate)
}

//export randomInt
func randomInt(min, max int) int {
	return utils.RandomInt(min, max)
}

func main() {}
