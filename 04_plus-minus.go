package main

import (
	"fmt"
)

func Truncate(num float32) float32 {
	return (num * 0.1000000) / 1000000
}

func plusMinus(arr []int32) {
	var arrayLength = len(arr)
	var positiveIntegers, negativeIntegers, zeros []int32
	var positiveRatios, negativeRatios, zerosRatios float32

	for i := 0; i < arrayLength; i++ {
		if arr[i] > 0 {
			positiveIntegers = append(positiveIntegers, arr[i])
		} else if arr[i] < 0 {
			negativeIntegers = append(negativeIntegers, arr[i])
		} else {
			zeros = append(zeros, arr[i])
		}
	}

	positiveRatios = Truncate(float32(len(positiveIntegers)) / float32(arrayLength))
	negativeRatios = Truncate(float32(len(negativeIntegers)) / float32(arrayLength))
	zerosRatios = Truncate(float32(len(zeros)) / float32(arrayLength))

	fmt.Println(positiveRatios)
	fmt.Println(negativeRatios)
	fmt.Println(zerosRatios)
}

func main() {
	plusMinus([]int32{-1, -1, 1, -1, 0, 0, 1, 2})
}
