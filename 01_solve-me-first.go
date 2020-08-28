package main

import "fmt"

func solveMeFirst(a uint32, b uint32) uint32 {
	return a + b
}

func main() {
	var res uint32 = solveMeFirst(1, 2)
	fmt.Println(res)
}
