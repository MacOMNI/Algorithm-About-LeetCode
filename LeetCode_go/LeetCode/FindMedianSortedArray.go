package LeetCode

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	ls1 := 0
	rs1 := 0
	rs2 := m - 1
	ls2 := n - 1
	var mids1 int
	var mids2 int
	var mid int

	var res float64
	if m+n%2 == 1 { //情况 1
		mid = (m + n) / 2
		mids1 = (ls1 + rs1) / 2
		mids2 = (ls2 + rs2) / 2

		if nums1[mids1] < nums2[mids2] {
			if mids1+mids2 < mid {
				ls1 = mids1
			} else {
				rs2 = mids2
			}
		} else {

		}

	} else { // 情况 2 + 3

	}
	fmt.Println(mid, mids1, mids2)

	fmt.Println(ls1, ls2, rs1, rs2)
	return res
}
