package main

import "github.com/el-ideal-ideas/ellib/eldump"

func main() {
	eldump.V(map[string]interface{}{
		"key00": []int{1, 2, 3, 4},
		"key01": map[int]int{
			1: 2,
			2: 4,
			3: 6,
		},
		"key02": "some text",
	})
}
