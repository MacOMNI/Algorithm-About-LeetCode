package LeetCode

import (
	"math"
)

func myAtoi(str string) int {
	isPostive := -1
	isNumber := 0
	ans := 0
	min := -math.MinInt32 / 10.0
	max := math.MaxInt32 / 10.0

	for _, val := range str {
		if val == ' ' {
			if isNumber == 1 || isPostive >= 0 {
				break
			}
		} else if val == '+' {
			if isPostive != -1 || isNumber == 1 {
				break
			}
			isPostive = 1
		} else if val == '-' {
			if isPostive != -1 || isNumber == 1 {
				break
			}
			isPostive = 0
		} else if val >= '0' && val <= '9' {
			isNumber = 1
			res := float64(ans) + float64(int(val)-48)/float64(10)
			if isPostive != 0 && res >= max {
				return math.MaxInt32
			}
			if isPostive == 0 && res > min {
				return math.MinInt32
			}
			ans = ans*10 + (int(val) - 48)
		} else {
			break
		}
	}
	if isNumber == 1 {
		if isPostive == 0 {
			return -ans
		}
		return ans
	}
	return 0
}
