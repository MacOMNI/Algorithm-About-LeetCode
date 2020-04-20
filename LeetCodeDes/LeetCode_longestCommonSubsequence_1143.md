# LeetCode No.1143  [Longest Common Subsequence](https://leetcode.com/problems/longest-common-subsequence/)

### 题目描述
给定两个字符串 `text1` 和 `text2`，返回这两个字符串的最长公共子序列。

一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。

若这两个字符串没有公共子序列，则返回 0。

**示例:**

```
输入：text1 = "abcde", text2 = "ace" 
输出：3  
解释：最长公共子序列是 "ace"，它的长度为 3。

```

```
输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc"，它的长度为 3。

```
1 <= text1.length <= 1000
1 <= text2.length <= 1000
输入的字符串只含有小写英文字符。

### 题目解析
关键因素分析:
1. 最长公共子序列;
2. 因为小于 1000,可以 二维数组存储状态;
3. 假设 dp[i][j] 代表 以 text1[i] 结尾,text2[j] 结尾 最长的公共字串;
4. 那么 if text1[i] == text2[j] dp[i][j] = max(dp[i-1][j-1] + 1,dp[i-1][j],dp[i][j-1]) else  dp[i][j] = max(dp[i-1][j],dp[i][j-1])


时间复杂度：O(n^2)
空间复杂度：O(n^2)


### 代码实现

`Golang` 版本实现:

```golang

func longestCommonSubsequence(text1 string, text2 string) int {

	n := len(text1)
	m := len(text2)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = 1
				if i-1 >= 0 && j-1 >= 0 {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			} else {
				if i-1 >= 0 && dp[i][j] < dp[i-1][j] {
					dp[i][j] = dp[i-1][j]
				}
				if j-1 >= 0 && dp[i][j] < dp[i][j-1] {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[n-1][m-1]
}


```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|4 ms|9.9 MB	 |golang|
