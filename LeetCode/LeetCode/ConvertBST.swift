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
    var sum = 0
    //https://leetcode-cn.com/problems/convert-bst-to-greater-tree/
    func ConverBST(root: TreeNode?) -> TreeNode? {
        guard let root = root else {
            return nil
        }
        ConverBST(root: root.right)
        // print(root.val,"sum = ",sum)
        sum += root.val
        root.val = sum
        ConverBST(root: root.left)
        return root
        
    }
}
