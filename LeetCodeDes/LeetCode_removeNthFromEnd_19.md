# LeetCode No.19 [RemoveNthFromEnd](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)

### 题目描述

给定一个链表，删除链表的倒数第 *n* 个节点，并且返回链表的头结点。

**示例：**

```
给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
```

### 题目解析
关键因素分析:
1. 复制一条链,让原始链先跑 n 个位置,第二条链 再跑;
2. 那么 先跑的链 跑完后,第二条链剩余链表 tail 即为 倒数第n个后面的链表;
3. 要求 返回 删除 第 n 个元素的 头节点;
4. 需要组装出 front + tail 

时间复杂度：O(n)
空间复杂度：O(n)


### 代码实现

`Golang` 版本实现:

```golang

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nHead := head
	var tailHeadTmp *ListNode
	var headLastTmp *ListNode
	var resNode *ListNode
	i := 0
	count := 0
	for head.Next != nil {

		head = head.Next
		i++
		count++
		if i == n {
			resNode = nHead
			headLastTmp = nHead
			//head = head.Next
			//	head = head.Next

			//tailHead = nHead.Next
			break
		}
	}
	for head != nil {
		count++
		headLastTmp = nHead
		// headTmp = nHead
		nHead = nHead.Next
		tailHeadTmp = nHead
		head = head.Next

		//tailHead = tailHead.Next
	}
	if n == 1 {
		headLastTmp.Next = nil
		return resNode
	}
	if count-n == 0 {
		return tailHeadTmp
	}
	if tailHeadTmp != nil {
		headLastTmp.Next = tailHeadTmp.Next
	}

	return resNode
}


```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|0 ms|2.2 MB	 |golang|
