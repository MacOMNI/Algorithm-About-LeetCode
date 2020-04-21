# LeetCode No.59  [Spiral Matrix II](https://leetcode.com/problems/spiral-matrix-ii/)

### 题目描述

给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

**示例:**

```
输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]

```

### 题目解析
1. 解题思路 模拟时针变动 右->下->左->上;

测试 cases:

```
1. nums = [3,3], target = 6
2. nums = [3, 2, 5, 8, -2], target = 6
```
时间复杂度：O(n^2)
空间复杂度：O(n^2)


### 代码实现

`Golang` 版本实现:

```golang

func generateMatrix(n int) [][]int {
	var matrix = make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	sum := 1
	left := 0
	right := n - 1
	up := 0
	down := n - 1
	// if n == 1 {
	// 	matrix[0][0] = 1
	// 	return matrix
	// }
	for sum <= n*n {
		for j := left; j <= right; j++ { //右
			matrix[up][j] = sum
			sum++
		}
		up++
		for j := up; j <= down; j++ { //下
			matrix[j][right] = sum

			sum++
		}
		right--
		for j := right; j >= left; j-- { //左
			matrix[down][j] = sum
			sum++
		}
		down--
		for j := down; j >= up; j-- { //上
			matrix[j][left] = sum
			sum++
		}
		left++
	}
	return matrix
}

```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|0 ms|2.1 MB	 |golang|
