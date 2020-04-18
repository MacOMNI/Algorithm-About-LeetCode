package LeetCode

import (
	"fmt"
	"testing"
)

func Test_RemoveNthFromEnd(t *testing.T) {
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
	listHead := removeNthFromEnd(head1, 2)
	for listHead != nil {
		tmp := listHead.Next
		fmt.Printf("%d->", listHead.Val)
		listHead = tmp
	}
	fmt.Println()
}

func Test_RemoveNthFromEnd1(t *testing.T) {
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
	listHead := removeNthFromEnd(head1, 1)
	for listHead != nil {
		tmp := listHead.Next
		fmt.Printf("%d->", listHead.Val)
		listHead = tmp
	}
	fmt.Println()
}

func Test_RemoveNthFromEnd2(t *testing.T) {
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
	listHead := removeNthFromEnd(head1, 5)
	for listHead != nil {
		tmp := listHead.Next
		fmt.Printf("%d->", listHead.Val)
		listHead = tmp
	}
	fmt.Println()
}

func Test_RemoveNthFromEnd3(t *testing.T) {
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
	listHead := removeNthFromEnd(head1, 4)
	for listHead != nil {
		tmp := listHead.Next
		fmt.Printf("%d->", listHead.Val)
		listHead = tmp
	}
	fmt.Println()
}
