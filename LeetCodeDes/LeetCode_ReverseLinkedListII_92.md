# LeetCode No.92  [Reverse Linked List II](https://leetcode.com/problems/reverse-linked-list-ii/)

### 题目描述

反转从位置 *m* 到 *n* 的链表。请使用一趟扫描完成反转。

**说明:**
1 ≤ *m* ≤ *n* ≤ 链表长度。

**示例:**

```
输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
```

### 题目解析
关键因素分析:
1. 反转链表(如何反转);
2. 区间操作(m,n);
3. 一趟扫描;

使用 Map字典映射来解决问题:

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
class Solution {
        func twoSum(_ nums: [Int], _ target: Int) -> [Int] {
            var map = [Int:Int]()
            for index in (0..<nums.count).reversed() {
                map[nums[index]] = index
            }
            for index in 0..<nums.count {
                if map[target - nums[index]] != nil  && index != map[target - nums[index]]{
                    return [index,map[target - nums[index]]!]
                }
            }
            return []
        }
}

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
|Accepted|40 ms|21.3 MB	 |swift|