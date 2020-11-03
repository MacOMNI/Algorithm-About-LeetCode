//
//  TargetSumWays.swift
//  LeetCode
//
//  Created by Mackun on 2020/11/3.
//  Copyright © 2020 Xcode. All rights reserved.
//

import Foundation

class TargetSumWays {
    //https://leetcode-cn.com/problems/target-sum/
    var ans = 0
    var sums = Array<Int>()
    func findTargetSumWays(_ nums: [Int], _ S: Int) -> Int {
        let n = nums.count
        sums  = Array.init(repeating: 0, count: n + 1)
        for index in 0..<n {
            sums[index+1] += sums[index] + nums[index]
        }
        targetSums(nums: nums,start: 0,target: S)
        return ans
    }
    func targetSums(nums: [Int],start: Int,target: Int){
        let n = nums.count
        if start == n  {
            if target == 0 {
                 ans += 1
            }
        } else {
            targetSums(nums: nums,start: start+1,target: target - nums[start])
            targetSums(nums: nums,start: start+1,target: target + nums[start])
        }
//        if start < n && target >= -(sums[n] - sums[start]) && target <= (sums[n] - sums[start]) {
//
//        }
       
    }
}
