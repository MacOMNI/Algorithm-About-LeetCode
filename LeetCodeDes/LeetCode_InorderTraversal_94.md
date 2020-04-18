# LeetCode No.94  [Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/)

### 题目描述

给定一个二叉树，返回它的*中序* 遍历。

**示例:**

```
输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
```


### 题目解析
关键因素分析:
1. *中序* 遍历 , L->ROOT->R;
2. 递归 ,非递归;

时间复杂度：O(n)
空间复杂度：O(n)


### 代码实现

`Golang` 版本实现:

```golang
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, []int{root.Val}...)

	res = append(res, inorderTraversal(root.Right)...)
	// res = append(res)
	return res
}



```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|0 ms|2.1 MB	 |golang|
