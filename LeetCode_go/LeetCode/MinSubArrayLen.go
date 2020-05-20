package LeetCode

func minSubArrayLen(s int, nums []int) int {
    left := 0
    ans := 0
    res := len(nums) + 1
    for index,val := range nums {
        ans = ans + val
        for ans >= s {
            res = min(res,index - left + 1)
            ans = ans - nums[left]
            left++
        }
    }
    if res == len(nums) + 1 {
        return 0
    }
    return res
}
func min(a,b int) int {
    if a > b {
        return b
    }
    return a
}