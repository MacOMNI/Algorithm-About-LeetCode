# LeetCode No.3  [LongestSubstringWithoutRepeatingCharacters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

### 题目描述

给定一个字符串，请你找出其中不含有重复字符的 **最长子串** 的长度。

**示例:**

```
输入: "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
```

### 题目解析
关键因素分析:
1. 无重复字符串;
2. 连续子串;

解决问题:

1. 遍历字符串时候, 用hashmap 找到当前字节的前一个 index;
2. 找到当前下标的最长子串的起始下标 preIndex;
3. 所有 当前下标到 preIndex 差值最大的即为最优解;

测试 cases:

```
输入: "abcabcdefg"
输出: 7
输入: "abcadfeg"
输出: 7
输入: "tmmzuxt"
输出: 5
```
时间复杂度：O(1)
空间复杂度：O(1)


### 代码实现

`Swift` 版本实现:

```Swift


```

`Golang` 版本实现:

```golang
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

```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|0 ms|3.4 MB	 |golang|
