# LeetCode No.1  [TwoSum](https://leetcode.com/problems/two-sum/solution/)

### 题目描述

给出二叉 搜索 树的根节点，该二叉树的节点值各不相同，修改二叉树，使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。

提醒一下，二叉搜索树满足下列约束条件：

节点的左子树仅包含键 小于 节点键的节点。
节点的右子树仅包含键 大于 节点键的节点。
左右子树也必须是二叉搜索树。



**示例:**

```
输入：[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
输出：[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]

```

### 题目解析
1. 后序遍历;
2. count 统计;
3. 二分查找;
时间复杂度：O(nlgn)
空间复杂度：O(n)


### 代码实现

`Golang` 版本实现:

```golang
var nodes []int
var counts []int
func bstToGst(root *TreeNode) *TreeNode {
    
    //nodes := make([]int,0)
    counts = make([]int,0)
    nodes = treeOrder(root)
   // fmt.Println(nodes)
    counts = append(counts,nodes[0])
    for i:= 1;i< len(nodes);i++{
      counts = append(counts,counts[i-1]+nodes[i])
    }
  //  fmt.Println(counts)
    updateTree(root)
    return root
}
func updateTree(root *TreeNode) {
    if root == nil {
        return
    }
    beforeIndex := binarySearch(root.Val)
   // fmt.Println(root.Val," ",beforeIndex)

    root.Val = counts[beforeIndex - 1]

    updateTree(root.Left)
    updateTree(root.Right)

}
func treeOrder(root *TreeNode) []int{
    var res  []int

    if root == nil {
        return res
    }
    res = append(res,treeOrder(root.Right)...)
    res = append(res,root.Val)
    res = append(res,treeOrder(root.Left)...)
    return res
}
func binarySearch(val int) int {
    left := 0
    right := len(nodes)
    mid := (left + right)>>1
    for left < right {
        if nodes[mid] >= val {
            left = mid+1
        } else {
            right = mid
        }
        mid = (left + right)>>1
    }
    return mid
}
```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|0 ms|3.8 MB	 |golang|
