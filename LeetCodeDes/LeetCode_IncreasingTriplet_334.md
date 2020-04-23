# LeetCode No.334  [Increasing Triplet Subsequence](https://leetcode.com/problems/increasing-triplet-subsequence/)

### 题目描述
给定一个未排序的数组，判断这个数组中是否存在长度为 3 的递增子序列。

数学表达式如下:

如果存在这样的 i, j, k,  且满足 0 ≤ i < j < k ≤ n-1，
使得 arr[i] < arr[j] < arr[k] ，返回 true ; 否则返回 false 。
说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1) 。

**示例:**

```
输入: [1,2,3,4,5]
输出: true

输入: [5,4,3,2,1]
输出: false
```

### 题目解析
扫描一遍数组，用变量x1保存当前最小的值，变量x2保存当前第二小的值。如果右面能碰到一个数大于x2，说明必然存在一个递增的三元组。

时间复杂度：O(n)
空间复杂度：O(1)


### 代码实现

`Golang` 版本实现:

```golang
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


```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|4 ms|3.8 MB	 |golang|
