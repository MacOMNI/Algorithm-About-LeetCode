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
func findNext(root *Node) *Node {
    next := root.Next
    for next != nil {
        if next.Left != nil {
           return  next.Left
        } else if next.Right != nil {
           return next.Right
        }
        next = next.Next
    }		
    return nil
}

func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}
    if root.Right != nil {
        root.Right.Next = findNext(root)
    }
	if root.Left != nil {
		if root.Right == nil {
		    root.Left.Next = findNext(root)
		} else {
			root.Left.Next = root.Right
		}
    }
   	connect2(root.Right)
	connect2(root.Left)
	return root
}
