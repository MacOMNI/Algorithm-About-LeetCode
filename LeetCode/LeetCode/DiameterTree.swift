//
//  DiameterTree.swift
//  LeetCode
//
//  Created by Mackun on 2020/11/5.
//  Copyright © 2020 Xcode. All rights reserved.
//

import Foundation

class DiameterTree {
    
    var ans = 0
    func diameterOfBinaryTree(_ root: TreeNode?) -> Int {
        DiameterTree(root: root)
        return ans;
    }
    func DiameterTree(root: TreeNode?) -> Int{
        guard let root = root else {
            return 0
        }
        var left = DiameterTree(root: root.left)
        var right  = DiameterTree(root: root.right)
        ans = max(left + right,ans)
        return max(left,right) + 1
    }
}
