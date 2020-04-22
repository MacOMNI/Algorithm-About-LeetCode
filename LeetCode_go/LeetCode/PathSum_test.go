package LeetCode

import (
	"fmt"
	"testing"
)

func Test_PathSum(t *testing.T) {
	right1Node := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	left2Node := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	left1Node := &TreeNode{
		Val:   11,
		Left:  left2Node,
		Right: right1Node,
	}

	leftNode := &TreeNode{
		Val:   4,
		Left:  left1Node,
		Right: nil,
	}
	rightRLNode := &TreeNode{
		Val:   13,
		Left:  nil,
		Right: nil,
	}
	rightRRLNode := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	rightRRRNode := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	rightRRNode := &TreeNode{
		Val:   4,
		Left:  rightRRLNode,
		Right: rightRRRNode,
	}
	rightNode := &TreeNode{
		Val:   8,
		Left:  rightRLNode,
		Right: rightRRNode,
	}

	rootNode := &TreeNode{
		Val:   5,
		Right: rightNode,
		Left:  leftNode,
	}
	fmt.Println(levelOrder(rootNode))
	fmt.Println(pathSum(rootNode, 22))

}
