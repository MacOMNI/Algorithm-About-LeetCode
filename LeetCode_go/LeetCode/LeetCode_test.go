package LeetCode

import (
	"fmt"
	"testing"
)

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
