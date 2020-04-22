# LeetCode No.300  [Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/)

### 题目描述

给定一个无序的整数数组，找到其中最长上升子序列的长度。



**示例:**

```
输入: [10,9,2,5,3,7,101,18]
输出: 4 
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。

```
`进阶`: 你能将算法的时间复杂度降低到 O(n log n) 吗?


### 题目解析
关键因素分析:
1. 最长上升子序列;
2. 二分查找前面序列比自己小的;
3. 假设当前 dp[i],表示以 i 为 结尾的最长子序列,j 为数组中 比 nums[i]小的下标, 那么 dp[i] = max(dp[i-1],dp[j]+1)

时间复杂度：O(nlgn)
空间复杂度：O(n)

### 代码实现

`Golang` 版本实现:

```golang
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

```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|8 ms|2.4 MB	 |golang|
