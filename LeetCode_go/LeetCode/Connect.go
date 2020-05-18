package LeetCode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		root.Left.Next = root.Right
		if root.Next == nil {
			root.Right.Next = nil
		} else {
			root.Right.Next = root.Next.Left
		}
	}
	connect(root.Left)
	connect(root.Right)
	return root
}

func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		root.Left.Next = root.Right
		if root.Next == nil {
			root.Right.Next = nil
		} else {
			root.Right.Next = root.Next.Left
		}
	}
	connect2(root.Left)
	connect2(root.Right)
	return root
}
