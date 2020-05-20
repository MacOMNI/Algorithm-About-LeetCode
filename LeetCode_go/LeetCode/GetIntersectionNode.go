package LeetCode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }

    ptrA := headA
    ptrB := headB

    for ptrA != nil && ptrB != nil {
        if ptrA == ptrB {
            return ptrA
        }
        if ptrB == nil {
            ptrB = headA 
        }
        if ptrA == nil {
            ptrA = headB 
        }
        ptrB = ptrB.Next
        ptrA = ptrA.Next
    }

    return nil
}