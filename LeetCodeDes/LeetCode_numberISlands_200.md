# LeetCode No.200  [Number of Islands](https://leetcode.com/problems/number-of-islands/)

### 题目描述

给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

**示例:**

```
输入:
11110
11010
11000
00000
输出: 1

```

### 题目解析
1. 用 vis [i][j] 表示 i,j 被访问状态;
2. 对于 grid[i][j] == '1' , 上下左右 dfs 出 所有 状态相同的,并标记vis;

时间复杂度：O(n^2)
空间复杂度：O(n^2)


### 代码实现

`Golang` 版本实现:

```golang
var vis [][]int
var m int
var n int

func numIslands(grid [][]byte) int {
	res := 0
	m = len(grid)
	if m <= 0 {
		return res
	}
	n = len(grid[0])
	vis = make([][]int, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if vis[i][j] != 1 && grid[i][j] == '1' {
				res++
				dfsGrid(i, j, grid)
			}
			vis[i][j] = 1
		}
	}
	return res
}
func dfsGrid(i int, j int, grid [][]byte) {
	vis[i][j] = 1
	if j-1 >= 0 && grid[i][j-1] == '1' && vis[i][j-1] == 0 { //left
		dfsGrid(i, j-1, grid)
	}
	if j+1 < n && grid[i][j+1] == '1' && vis[i][j+1] == 0 { //right
		dfsGrid(i, j+1, grid)
	}
	if i-1 >= 0 && grid[i-1][j] == '1' && vis[i-1][j] == 0 { //up
		dfsGrid(i-1, j, grid)
	}
	if i+1 < m && grid[i+1][j] == '1' && vis[i+1][j] == 0 { //down
		dfsGrid(i+1, j, grid)
	}
}

```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|0 ms|3.8 MB	 |golang|
