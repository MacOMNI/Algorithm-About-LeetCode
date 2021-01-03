//
//  LowestCommonAncestor.swift
//  LeetCode
//
//  Created by Mackun on 2020/12/29.
//  Copyright © 2020 Xcode. All rights reserved.
//

import Foundation

class LowestCommonAncestor {
    var ans : Int = 0
    func lowestCommonAncestor( _ root:TreeNode?, _ o1: Int, _ o2: Int) -> Int {
        commonAncestor(root: root, o1: o1, o2: o2, cnt: 0)
        return ans
    }
    
    func commonAncestor( root: TreeNode?, o1: Int, o2: Int, cnt: Int) {
      
        guard let root = root, cnt < 2 else {
            return
        }
        var res = cnt
        if root.val == o1 || root.val == o2 {
            res += 1
        }
        if res == 2 {
            ans = root.val
        } else {
            commonAncestor(root: root.left, o1: o1, o2: o2, cnt: res)
        }
       
    }
}
