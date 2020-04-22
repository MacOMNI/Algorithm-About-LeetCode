# LeetCode No.72  [Edit Distance](https://leetcode.com/problems/edit-distance/)

### 题目描述

给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
 
**示例:**

```
输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')

```

### 题目解析
1. dp[i][j] 表示 w1 i结尾 w2 j 结尾的字符串 操作数;
2. 那么
for  i each w2
For j  each w1 
 w1[I] == w1[j] 
dp[i][j]  = min(dp[i-1][j-1],dp[i-1][j] +1,dp[i][j-1]+1)
 else  dp[i][j]  = min(dp[i-1][j-1]+1,dp[i-1][j] +1,dp[i][j-1]+1)


时间复杂度：O(n^2)
空间复杂度：O(n^2)


### 代码实现

`Golang` 版本实现:

```golang

func minDistance(word1 string, word2 string) int {
	m := len(word2)
	n := len(word1)
	var dp = make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	max := m
	if max < n {
		max = n
	}

	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word1[j] == word2[i] {
				dp[i+1][j+1] = min(dp[i][j], dp[i][j+1]+1, dp[i+1][j]+1)
			} else {
				dp[i+1][j+1] = min(dp[i][j]+1, dp[i][j+1]+1, dp[i+1][j]+1)
			}
		}
	}
	//fmt.Println(dp)
	return dp[m][n]
}
func min(a int, b int, c int) int {
	mm := a
	if mm > b {
		mm = b
	}
	if mm > c {
		mm = c
	}
	return mm
}


```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|8 ms|5.6 MB	 |golang|
