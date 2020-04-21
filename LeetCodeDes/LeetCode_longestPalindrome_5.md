# LeetCode No.5  [Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)

### 题目描述


给定一个字符串 `s`，找到 `s` 中最长的回文子串。你可以假设 `s` 的最大长度为 1000。



**示例:**

```
输入: s = "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
```

### 题目解析
关键因素分析:
1. 回文串s(对折相同);
2. 连续子串;

解题思路:
1.  n < 1000,空间上 可以 二维;
2. 假设 dp[i][j] 表示 i -> j的回文串长度;
3. 那么 if s[i] == s[j] 并且 i->j 之间为回文串, dp[i][j] = j - i + 1;
否则 dp[i][j] = max(dp[i][j-1],dp[i-1][j]) ;

时间复杂度：O(n^2)
空间复杂度：O(n^2)


### 代码实现

`Golang` 版本实现:

```golang

func longestPalindrome(s string) string {
	n := int(len(s))
	var dp [1000][1000]int
	mlen := 0

	mIndex := 0
	if n > 0 {
		mlen = 1
	}
	for i := 0; i < n; i++ {
		dp[i][i] = 1

		si := int(s[i])
		for j := i - 1; j >= 0; j-- {
			sj := int(s[j])
			if si == sj {
				if (i-j > 2 && dp[j+1][i-1] == i-j-1) || (i-j <= 2) {
					dp[j][i] = i - j + 1
					if mlen < i-j+1 {
						mlen = i - j + 1
						mIndex = j
					}
				} else {
					if dp[j][i] < dp[j][i-1] {
						dp[j][i] = dp[j][i-1]
					}
					if dp[j+1][i] > dp[j][i] {
						dp[j][i] = dp[j+1][i]
					}
				}

			} else {
				if dp[j][i] < dp[j][i-1] {
					dp[j][i] = dp[j][i-1]
				}
				if dp[j+1][i] > dp[j][i] {
					dp[j][i] = dp[j+1][i]
				}
			}
			//fmt.Println("dp[", j, "][", i, "] = ", dp[j][i])

		}
	}
	//	fmt.Println("mlen = ", mlen, " palindrome = ", s[mIndex:mIndex+mlen])
	return s[mIndex : mIndex+mlen]
}


```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|60 ms|10.2 MB	 |golang|
