# LeetCode No.480  [MedianSlidingWindow](https://leetcode.com/problems/sliding-window-median/)


### 题目描述

中位数是有序序列最中间的那个数。如果序列的大小是偶数，则没有最中间的数；此时中位数是最中间的两个数的平均数。

例如：

```
[2,3,4]，中位数是 3
[2,3]，中位数是 (2 + 3) / 2 = 2.5
给你一个数组 nums，有一个大小为 k 的窗口从最左端滑动到最右端。窗口中有 k 个数，每次窗口向右移动 1 位。你的任务是找出每次窗口移动后得到的新窗口中元素的中位数，并输出由它们组成的数组。
```

**示例:**

```
给出 nums = [1,3,-1,-3,5,3,6,7]，以及 k = 3。

窗口位置                      中位数
---------------               -----
[1  3  -1] -3  5  3  6  7       1
 1 [3  -1  -3] 5  3  6  7      -1
 1  3 [-1  -3  5] 3  6  7      -1
 1  3  -1 [-3  5  3] 6  7       3
 1  3  -1  -3 [5  3  6] 7       5
 1  3  -1  -3  5 [3  6  7]      6
 因此，返回该滑动窗口的中位数数组 [1,-1,-1,3,5,6]。
```

**注意：**

* 你可以假设 k 始终有效，即：k 始终小于输入的非空数组的元素个数。
* 与真实值误差在 10 ^ -5 以内的答案将被视作正确答案。

### 题目解析
关键因素分析:
1. 前 k 数组进行排序;
2. 对 k -> n 依次 二分查找到 删除到 i - k 元素,插入当前元素;
3. 计算 中位数;

时间复杂度：O(nlg(k))
空间复杂度 O(k)

### 代码实现

`Golang` 版本实现:

```golang

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


```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|44 ms|6.9 MB	 |golang|
