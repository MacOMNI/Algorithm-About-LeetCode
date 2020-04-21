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
func Test_Delete2Node(t *testing.T) {
	left1Node := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	leftNode := &TreeNode{
		Val:   2,
		Left:  left1Node,
		Right: nil,
	}
	rightNode := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	rootNode := &TreeNode{
		Val:   3,
		Right: rightNode,
		Left:  leftNode,
	}
	fmt.Println(levelOrder(rootNode))
	fmt.Println(levelOrder(deleteNode(rootNode, 3)))

}
func Test_Delete5Node(t *testing.T) {
	left1Node := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	leftNode := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: left1Node,
	}
	rightNode := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	rootNode := &TreeNode{
		Val:   3,
		Right: rightNode,
		Left:  leftNode,
	}
	// 3 (1,2), 4
	fmt.Println(levelOrder(rootNode))
	fmt.Println(levelOrder(deleteNode(rootNode, 3)))

}
func Test_Delete3Node(t *testing.T) {
	right1Node := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	left1Node := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: right1Node,
	}
	right2Node := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	leftNode := &TreeNode{
		Val:   3,
		Left:  left1Node,
		Right: right2Node,
	}

	rightNode := &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	rootNode := &TreeNode{
		Val:   5,
		Right: rightNode,
		Left:  leftNode,
	}
	fmt.Println(levelOrder(rootNode))
	fmt.Println(levelOrder(deleteNode(rootNode, 3)))

}
func Test_Delete1Node(t *testing.T) {

	levelLNode := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	levelRNode := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	level1RNode := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	leftNode := &TreeNode{
		Val:   3,
		Left:  levelLNode,
		Right: levelRNode,
	}
	rightNode := &TreeNode{
		Val:   6,
		Left:  nil,
		Right: level1RNode,
	}
	rootNode := &TreeNode{
		Val:   5,
		Right: rightNode,
		Left:  leftNode,
	}
	fmt.Println(levelOrder(rootNode))
	fmt.Println(levelOrder(deleteNode(rootNode, 7)))
	//fmt.Println(levelOrder(deleteNode(rootNode, 6)))

}
func Test_DeleteNode(t *testing.T) {

	levelLNode := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	levelRNode := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	level1RNode := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	leftNode := &TreeNode{
		Val:   3,
		Left:  levelLNode,
		Right: levelRNode,
	}
	rightNode := &TreeNode{
		Val:   6,
		Left:  nil,
		Right: level1RNode,
	}
	rootNode := &TreeNode{
		Val:   5,
		Right: rightNode,
		Left:  leftNode,
	}
	fmt.Println(levelOrder(rootNode))
	fmt.Println(levelOrder(deleteNode(rootNode, 3)))

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
