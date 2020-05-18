package LeetCode

func isMatch(s string, p string) bool {
	return dfs([]byte(s), []byte(p), 0, 0)
}

func dfs(s []byte, p []byte, sIndex int, pIndex int) bool {
	res := false
	if len(s) == sIndex {
		return true
	}
	if len(p) == pIndex {
		return false
	}
	for i := pIndex; i < len(p); i++ {
		if s[sIndex] == p[i] || p[i] == '.' {
			res = dfs(s, p, sIndex+1, i+1)
		}
		if res {
			return true
		}
		if p[i] == '*' {
			res = dfs(s, p, sIndex, i+1)
			if !res && ((i-1 >= 0 && p[i-1] == '.') || p[i-1] == s[sIndex]) {
				res = dfs(s, p, sIndex+1, i+1)
			}
			if res {
				return true
			}
		}
		res = dfs(s, p, sIndex, i+1)
		if res {
			return true
		}
	}
	return res
}
