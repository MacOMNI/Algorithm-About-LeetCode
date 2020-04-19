# LeetCode No.239  [Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/)


### 题目描述

给定一个数组 *nums*，有一个大小为 *k* 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口 *k* 内的数字。滑动窗口每次只向右移动一位。

返回滑动窗口最大值。

**示例:**

```
输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7] 
解释: 

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
```

**注意：**

你可以假设 *k* 总是有效的，1 ≤ k ≤ 输入数组的大小，且输入数组不为空。

**进阶：**

你能在线性时间复杂度内解决此题吗？

### 题目解析
关键因素分析:
1. 大顶堆;
2. 双端队列实现 大顶堆;

时间复杂度：O(n)


### 代码实现

`Golang` 版本实现:

```golang

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	queue := list.New()
	n := len(nums)
	for i := 0; i < k; i++ {
		for queue.Len() > 0 {
			index, _ := queue.Back().Value.(int)
			if nums[index] <= nums[i] {
				queue.Remove(queue.Back())
			} else {
				break
			}
		}
		queue.PushBack(i)
	}
	index, ok := queue.Front().Value.(int)
	if ok {
		res = append(res, nums[index])
	}
	for i := k; i < n; i++ {
		index, _ := queue.Front().Value.(int)
		if i-index >= k {
			queue.Remove(queue.Front())
		}
		for queue.Len() > 0 {
			index, _ := queue.Back().Value.(int)
			if nums[index] <= nums[i] {
				queue.Remove(queue.Back())
			} else {
				break
			}
		}
		queue.PushBack(i)

		index, _ = queue.Front().Value.(int)
		res = append(res, nums[index])

	}

	//res = append(res, nums[mIndex])
	return res
}


```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|20 ms|6.4 MB	 |golang|
