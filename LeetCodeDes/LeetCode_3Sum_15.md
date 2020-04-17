# LeetCode No.15  [3Sum](https://leetcode.com/problems/3sum/)

### 题目描述

给定一个整数数组 `nums` 和一个目标值 `target`，判断 nums 中是否存在三个元素 a，b，c  ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。

**示例:**

```
给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

```

### 题目解析
关键因素分析:
1. 3数之和 为 0 ;
2. 三元组,不重复;
3. 如果3个数字,不全为 0 ,那么至少会有一个负数,一个正数;

使用 Map字典映射来解决问题:

1. 倒序 遍历 `nums`, 设置元素数值与下标 map[nums[index]] = index;
测试 cases:

```
Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
```
时间复杂度：O(n^2)
空间复杂度：O(n^2)


### 代码实现

`Swift` 版本实现:

```Swift

```

`Golang` 版本实现:

```golang

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	//fmt.Println(nums)
	var res [][]int

	n := len(nums)
	if n <= 0 {
		return res
	}
	rIndex := n - 1

	var hashMap = make(map[int]int, 0)
	var countMap = make(map[int]int, 0)
	for i := n - 1; i >= 0; i-- {
		if nums[i] > 0 {
			rIndex = i
		}
		hashMap[nums[i]] = i
	}
	for i := n - 1; i >= 0; i-- {
		countMap[nums[i]] = countMap[nums[i]] + 1
	}
	if nums[rIndex] > 0 { //有大于0的
		i := 0
		for i < rIndex {
			target := -nums[i]
			tail := n - 1
			for j := i + 1; j <= tail; j++ {
				if val, ok := hashMap[target-nums[j]]; ok {

					if j < val && val <= tail {
						res = append(res, []int{nums[i], nums[j], nums[val]})
						for nums[tail] >= nums[val] {
							tail--
						}

					} else if j == val && countMap[nums[j]] >= 2 {
						res = append(res, []int{nums[i], nums[j], nums[val]})
						preValue := nums[j]

						for k := j + 1; k <= tail; k++ {
							if nums[k] == preValue {
								j++
								break
							}
						}
					} else if i == val && nums[j] == nums[val] && countMap[nums[j]] >= 3 {
						res = append(res, []int{nums[i], nums[j], nums[val]})
						for nums[tail] >= nums[val] {
							tail--
						}
					}
					//	fmt.Println(i, " ", j, " ", val, " res = ", res)

				}
			}
			for target == -nums[i+1] {
				i++
			}
			i++

		}

	} else if nums[rIndex] == 0 {
		if countMap[0] >= 3 {
			res = append(res, []int{0, 0, 0})
		}
	}
	//fmt.Println("rIndex = ", mid)

	return res
}

```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|48 ms|7 MB	 |golang|
