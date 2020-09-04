package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int) {
	var _arr []int
	var i, j int

	arrayLength := int(len(arr))

	for i = 0; i < arrayLength; i++ {
		var sum int
		for j = 0; j < arrayLength; j++ {
			if i != j {
				sum += arr[j]
			}
		}
		_arr = append(_arr, sum)
	}

	sort.Ints(_arr)
	fmt.Println(_arr[0], _arr[arrayLength-1])
}

func main() {
	miniMaxSum([]int{1, 3, 5, 7, 9})
}
