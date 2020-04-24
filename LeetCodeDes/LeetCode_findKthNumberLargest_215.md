# LeetCode No.215  [Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/)

### 题目描述

在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

**示例:**

```
输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。


```

### 题目解析

时间复杂度：O(nlgn)
空间复杂度：O(n)


### 代码实现

`Golang` 版本实现:

```golang
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

```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|4 ms|3.8 MB	 |golang|
