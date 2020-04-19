# LeetCode No.102  [Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/)

### 题目描述

给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

例如:
给定二叉树: `[3,9,20,null,null,15,7]`,

```
    3
   / \
  9  20
    /  \
   15   7
```

返回其层次遍历结果：

```
[
  [3],
  [9,20],
  [15,7]
]
```


### 题目解析
关键因素分析:

1. 层序 遍历出 二维数组;
2. 确定一层的关系,用队列,把当前层添加进入,并在末尾 新增nil,表示这层最后一个元素;


时间复杂度：O(n)
空间复杂度：O(n)


### 代码实现

`Golang` 版本实现:

```golang


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


```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|0 ms|3.2 MB	 |golang|
