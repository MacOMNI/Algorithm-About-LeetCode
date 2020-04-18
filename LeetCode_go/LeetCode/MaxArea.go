package LeetCode

func maxArea(height []int) int {
	res := 0
	left := 0
	low := 0
	right := len(height) - 1
	for left <= right {
		low = height[left]
		if height[left] > height[right] {
			low = height[right]
		}
		if res < low*(right-left) {
			res = low * (right - left)
		}
		if low == height[left] {
			left++
		} else {
			right--
		}
		//	fmt.Println("left ", left, " right ", right, " res ", res)
	}
	return res
}
