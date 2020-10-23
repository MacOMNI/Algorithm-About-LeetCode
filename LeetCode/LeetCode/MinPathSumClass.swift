//
//  MinPathSumClass.swift
//  LeetCode
//
//  Created by Mackun on 2020/10/23.
//  Copyright © 2020 Xcode. All rights reserved.
//

import Foundation

class MinPathSum {
    func minPathSum(grid: [[Int]]) -> Int {
        let m = grid.count
        if let n = grid.first?.count {
            var dp = Array.init(repeating: Array.init(repeating:Int.max, count: n + 1), count: m+1)
            dp[0][0] = 0
            dp[0][1] = 0
            dp[1][0] = 0
            for i in 1...m {
                for j in 1...n {
                    dp[i][j] = min(dp[i][j-1],dp[i-1][j])+grid[i-1][j-1]
                }
            }
            return dp[m][n]
        }
        return 0
        
    }
}
