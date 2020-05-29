package LeetCode

var result []string

func wordBreakII(s string, wordDict []string) []string {
	//result = make([]string, 0)
	vis := make(map[string]string, len(wordDict))
	hashMap := make(map[string]bool, 0)
	for _, val := range wordDict {
		hashMap[val] = true
	}

	return breakDfs(s, hashMap, vis, "")
}

func breakDfs(s string, hashMap map[string]bool, vis map[string]string, str string) []string {
	var ans = make([]string, 0)
	if len(s) == 0 {
		return ans
	}
	return ans
}

// func wordBreakII(s string, wordDict []string) bool {

// 	return wordDfsII(s, hashMap, vis)
// }
// func wordDfsII(s string, hashMap map[string]bool, vis map[string]bool) [][]string {
// 	var res [][]string
// 	if len(s) == 0 {
// 		return true
// 	}
// 	if val, ok := vis[s]; ok {
// 		return val
// 	}
// 	res := false
// 	for i := len(s) - 1; i >= 0; i-- {
// 		if _, ok := hashMap[s[i:]]; ok {
// 			res = wordDfsII(s[:i], hashMap, vis)
// 			vis[s[i:]] = res
// 			if res {
// 				return true
// 			}
// 		}
// 	}
// 	vis[s] = false
// 	return false
// }
