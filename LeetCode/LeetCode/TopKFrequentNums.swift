//
//  TopKFrequentNums.swift
//  LeetCode
//
//  Created by Mackun on 2020/10/25.
//  Copyright © 2020 Xcode. All rights reserved.
//

import Foundation

class TopKNumbers {
    // hashmap 存 cnt
    // 滑块 l -> k
    func topKNumbers(nums: [Int], k: Int) -> [Int] {
        var topKnumbers = [Int]()
        var hashMaps = [Int:Int]()
        for item in nums {
            if let cnt = hashMaps[item] {
                hashMaps[item] = cnt + 1
            } else {
                hashMaps[item] = 1
            }
        }
        let hashDic = hashMaps.sorted { (item1, item2) -> Bool in
            return item1.value > item2.value
        }
        var itemCnt = 0
        for (key,_) in hashDic {

            if itemCnt + 1 <= k {
                topKnumbers.append(key)
                itemCnt = itemCnt + 1
            }
        }
        
        return topKnumbers
    }
    //    func insertBinarySearch(nums: [Int],value: Int) -> Int  {
    //
    //        return nums;
    //    }
}
