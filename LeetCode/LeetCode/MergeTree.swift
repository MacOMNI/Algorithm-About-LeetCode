//
//  MergeTree.swift
//  LeetCode
//
//  Created by Mackun on 2020/10/30.
//  Copyright © 2020 Xcode. All rights reserved.
//

import Foundation

class MergeTree {
    func mergeTrees(_ t1 : TreeNode?, _ t2: TreeNode?) -> TreeNode? {
          guard let t1 = t1 else {
              return t2
          }
          guard let t2 = t2 else {
              return t1
          }
          let head = TreeNode.init()
          head.val = t1.val + t2.val
          head.left = mergeTrees(t1.left, t2.left)
          head.right = mergeTrees(t1.right, t2.right)
          return head
    }
}
