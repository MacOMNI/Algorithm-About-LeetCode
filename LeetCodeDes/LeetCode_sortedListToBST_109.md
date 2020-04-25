# LeetCode No.109  [Convert Sorted List to Binary Search Tree](https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/)

### 题目描述
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
**示例:**

```
给定的有序链表： [-10, -3, 0, 5, 9],

一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5

```

### 题目解析
关键因素分析:
1. 遍历有序链表;
2. 递归构建平衡二叉树;

时间复杂度：O(n)
空间复杂度：O(n)


### 代码实现

`Golang` 版本实现:

```golang
var nodeList []int

func sortedListToBST(head *ListNode) *TreeNode {
	nodeList = make([]int, 0)
	for head != nil {
		nodeList = append(nodeList, head.Val)
		head = head.Next
	}
	//fmt.Println(nodeList)
	return buildTree(nodeList)
}
func buildTree(array []int) *TreeNode {
	n := len(array)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return &TreeNode{
			Val:   array[0],
			Left:  nil,
			Right: nil,
		}
	}
	rootIndex := n / 2
	// fmt.Println(rootIndex)
	// fmt.Println(rootIndex)

	return &TreeNode{
		Val:   array[rootIndex],
		Left:  buildTree(array[:rootIndex]),
		Right: buildTree(array[rootIndex+1:]),
	}
}


```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|480 ms|272 MB	 |golang|
