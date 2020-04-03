//
//  TowSum.swift
//  LeetCode
//
//  Created by MacOMNI on 2020/4/3.
//  Copyright © 2020 Xcode. All rights reserved.
//

import UIKit

class Solution {
        func twoSum(_ nums: [Int], _ target: Int) -> [Int] {
            var map = [Int:Int]()
            for index in (0..<nums.count).reversed() {
                map[nums[index]] = index
            }
            for index in 0..<nums.count {
                if map[target - nums[index]] != nil  && index != map[target - nums[index]]{
                    return [index,map[target - nums[index]]!]
                }
            }
            return []
        }
}
