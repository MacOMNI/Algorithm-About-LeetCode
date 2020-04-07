# LeetCode No.92  [Reverse Linked List II](https://leetcode.com/problems/reverse-linked-list-ii/)

### 题目描述

反转从位置 *m* 到 *n* 的链表。请使用一趟扫描完成反转。

**说明:**
1 ≤ *m* ≤ *n* ≤ 链表长度。

**示例:**

```
输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
输入: 1->2->3->4->5->NULL, m = 1, n = 1
输出: 1->2->3->4->5->NULL
```

### 题目解析
关键因素分析:
1. 反转链表(如何反转);
2. 区间操作(m,n);
3. 一趟扫描;

解决问题:

1. 划分三部分 , head ,between,tail;

2. 对 between 进行链表反转;
3. 拼接 head + between + tail;

测试 cases:

```
1. 1->2->3->4->5->NULL, m = 1, n = 5
2. 1->2->3->4->5->NULL, m = 5, n = 5
```
时间复杂度：O(n)
空间复杂度：O(n)


### 代码实现

`Swift` 版本实现:

```Swift


```

`Golang` 版本实现:

```golang

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	listHead := head
	var resNode *ListNode //结果 res
	var headTmp *ListNode
	var headLastTmp *ListNode
	//	var betweenLastTmp *ListNode
	var tailFirstTmp *ListNode

	if m-1 > 0 { // 有 head
		resNode = listHead
		for i := 1; i < m; i++ { // head
			headLastTmp = listHead
			listHead = listHead.Next
		}
		headLastTmp.Next = nil
		//resNode.Next = nil

	}
	headTmp = listHead.Next

	rverseHead := listHead
	rverseHead.Next = nil
	tailFirstTmp = headTmp
	for i := m; i < n; i++ { //btw
		tmp := headTmp.Next
		headTmp.Next = rverseHead
		tailFirstTmp = tmp
		rverseHead = headTmp
		headTmp = tmp

	}

	if headLastTmp != nil { //head + btw
		headLastTmp.Next = rverseHead
	} else {
		resNode = rverseHead
	}
	if tailFirstTmp != nil { //tail
		tmp := resNode
		for i := 1; i < n; i++ { // head
			tmp = tmp.Next
		}
		tmp.Next = tailFirstTmp
	}

	return resNode
}

```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|0 ms|2.2 MB	 |golang|
|Accepted|40 ms|21.3 MB	 |swift|