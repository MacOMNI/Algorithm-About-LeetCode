# LeetCode No.3  [LongestSubstringWithoutRepeatingCharacters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

### 题目描述

19. 删除链表的倒数第N个节点
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
**示例:**

```
给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
```

### 题目解析
关键因素分析:
1. 链表操作;

解决问题:


测试 cases:

```
输入: "abcabcdefg"
输出: 7
输入: "abcadfeg"
输出: 7
输入: "tmmzuxt"
输出: 5
```
时间复杂度：O(n)
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
