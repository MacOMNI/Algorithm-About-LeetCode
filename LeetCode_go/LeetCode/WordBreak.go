package LeetCode

func wordBreak(s string, wordDict []string) bool {
	vis := make(map[string]bool, len(wordDict))
	hashMap := make(map[string]bool, 0)
	for _, val := range wordDict {
		hashMap[val] = true
	}
	return wordDfs(s, hashMap, vis)
}
func wordDfs(s string, hashMap map[string]bool, vis map[string]bool) bool {
	if len(s) == 0 {
		return true
	}
	if val, ok := vis[s]; ok {
		return val
	}
	res := false
	for i := len(s) - 1; i >= 0; i-- {
		if _, ok := hashMap[s[i:]]; ok {
			res = wordDfs(s[:i], hashMap, vis)
			vis[s[i:]] = res
			if res {
				return true
			}
		}
	}
	vis[s] = false
	return false
}
