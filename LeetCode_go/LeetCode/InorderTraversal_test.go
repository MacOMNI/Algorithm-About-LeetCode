package LeetCode

import (
	"fmt"
	"testing"
)

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func Test_InorderTraversal(t *testing.T) {

	leftNode := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	rightNode := &TreeNode{
		Val:   2,
		Left:  leftNode,
		Right: nil,
	}
	rootNode := &TreeNode{
		Val:   1,
		Right: rightNode,
		Left:  nil,
	}
	fmt.Println(inorderTraversal(rootNode))

}

func Test_LeverOrder(t *testing.T) {

	leftNode := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	rightNode := &TreeNode{
		Val:   2,
		Left:  leftNode,
		Right: nil,
	}
	rootNode := &TreeNode{
		Val:   1,
		Right: rightNode,
		Left:  nil,
	}
	fmt.Println(levelOrder(rootNode))

}
