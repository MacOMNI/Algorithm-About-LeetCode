package LeetCode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    needUp := 0
    node := &ListNode{
        Next:nil,
	}
	tmp := node
    for l1 != nil && l2 != nil {
		ans :=  l1.Val + l2.Val + needUp
        if ans >= 10 {
			needUp = 1
		} else {
			needUp = 0	
		}
		ans = ans % 10

		tmp.Next = &ListNode{
			Val:ans,
			Next:nil,
		}
		tmp = tmp.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil{
		ans :=  l1.Val + needUp
        if ans >= 10 {
			needUp = 1
		} else {
			needUp = 0	
		}
		ans = ans % 10
		tmp.Next = &ListNode{
			Val:ans,
			Next:nil,
		}
		tmp = tmp.Next
		l1 = l1.Next
	}
	for l2 != nil{
        ans :=  l2.Val + needUp
        if ans >= 10 {
			needUp = 1
		} else {
			needUp = 0	
		}
		ans = ans % 10
		tmp.Next = &ListNode{
			Val:ans,
			Next:nil,
		}
		tmp = tmp.Next
		l2 = l2.Next
	}
	if needUp == 1 {
		tmp.Next = &ListNode{
			Val:needUp,
			Next:nil,
		}
	}
    return node.Next
}