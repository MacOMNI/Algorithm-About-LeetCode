package LeetCode

func ladderLength(beginWord string, endWord string, wordList []string) int {
	que := make([]string, 0)
	que = append(que, beginWord)
	hashMap := make(map[string]int, 0)
	// n := len(wordList)
	for _, val := range wordList {
		hashMap[val] = 0
	}
	ans := 0
	hashMap[beginWord] = 0
	if _, ok := hashMap[endWord]; ok {
		for len(que) > 0 {
			word := que[0]
			que = que[1:]
			res, _ := hashMap[word]
			for i, val := range word {
				for c := 'a'; c <= 'z'; c++ {
					if val != c {
						next := word[:i] + string(c) + word[i+1:]
						if next == endWord {
							return res + 2

						}
						if val, ok := hashMap[next]; ok {
							if val == 0 {
								hashMap[next] = res + 1
								que = append(que, next)
							} else if val > res+1 {
								hashMap[next] = res + 1
								que = append(que, next)
							}
						}
					}
				}
			}
		}
	}
	return ans
}
