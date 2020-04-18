# LeetCode No.11  [Container With Most Water](https://leetcode.com/problems/container-with-most-water/)

### 题目描述

给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

**示例:**

```
输入：[1,8,6,2,5,4,8,3,7]
输出：49

```

### 题目解析
关键因素分析:
1. 面试是 left right 两个高度中 较小高度  low * (right - left) 计算而来;
2. 那么 left right 中 较小的决定了 池子的面积,较大的决定了面积的上限;
3. 那么 left right 两个指针中 保留较大的作为上限,依次扫描计算 所有 low * (right - left) 面积中最大的;

时间复杂度：O(n)
空间复杂度：O(1)


### 代码实现

`Golang` 版本实现:

```golang


func maxArea(height []int) int {
	res := 0
	left := 0
	low := 0
	right := len(height) - 1
	for left <= right {
		low = height[left]
		if height[left] > height[right] {
			low = height[right]
		}
		if res < low*(right-left) {
			res = low * (right - left)
		}
		if low == height[left] {
			left++
		} else {
			right--
		}
		//	fmt.Println("left ", left, " right ", right, " res ", res)
	}
	return res
}


```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|8 ms|5.8 MB	 |golang|
