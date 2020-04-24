package LeetCode

import "sort"

func findKthNumberLargest(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	knums := make([]int, k)
	//sort.Reverse()
	sort.Ints(nums[:k])
	copy(knums[:], nums[:k])

	for i := 0; i < k/2; i++ {
		knums[i], knums[k-i-1] = knums[k-i-1], knums[i]
	}
	for i := k; i < n; i++ {
		index := binarySearchArray(knums, nums[i])
		if index < k {
			copy(knums[index+1:], knums[index:])
			knums[index] = nums[i]
		}
	}
	return knums[k-1]
}

func binarySearchArray(knums []int, val int) int {
	left := 0
	right := len(knums)
	mid := (left + right) >> 1
	for left < right {
		if knums[mid] < val {
			right = mid
		} else {
			left = mid + 1
		}
		mid = (left + right) >> 1
	}
	return mid
}
