//
//  ListTree.swift
//  LeetCode
//
//  Created by Mackun on 2020/11/20.
//  Copyright © 2020 Xcode. All rights reserved.
//

import Foundation

class ListTree {
    var next: TreeNode?
    func flatten(_ root: TreeNode?) {
        guard let root = root else {
            return
        }
        // let left = node.left
        flatten(root.right)
        flatten(root.left)
        
        root.right = next
        // flatten(right)
        root.left = nil
        next = root
    }
}
