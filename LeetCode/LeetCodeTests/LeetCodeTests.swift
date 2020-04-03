//
//  LeetCodeTests.swift
//  LeetCodeTests
//
//  Created by 徐坤坤 on 2020/4/3.
//  Copyright © 2020 Xcode. All rights reserved.
//

import XCTest
@testable import LeetCode

class LeetCodeTests: XCTestCase {

    override func setUp() {
        // Put setup code here. This method is called before the invocation of each test method in the class.
    }

    override func tearDown() {
        // Put teardown code here. This method is called after the invocation of each test method in the class.
    }

    func testExample() {
       // print(Solution.init().twoSum([-2, 7, 11, 15],9))
        print(Solution.init().twoSum([3, 2, 4],6))
        print(Solution.init().twoSum([3, 3],6))

        // This is an example of a functional test case.
        // Use XCTAssert and related functions to verify your tests produce the correct results.
    }

    func testPerformanceExample() {
        // This is an example of a performance test case.
        self.measure {
            // Put the code you want to measure the time of here.
        }
    }

}
