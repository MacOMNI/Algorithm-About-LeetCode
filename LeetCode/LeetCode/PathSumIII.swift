//
//  PathSumIII.swift
//  LeetCode
//
//  Created by Mackun on 2020/10/30.
//  Copyright © 2020 Xcode. All rights reserved.
//

import Foundation

class PathSum {
//    func pathSum(_ root: TreeNode?, _ sum: Int) -> Int {
//        guard let root = root else {
//            return 0
//        }
//        let cnt = (root.val == sum ? 1 : 0) // 满足当前节点
//        return cnt  + pathSum(root.left,sum) + pathSum(root.right, sum) + pathNodeSum(root.left, sum - root.val) + pathNodeSum(root.right, sum - root.val)
//    }
//    func pathNodeSum(_ root: TreeNode?, _ sum: Int) -> Int {
//        guard let root = root else {
//            return 0
//        }
//        let cnt = (root.val == sum ? 1 : 0) // 满足当前节点
//
//       // print(root.val,sum,cnt)
//        return cnt  + pathNodeSum(root.left, sum - root.val) + pathNodeSum(root.right, sum - root.val)
//    }
    func pathSum(_ root: TreeNode?, _ sum: Int) -> Int {
        guard let root = root else {
            return 0
        }
        let cnt = (root.val == sum ? 1 : 0) // 满足当前节点
        return cnt  + pathNodeSum(root.left,sum) + pathNodeSum(root.right, sum) + pathSum(root.left, sum - root.val) + pathSum(root.right, sum - root.val)
    }
    func pathNodeSum(_ root: TreeNode?, _ sum: Int) -> Int {
        guard let root = root else {
            return 0
        }
        let cnt = (root.val == sum ? 1 : 0) // 满足当前节点
        // print(root.val,sum,cnt)
        return cnt  + pathNodeSum(root.left, sum - root.val) + pathNodeSum(root.right, sum - root.val)
    }
}
