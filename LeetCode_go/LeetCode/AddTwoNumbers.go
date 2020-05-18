package LeetCode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    needUp := 1
    node := &ListNode{
        Next:nil
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
		node.Next = &ListNode{
			Val:ans,
			Next:nil,
		}

	}
	for l1 != nil{

	}
	for l2 != nil{
        
	}
    return node.Next
}