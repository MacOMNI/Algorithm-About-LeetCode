package LeetCode

import "container/list"

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	queue := list.New()
	queue.PushBack(root)
	queue.PushBack(nil)
	level := []int{}
	preNode := root
	for queue.Len() > 0 {
		node, _ := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())
		if node == nil {
			if len(level) > 0 {
				res = append(res, level)
			}
			level = []int{}
			if preNode != nil {
				queue.PushBack(nil)
			}
		} else {
			level = append(level, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		preNode = node

	}
	return res
}
