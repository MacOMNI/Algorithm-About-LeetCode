# LeetCode No.1  [LongestSubstringWithoutRepeatingCharacters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

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
1. 两个数之和(当前下标数字如果是目标结果,那么另外一张下标数字数值已经确认);
2. 有切只有一个答案(不用考虑多种情况);
3. 同一位置元素只能一次;

解决问题:

1. 倒序 遍历 `nums`, 设置元素数值与下标 map[nums[index]] = index;

2. 正序遍历 `nums` , 判断 target - nums[index] 是否 在 map 对象上有存储;
3. 如果 `步骤 2` 存在,并且符合`因素3` 即 target - nums[index] != index;
4. 满足`步骤 3`的 index 以及 target - nums[index], 即是目标结果

测试 cases:

```
1. nums = [3,3], target = 6
2. nums = [3, 2, 5, 8, -2], target = 6
```
时间复杂度：O(n)
空间复杂度：O(n)


### 代码实现

`Swift` 版本实现:

```Swift


```

`Golang` 版本实现:

```golang
func twoSum(nums []int, target int) []int {
	var hashMap = make(map[int]int, 0)

	for i := len(nums) - 1; i > 0; i-- {
		hashMap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		if val, ok := hashMap[target-nums[i]]; ok {
			if val != i {
				return []int{i, val}
			}
		}

	}
	return []int{}
}

```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|4 ms|3.8 MB	 |golang|
