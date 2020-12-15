//
//  LinkList.swift
//  LeetCode
//
//  Created by Mackun on 2020/12/2.
//  Copyright © 2020 Xcode. All rights reserved.
//

import UIKit

class LinkList: NSObject {

    func isPalindrome(_ head: ListNode?) -> Bool {
        var head = head
        if head == nil {
            return true
        }
        var i = 0
        var fast = head
        var slow = head
        while fast?.next != nil  {//走一半
            fast = fast?.next!.next
            slow = slow?.next!
            i += 1
            print(slow!.val)
        }
        print("i = ",i)
        // 反转链表
        // var reverHead = ListNode.init(-1)
        let reverHead = slow
        var tmpHead = reverHead
        var preHead: ListNode?
        while tmpHead != nil {
            let next = tmpHead?.next
            tmpHead?.next = preHead
            preHead = tmpHead
            tmpHead = next
        }
        
        while preHead != nil  {
            print(head?.val)
            print(preHead!.val)
            if preHead!.val != head?.val {
                return false
            }
            preHead = preHead!.next
            head = head?.next
        }
        
        return true
    }
    
    func rotateRight(_ head: ListNode?, _ k: Int) -> ListNode? {
            guard let head = head else {
                return nil
            }
            if k == 0 {
                return head
            }
            var headTmp = head
            var cnt = 1
            while headTmp.next != nil {
                cnt += 1
                headTmp = headTmp.next!
            }
            
            var res = cnt - k%cnt - 1
           // print("res = ",res)
            if  k%cnt == 0 {
                return head
            }
            
            var endTmp = head
            while res > 0 {
                res -= 1
                endTmp = endTmp.next!
            }
            let resHead = endTmp.next
            endTmp.next = nil
            headTmp.next = head
            
            return resHead
        }
}
