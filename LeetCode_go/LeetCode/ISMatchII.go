package LeetCode

func isMatchII(s string, p string) bool {
	//dp i , j
	ptr := []byte{}
	for idx, val := range []byte(p) {
		if val == '*' {
			if idx+1 < len(p) && p[idx+1] == '*' {
				continue
			}
		}
		ptr = append(ptr, val)
	}
	p = string(ptr)
	m := len(s)
	n := len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		dp[0][i] = (dp[0][i-1] && p[i-1] == '*')
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == s[i-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}
		}
	}
	return dp[m][n]
}
