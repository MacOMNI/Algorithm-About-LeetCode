package LeetCode

import (
	"fmt"
	"testing"
)

func Test_NumIslands(t *testing.T) {
	var matrix [][]byte
	//	["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]

	matrix = append(matrix, []byte{1, 1, 1, 1, 0})
	matrix = append(matrix, []byte{1, 1, 0, 1, 0})

	matrix = append(matrix, []byte{1, 1, 0, 0, 0})
	matrix = append(matrix, []byte{0, 0, 0, 0, 0})

	fmt.Println(numIslands(matrix))
}

func Test_NumIslands2(t *testing.T) {
	var matrix [][]byte
	matrix = append(matrix, []byte{1, 1, 0, 0, 0})
	matrix = append(matrix, []byte{1, 1, 0, 0, 0})

	matrix = append(matrix, []byte{0, 0, 1, 0, 0})
	matrix = append(matrix, []byte{0, 0, 0, 1, 1})

	fmt.Println(numIslands(matrix))
}
