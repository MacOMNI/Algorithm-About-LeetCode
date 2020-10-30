//
//  InvertTree.swift
//  LeetCode
//
//  Created by Mackun on 2020/10/30.
//  Copyright © 2020 Xcode. All rights reserved.
//

import Foundation

class InvertTree {
    func invertTree(_ root: TreeNode?) -> TreeNode? {
        guard let root = root else {
            return nil
        }
        let tmp = root.right
        root.right = invertTree(root.left)
        root.left = invertTree(tmp)
        return root
    }}
