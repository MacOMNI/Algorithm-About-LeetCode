package LeetCode

import (
	"fmt"
	"testing"
)

func Test_NLIS(t *testing.T) {
	//	fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7}))
	//	fmt.Println(findNumberOfLIS([]int{7, 9}))
	fmt.Println(findNumberOfLIS([]int{7, 7, 7, 101, 18}))

	//	fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))

}
func Test_LIS(t *testing.T) {
	fmt.Println(lengthOfLIS([]int{}))

	fmt.Println(lengthOfLIS([]int{7}))

	fmt.Println(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))

	fmt.Println(lengthOfLIS([]int{7, 9}))

	fmt.Println(lengthOfLIS([]int{10, 9, 8, 7, 7, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 101, 18}))

}
