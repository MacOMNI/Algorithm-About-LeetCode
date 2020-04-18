/**
 * Definition for singly-linked list.

 */
package LeetCode

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nHead := head
	var tailHeadTmp *ListNode
	var headLastTmp *ListNode
	var resNode *ListNode
	i := 0
	count := 0
	for head.Next != nil {

		head = head.Next
		i++
		count++
		if i == n {
			resNode = nHead
			headLastTmp = nHead
			//head = head.Next
			//	head = head.Next

			//tailHead = nHead.Next
			break
		}
	}
	for head != nil {
		count++
		headLastTmp = nHead
		// headTmp = nHead
		nHead = nHead.Next
		tailHeadTmp = nHead
		head = head.Next

		//tailHead = tailHead.Next
	}
	if n == 1 {
		headLastTmp.Next = nil
		return resNode
	}
	if count-n == 0 {
		return tailHeadTmp
	}
	if tailHeadTmp != nil {
		headLastTmp.Next = tailHeadTmp.Next
	}

	return resNode
}
