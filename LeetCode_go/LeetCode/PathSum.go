package LeetCode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

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
