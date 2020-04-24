package LeetCode

import (
	"fmt"
	"testing"
)

func Test_findKthLargest(t *testing.T) {
	fmt.Println(findKthNumberLargest([]int{-1, 2, 0}, 3))

	fmt.Println(findKthNumberLargest([]int{2, 1}, 1))
	fmt.Println(findKthNumberLargest([]int{-1, 2, 0}, 2))

	fmt.Println(findKthNumberLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
