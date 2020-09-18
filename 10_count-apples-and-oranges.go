package main

import "fmt"

func countApplesAndOranges(s int, t int, a int, b int, apples []int, oranges []int) {
	houseDistanceStart := s
	houseDistanceEnd := t
	appleTreePosition := a
	orangeTreePosition := b

	var _apples []int
	var _oranges []int

	for i := 0; i < len(apples); i++ {
		sum := appleTreePosition + apples[i]

		if sum >= houseDistanceStart && sum <= houseDistanceEnd {
			_apples = append(_apples, sum)
		}
	}

	for i := 0; i < len(oranges); i++ {
		sum := orangeTreePosition + oranges[i]

		if sum >= houseDistanceStart && sum <= houseDistanceEnd {
			_oranges = append(_oranges, sum)
		}
	}

	fmt.Println(len(_apples))
	fmt.Println(len(_oranges))
}

func main() {
	countApplesAndOranges(7, 11, 5, 15, []int{-2, 2, 1}, []int{5, -6})
}
