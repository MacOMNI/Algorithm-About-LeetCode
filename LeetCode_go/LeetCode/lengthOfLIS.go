package LeetCode

import "fmt"

type LIS struct {
	Index int
	Value int
}

func lengthOfLIS(nums []int) int {
	len := len(nums)
	dp := make([]int, len)
	res := 0

	for i := 0; i < len; i++ {
		dp[i] = 1
		res = 1
	}
	for i := len - 1; i >= 0; i-- {
		for j := i + 1; j < len; j++ {
			if nums[i] < nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
			res = max(res, dp[i])
		}
	}

	return res
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func lengthOfLISlgn(nums []int) int {
	len := len(nums)
	// if len == 0 {
	// 	return 0
	// }
	dp := make([]int, len)
	max := 0
	sortNums := make([]*LIS, len)
	for i := 0; i < len; i++ {
		dp[i] = 1
		insertIndex := preLargeIndex(sortNums, nums[i], i)
		cnt := 0
		if insertIndex >= 0 && insertIndex < i {
			if insertIndex > 0 {
				// if sortNums[insertIndex-1].Value < nums[i] {
				// 	cnt = 1
				// }
				dp[i] = dp[insertIndex-1] + 1
			}
			if i-1 >= 0 && dp[i-1] > dp[i] {
				dp[i] = dp[i-1]
			}
			if max < dp[i] {
				max = dp[i]
			}
			copy(sortNums[insertIndex+1:], sortNums[insertIndex:])
			sortNums[insertIndex] = &LIS{
				Index: i,
				Value: nums[i],
			}
		} else {
			sortNums[i] = &LIS{
				Index: i,
				Value: nums[i],
			}
			if i-1 >= 0 && sortNums[i-1].Index < nums[i] {
				cnt = 1
			}
			if i-1 >= 0 && dp[i-1]+cnt > dp[i] {
				dp[i] = dp[i-1] + cnt
			}
			if max < dp[i] {
				max = dp[i]
			}
		}
	}
	fmt.Println(dp)
	return max
}
func preLargeIndex(nums []*LIS, num int, index int) int {
	left := 0
	right := index
	mid := (left + right) >> 1
	for left < right {
		if nums[mid].Value < num {
			left = mid + 1
		} else {
			right = mid
		}
		mid = (left + right) >> 1

	}
	return mid
}
