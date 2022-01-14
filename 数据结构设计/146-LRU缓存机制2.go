/*
	lru思路:
		1. 要满足查询的时间复杂O(1)，需要用map记录
		2. 要满足最少使用删除原则，需要用双向链表记录每个key的使用时间顺序
		3. 最后lru结构为 哈希表 + 双向链表
*/
package main

func main() {

}

type LRUCache struct {
	capacity   int                // lru容量
	dict       map[int]*DLinkNode // map
	head, tail *DLinkNode
}

type DLinkNode struct {
	key        int
	value      int
	prev, next *DLinkNode
}

func Constructor(capacity int) *LRUCache {
	lru := &LRUCache{
		capacity: capacity,
		head:     initDLinkNode(0, 0),
		tail:     initDLinkNode(0, 0),
		dict:     make(map[int]*DLinkNode),
	}

	lru.head.next = lru.tail
	lru.tail.prev = lru.head

	return lru
}

func initDLinkNode(key, val int) *DLinkNode {
	return &DLinkNode{
		key:   key,
		value: val,
	}
}

func (this *LRUCache) Get(key int) int {
	dlinkNode, ok := this.dict[key]
	if !ok {
		return -1
	}

	// 更新当前节点位置，移至头结点
	this.moveToHead(dlinkNode)
	return dlinkNode.value
}

func (this *LRUCache) Put(key, val int) {
	// 是否在map中
	if dLinkNode, ok := this.dict[key]; !ok {
		// 不在map中，需写入
		newNode := initDLinkNode(key, val)
		this.dict[key] = newNode
		// 写到链表头结点
		this.addToHead(newNode)
		// 如果lru已满
		if this.capacity < len(this.dict) {
			// 从尾结点删除元素
			oldNode := this.removeTail()
			delete(this.dict, oldNode.key)
		}

	} else {
		// 在map中
		dLinkNode.value = val
		// 更新lru
		this.moveToHead(dLinkNode)
	}
}

// 头插法
func (this *LRUCache) addToHead(node *DLinkNode) {
	// 当前节点prev, next与头部拼接
	node.prev = this.head
	node.next = this.head.next

	// 头结点的下一节点的prev拼接该节点
	this.head.next.prev = node
	// 头结点next拼接该节点
	this.head.next = node

}

// 当前节点移至头结点
func (this *LRUCache) moveToHead(node *DLinkNode) {
	// 删除当前节点
	this.removeNode(node)

	// 头插法移至头结点
	this.addToHead(node)
}

// 删除节点
func (this *LRUCache) removeNode(node *DLinkNode) {
	// 当前节点的前后节点与其断连
	node.next.prev = node.prev
	node.prev.next = node.next
}

func (this *LRUCache) removeTail() *DLinkNode {
	// 获取尾部节点
	node := this.tail.prev
	// 删除节点
	this.removeNode(node)
	return node
}
