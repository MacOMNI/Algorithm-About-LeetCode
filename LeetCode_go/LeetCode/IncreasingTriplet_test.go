package LeetCode

import (
	"fmt"
	"testing"
)

func Test_increasingTriplet(t *testing.T) {
	fmt.Println(increasingTriplet([]int{1, 2, 3, 1, 2, 1}))

	fmt.Println(increasingTriplet([]int{1, 2, -10, -8, -7}))
	fmt.Println(increasingTriplet([]int{1, 1, 1, 2, 31}))
	fmt.Println(increasingTriplet([]int{1, 1, 1, 2, 3, 1, 1, 2, 1, 1}))

	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))

	fmt.Println(increasingTriplet([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 9, 4, 10, 5, 6}))

}
