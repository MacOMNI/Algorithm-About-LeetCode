//
//  ConvertBST.swift
//  LeetCode
//
//  Created by Mackun on 2020/10/25.
//  Copyright © 2020 Xcode. All rights reserved.
//

import Foundation

public class TreeNode {
    public var val: Int
    public var left: TreeNode?
    public var right: TreeNode?
    init() {
        self.val = 0
        self.right = nil
        self.left = nil
    }
    init(val :Int) {
        self.val = val
        self.left = nil
        self.right = nil
    }
    init(val: Int,left: TreeNode?,right: TreeNode?) {
        self.val = val
        self.left = left
        self.right = right
    }
}

class ConverBST {
    //https://leetcode-cn.com/problems/convert-bst-to-greater-tree/
    func ConverBST(root: TreeNode?) -> TreeNode? {
        //var sum = 0
        guard let root = root else {
            return nil
        }
        if root.right != nil {
            root.right?.val += (ConverBST(root: root.right?.right)?.val ?? 0)
        }
        root.val += root.right?.val ?? 0
        if root.left != nil {
            root.left?.val += root.val + (ConverBST(root: root.left?.right)?.val ?? 0)
        }
        return root

    }
}
