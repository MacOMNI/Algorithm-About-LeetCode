package LeetCode

import (
	"sort"
)

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
