package LeetCode

func maxProfit(prices []int) int {
	ans := 0
	len := len(prices)
	// dp[i][j] i - > j  之间的最大值
	var dp = make([][]int, len+1)
	//var dp1[i][0] dp[i][1]
	for i := 0; i < len+1; i++ {
		dp[i] = make([]int, 2)
	}
	for i := len - 1; i >= 0; i-- {
		dp[i][0] = dp[i+1][0]
		dp[i][1] = dp[i+1][1]
		for j := i + 1; j <= len-1; j++ {
			if prices[i] < prices[j] {
				dp[i][0] = max2(prices[j]-prices[i], dp[i][0])
				dp[i][1] = max2(prices[j]-prices[i]+dp[j+1][0], dp[i][1])
			}
			ans = max2(ans, dp[i][1])
			ans = max2(ans, dp[i][0])

		}
	}
	//fmt.Println(dp)
	//fmt.Println(dp2)

	return ans
}
func max2(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
