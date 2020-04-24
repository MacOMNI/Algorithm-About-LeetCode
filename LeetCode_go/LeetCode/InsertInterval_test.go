package LeetCode

import (
	"fmt"
	"testing"
)

func Test_MergeInterval(t *testing.T) {
	var matrix [][]int
	matrix = append(matrix, []int{1, 3})
	matrix = append(matrix, []int{2, 6})
	matrix = append(matrix, []int{8, 10})
	matrix = append(matrix, []int{15, 18})

	resMatrix := merge(matrix)
	for i := 0; i < len(resMatrix); i++ {
		fmt.Println(resMatrix[i])

	}
}
func Test_InsertInterval(t *testing.T) {
	var matrix [][]int
	matrix = append(matrix, []int{1, 3})
	matrix = append(matrix, []int{6, 9})
	resMatrix := insert(matrix, []int{2, 5})
	for i := 0; i < len(resMatrix); i++ {
		fmt.Println(resMatrix[i])

	}
}
func Test_InsertInterval1(t *testing.T) {
	var matrix [][]int
	matrix = append(matrix, []int{1, 2})
	matrix = append(matrix, []int{3, 5})
	matrix = append(matrix, []int{6, 7})

	matrix = append(matrix, []int{8, 10})

	matrix = append(matrix, []int{12, 16})

	resMatrix := insert(matrix, []int{4, 8})
	for i := 0; i < len(resMatrix); i++ {
		fmt.Println(resMatrix[i])

	}
}
