package main

import "fmt"

func DedupeInt(list []int) []int {
	var newList []int

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && list[i] == list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

func main() {
	fmt.Println(DedupeInt([]int{1, 2, 3, 2, 3, 4}))
}
