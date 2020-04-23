package LeetCode

import (
	"fmt"
	"testing"
)

func Test_RotateImage(t *testing.T) {
	var matrix [][]int
	matrix = append(matrix, []int{1, 2, 3})
	matrix = append(matrix, []int{4, 5, 6})
	matrix = append(matrix, []int{7, 8, 9})

	rotate(matrix)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])

	}
}
func Test_RotateImage1(t *testing.T) {
	var matrix [][]int
	matrix = append(matrix, []int{5, 1, 9, 11})
	matrix = append(matrix, []int{2, 4, 8, 10})
	matrix = append(matrix, []int{13, 3, 6, 7})
	matrix = append(matrix, []int{15, 14, 12, 16})

	rotate(matrix)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])

	}
}
