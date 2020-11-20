//
//  SymmetricTree.swift
//  LeetCode
//
//  Created by Mackun on 2020/11/5.
//  Copyright © 2020 Xcode. All rights reserved.
//

import Foundation

class SymmetricTree {
    func isSymmetric(_ root: TreeNode?) -> Bool {
        return symmetricTree(root: root?.left, newRoot: root?.right)
    }
    func symmetricTree(root: TreeNode?,newRoot: TreeNode?) -> Bool {
        if root == nil && newRoot == nil {
            return true
        }
        guard let root = root , let newRoot = newRoot else {
            return false
        }
        if root.val != newRoot.val {
            return false
        }
        
        return symmetricTree(root: root.left, newRoot: newRoot.right) && symmetricTree(root: root.right, newRoot: newRoot.left)
    }
}
