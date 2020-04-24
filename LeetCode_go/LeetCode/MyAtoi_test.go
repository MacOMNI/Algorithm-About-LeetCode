package LeetCode

import (
	"fmt"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	fmt.Println(myAtoi("0-1"))

	fmt.Println(myAtoi("+ 123"))

	fmt.Println(myAtoi("2147483648"))
	fmt.Println(myAtoi("+0 123"))

	fmt.Println(myAtoi("+ 1"))

	fmt.Println(myAtoi("   +-1-++"))
	fmt.Println(myAtoi("0   --++"))

	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))

}
