package main

import (
	"fmt"
	"github.com/el-ideal-ideas/ellib/elconv"
)

func main() {
	s := "100"
	v := &s
	fmt.Println(elconv.AsInt(v))
}
