package LeetCode

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	ls1 := 0
	ls2 := 0
	rs1 := m - 1
	rs2 := n - 1
	mids1 := (ls1 + rs1) >> 1
	mids2 := (ls2 + rs2) >> 1
	mid := (m + n) >> 1
	var res float64
	if m+n%2 == 1 { //情况 1
		for ls1 <= rs1 && ls2 <= rs2 {
			if nums1[mids1] <= nums2[mids2] {
				if mids1 + mids2 < mid {
					ls1 = mids1 + 1
				} else if mids1 + mids2 > mid {
					rs2 = mids2 - 1
				} else {
					return float64(nums2[mids2])
				}
			} else {
				if mids1 + mids2 < mid {
					ls2 = mids2 + 1
				} else if mids1 + mids2 > mid {
					rs1 = mids1 - 1
				} else {
					return float64(nums1[mids1])
				}
			}
		}
		for ls1 <= rs1 {

		}
		for ls2 <= rs2 {
			
		}

	} else { // 情况 2 + 3

	}
	fmt.Println(mid, mids1, mids2)

	fmt.Println(ls1, ls2, rs1, rs2)
	return res
}
