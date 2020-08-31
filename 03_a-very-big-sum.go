package main

import (
	"fmt"
)

func aVeryBigSum(ar []int64) int64 {
	var sum int64 = 0

	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	}

	return sum

}

func main() {
	res := aVeryBigSum([]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005})
	fmt.Println(res)
}
