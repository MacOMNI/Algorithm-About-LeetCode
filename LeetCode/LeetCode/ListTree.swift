//
//  ListTree.swift
//  LeetCode
//
//  Created by Mackun on 2020/11/20.
//  Copyright © 2020 Xcode. All rights reserved.
//

import Foundation

class ListTree {
    func flatten(node: TreeNode?) -> Void {
        guard let node = node else {
            return
        }
       // let left = node.left
        let  right = node.right
        flatten(node: node.left)
        node.right = node.left
        flatten(node: right)
        node.left = nil
       // node.right = left
    }
}
