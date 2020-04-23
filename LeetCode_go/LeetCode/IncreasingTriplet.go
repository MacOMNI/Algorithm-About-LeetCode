package LeetCode

func increasingTriplet(nums []int) bool {

	len := len(nums)
	if len <= 2 {
		return false
	}
	min := 2<<31 - 1
	max := 2<<31 - 1
	for i := 0; i < len; i++ {
		if nums[i] <= min {
			min = nums[i]
		} else if nums[i] <= max {
			max = nums[i]
		} else {
			return true
		}
	}
	return false
}
