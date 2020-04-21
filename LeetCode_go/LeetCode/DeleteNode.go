package LeetCode

func deleteNode(root *TreeNode, key int) *TreeNode {
	var findNode *TreeNode
	var preNode *TreeNode //:= root
	node := root
	for node != nil {
		if node.Val == key {
			findNode = node
			break
		}
		preNode = node
		if node.Val > key {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	if findNode != nil {
		if findNode.Left == nil && findNode.Right == nil {
			if preNode == nil {
				return nil
			}
			if preNode.Val > key {
				preNode.Left = nil
			} else {
				preNode.Right = nil
			}
			return root
		}
		if findNode.Left == nil {
			if preNode == nil {
				return findNode.Right
			}
			if preNode.Val > key {
				preNode.Left = findNode.Right
			} else {
				preNode.Right = findNode.Right
			}
			return root
		}
		if findNode.Right == nil {
			if preNode == nil {
				return findNode.Left
			}
			if preNode.Val > key {
				preNode.Left = findNode.Left
			} else {
				preNode.Right = findNode.Left
			}
			return root
		}
		var preMaxNode *TreeNode
		node = node.Left

		for node != nil {
			if node.Right == nil {
				break
			}
			preMaxNode = node
			node = node.Right
		}
		if preMaxNode != nil {
			preMaxNode.Right = nil
		} else {
			findNode.Left = nil
		}
		if preNode == nil {
			if findNode.Left != nil {
				node.Left = findNode.Left
			}
			node.Right = findNode.Right
			return node
		} else {
			if findNode.Left != nil {
				node.Left = findNode.Left
			}
			node.Right = findNode.Right
			if preNode.Val > key {
				preNode.Left = node
			} else {
				preNode.Right = node
			}
			return root
		}

	}
	return root
}
