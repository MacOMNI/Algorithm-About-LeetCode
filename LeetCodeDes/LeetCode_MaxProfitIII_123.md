# LeetCode No.123  [MaxProfit](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/)

### 题目描述
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
**示例:**

```

注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
输入: [3,3,5,0,0,3,1,4]
输出: 6
解释: 在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
     随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。

输入: [1,2,3,4,5]
输出: 4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。   
     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。   
     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。

输入: [7,6,4,3,1] 
输出: 0 
解释: 在这个情况下, 没有交易完成, 所以最大利润为 0。

```

### 题目解析



### 代码实现

`Golang` 版本实现:

```golang
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


```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|436 ms|4.3 MB	 |golang|
