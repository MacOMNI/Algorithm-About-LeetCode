# LeetCode No.113  [Path Sum II](https://leetcode.com/problems/path-sum-ii/)

### 题目描述

给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明: 叶子节点是指没有子节点的节点。

**示例:**

```
给定如下二叉树，以及目标和 sum = 22，

   			  5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1

返回:
[
   [5,4,11,2],
   [5,8,4,5]
]

```

### 题目解析
关键因素分析:
1. 所有集合;
2. 递归求解;


时间复杂度：O(n)
空间复杂度：O(n)


### 代码实现

`Golang` 版本实现:

```golang

func pathSum(root *TreeNode, sum int) [][]int {
	var path []int
	var res [][]int

	return dfs(res, path, root, sum)
}

func dfs(res [][]int, path []int, root *TreeNode, sum int) [][]int {
	if root == nil {
		return res
	}
	path = append(path, root.Val)

	if nil == root.Left && nil == root.Right { //叶子节点
		if sum == root.Val {
			slice := make([]int, len(path))
			copy(slice[0:], path[0:])
			res = append(res, slice)
		}
	}

	if root.Left != nil {
		res = dfs(res, path, root.Left, sum-root.Val)
	}
	if root.Right != nil {
		res = dfs(res, path, root.Right, sum-root.Val)

	}
	return res
}


```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|4 ms|4.6 MB	 |golang|
