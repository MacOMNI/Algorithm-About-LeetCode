# LeetCode No.127  [Word Ladder](https://leetcode.com/problems/word-ladder/)

### 题目描述
给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：

每次转换只能改变一个字母。
转换过程中的中间单词必须是字典中的单词。
说明:

如果不存在这样的转换序列，返回 0。
所有单词具有相同的长度。
所有单词只由小写字母组成。
字典中不存在重复的单词。
你可以假设 beginWord 和 endWord 是非空的，且二者不相同。

**示例:**

```
输入:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

输出: 5

解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
     返回它的长度 5。

```

### 题目解析
1. bfs;
2. 从 endword 剪枝;


### 代码实现

`Golang` 版本实现:

```golang
func ladderLength(beginWord string, endWord string, wordList []string) int {
    que := make([]string,0)
    que = append(que,beginWord)
    hashMap := make(map[string]int,0)
   // n := len(wordList)
    for _,val := range wordList {
        hashMap[val] = 0
    }
    ans := 0
    hashMap[beginWord] = 0
    if _,ok := hashMap[endWord];ok {
        for len(que) >0 {
            word := que[0]
            que = que[1:]
            res ,_ := hashMap[word]
            for i,val := range word {
                for c:= 'a' ; c <= 'z';c++ {
                    if val != c {
                        next := word[:i] + string(c) + word[i+1:]
                        if next == endWord {
                            return res + 2
                            
                        }
                        if val,ok := hashMap[next];ok{
                            if val == 0 {
                                hashMap[next] = res + 1
                                que = append(que,next)
                             } else if val > res+1 {
                                hashMap[next] = res + 1
                                que = append(que,next)
                            }
                        }
                    }
                }
            }
        }
    }
    return ans
}

```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|128 ms|3.8 MB	 |golang|
