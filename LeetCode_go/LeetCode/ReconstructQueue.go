package LeetCode

import (
	"fmt"
	"sort"
)

type Reconstruct [][]int

func (ques Reconstruct) Less(i, j int) bool {
	if ques[i][0] > ques[j][0] {
		return true
	} else if ques[i][0] == ques[j][0] {
		if ques[i][1] < ques[j][1] {
			return true
		}
	}
	return false
}
func (ques Reconstruct) Len() int {
	return len(ques)
}
func (ques Reconstruct) Swap(i, j int) {
	ques[i], ques[j] = ques[j], ques[i]
}
func reconstructQueue(people [][]int) [][]int {
	n := len(people)
	if n < 0 {
		return nil
	}
	res := make([][]int, n)
	ques := Reconstruct(people)
	sort.Sort(ques)
	fmt.Println(ques)

	for i := 0; i < n; i++ {
		res[i] = people[i]
		j := i - 1
		for j >= 0 {
			if people[i][1] > res[j][1] {
				break
			}
			if people[i][0] == res[j][0] {
				break
			}
			j--
		}
		if j+2 < i {
			copy(res[j+2:i+1], res[j+1:i])
		}
		res[j+1] = people[i]
		fmt.Println(res)

	}
	return res
}
