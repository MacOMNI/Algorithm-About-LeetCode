package LeetCode

import (
	"fmt"
	"testing"
)

func Test_SortedListToBST(t *testing.T) {
	head5 := &ListNode{
		Val:  9,
		Next: nil,
	}
	head4 := &ListNode{
		Val:  5,
		Next: head5,
	}
	head3 := &ListNode{
		Val:  0,
		Next: head4,
	}
	head2 := &ListNode{
		Val:  -3,
		Next: head3,
	}
	head1 := &ListNode{
		Val:  -10,
		Next: head2,
	}
	tree := sortedListToBST(head1)
	fmt.Println(inorderTraversal(tree))
}
