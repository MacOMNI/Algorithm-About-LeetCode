# LeetCode No.3  [3Sum](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)

### 题目描述

给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

**示例:**

```
输入: "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
```

### 题目解析
关键因素分析:
1. 窗口滑动;
2. 用 hashmap 记录前缀;
时间复杂度：O(n)
空间复杂度：O(n)


### 代码实现

`Swift` 版本实现:

```Swift
 func lengthOfLongSubString(str:String) -> Int {
        var subLength = 0 ;
        let len = str.count
        let array = Array<Character>(str)
        
        var hashMap = [Character:Int]()
        var preIndex = -1
        for index in 0..<len {
            let curCharacter =  array[index]
            if let curPrexIndex = hashMap[curCharacter] {
                preIndex = max(curPrexIndex, preIndex)
            } else {
                
            }
            hashMap[curCharacter] = index
            subLength = max(subLength, index - preIndex)
        }
        return subLength
    }
```

`Golang` 版本实现:

```golang

```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|48 ms|7 MB	 |Swift|
