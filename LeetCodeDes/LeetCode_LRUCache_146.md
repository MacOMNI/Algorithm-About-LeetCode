# LeetCode No.146  [LRUCache](https://leetcode.com/problems/lru-cache/)

### 题目描述

运用你所掌握的数据结构，设计和实现一个  [LRU (最近最少使用) 缓存机制](https://baike.baidu.com/item/LRU)。它应该支持以下操作： 获取数据 `get` 和 写入数据 `put` 。

获取数据 `get(key)` - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
写入数据 `put(key, value)` - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。

**进阶:**

你是否可以在 **O(1)** 时间复杂度内完成这两种操作？

**示例:**

```
LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // 返回  1
cache.put(3, 3);    // 该操作会使得密钥 2 作废
cache.get(2);       // 返回 -1 (未找到)
cache.put(4, 4);    // 该操作会使得密钥 1 作废
cache.get(1);       // 返回 -1 (未找到)
cache.get(3);       // 返回  3
cache.get(4);       // 返回  4
```


### 题目解析
关键因素分析:
1. hashmap key value;
2. linkList 双向链表 get put ;

时间复杂度：O(1)
空间复杂度：O(n)


### 代码实现

`Golang` 版本实现:

```golang


type LRUCache struct {
	Capacity int
	Hashmap  map[int]*list.Element
	LinkList *list.List
}
type Link struct {
	Key   int
	Value int
}

func Constructor(capacity int) LRUCache {

	return LRUCache{
		Capacity: capacity,
		Hashmap:  make(map[int]*list.Element, capacity),
		LinkList: list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if element, ok := this.Hashmap[key]; ok {
		//	delete(this.Hashmap, key)
		hashKey, _ := element.Value.(*Link)
		this.LinkList.MoveToFront(element)
		element = this.LinkList.Front()
		this.Hashmap[key] = element
		return hashKey.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	link := &Link{
		Key:   key,
		Value: value,
	}
	var element *list.Element

	if hashElement, ok := this.Hashmap[key]; ok {
		hashElement.Value.(*Link).Value = value
		this.LinkList.MoveToFront(hashElement)
		element = this.LinkList.Front()
	} else {
		if this.LinkList.Len() < this.Capacity {
			element = this.LinkList.PushFront(link)

		} else {
			lastElement := this.LinkList.Back()
			delete(this.Hashmap, lastElement.Value.(*Link).Key)
			this.LinkList.Remove(lastElement)
			element = this.LinkList.PushFront(link)
		}

	}
	this.Hashmap[key] = element

}


```

| Status | Runtime | Memory |Language|
|:-------:|:-------:|:------|:------|
|Accepted|128 ms|17.3 MB	 |golang|
