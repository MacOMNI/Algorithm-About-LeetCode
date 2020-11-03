//
//  UniquePaths.swift
//  LeetCode
//
//  Created by Mackun on 2020/11/3.
//  Copyright © 2020 Xcode. All rights reserved.
//

import Foundation

class UniquePaths {
    func uniquePaths(_ m: Int, _ n: Int) -> Int {
        var dp = Array.init(repeating: Array.init(repeating: 0, count: m+1), count: n+1)
        for i in 0..<n {
           dp[i][0] = 1
        }
        for i in 0..<m {
           dp[0][i] = 1
        }
        for i in 0..<n {
            for j in 0..<m {
                dp[i+1][j+1] = dp[i][j+1] + dp[i+1][j]
               // print("dp = ",i+1,j+1,dp[i+1][j+1])
            }
        }
        return dp[n-1][m-1]
    }
}
