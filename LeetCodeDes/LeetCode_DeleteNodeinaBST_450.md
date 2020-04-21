# LeetCode No.450  [Delete Node in a BST](https://leetcode.com/problems/delete-node-in-a-bst/)

### 题目描述

给定一个二叉搜索树的根节点 `root` 和一个值 `key`，删除二叉搜索树中的 `key` 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。
说明： 要求算法时间复杂度为 O(h)，h 为树的高度。

**示例:**

```
root = [5,3,6,2,4,null,7]
key = 3

    5
   / \
  3   6
 / \   \
2   4   7

给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。

一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。

    5
   / \
  4   6
 /     \
2       7

另一个正确答案是 [5,2,6,null,4,null,7]。

    5
   / \
  2   6
   \   \
    4   7

```

### 题目解析
1. find keyNode 以及 preKeyNode ;
2. 如果 无 left right 节点, 即为叶子节点直接删除;
3. 如果 有 left 无 right 节点,  preKeyNode的left or right 指向 keyNode.Left;
4. 如果 无 left 有 right 节点,  preKeyNode的left or right 指向 keyNode.Right;
5. 如果 有 left 有 right 节点,  preKeyNode的left or right 指向 keyNode.Left 节点中的最大数值 preMaxNode 以及 maxNode(包含 是否有preKeyNode, maxNode 是否有 left等边界情况考量);
时间复杂度：O(n)
空间复杂度：O(n)


### 代码实现

`Golang` 版本实现:

```golang
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
			if node.Left == nil {
				preMaxNode.Right = nil
			} else {
				preMaxNode.Right = node.Left
			}
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


```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|20 ms|6.7 MB	 |golang|
