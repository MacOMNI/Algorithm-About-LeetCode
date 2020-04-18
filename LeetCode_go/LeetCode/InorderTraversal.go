package LeetCode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
