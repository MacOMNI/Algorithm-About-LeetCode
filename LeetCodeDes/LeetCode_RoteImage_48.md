# LeetCode No.48  [Rotate Image](https://leetcode.com/problems/rotate-image/)

### 题目描述

旋转图像
给定一个 n × n 的二维矩阵表示一个图像。
将图像顺时针旋转 90 度。
说明：
你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。

**示例:**

```
给定 matrix = 
 [
   [1,2,3],
   [4,5,6],
   [7,8,9]
 ],
 原地旋转输入矩阵，使其变为:
 [
   [7,4,1],
   [8,5,2],
   [9,6,3]
 ]
```

### 题目解析
1. 记录 圈内的 四个位置;
2. 依次 替换;

时间复杂度：O(n^2)
空间复杂度：O(1)


### 代码实现

`Golang` 版本实现:

```golang


func rotate(matrix [][]int) {
	n := len(matrix)
	//len := len(matrix) - 1

	for i := 0; i < n-1; i++ {
		for j := i; j < n-i-1; j++ {
			ttmp := matrix[i][j]
			rtmp := matrix[j][n-i-1]
			btmp := matrix[n-i-1][n-1-j]
			ltmp := matrix[n-1-j][i]
			matrix[j][n-i-1] = ttmp
			matrix[n-i-1][n-1-j] = rtmp
			matrix[n-1-j][i] = btmp
			matrix[i][j] = ltmp
		}

	}
}



```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|0 ms|2.2 MB	 |golang|
