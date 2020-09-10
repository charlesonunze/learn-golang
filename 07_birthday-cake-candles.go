package main

import (
	"fmt"
	"sort"
)

func birthdayCakeCandles(ar []int) int {
	var highestNumbers []int
	highestNumber := 0

	sort.Ints(ar)

	for i := len(ar) - 1; i >= 0; i-- {
		if ar[i] >= highestNumber {
			highestNumber = ar[i]
			highestNumbers = append(highestNumbers, ar[i])
		}
	}

	return len(highestNumbers)
}

func main() {
	res := birthdayCakeCandles([]int{1, 3, 5, 7, 9, 9, 9})
	fmt.Println(res)
}
