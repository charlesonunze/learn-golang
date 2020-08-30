package main

import (
	"fmt"
)

func simpleArraySum(array []int32) int32 {
	var sum int32 = 0

	for i := 0; i < len(array); i++ {
		sum += array[i]
	}

	return sum
}

func main() {
	res := simpleArraySum([]int32{1, 2, 3, 4})
	fmt.Println(res)
}
