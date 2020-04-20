package LeetCode

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
