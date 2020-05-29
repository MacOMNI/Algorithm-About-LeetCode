package LeetCode

import (
	"fmt"
	"testing"
)

func Test_WordBreak(T *testing.T) {
	dict := make([]string, 0)
	dict = append(dict, "leet", "code")
	fmt.Println(wordBreak("leetcode", dict))
}
func Test_WordBreakII(T *testing.T) {
	dict := make([]string, 0)
	dict = append(dict, "leet", "code")
	fmt.Println(wordBreakII("leetcode", dict))
}
