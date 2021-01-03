//
//  TowSum.swift
//  LeetCode
//
//  Created by MacOMNI on 2020/4/3.
//  Copyright © 2020 Xcode. All rights reserved.
//

import UIKit
//class ListNode {
//    var val:Int = 0
//    var next:ListNode?
//    init(val:Int) {
//        self.val = val
//        self.next = nil
//    }
//    init(val:Int,next:ListNode?) {
//        self.val = val
//        self.next = next
//    }
//}
public class ListNode {
    public var val: Int
    public var next: ListNode?
    public init(_ val: Int) {
        self.val = val
        self.next = nil
    }
    init(val:Int,next:ListNode?) {
        self.val = val
        self.next = next
    }
}
class Solution {
    func lowestCommonAncestor(_ root: TreeNode?, _ p: TreeNode?, _ q: TreeNode?) -> TreeNode? {
        if root == nil || root === p || root === q {
            return root
        }
        let right = lowestCommonAncestor(root?.right, p, q)
        let left = lowestCommonAncestor(root?.left, p, q)
        if right == nil {
            return left
        }
        if left == nil {
            return right
        }
        return root
    }
    func maxProduct(nums: [Int]) -> Int {
        var ans = -Int.max
        var current = 1
        var currentIndex = -1
        var dpos = [Int]()
        for (index,item) in nums.enumerated() {
            ans = max(ans,item)
            if item != 0 {
                current *= item
                if item < 0 {
                   dpos.append(current)
                }
                if index - currentIndex > 1 {
                    ans = max(ans,current)
                }
                print(ans)
            } else {
                if index - currentIndex > 1 && dpos.count % 2 == 1 {
                    ans = max(ans,current/dpos[0])
                }
                dpos = [Int]()
                current = 1
                currentIndex = index
                print(index,current,ans)
            }
        }
        if nums.count - 1 - currentIndex > 1 && dpos.count % 2 == 1 {
            ans = max(ans,current/dpos[0])
        }
        return ans
    }
    
    func sortList(_ head: ListNode?) -> ListNode? {
        if head == nil || head?.next == nil  {
            return head
        }
        var slowHead = head
        var fastHead = head?.next
        while fastHead != nil && fastHead?.next != nil {
            slowHead = slowHead!.next
            fastHead = fastHead!.next?.next
        }
//        if fastHead === slowHead {
//            slowHead = slowHead!.next
//        }
        let tmpHead = slowHead?.next;
        slowHead?.next = nil
        return mergeSort(leftHead: sortList(head), rightHead: sortList(tmpHead))
    }
    func ReveveList(_ head: ListNode?) -> ListNode?{
//        guard let head = head else {
//            return nil
//        }
        var tmp = head
        var pre: ListNode?
        while tmp != nil {
            let next = tmp?.next
            tmp?.next = pre
            pre = tmp
            tmp = next
            
        }
        return pre
    }
    func mergeSort(leftHead: ListNode? , rightHead: ListNode?) -> ListNode? {
        var leftHead = leftHead
        var rightHead = rightHead
        let head = ListNode.init(0)
        var tmpHead = head
        while leftHead != nil && rightHead != nil {
            if leftHead!.val <= rightHead!.val {
                tmpHead.next = leftHead
                leftHead  = leftHead!.next
            } else {
                tmpHead.next = rightHead
                rightHead  = rightHead!.next
            }
            tmpHead = tmpHead.next!
        }
        if leftHead != nil {
            tmpHead.next = leftHead
        }
        if rightHead != nil {
            tmpHead.next = rightHead
        }
        return head.next;
    }
    func detectCycle(_ head: ListNode?) -> ListNode? {
        // var head = head
        var slow = head
        var fast = head
        while slow != nil && fast != nil {
            if fast?.next != nil {
                slow = slow!.next
                fast = fast!.next?.next
            } else {
                return nil
            }
            if slow === fast {
                var ptr = head
                while(!(ptr === slow)) {
                    ptr = ptr!.next
                    slow = slow!.next
                }
                return ptr
            }
        }
        return nil
        
    }
    func maxArea(_ height: [Int]) -> Int {
        var r = height.count - 1,l = 0, ans = 0
        while l < r {
            let low = min(height[l],height[r])
            let res = low * (r-l)
            ans = max(ans,res)
            if low == height[l] {
                l = l + 1
            }else {
                r = r - 1
            }
        }
        return ans
    }
    func removeNthNumber( head:ListNode?, n:Int) -> ListNode? {
        var headTmp = head
        var preTmp = head
        for _ in 0..<n {
            headTmp = headTmp?.next
        }
        if headTmp == nil {
            return head?.next
        }
        while headTmp?.next != nil {
            headTmp = headTmp?.next
            preTmp = preTmp?.next
        }
        
        preTmp?.next = preTmp?.next?.next
        return head
    }
    
    func twoSum(_ nums: [Int], _ target: Int) -> [Int] {
        var map = [Int:Int]()
        for index in (0..<nums.count).reversed() {
            map[nums[index]] = index
        }
        for index in 0..<nums.count {
            if map[target - nums[index]] != nil  && index != map[target - nums[index]]{
                return [index,map[target - nums[index]]!]
            }
        }
        return []
    }
    
    func lengthOfLongSubString(str:String) -> Int {
        var subLength = 0 ;
        let len = str.count
        let array = Array<Character>(str)
        var hashMap = [Character:Int]()
        var preIndex = -1
        for index in 0..<len {
            let curCharacter =  array[index]
            if let curPrexIndex = hashMap[curCharacter] {
                preIndex = max(curPrexIndex, preIndex)
            } else {
                
            }
            hashMap[curCharacter] = index
            subLength = max(subLength, index - preIndex)
        }
        return subLength
    }
}
