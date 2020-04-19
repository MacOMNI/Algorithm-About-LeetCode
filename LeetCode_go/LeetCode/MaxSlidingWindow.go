package LeetCode

import "container/list"

// type Slide struct {
// 	Val   int
// 	Index int
// }

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
