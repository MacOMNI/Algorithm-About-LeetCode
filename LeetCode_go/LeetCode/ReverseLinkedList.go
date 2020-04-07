package LeetCode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	listHead := head
	var resNode *ListNode //结果 res
	var headTmp *ListNode
	var headLastTmp *ListNode
	//	var betweenLastTmp *ListNode
	var tailFirstTmp *ListNode

	if m-1 > 0 { // 有 head
		resNode = listHead
		for i := 1; i < m; i++ { // head
			headLastTmp = listHead
			listHead = listHead.Next
		}
		headLastTmp.Next = nil
		//resNode.Next = nil

	}
	headTmp = listHead.Next

	rverseHead := listHead
	rverseHead.Next = nil
	tailFirstTmp = headTmp
	for i := m; i < n; i++ { //btw
		tmp := headTmp.Next
		headTmp.Next = rverseHead
		tailFirstTmp = tmp
		rverseHead = headTmp
		headTmp = tmp

	}

	if headLastTmp != nil { //head + btw
		headLastTmp.Next = rverseHead
	} else {
		resNode = rverseHead
	}
	if tailFirstTmp != nil { //tail
		tmp := resNode
		for i := 1; i < n; i++ { // head
			tmp = tmp.Next
		}
		tmp.Next = tailFirstTmp
	}

	return resNode
}

// 反转链表
func reverseListI(head *ListNode) *ListNode {
	listHead := head.Next
	rverseHead := head
	rverseHead.Next = nil
	for listHead != nil {
		tmp := listHead.Next
		listHead.Next = rverseHead
		rverseHead = listHead
		listHead = tmp
	}
	return rverseHead
}
