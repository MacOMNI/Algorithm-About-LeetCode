# LeetCode No.56  [Merge Intervals](https://leetcode.com/problems/merge-intervals/)

### 题目描述
给出一个区间的集合，请合并所有重叠的区间。


**示例:**

```
输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间
```

### 题目解析
关键因素分析:
1. 二维排序;
2. 合并 startvalue & rightvalue;

时间复杂度：O(nlgn)
空间复杂度：O(n^n)


### 代码实现

`Golang` 版本实现:

```golang

type InterVal [][]int

func (intervals InterVal) Len() int {
	return len(intervals)
}
func (intervals InterVal) Less(i, j int) bool {
	if intervals[i][0] < intervals[j][0] {
		return true
	} else if intervals[i][0] == intervals[j][0] {
		return intervals[i][1] <= intervals[j][1]
	}
	return false
}
func (intervals InterVal) Swap(i, j int) {
	intervals[i], intervals[j] = intervals[j], intervals[i]
}

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n <= 0 {
		return intervals
	}
	interval := InterVal(intervals)
	sort.Sort(interval)
	var mergeIntervals = make([][]int, 0)
	startValue := interval[0][0]
	rightValue := interval[0][1]
	for i := 1; i < n; i++ {
		if rightValue < interval[i][0] {
			mergeIntervals = append(mergeIntervals, []int{startValue, rightValue})
			startValue = interval[i][0]
			rightValue = interval[i][1]
		} else if rightValue <= interval[i][1] {
			rightValue = interval[i][1]
		}
	}
	mergeIntervals = append(mergeIntervals, []int{startValue, rightValue})

	return mergeIntervals
}


```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|8 ms|3.8 MB	 |golang|
