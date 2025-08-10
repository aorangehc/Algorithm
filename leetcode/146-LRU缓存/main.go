package main

import (
	"container/list"
	"fmt"
)

type entry struct {
	key   int
	value int
}

type LRUCache struct {
	capacity int
	list     *list.List
	cache    map[int]*list.Element
}

// Constructor 创建缓存空间
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		cache:    make(map[int]*list.Element),
	}
}

// Get 查询缓存，如果找不到返回-1，将查询到的node放到最前面，返回查询到的值
func (this *LRUCache) Get(key int) int {
	node := this.cache[key]

	if node == nil {
		return -1
	}
	this.list.MoveToFront(node)

	return node.Value.(entry).value
}

// Put 先判断key存不存在，如果存在更新值并将node放到缓存前面；如果存不存在，就建立新的节点，并检查是否要进行扩容操作
func (this *LRUCache) Put(key int, value int) {
	// 先查看key存不存在，如果存在覆盖原值
	if node := this.cache[key]; node != nil {
		node.Value = entry{key, value}

		this.list.MoveToFront(node)
		return
	}

	// 建立新的节点，放到list和缓存中
	this.cache[key] = this.list.PushFront(entry{key, value})

	if len(this.cache) > this.capacity {
		removed := this.list.Back()

		if removed != nil {
			this.list.Remove(removed)
			delete(this.cache, removed.Value.(entry).key)
		}
	}
}

func main() {
	obj := Constructor(2)

	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))

	obj.Put(3, 3)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(3))

	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))

}
