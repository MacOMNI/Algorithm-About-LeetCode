package LeetCode

import (
	"fmt"
	"testing"
)

func Test_LongestCommonSubsequence(t *testing.T) {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
	//fmt.Println(medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
func Test_MedianSlidingWindow(t *testing.T) {
	fmt.Println(medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	//fmt.Println(medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
func Test_MaxSlidingWindow(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, -5, -3, -6, -7}, 3))

	//fmt.Println(maxSlidingWindow([]int{1, 8, 6, 2, 100, 114, 8, 3, 7}, 3))

}
func Test_MaxArea(t *testing.T) {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 8, 6, 2, 100, 114, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 80, 6, 2, 100, 114, 8, 3, 7}))
	fmt.Println(maxArea([]int{100, 80, 6, 2, 100, 114, 8, 3, 7}))

}

func Test_3Sum(t *testing.T) {
	fmt.Println(threeSum([]int{4, 4, 3, -5, 0, 0, 0, -2, 3, -5, -5, 0}))

	fmt.Println(threeSum([]int{-4, -1, -4, 0, 2, -2, -4, -3, 2, -3, 2, 3, 3, -4}))

	fmt.Println(threeSum([]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}))
	fmt.Println(threeSum([]int{1, 1, 0, 0, -2}))

	fmt.Println(threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}))

	fmt.Println(threeSum([]int{-1, -4, 0, 2, 2}))
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))

	fmt.Println(threeSum([]int{0, 0, 0}))
	fmt.Println(threeSum([]int{}))

}
func Test_LongestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome("babadada"))

	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("acdffa"))
	fmt.Println(longestPalindrome("bbcb"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("bb"))
	fmt.Println(longestPalindrome("bbb"))

}
func Test_TwoSum(t *testing.T) {
	fmt.Println(twoSum([]int{3, 3}, 6))
	fmt.Println(twoSum([]int{3, 2, 5, 8, -2}, 6))

}
func Test_LengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))

	fmt.Println(lengthOfLongestSubstring("aabaab!bb"))

	fmt.Println(lengthOfLongestSubstring("abba"))

	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("au"))

	fmt.Println(lengthOfLongestSubstring("abcabcdefg"))
	fmt.Println(lengthOfLongestSubstring("abcadfeg"))

	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("cdd"))
}

func Test_ReverseList(t *testing.T) {
	head5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	head4 := &ListNode{
		Val:  4,
		Next: head5,
	}
	head3 := &ListNode{
		Val:  3,
		Next: head4,
	}
	head2 := &ListNode{
		Val:  2,
		Next: head3,
	}
	head1 := &ListNode{
		Val:  1,
		Next: head2,
	}
	// listHead := reverseBetween(head1, 2, 4)
	// for listHead != nil {
	// 	tmp := listHead.Next
	// 	fmt.Printf("%d->", listHead.Val)
	// 	listHead = tmp
	// }
	// fmt.Println()
	// listHead := reverseBetween(head1, 1, 1)
	// for listHead != nil {
	// 	tmp := listHead.Next
	// 	fmt.Printf("%d->", listHead.Val)
	// 	listHead = tmp
	// }
	// fmt.Println()
	listHead := reverseBetween(head1, 1, 5)
	//listHead := reverseBetween(head1, 2, 4)

	//listHead := reverseBetween(head1, 1, 1)
	for listHead != nil {
		tmp := listHead.Next
		fmt.Printf("%d->", listHead.Val)
		listHead = tmp
	}
	fmt.Println()
	// listHead := reverseBetween(head1, 1, 5)
	// for listHead != nil {
	// 	tmp := listHead.Next
	// 	fmt.Printf("%d->", listHead.Val)
	// 	listHead = tmp
	// }
	// fmt.Println()
}
