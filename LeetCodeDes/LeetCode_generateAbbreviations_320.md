# LeetCode No.320  [列举单词的全部缩写](https://leetcode-cn.com/problems/generalized-abbreviation/)

### 题目描述

请你写出一个能够举单词全部缩写的函数。

注意：输出的顺序并不重要。

**示例:**

```
输入: "word"
输出:
["word", "1ord", "w1rd", "wo1d", "wor1", "2rd", "w2d", "wo2", "1o1d", "1or1", "w1r1", "1o2", "2r1", "3d", "w3", "4"]

```

### 题目解析

想法

对于一个长度为 nn 的单次有多少种缩写？答案是 2^n2 
n
  种，因为每个字符可以被缩写或者不被缩写，都将导致不同的缩写。

算法

回溯算法穷举了问题一系列可能的部分候选串。抽象地说，部分候选串可以被看作是一棵潜在的搜索树的节点。每一个部分候选串都是某些候选串的父节点，树的叶子节点就是部分候选串没法继续扩展的结果。

部分候选串可以被扩展应该是以下两种情况之一：

保留下一个字符
将下一个字符缩写
我们按深度优先的顺序扩展所有的潜在候选串。当在搜索树中遇到一个叶子节点的时候开始回溯。搜索树中所有的椰子节点都是有效的缩写并被放到全局的列表中，以便最后返回。

时间复杂度：O(n2^N)
空间复杂度：O(2^N)


### 代码实现

`Golang` 版本实现:

```golang
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

```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|62 ms|3.8 MB	 |golang|
