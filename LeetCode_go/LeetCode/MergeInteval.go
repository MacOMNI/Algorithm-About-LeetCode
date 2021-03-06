package LeetCode

import "sort"

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
