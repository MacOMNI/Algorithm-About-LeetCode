# LeetCode  [TwoSum](https://leetcode.com/problems/two-sum/solution/)

### 题目描述

给定一个整数数组 `nums` 和一个目标值 `target`，请你在该数组中找出和为目标值的那 **两个** 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

**示例:**

```
给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
```

### 题目解析

使用 Map字典映射来解决问题。

设置一个 map 容器 record 用来记录元素的值与索引，然后遍历数组 nums。

* 每次遍历时使用临时变量 complement 用来保存目标值与当前值的差值
* 在此次遍历中查找 record ，查看是否有与 complement 一致的值，如果查找成功则返回查找值的索引值与当前变量的值 i
* 如果未找到，则在 record 保存该元素与索引值 i

时间复杂度：O(n)
空间复杂度：O(n)

### 代码实现

`Swift` 版本实现:

```Swift


```

`Golang` 版本实现:

```golang


```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|4 ms|3.8 MB	 |golang|
|Accepted|40 ms|21.3 MB	 |swift|