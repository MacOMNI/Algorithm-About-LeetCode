package LeetCode

import "container/list"

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
