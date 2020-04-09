# LeetCode No.4  [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/)

### 题目描述

给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

**示例:**

```
nums1 = [1, 3]
nums2 = [2]
则中位数是 2.0
```

### 题目解析
关键因素分析:
1. nums1 , nums2 都是有序的;
2. 时间复杂度 O(log(m + n)),二分查找;
3. 中位数 为 s1 s2 排序之后,下标为 (m+n)/2的一个或两个数值平均值;


思考分析:
1. 情况1:如果 m+n == 奇数  中位数为 nums1 或则 nums2的其中一个数字下标数值;
2. 情况2:如果 m+n == 偶数 中位数位于 nums1 或者 nums2 中 两个临近下标的中位数;
3. 情况3:如果 m+n == 偶数 中位数位于 nums1(或nums2) 末尾 , nums2(或nums1) 首位 两个下标数值的中位数;
4. 对于情况1: 如果 mid(s1) < mid(s2) 
 1,2,3
 1,2,5,5
1,1,2,2,3,5,5
 
测试 cases:

```
1.
nums1 = [1, 2]
nums2 = [3, 4]
则中位数是 (2 + 3)/2 = 2.5
2.

```
时间复杂度：O(log(m + n))
空间复杂度：O(1)


### 代码实现

`Swift` 版本实现:

```Swift


```

`Golang` 版本实现:

```golang


```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|0 ms|2.2 MB	 |golang|
|Accepted|40 ms|21.3 MB	 |swift|