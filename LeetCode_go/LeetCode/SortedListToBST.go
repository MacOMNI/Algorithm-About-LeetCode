package LeetCode

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var nodeList []int

func sortedListToBST(head *ListNode) *TreeNode {
	nodeList = make([]int, 0)
	for head != nil {
		nodeList = append(nodeList, head.Val)
		head = head.Next
	}
	//fmt.Println(nodeList)
	return buildTree(nodeList)
}
func buildTree(array []int) *TreeNode {
	n := len(array)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return &TreeNode{
			Val:   array[0],
			Left:  nil,
			Right: nil,
		}
	}
	rootIndex := n / 2
	// fmt.Println(rootIndex)
	// fmt.Println(rootIndex)

	return &TreeNode{
		Val:   array[rootIndex],
		Left:  buildTree(array[:rootIndex]),
		Right: buildTree(array[rootIndex+1:]),
	}
}
