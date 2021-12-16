/*
	lru思路
	1. 要满足查询时间复杂度为O(1)，需要用map记录
	2. 要满足最近最少使用删除原则，需要用双向链表记录每个key的使用时间顺序
	3. 故最后lru结构为 哈希表+双链表
*/
package main

func main() {

}

type LRUCache struct {
	capacity   int // LRU容量
	dict       map[int]*DLinkNode
	head, tail *DLinkNode
}

type DLinkNode struct {
	key        int
	value      int
	prev, next *DLinkNode
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		head:     initDLinkNode(0, 0),
		tail:     initDLinkNode(0, 0),
		dict:     make(map[int]*DLinkNode),
	}

	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func initDLinkNode(key int, value int) *DLinkNode {
	return &DLinkNode{
		key:   key,
		value: value,
	}
}

// 需要O(1)时间复杂度
func (this *LRUCache) Get(key int) int {
	// 从dict中获取，不存在则返回-1
	dLinkNode, ok := this.dict[key]
	if !ok {
		return -1
	}

	// 更新lru链表
	this.moveToHead(dLinkNode)
	return dLinkNode.value
}

// 需要O(1)时间复杂度
func (this *LRUCache) Put(key int, value int) {
	// 先check是否存在于dict中
	if dlinkNode, ok := this.dict[key]; !ok {
		// 不在dict，需写入
		newNode := initDLinkNode(key, value)
		this.dict[key] = newNode
		// 写到链表头结点
		this.addHeadNode(newNode)
		if this.capacity < len(this.dict) {
			// lru已满，从尾结点移除一元素
			oldNode := this.removeTail()
			delete(this.dict, oldNode.key)
		}
	} else {
		// 在dict中
		dlinkNode.value = value
		// 更新lru
		this.moveToHead(dlinkNode)
	}
}

// 节点插入到链表头
func (this *LRUCache) addHeadNode(node *DLinkNode) {
	// 修改当前节点prev, next指针
	node.prev = this.head
	node.next = this.head.next

	// 修改头结点后一个节点的prev指针
	this.head.next.prev = node
	// 修改头结点next指针
	this.head.next = node
}

// 旧节点移动到链表头
func (this *LRUCache) moveToHead(node *DLinkNode) {
	// !这里顺序很重要，必须先删除原节点在链表中的位置，再移至链表头
	// 删除当前节点原始位置
	this.removeNode(node)
	// 移至链表头
	this.addHeadNode(node)
}

func (this *LRUCache) removeNode(node *DLinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) removeTail() *DLinkNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
