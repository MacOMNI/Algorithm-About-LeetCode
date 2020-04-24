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
1. 排序 前 k 个数数组;
2. 维护插入的 最大k个数据;

时间复杂度：O(nlgn)
空间复杂度：O(n)


### 代码实现

`Golang` 版本实现:

```golang

func findKthLargest(nums []int, k int) int {
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


```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|4 ms|3.8 MB	 |golang|
