//
//  PathSumIII.swift
//  LeetCode
//
//  Created by Mackun on 2020/10/30.
//  Copyright © 2020 Xcode. All rights reserved.
//

import Foundation

class PathSum {
    func pathSum(_ root: TreeNode?, _ sum: Int) -> Int {
        var cnt = 0
        if sum == 0 {
           cnt = 1
        }
        guard let root = root else {
            return 0
        }
        return cnt + pathSum(root.left,sum) + pathSum(root.right,sum) + pathSum(root.left,sum - root.val) + pathSum(root.right,sum - root.val)
    }
}
