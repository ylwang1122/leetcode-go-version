//运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。 
//
// 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。 
//写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的
//数据值留出空间。 
//
// 进阶: 
//
// 你是否可以在 O(1) 时间复杂度内完成这两种操作？ 
//
// 示例: 
//
// LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得密钥 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得密钥 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4
// 
// Related Topics 设计

package cn

import "container/list"

//leetcode submit region begin(Prohibit modification and deletion)
type cacheValue struct {
	ele *list.Element
	val int
}
type LRUCache struct {
	dataMap map[int]*cacheValue
	list *list.List
	count int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		dataMap: make(map[int]*cacheValue),
		list:    list.New(),
		count:	capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	// 判定是否存在，如果存在则调整list顺序，并返回，否则返回-1
	if val, ok := this.dataMap[key]; ok {
		this.list.MoveToBack(val.ele)
		return val.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.dataMap[key]; ok {
		ele := this.dataMap[key].ele
		this.list.MoveToBack(ele)
		this.dataMap[key].val = value
		return
	}

	this.list.PushBack(key)
	this.dataMap[key] = &cacheValue{
		ele: this.list.Back(),
		val: value,
	}
	if this.list.Len() > this.count {
		front := this.list.Front()
		this.list.Remove(front)
		delete(this.dataMap, front.Value.(int))
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
