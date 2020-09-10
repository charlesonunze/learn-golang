package main

import (
	"fmt"
	"math"
)

func gradingStudents(grades []int) []int {
	var finalGrades []int

	for i := 0; i < len(grades); i++ {
		quotient := float64(grades[i]) / float64(5)
		rounded := int(math.Ceil(quotient) * 5)
		difference := rounded - grades[i]

		if grades[i] > 34 {
			if difference < 3 {
				finalGrades = append(finalGrades, rounded)
			} else {
				finalGrades = append(finalGrades, grades[i])
			}
		}

		if grades[i] < 35 {
			finalGrades = append(finalGrades, grades[i])
		}
	}

	return finalGrades
}

func main() {
	res := gradingStudents([]int{
		22,
		86,
		30,
		0,
		16,
		51,
		53,
		42,
		48,
		22,
		69,
		12,
		27,
		34,
		24,
		95,
		16,
		32,
		22,
		52,
		56,
		71,
		95,
	})
	fmt.Println(" ")
	fmt.Println(res)
}
