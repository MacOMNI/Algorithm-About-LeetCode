package LeetCode

func isMatch(s string, p string) bool {
	return dfsMatch([]byte(s), []byte(p), 0, 0)
}

func dfsMatch(s []byte, p []byte, sIndex int, pIndex int) bool {
	res := false
	if len(s) == sIndex {
		return true
	}
	if len(p) == pIndex {
		return false
	}
	for i := pIndex; i < len(p); i++ {
		if s[sIndex] == p[i] || p[i] == '.' {
			res = dfsMatch(s, p, sIndex+1, i+1)
		}
		if res {
			return true
		}
		if p[i] == '*' {
			res = dfsMatch(s, p, sIndex, i+1)
			if !res && ((i-1 >= 0 && p[i-1] == '.') || p[i-1] == s[sIndex]) {
				res = dfsMatch(s, p, sIndex+1, i+1)
			}
			if res {
				return true
			}
		}
		res = dfsMatch(s, p, sIndex, i+1)
		if res {
			return true
		}
	}
	return res
}
