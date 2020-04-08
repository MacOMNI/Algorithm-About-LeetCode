package LeetCode

func lengthOfLongestSubstring(s string) int {
	var hashMap = make(map[int]int, 0)
	lcs := 0
	preIndex := 0
	for i, ch := range s {

		if val, ok := hashMap[int(ch)]; ok {
			if preIndex <= val {
				if i-preIndex > lcs {
					lcs = i - preIndex
				}
				preIndex = val + 1
			} else {
				if i-preIndex+1 > lcs {
					lcs = i - preIndex + 1
				}
			}

		} else {
			if i-preIndex+1 > lcs {
				lcs = i - preIndex + 1
			}
		}
		hashMap[int(ch)] = i
		//	fmt.Println("preIndex ", preIndex, " ", string(ch), lcs)
	}
	return lcs
}
