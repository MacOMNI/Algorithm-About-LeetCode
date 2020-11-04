//
//  FindKthLargest.swift
//  LeetCode
//
//  Created by Mackun on 2020/11/4.
//  Copyright © 2020 Xcode. All rights reserved.
//

import Foundation

class FindKthLargest {
    // 3,2,1,5,6,4  k = 2
    func findKthLargest(nums: [Int],k: Int) -> Int {
        var nums = nums
        let cnt = nums.count
        var l = 0, r = cnt - 1
        let tmp:Int = nums[0]
        if cnt == 1 {
            return nums[l]
        }
        if cnt == 2 {
            return k == 1 ? max(nums[0], nums[1]) :min(nums[0],nums[1])
        }
        while l < r {
            while l < r && tmp >= nums[r] {
                          r -= 1
            }
            while l < r && tmp <= nums[l] {
                l += 1
            }
            if l < r {
                (nums[l],nums[r]) = (nums[r],nums[l])
            }
        }
        nums[0] = nums[l]
        nums[l] = tmp
        print(nums)
        if l <  k - 1 {
           return findKthLargest(nums: Array(nums[l+1..<cnt]), k: k - l - 1)
        } else if l >  k - 1 {
           return findKthLargest(nums: Array(nums[0..<l]), k: k)
        }
        return nums[k - 1]
    }
    
}
