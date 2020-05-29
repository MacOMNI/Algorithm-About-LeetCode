package LeetCode

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
				dp[i+1][j+1] = min3(dp[i][j], dp[i][j+1]+1, dp[i+1][j]+1)
			} else {
				dp[i+1][j+1] = min3(dp[i][j]+1, dp[i][j+1]+1, dp[i+1][j]+1)
			}
		}
	}
	//fmt.Println(dp)
	return dp[m][n]
}
func min3(a int, b int, c int) int {
	mm := a
	if mm > b {
		mm = b
	}
	if mm > c {
		mm = c
	}
	return mm
}
