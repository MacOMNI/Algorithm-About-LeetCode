package LeetCode

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
