package LeetCode

import "fmt"

func findNumberOfLIS(nums []int) int {
	len := len(nums)
	dp := make([]int, len)
	res := 0
	cnt := 0
	for i := 0; i < len; i++ {
		dp[i] = 1
		res = 1
		cnt = 1
	}
	for i := len - 1; i >= 0; i-- {
		for j := i + 1; j < len; j++ {
			if nums[i] < nums[j] {
				if dp[j]+1 == dp[i] {
					cnt++
				}
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					cnt = 1
				}

			}
			if res < dp[i] {
				res = dp[i]
				cnt = 1
			}
		}
	}
	fmt.Println(res)
	fmt.Println(dp)

	return cnt
}
