package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	time := strings.Split(s, "")
	firstSlice := strings.Split(s, ":")[0]

	meridiem := time[len(time)-2]
	var left string
	var i int

	hour, err := strconv.Atoi(firstSlice)

	if meridiem == "P" {
		if err == nil {
			if hour < 12 {
				i = hour + 12
				left = strconv.Itoa(i)
			} else if hour >= 24 {
				i = hour - 24
				if i < 10 {
					left += "0"
				}
				left += strconv.Itoa(i)
			} else {
				left = firstSlice
			}
		}
	}

	if meridiem == "A" {
		if err == nil {
			if hour >= 12 {
				i = hour - 12
				left += "0"
				left += strconv.Itoa(i)
			} else {
				left = firstSlice
			}
		}
	}

	trim := strings.Join(time[:len(time)-2], "")
	right := strings.Split(trim, firstSlice)[1]
	final := left + right

	return final
}

func main() {
	res := timeConversion("06:05:45AM")
	fmt.Println(res)
}
