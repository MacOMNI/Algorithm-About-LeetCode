package LeetCode

import "sort"

func medianSlidingWindow(nums []int, k int) []float64 {
	var res []float64
	n := len(nums)
	slides := make([]int, k)
	copy(slides[:], nums[:k])

	sort.Ints(slides[:k])
	res = append(res, calMedianValue(slides))
	for i := k; i < n; i++ {
		index := binarySearch(slides, nums[i-k]) // need remove index
		slides = append(slides[:index], slides[index+1:]...)

		insertIndex := binarySearch(slides, nums[i]) // need insert index

		slides = append(slides[:insertIndex], append([]int{nums[i]}, slides[insertIndex:]...)...)
		res = append(res, calMedianValue(slides))
	}
	return res
}
func binarySearch(nums []int, val int) int {
	left := 0
	right := len(nums)
	mid := (left + right) >> 1
	for left < right {
		if nums[mid] < val {
			left = mid + 1
		} else {
			right = mid
		}
		mid = (left + right) >> 1
	}
	return mid
}
func calMedianValue(nums []int) float64 {
	k := len(nums)
	if k%2 == 0 {
		return float64(nums[k/2-1]+nums[k/2]) / 2.0
	} else {
		return float64(nums[(k)/2])
	}
}
