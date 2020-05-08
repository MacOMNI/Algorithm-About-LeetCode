package LeetCode

import "strconv"

// 部分候选串可以被扩展应该是以下两种情况之一：

// 保留下一个字符
// 将下一个字符缩写
var res []string

func generateAbbreviations(word string) []string {
	res = make([]string, 0)
	dfsGenerate("", word, 0, 0)
	return res
}

//回溯   index currentindex cnt 区间分割
func dfsGenerate(preword string, word string, index int, cnt int) {
	if index == len(word) {
		if cnt != 0 {
			preword = preword + strconv.FormatInt(int64(cnt), 10)
		}
		res = append(res, preword)
	} else {
		dfsGenerate(preword, word, index+1, cnt+1)

		if cnt > 0 {
			preword = preword + strconv.FormatInt(int64(cnt), 10)
		}
		preword = preword + string(word[index])

		dfsGenerate(preword, word, index+1, 0)
	}
	preword = ""

}
