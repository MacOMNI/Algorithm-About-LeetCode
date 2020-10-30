//
//  SortList.swift
//  LeetCode
//
//  Created by Mackun on 2020/10/30.
//  Copyright © 2020 Xcode. All rights reserved.
//

import Foundation

class SortList {
    func sortList(_ head: ListNode?) -> ListNode? {
//        guard let head = head else {
//            return nil
//        }
        var fastHead = head
        var slowHead = head
        while fastHead != nil && slowHead != nil {
            if fastHead?.next != nil {
                fastHead = fastHead?.next?.next
                slowHead = slowHead?.next
            } else {
                break
            }
        }
        // merge
        
        sortList(fastHead)
        sortList(slowHead)
        var resHead =
        
        return head
    }
}
