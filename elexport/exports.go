package main

import "C"
import (
	"github.com/el-ideal-ideas/ellib/elrand"
	"github.com/el-ideal-ideas/ellib/elstr"
)

//export similarity
func similarity(s, t *C.char) float64 {
	rate, _ := elstr.Similarity(C.GoString(s), C.GoString(t), 0)
	return float64(rate)
}

//export randomInt
func randomInt(min, max int) int {
	return elrand.RandomInt(min, max)
}

func main() {}
