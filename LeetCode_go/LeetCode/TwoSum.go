package LeetCode

// two sum
func twoSum(nums []int, target int) []int {
	var hashMap = make(map[int]int, 0)

	for i := len(nums) - 1; i > 0; i-- {
		hashMap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		if val, ok := hashMap[target-nums[i]]; ok {
			if val != i {
				return []int{i, val}
			}
		}

	}
	return []int{}
}
