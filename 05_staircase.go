package main

import (
	"bytes"
	"fmt"
)

func staircase(n int32) {
	var i, j int32

	for i = 1; i <= n; i++ {
		var b bytes.Buffer

		for j = 1; j <= n; j++ {
			str := "#"

			if n-j >= i {
				str = " "
			}

			b.WriteString(str)
		}

		fmt.Println(b.String())
	}
}

func main() {
	staircase(10)
}
