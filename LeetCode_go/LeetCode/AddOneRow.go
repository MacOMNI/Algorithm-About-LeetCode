package LeetCode

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		tmp := &TreeNode{
			Val:  v,
			Left: root,
		}
		return tmp
	}
	level := 1
	q := make([]*TreeNode, 0)
	q = append(q, root)
	q = append(q, nil)
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			if level+1 > d {
				break
			}
			if len(q) > 0 {
				level++
				q = append(q, nil)
			}
			continue
		}
		if level+1 == d {
			node.Left = &TreeNode{
				Val:  v,
				Left: node.Left,
			}
			node.Right = &TreeNode{
				Val:   v,
				Right: node.Right,
			}
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return root
}
