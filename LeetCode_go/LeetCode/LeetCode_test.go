package LeetCode

import (
	"fmt"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	fmt.Println(twoSum([]int{3, 3}, 6))
	fmt.Println(twoSum([]int{3, 2, 5, 8, -2}, 6))

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
